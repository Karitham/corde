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
	"int64",
	"uint",
	"uint64",
	"float32",
	"float64",
	"bool",
	"Snowflake",
}

var body string = `	v, _ := o[k].(%s)
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
