package main

import (
	"flag"
	"os"
	"strings"

	genial "github.com/karitham/go-genial"
)

var types = []string{
	"string",
	"int",
	"int32",
	"int64",
	"uint",
	"uint32",
	"uint64",
	"float32",
	"float64",
	"bool",
	"Snowflake",
	"any",
}

var body string = `	var v %s
	_ = o[k].UnmarshalTo(&v)
	return v
`

var file = flag.String("file", "../../../interaction-opt.gen.go", "output file")

//go:generate go run .
func main() {
	flag.Parse()

	p := &genial.PackageB{}
	p.Name("corde").License("GENERATED BY ./internal/cmd/gen-opt/ DO NOT EDIT.")

	for _, t := range types {
		name := strings.Title(t)

		f := &genial.FuncB{}
		f.Receiver("o", "OptionsInteractions").
			Commentf("%s returns the option with key k of type %s", name, t).
			Name(name).
			Parameter("k", "string").
			ReturnTypes(t).
			Writef(body, t)

		p.Declarations(f)
	}

	os.WriteFile(*file, p.Bytes(), 0o644)
}
