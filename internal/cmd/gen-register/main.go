package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	genial "github.com/karitham/go-genial"
)

var types = map[string]opt{
	"OPTION_STRING":      {Type: "string", Name: "string", CanAutocomplete: true},
	"OPTION_INTEGER":     {Type: "int", Name: "int", CanAutocomplete: true},
	"OPTION_NUMBER":      {Type: "float64", Name: "number", CanAutocomplete: true},
	"OPTION_BOOLEAN":     {Type: "bool", Name: "bool"},
	"OPTION_USER":        {Type: "Snowflake", Name: "user"},
	"OPTION_CHANNEL":     {Type: "Snowflake", Name: "channel"},
	"OPTION_ROLE":        {Type: "Snowflake", Name: "role"},
	"OPTION_MENTIONABLE": {Type: "Snowflake", Name: "mentionable"},
}

type opt struct {
	Type            string
	Name            string
	CanAutocomplete bool
}

var ConstructorBody string = `	o := &%s{
		Name:        name,
		Description: description,
		Required:    required,
		Choices:     []Choice[any]{},
		ChannelTypes: []ChannelType{},
	}

	for _, choice := range choices {
		o.Choices = append(
			o.Choices,
			Choice[any]{Name: choice.Name, Value: choice.Value},
		)
	}

	return o
`

var CreateOptionBody string = `	return CreateOption{
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
		Choices:     o.Choices,
		ChannelTypes: o.ChannelTypes,
		Autocomplete: o.Autocomplete,
		Type:        %s,
	}
`

var file = flag.String("file", "../../../register-cmd.gen.go", "output file")

//go:generate go run .
func main() {
	flag.Parse()

	p := &genial.PackageB{}
	p.Name("corde").
		Imports("encoding/json").
		License(" GENERATED BY ./internal/cmd/gen-opt/ DO NOT EDIT.")

	for k, v := range types {
		typeName := strings.Title(fmt.Sprintf("%sOption", v.Name))

		typeOpt := &genial.StructB{}
		typeOpt.Name(typeName).
			Commentf("%s represents a %s option", typeName, v.Type).
			Field("Name", "string").
			Field("Description", "string").
			Field("Required", "bool").
			Field("Choices", "[]Choice[any]").
			Field("ChannelTypes", "[]ChannelType").
			Field("Autocomplete", "bool")

		constructorF := &genial.FuncB{}
		constructorF.Namef("New%s", typeName).
			Commentf("New%s returns a new %s", typeName, typeName).
			Parameter("name", "string").
			Parameter("description", "string").
			Parameter("required", "bool").
			Parameter("choices", fmt.Sprintf("...Choice[%s]", v.Type)).
			ReturnTypes("*"+typeName).
			Writef(ConstructorBody, typeName)

		channelTypesF := &genial.FuncB{}
		channelTypesF.Name("ChanTypes").
			Commentf("ChanTypes sets the options channel types").
			Receiver("o", "*"+typeName).
			ReturnTypes("*"+typeName).
			Parameter("typs", "...ChannelType").
			WriteString("\to.ChannelTypes = append(o.ChannelTypes, typs...)\n\treturn o\n")

		createOptionF := &genial.FuncB{}
		createOptionF.Name("createOption").
			Comment("createOption returns the CreateOption of the type").
			Receiver("o", "*"+typeName).
			ReturnTypes("CreateOption").
			Writef(CreateOptionBody, k)

		marshalF := &genial.FuncB{}
		marshalF.Name("MarshalJSON").
			Comment("MarshalJSON returns the JSON representation of the option").
			Receiver("o", "*"+typeName).
			ReturnTypes("[]byte", "error").
			WriteString("\treturn json.Marshal(o.createOption())\n")

		p.Declarations(typeOpt, constructorF, createOptionF, channelTypesF, marshalF)

		if v.CanAutocomplete {
			autocompleteF := &genial.FuncB{}
			autocompleteF.Name("CanAutocomplete").
				Commentf("CanAutocomplete sets the option as autocomplete-able").
				Receiver("o", "*"+typeName).
				ReturnTypes("*" + typeName).
				WriteString("\to.Autocomplete = true\n\treturn o\n")

			p.Declarations(autocompleteF)
		}

	}

	os.WriteFile(*file, []byte(p.String()), 0o644)
}
