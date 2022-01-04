package main

import (
	"flag"
	"fmt"
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
	p.Name("corde").Comment("retrieve options from interaction data")

	for _, t := range types {
		f := &genial.FuncB{}

		name := strings.Title(t)

		f.Receiver(genial.Parameter{Name: "o", Type: "OptionsInteractions"}).
			Comment(fmt.Sprintf("%s returns the option with key k of type %s", name, t)).
			Name(name).
			Parameters(genial.Parameter{
				Name: "k",
				Type: "string",
			}).
			ReturnTypes(genial.Parameter{
				Type: t,
			})

		f.Write([]byte(fmt.Sprintf(body, t)))

		p.Declarations(f)
	}

	os.WriteFile(*file, []byte(p.String()), 0o644)
}
