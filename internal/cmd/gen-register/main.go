package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	genial "github.com/karitham/go-genial"
)

var types = []options{
	{DiscordType: "OPTION_STRING", Type: "string", Name: "string", CanAutocomplete: true},
	{DiscordType: "OPTION_INTEGER", Type: "int", Name: "int", CanAutocomplete: true},
	{DiscordType: "OPTION_NUMBER", Type: "float64", Name: "number", CanAutocomplete: true},
	{DiscordType: "OPTION_BOOLEAN", Type: "bool", Name: "bool"},
	{DiscordType: "OPTION_USER", Type: "snowflake.Snowflake", Name: "user"},
	{DiscordType: "OPTION_CHANNEL", Type: "snowflake.Snowflake", Name: "channel"},
	{DiscordType: "OPTION_ROLE", Type: "snowflake.Snowflake", Name: "role"},
	{DiscordType: "OPTION_MENTIONABLE", Type: "snowflake.Snowflake", Name: "mentionable"},
}

type options struct {
	DiscordType     string
	Type            string
	Name            string
	CanAutocomplete bool
}

var ConstructorBody string = `	o := &%s{
		Name:         name,
		Description:  description,
		Required:     required,
		Choices:      []components.Choice[any]{},
		ChannelTypes: []components.ChannelType{},
	}

	for _, choice := range choices {
		o.Choices = append(
			o.Choices,
			components.Choice[any]{Name: choice.Name, Value: choice.Value},
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
		Type:         components.%s,
	}
`

var file = flag.String("file", "../../../register-cmd.gen.go", "output file")

//go:generate go run .
func main() {
	flag.Parse()

	p := &genial.PackageB{}
	p.Name("corde").
		Imports("encoding/json", "github.com/Karitham/corde/snowflake", "github.com/Karitham/corde/components").
		License("GENERATED BY ./internal/cmd/gen-opt/ DO NOT EDIT.")

	for _, v := range types {
		typeName := strings.Title(fmt.Sprintf("%sOption", v.Name))

		typeOpt := &genial.StructB{}
		typeOpt.Name(typeName).
			Commentf("%s represents a %s option", typeName, v.Type).
			Field("Name", "string").
			Field("Description", "string").
			Field("Required", "bool").
			Field("Choices", "[]components.Choice[any]").
			Field("ChannelTypes", "[]components.ChannelType").
			Field("Autocomplete", "bool")

		constructorF := &genial.FuncB{}
		constructorF.Namef("New%s", typeName).
			Commentf("New%s returns a new %s", typeName, typeName).
			Parameter("name", "string").
			Parameter("description", "string").
			Parameter("required", "bool").
			Parameter("choices", fmt.Sprintf("...components.Choice[%s]", v.Type)).
			ReturnTypes("*"+typeName).
			Writef(ConstructorBody, typeName)

		channelTypesF := &genial.FuncB{}
		channelTypesF.Name("ChanTypes").
			Commentf("ChanTypes sets the options channel types").
			Receiver("o", "*"+typeName).
			ReturnTypes("*"+typeName).
			Parameter("typs", "...components.ChannelType").
			WriteString("\to.ChannelTypes = append(o.ChannelTypes, typs...)\n\treturn o\n")

		createOptionF := &genial.FuncB{}
		createOptionF.Name("createOption").
			Comment("createOption returns the CreateOption of the type").
			Receiver("o", "*"+typeName).
			ReturnTypes("CreateOption").
			Writef(CreateOptionBody, v.DiscordType)

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

	os.WriteFile(*file, p.Bytes(), 0o644)
}
