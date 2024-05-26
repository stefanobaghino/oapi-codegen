// Copyright 2019 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"golang.org/x/tools/imports"

	"github.com/stefanobaghino/oapi-codegen/v2/pkg/util"
)

// Embed the templates directory
//
//go:embed templates
var templates embed.FS

// globalState stores all global state. Please don't put global state anywhere
// else so that we can easily track it.
var globalState struct {
	options       Configuration
	spec          *libopenapi.DocumentModel[v3.Document]
	importMapping importMap
}

// goImport represents a go package to be imported in the generated code
type goImport struct {
	Name string // package name
	Path string // package path
}

// String returns a go import statement
func (gi goImport) String() string {
	if gi.Name != "" {
		return fmt.Sprintf("%s %q", gi.Name, gi.Path)
	}
	return fmt.Sprintf("%q", gi.Path)
}

// importMap maps external OpenAPI specifications files/urls to external go packages
type importMap map[string]goImport

// GoImports returns a slice of go import statements
func (im importMap) GoImports() []string {
	goImports := make([]string, 0, len(im))
	for _, v := range im {
		goImports = append(goImports, v.String())
	}
	return goImports
}

func constructImportMapping(importMapping map[string]string) importMap {
	var (
		pathToName = map[string]string{}
		result     = importMap{}
	)

	{
		var packagePaths []string
		for _, packageName := range importMapping {
			packagePaths = append(packagePaths, packageName)
		}
		sort.Strings(packagePaths)

		for _, packagePath := range packagePaths {
			if _, ok := pathToName[packagePath]; !ok {
				pathToName[packagePath] = fmt.Sprintf("externalRef%d", len(pathToName))
			}
		}
	}
	for specPath, packagePath := range importMapping {
		result[specPath] = goImport{Name: pathToName[packagePath], Path: packagePath}
	}
	return result
}

