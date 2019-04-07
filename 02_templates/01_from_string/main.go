package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	m := map[string]interface{}{}
	m["Material"] = "Wood"
	m["Count"] = 10

	var tpl bytes.Buffer
	// sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tpl, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("result: ", tpl.String())
}