// Generate uses the Go templating engine to generate all of our server wrappers from
// the descriptions we've built up above from the schema objects.
// opts defines
func Generate(spec *libopenapi.DocumentModel[v3.Document], opts Configuration) (string, error) {
	// This is global state
	globalState.options = opts
	globalState.spec = spec
	globalState.importMapping = constructImportMapping(opts.ImportMapping)

	filterOperationsByTag(spec, opts)
	if !opts.OutputOptions.SkipPrune {
		pruneUnusedComponents(spec)
	}

	// if we are provided an override for the response type suffix update it
	if opts.OutputOptions.ResponseTypeSuffix != "" {
		responseTypeSuffix = opts.OutputOptions.ResponseTypeSuffix
	}

	if globalState.options.OutputOptions.ClientTypeName == "" {
		globalState.options.OutputOptions.ClientTypeName = defaultClientTypeName
	}

	// This creates the golang templates text package
	TemplateFunctions["opts"] = func() Configuration { return globalState.options }
	t := template.New("oapi-codegen").Funcs(TemplateFunctions)
	// This parses all of our own template files into the template object
	// above
	err := LoadTemplates(templates, t)
	if err != nil {
		return "", fmt.Errorf("error parsing oapi-codegen templates: %w", err)
	}

	// load user-provided templates. Will Override built-in versions.
	for name, template := range opts.OutputOptions.UserTemplates {
		utpl := t.New(name)

		txt, err := GetUserTemplateText(template)
		if err != nil {
			return "", fmt.Errorf("error loading user-provided template %q: %w", name, err)
		}

		_, err = utpl.Parse(txt)
		if err != nil {
			return "", fmt.Errorf("error parsing user-provided template %q: %w", name, err)
		}
	}

	ops, err := OperationDefinitions(spec, opts.OutputOptions.InitialismOverrides)
	if err != nil {
		return "", fmt.Errorf("error creating operation definitions: %w", err)
	}

	xGoTypeImports, err := OperationImports(ops)
	if err != nil {
		return "", fmt.Errorf("error getting operation imports: %w", err)
	}

	var typeDefinitions, constantDefinitions string
	if opts.Generate.Models {
		typeDefinitions, err = GenerateTypeDefinitions(t, spec, ops, opts.OutputOptions.ExcludeSchemas)
		if err != nil {
			return "", fmt.Errorf("error generating type definitions: %w", err)
		}

		constantDefinitions, err = GenerateConstants(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating constants: %w", err)
		}

		imprts, err := GetTypeDefinitionsImports(spec, opts.OutputOptions.ExcludeSchemas)
		if err != nil {
			return "", fmt.Errorf("error getting type definition imports: %w", err)
		}
		MergeImports(xGoTypeImports, imprts)
	}

	var irisServerOut string
	if opts.Generate.IrisServer {
		irisServerOut, err = GenerateIrisServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var echoServerOut string
	if opts.Generate.EchoServer {
		echoServerOut, err = GenerateEchoServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var chiServerOut string
	if opts.Generate.ChiServer {
		chiServerOut, err = GenerateChiServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var fiberServerOut string
	if opts.Generate.FiberServer {
		fiberServerOut, err = GenerateFiberServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var ginServerOut string
	if opts.Generate.GinServer {
		ginServerOut, err = GenerateGinServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var gorillaServerOut string
	if opts.Generate.GorillaServer {
		gorillaServerOut, err = GenerateGorillaServer(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var strictServerOut string
	if opts.Generate.Strict {
		var responses []ResponseDefinition
		if spec.Model.Components != nil {
			responses, err = GenerateResponseDefinitions("", spec.Model.Components.Responses)
			if err != nil {
				return "", fmt.Errorf("error generation response definitions for schema: %w", err)
			}
		}
		strictServerResponses, err := GenerateStrictResponses(t, responses)
		if err != nil {
			return "", fmt.Errorf("error generation response definitions for schema: %w", err)
		}
		strictServerOut, err = GenerateStrictServer(t, ops, opts)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
		strictServerOut = strictServerResponses + strictServerOut
	}

	var clientOut string
	if opts.Generate.Client {
		clientOut, err = GenerateClient(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating client: %w", err)
		}
	}

	var clientWithResponsesOut string
	if opts.Generate.Client {
		clientWithResponsesOut, err = GenerateClientWithResponses(t, ops)
		if err != nil {
			return "", fmt.Errorf("error generating client with responses: %w", err)
		}
	}

	var inlinedSpec string
	if opts.Generate.EmbeddedSpec {
		inlinedSpec, err = GenerateInlinedSpec(t, globalState.importMapping, spec)
		if err != nil {
			return "", fmt.Errorf("error generating Go handlers for Paths: %w", err)
		}
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	externalImports := append(globalState.importMapping.GoImports(), importMap(xGoTypeImports).GoImports()...)
	importsOut, err := GenerateImports(
		t,
		externalImports,
		opts.PackageName,
		opts.NoVCSVersionOverride,
	)
	if err != nil {
		return "", fmt.Errorf("error generating imports: %w", err)
	}

	_, err = w.WriteString(importsOut)
	if err != nil {
		return "", fmt.Errorf("error writing imports: %w", err)
	}

	_, err = w.WriteString(constantDefinitions)
	if err != nil {
		return "", fmt.Errorf("error writing constants: %w", err)
	}

	_, err = w.WriteString(typeDefinitions)
	if err != nil {
		return "", fmt.Errorf("error writing type definitions: %w", err)
	}

	if opts.Generate.Client {
		_, err = w.WriteString(clientOut)
		if err != nil {
			return "", fmt.Errorf("error writing client: %w", err)
		}
		_, err = w.WriteString(clientWithResponsesOut)
		if err != nil {
			return "", fmt.Errorf("error writing client: %w", err)
		}
	}

	if opts.Generate.IrisServer {
		_, err = w.WriteString(irisServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}

	}

	if opts.Generate.EchoServer {
		_, err = w.WriteString(echoServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.ChiServer {
		_, err = w.WriteString(chiServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.FiberServer {
		_, err = w.WriteString(fiberServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.GinServer {
		_, err = w.WriteString(ginServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.GorillaServer {
		_, err = w.WriteString(gorillaServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.Strict {
		_, err = w.WriteString(strictServerOut)
		if err != nil {
			return "", fmt.Errorf("error writing server path handlers: %w", err)
		}
	}

	if opts.Generate.EmbeddedSpec {
		_, err = w.WriteString(inlinedSpec)
		if err != nil {
			return "", fmt.Errorf("error writing inlined spec: %w", err)
		}
	}

	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer: %w", err)
	}

	// remove any byte-order-marks which break Go-Code
	goCode := SanitizeCode(buf.String())

	// The generation code produces unindented horrors. Use the Go Imports
	// to make it all pretty.
	if opts.OutputOptions.SkipFmt {
		return goCode, nil
	}

	outBytes, err := imports.Process(opts.PackageName+".go", []byte(goCode), nil)
	if err != nil {
		return "", fmt.Errorf("error formatting Go code %s: %w", goCode, err)
	}
	return string(outBytes), nil
}

func GenerateTypeDefinitions(t *template.Template, swagger *libopenapi.DocumentModel[v3.Document], ops []OperationDefinition, excludeSchemas []string) (string, error) {
	var allTypes []TypeDefinition
	if swagger.Model.Components != nil {
		schemaTypes, err := GenerateTypesForSchemas(t, swagger.Model.Components.Schemas, excludeSchemas)
		if err != nil {
			return "", fmt.Errorf("error generating Go types for component schemas: %w", err)
		}

		paramTypes, err := GenerateTypesForParameters(t, swagger.Model.Components.Parameters)
		if err != nil {
			return "", fmt.Errorf("error generating Go types for component parameters: %w", err)
		}
		allTypes = append(schemaTypes, paramTypes...)

		responseTypes, err := GenerateTypesForResponses(t, swagger.Model.Components.Responses)
		if err != nil {
			return "", fmt.Errorf("error generating Go types for component responses: %w", err)
		}
		allTypes = append(allTypes, responseTypes...)

		bodyTypes, err := GenerateTypesForRequestBodies(t, swagger.Model.Components.RequestBodies)
		if err != nil {
			return "", fmt.Errorf("error generating Go types for component request bodies: %w", err)
		}
		allTypes = append(allTypes, bodyTypes...)
	}

	// Go through all operations, and add their types to allTypes, so that we can
	// scan all of them for enums. Operation definitions are handled differently
	// from the rest, so let's keep track of enumTypes separately, which will contain
	// all types needed to be scanned for enums, which includes those within operations.
	enumTypes := allTypes
	for _, op := range ops {
		enumTypes = append(enumTypes, op.TypeDefinitions...)
	}

	operationsOut, err := GenerateTypesForOperations(t, ops)
	if err != nil {
		return "", fmt.Errorf("error generating Go types for component request bodies: %w", err)
	}

	enumsOut, err := GenerateEnums(t, enumTypes)
	if err != nil {
		return "", fmt.Errorf("error generating code for type enums: %w", err)
	}

	typesOut, err := GenerateTypes(t, allTypes)
	if err != nil {
		return "", fmt.Errorf("error generating code for type definitions: %w", err)
	}

	allOfBoilerplate, err := GenerateAdditionalPropertyBoilerplate(t, allTypes)
	if err != nil {
		return "", fmt.Errorf("error generating allOf boilerplate: %w", err)
	}

	unionBoilerplate, err := GenerateUnionBoilerplate(t, allTypes)
	if err != nil {
		return "", fmt.Errorf("error generating union boilerplate: %w", err)
	}

	unionAndAdditionalBoilerplate, err := GenerateUnionAndAdditionalProopertiesBoilerplate(t, allTypes)
	if err != nil {
		return "", fmt.Errorf("error generating boilerplate for union types with additionalProperties: %w", err)
	}

	typeDefinitions := strings.Join([]string{enumsOut, typesOut, operationsOut, allOfBoilerplate, unionBoilerplate, unionAndAdditionalBoilerplate}, "")
	return typeDefinitions, nil
}

// GenerateConstants generates operation ids, context keys, paths, etc. to be exported as constants
func GenerateConstants(t *template.Template, ops []OperationDefinition) (string, error) {
	constants := Constants{
		SecuritySchemeProviderNames: []string{},
	}

	providerNameMap := map[string]struct{}{}
	for _, op := range ops {
		for _, def := range op.SecurityDefinitions {
			providerName := SanitizeGoIdentity(def.ProviderName)
			providerNameMap[providerName] = struct{}{}
		}
	}

	var providerNames []string
	for providerName := range providerNameMap {
		providerNames = append(providerNames, providerName)
	}

	sort.Strings(providerNames)

	constants.SecuritySchemeProviderNames = append(constants.SecuritySchemeProviderNames, providerNames...)

	return GenerateTemplates([]string{"constants.tmpl"}, t, constants)
}

// GenerateTypesForSchemas generates type definitions for any custom types defined in the
// components/schemas section of the Swagger spec.
func GenerateTypesForSchemas(t *template.Template, schemas *orderedmap.Map[string, *base.SchemaProxy], excludeSchemas []string) ([]TypeDefinition, error) {
	excludeSchemasMap := make(map[string]bool)
	for _, schema := range excludeSchemas {
		excludeSchemasMap[schema] = true
	}
	types := make([]TypeDefinition, 0)
	// We're going to define Go types for every object under components/schemas
	for _, schemaName := range SortedKeys(schemas) {
		if _, ok := excludeSchemasMap[schemaName]; ok {
			continue
		}
		schemaRef := schemas.Value(schemaName)

		goSchema, err := GenerateGoSchema(schemaRef, []string{schemaName})
		if err != nil {
			return nil, fmt.Errorf("error converting Schema %s to Go type: %w", schemaName, err)
		}

		goTypeName, err := renameSchema(schemaName, schemaRef)
		if err != nil {
			return nil, fmt.Errorf("error making name for components/schemas/%s: %w", schemaName, err)
		}

		types = append(types, TypeDefinition{
			JsonName: schemaName,
			TypeName: goTypeName,
			Schema:   goSchema,
		})

		types = append(types, goSchema.GetAdditionalTypeDefs()...)
	}
	return types, nil
}

// GenerateTypesForParameters generates type definitions for any custom types defined in the
// components/parameters section of the Swagger spec.
func GenerateTypesForParameters(t *template.Template, params *orderedmap.Map[string, *v3.Parameter]) ([]TypeDefinition, error) {
	var types []TypeDefinition
	for _, paramName := range SortedKeys(params) {
		paramOrRef := params.Value(paramName)

		goType, err := paramToGoType(paramOrRef, nil)
		if err != nil {
			return nil, fmt.Errorf("error generating Go type for schema in parameter %s: %w", paramName, err)
		}

		goTypeName, err := renameParameter(paramName, paramOrRef)
		if err != nil {
			return nil, fmt.Errorf("error making name for components/parameters/%s: %w", paramName, err)
		}

		typeDef := TypeDefinition{
			JsonName: paramName,
			Schema:   goType,
			TypeName: goTypeName,
		}

		if paramOrRef.Schema != nil && paramOrRef.Schema.IsReference() && paramOrRef.Schema.GetReference() != "" { // TODO jvt is this the right logic?
			// Generate a reference type for referenced parameters
			refType, err := RefPathToGoType(paramOrRef.Schema.GetReference())
			if err != nil {
				return nil, fmt.Errorf("error generating Go type for (%s) in parameter %s: %w", paramOrRef.Schema.GetReference(), paramName, err)
			}
			typeDef.TypeName = SchemaNameToTypeName(refType)
		}

		types = append(types, typeDef)
	}
	return types, nil
}

// GenerateTypesForResponses generates type definitions for any custom types defined in the
// components/responses section of the Swagger spec.
func GenerateTypesForResponses(t *template.Template, responses *orderedmap.Map[string, *v3.Response]) ([]TypeDefinition, error) {
	var types []TypeDefinition

	for _, responseName := range SortedKeys(responses) {
		responseOrRef := responses.Value(responseName)

		// We have to generate the response object. We're only going to
		// handle media types that conform to JSON. Other responses should
		// simply be specified as strings or byte arrays.
		response := responseOrRef

		jsonCount := 0

		for mediaType := range ToMap(response.Content) {
			if util.IsMediaTypeJson(mediaType) {
				jsonCount++
			}
		}

		for _, mediaType := range SortedKeys(response.Content) {
			response := response.Content.Value(mediaType)
			if !util.IsMediaTypeJson(mediaType) {
				continue
			}

			goType, err := GenerateGoSchema(response.Schema, []string{responseName})
			if err != nil {
				return nil, fmt.Errorf("error generating Go type for schema in response %s: %w", responseName, err)
			}

			goTypeName, err := renameResponse(responseName, responseOrRef)
			if err != nil {
				return nil, fmt.Errorf("error making name for components/responses/%s: %w", responseName, err)
			}

			typeDef := TypeDefinition{
				JsonName: responseName,
				Schema:   goType,
				TypeName: goTypeName,
			}

			ref := responseOrRef.GoLow().GetReference()
			if ref != "" {
				// Generate a reference type for referenced parameters
				refType, err := RefPathToGoType(ref)
				if err != nil {
					return nil, fmt.Errorf("error generating Go type for (%s) in parameter %s: %w", ref, responseName, err)
				}
				typeDef.TypeName = SchemaNameToTypeName(refType)
			}

			if jsonCount > 1 {
				typeDef.TypeName = typeDef.TypeName + mediaTypeToCamelCase(mediaType)
			}

			types = append(types, typeDef)
		}
	}
	return types, nil
}

// GenerateTypesForRequestBodies generates type definitions for any custom types defined in the
// components/requestBodies section of the Swagger spec.
func GenerateTypesForRequestBodies(t *template.Template, bodies *orderedmap.Map[string, *v3.RequestBody]) ([]TypeDefinition, error) {
	var types []TypeDefinition

	for _, requestBodyName := range SortedKeys(bodies) {
		requestBodyRef := bodies.Value(requestBodyName)

		// As for responses, we will only generate Go code for JSON bodies,
		// the other body formats are up to the user.
		response := requestBodyRef
		for mediaType, body := range ToMap(response.Content) {
			if !util.IsMediaTypeJson(mediaType) {
				continue
			}

			goType, err := GenerateGoSchema(body.Schema, []string{requestBodyName})
			if err != nil {
				return nil, fmt.Errorf("error generating Go type for schema in body %s: %w", requestBodyName, err)
			}

			goTypeName, err := renameRequestBody(requestBodyName, requestBodyRef)
			if err != nil {
				return nil, fmt.Errorf("error making name for components/schemas/%s: %w", requestBodyName, err)
			}

			typeDef := TypeDefinition{
				JsonName: requestBodyName,
				Schema:   goType,
				TypeName: goTypeName,
			}

			if requestBodyRef.GoLow().IsReference() {
				// Generate a reference type for referenced bodies
				refType, err := RefPathToGoType(requestBodyRef.GoLow().GetReference())
				if err != nil {
					return nil, fmt.Errorf("error generating Go type for (%s) in body %s: %w", requestBodyRef.GoLow().GetReference(), requestBodyName, err)
				}
				typeDef.TypeName = SchemaNameToTypeName(refType)
			}
			types = append(types, typeDef)
		}
	}
	return types, nil
}

// GenerateTypes passes a bunch of types to the template engine, and buffers
// its output into a string.
func GenerateTypes(t *template.Template, types []TypeDefinition) (string, error) {
	m := map[string]TypeDefinition{}
	var ts []TypeDefinition

	for _, typ := range types {
		if prevType, found := m[typ.TypeName]; found {
			// If type names collide, we need to see if they refer to the same
			// exact type definition, in which case, we can de-dupe. If they don't
			// match, we error out.
			if TypeDefinitionsEquivalent(prevType, typ) {
				continue
			}
			// We want to create an error when we try to define the same type twice.
			return "", fmt.Errorf("duplicate typename '%s' detected, can't auto-rename, "+
				"please use x-go-name to specify your own name for one of them", typ.TypeName)
		}

		m[typ.TypeName] = typ

		ts = append(ts, typ)
	}

	context := struct {
		Types []TypeDefinition
	}{
		Types: ts,
	}

	return GenerateTemplates([]string{"typedef.tmpl"}, t, context)
}

func GenerateEnums(t *template.Template, types []TypeDefinition) (string, error) {
	enums := []EnumDefinition{}

	// Keep track of which enums we've generated
	m := map[string]bool{}

	// These are all types defined globally
	for _, tp := range types {
		if found := m[tp.TypeName]; found {
			continue
		}

		m[tp.TypeName] = true

		if len(tp.Schema.EnumValues) > 0 {
			wrapper := ""
			if tp.Schema.GoType == "string" {
				wrapper = `"`
			}
			enums = append(enums, EnumDefinition{
				Schema:         tp.Schema,
				TypeName:       tp.TypeName,
				ValueWrapper:   wrapper,
				PrefixTypeName: globalState.options.Compatibility.AlwaysPrefixEnumValues,
			})
		}
	}

	// Now, go through all the enums, and figure out if we have conflicts with
	// any others.
	for i := range enums {
		// Look through all other enums not compared so far. Make sure we don't
		// compare against self.
		e1 := enums[i]
		for j := i + 1; j < len(enums); j++ {
			e2 := enums[j]

			for e1key := range e1.GetValues() {
				_, found := e2.GetValues()[e1key]
				if found {
					e1.PrefixTypeName = true
					e2.PrefixTypeName = true
					enums[i] = e1
					enums[j] = e2
					break
				}
			}
		}

		// now see if this enum conflicts with any global type names.
		for _, tp := range types {
			// Skip over enums, since we've handled those above.
			if len(tp.Schema.EnumValues) > 0 {
				continue
			}
			_, found := e1.Schema.EnumValues[tp.TypeName]
			if found {
				e1.PrefixTypeName = true
				enums[i] = e1
			}
		}

		// Another edge case is that an enum value can conflict with its own
		// type name.
		_, found := e1.GetValues()[e1.TypeName]
		if found {
			e1.PrefixTypeName = true
			enums[i] = e1
		}
	}

	// Now see if enums conflict with any non-enum typenames

	return GenerateTemplates([]string{"constants.tmpl"}, t, Constants{EnumDefinitions: enums})
}

// GenerateImports generates our import statements and package definition.
func GenerateImports(t *template.Template, externalImports []string, packageName string, versionOverride *string) (string, error) {
	// Read build version for incorporating into generated files
	// Unit tests have ok=false, so we'll just use "unknown" for the
	// version if we can't read this.

	modulePath := "unknown module path"
	moduleVersion := "unknown version"
	if bi, ok := debug.ReadBuildInfo(); ok {
		if bi.Main.Path != "" {
			modulePath = bi.Main.Path
		}
		if bi.Main.Version != "" {
			moduleVersion = bi.Main.Version
		}
		if versionOverride != nil {
			moduleVersion = *versionOverride
		}
	}

	context := struct {
		ExternalImports   []string
		PackageName       string
		ModuleName        string
		Version           string
		AdditionalImports []AdditionalImport
	}{
		ExternalImports:   externalImports,
		PackageName:       packageName,
		ModuleName:        modulePath,
		Version:           moduleVersion,
		AdditionalImports: globalState.options.AdditionalImports,
	}

	return GenerateTemplates([]string{"imports.tmpl"}, t, context)
}

// GenerateAdditionalPropertyBoilerplate generates all the glue code which provides
// the API for interacting with additional properties and JSON-ification
func GenerateAdditionalPropertyBoilerplate(t *template.Template, typeDefs []TypeDefinition) (string, error) {
	var filteredTypes []TypeDefinition

	m := map[string]bool{}

	for _, t := range typeDefs {
		if found := m[t.TypeName]; found {
			continue
		}

		m[t.TypeName] = true

		if t.Schema.HasAdditionalProperties {
			filteredTypes = append(filteredTypes, t)
		}
	}

	context := struct {
		Types []TypeDefinition
	}{
		Types: filteredTypes,
	}

	return GenerateTemplates([]string{"additional-properties.tmpl"}, t, context)
}

func GenerateUnionBoilerplate(t *template.Template, typeDefs []TypeDefinition) (string, error) {
	var filteredTypes []TypeDefinition
	for _, t := range typeDefs {
		if len(t.Schema.UnionElements) != 0 {
			filteredTypes = append(filteredTypes, t)
		}
	}

	if len(filteredTypes) == 0 {
		return "", nil
	}

	context := struct {
		Types []TypeDefinition
	}{
		Types: filteredTypes,
	}

	return GenerateTemplates([]string{"union.tmpl"}, t, context)
}

func GenerateUnionAndAdditionalProopertiesBoilerplate(t *template.Template, typeDefs []TypeDefinition) (string, error) {
	var filteredTypes []TypeDefinition
	for _, t := range typeDefs {
		if len(t.Schema.UnionElements) != 0 && t.Schema.HasAdditionalProperties {
			filteredTypes = append(filteredTypes, t)
		}
	}

	if len(filteredTypes) == 0 {
		return "", nil
	}
	context := struct {
		Types []TypeDefinition
	}{
		Types: filteredTypes,
	}

	return GenerateTemplates([]string{"union-and-additional-properties.tmpl"}, t, context)
}

// SanitizeCode runs sanitizers across the generated Go code to ensure the
// generated code will be able to compile.
func SanitizeCode(goCode string) string {
	// remove any byte-order-marks which break Go-Code
	// See: https://groups.google.com/forum/#!topic/golang-nuts/OToNIPdfkks
	return strings.ReplaceAll(goCode, "\uFEFF", "")
}

// GetUserTemplateText attempts to retrieve the template text from a passed in URL or file
// path when inputData is more than one line.
// This function will attempt to load a file first, and if it fails, will try to get the
// data from the remote endpoint.
// The timeout for remote download file is 30 seconds.
func GetUserTemplateText(inputData string) (template string, err error) {
	// if the input data is more than one line, assume its a template and return that data.
	if strings.Contains(inputData, "\n") {
		return inputData, nil
	}

	// load data from file
	data, err := os.ReadFile(inputData)
	// return data if found and loaded
	if err == nil {
		return string(data), nil
	}

	// check for non "not found" errors
	if !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to open file %s: %w", inputData, err)
	}

	// attempt to get data from url with timeout
	const downloadTimeout = 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), downloadTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, inputData, http.NoBody)
	if err != nil {
		return "", fmt.Errorf("failed to create request GET %s: %w", inputData, err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute GET request data from %s: %w", inputData, err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("got non %d status code on GET %s", resp.StatusCode, inputData)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body from GET %s: %w", inputData, err)
	}

	return string(data), nil
}

// LoadTemplates loads all of our template files into a text/template. The
// path of template is relative to the templates directory.
func LoadTemplates(src embed.FS, t *template.Template) error {
	return fs.WalkDir(src, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		buf, err := src.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", path, err)
		}

		templateName := strings.TrimPrefix(path, "templates/")
		tmpl := t.New(templateName)
		_, err = tmpl.Parse(string(buf))
		if err != nil {
			return fmt.Errorf("parsing template '%s': %w", path, err)
		}
		return nil
	})
}

func OperationSchemaImports(s *Schema) (map[string]goImport, error) {
	res := map[string]goImport{}

	for _, p := range s.Properties {
		imprts, err := GoSchemaImports(base.CreateSchemaProxy(p.Schema.OAPISchema))
		if err != nil {
			return nil, err
		}
		MergeImports(res, imprts)
	}

	if s.OAPISchema != nil {
		imprts, err := GoSchemaImports(base.CreateSchemaProxy(s.OAPISchema))
		if err != nil {
			return nil, err
		}
		MergeImports(res, imprts)
	}
	return res, nil
}

func OperationImports(ops []OperationDefinition) (map[string]goImport, error) {
	res := map[string]goImport{}
	for _, op := range ops {
		for _, pd := range [][]ParameterDefinition{op.PathParams, op.QueryParams} {
			for _, p := range pd {
				imprts, err := OperationSchemaImports(&p.Schema)
				if err != nil {
					return nil, err
				}
				MergeImports(res, imprts)
			}
		}

		for _, b := range op.Bodies {
			imprts, err := OperationSchemaImports(&b.Schema)
			if err != nil {
				return nil, err
			}
			MergeImports(res, imprts)
		}

		for _, b := range op.Responses {
			for _, c := range b.Contents {
				imprts, err := OperationSchemaImports(&c.Schema)
				if err != nil {
					return nil, err
				}
				MergeImports(res, imprts)
			}
		}

	}
	return res, nil
}

func GetTypeDefinitionsImports(swagger *libopenapi.DocumentModel[v3.Document], excludeSchemas []string) (map[string]goImport, error) {
	res := map[string]goImport{}
	if swagger.Model.Components == nil {
		return res, nil
	}

	schemaImports, err := GetSchemaImports(swagger.Model.Components.Schemas, excludeSchemas)
	if err != nil {
		return nil, err
	}

	reqBodiesImports, err := GetRequestBodiesImports(swagger.Model.Components.RequestBodies)
	if err != nil {
		return nil, err
	}

	responsesImports, err := GetResponsesImports(swagger.Model.Components.Responses)
	if err != nil {
		return nil, err
	}

	parametersImports, err := GetParametersImports(swagger.Model.Components.Parameters)
	if err != nil {
		return nil, err
	}

	for _, imprts := range []map[string]goImport{schemaImports, reqBodiesImports, responsesImports, parametersImports} {
		MergeImports(res, imprts)
	}
	return res, nil
}

func GoSchemaImports(schemas ...*base.SchemaProxy) (map[string]goImport, error) {
	res := map[string]goImport{}
	for _, sref := range schemas {
		if sref == nil || (sref.IsReference() && IsGoTypeReference(sref.GetReference())) {
			return nil, nil
		}
		if gi, err := ParseGoImportExtension(sref); err != nil {
			return nil, err
		} else {
			if gi != nil {
				res[gi.String()] = *gi
			}
		}
		schemaVal := sref.Schema()

		t := schemaVal.Type
		// handle cases where there isn't a `type` set
		if len(t) == 0 {
			t = append(t, "")
		}

		switch t[0] { // TODO multi-value with OpenAPI 3.1
		case "", "object":
			for _, v := range ToMap(schemaVal.Properties) {
				imprts, err := GoSchemaImports(v)
				if err != nil {
					return nil, err
				}
				MergeImports(res, imprts)
			}
		case "array":
			if schemaVal.Items != nil && schemaVal.Items.IsA() {
				imprts, err := GoSchemaImports(schemaVal.Items.A)
				if err != nil {
					return nil, err
				}
				MergeImports(res, imprts)
			}
		}
	}
	return res, nil
}

func GetSchemaImports(schemas *orderedmap.Map[string, *base.SchemaProxy], excludeSchemas []string) (map[string]goImport, error) {
	res := map[string]goImport{}
	excludeSchemasMap := make(map[string]bool)
	for _, schema := range excludeSchemas {
		excludeSchemasMap[schema] = true
	}
	for schemaName, schema := range ToMap(schemas) {
		if _, ok := excludeSchemasMap[schemaName]; ok {
			continue
		}

		imprts, err := GoSchemaImports(schema)
		if err != nil {
			return nil, err
		}
		MergeImports(res, imprts)
	}
	return res, nil
}

func GetRequestBodiesImports(bodies *orderedmap.Map[string, *v3.RequestBody]) (map[string]goImport, error) {
	res := map[string]goImport{}

	for _, r := range ToMap(bodies) {
		response := r

		for mediaType, body := range ToMap(response.Content) {
			if !util.IsMediaTypeJson(mediaType) {
				continue
			}

			imprts, err := GoSchemaImports(body.Schema)
			if err != nil {
				return nil, err
			}
			MergeImports(res, imprts)
		}
	}
	return res, nil
}

func GetResponsesImports(responses *orderedmap.Map[string, *v3.Response]) (map[string]goImport, error) {
	res := map[string]goImport{}
	for _, r := range ToMap(responses) {
		response := r

		for mediaType, body := range ToMap(response.Content) {
			if !util.IsMediaTypeJson(mediaType) {
				continue
			}

			imprts, err := GoSchemaImports(body.Schema)
			if err != nil {
				return nil, err
			}
			MergeImports(res, imprts)
		}
	}
	return res, nil
}

func GetParametersImports(params *orderedmap.Map[string, *v3.Parameter]) (map[string]goImport, error) {
	res := map[string]goImport{}

	for _, param := range ToMap(params) {
		if param == nil {
			continue
		}
		imprts, err := GoSchemaImports(param.Schema)
		if err != nil {
			return nil, err
		}
		MergeImports(res, imprts)
	}
	return res, nil
}

func SetGlobalStateSpec(spec *libopenapi.DocumentModel[v3.Document]) {
	globalState.spec = spec
}