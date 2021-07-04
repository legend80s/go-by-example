package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

var log = fmt.Println

func main() {
	coffee := &Plant{Id: 1, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")

	log(string(out))
	log()
	log(xml.Header + string(out))
	log(strings.Repeat("-", 10))

	var p Plant

	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}

	log(p)

	// Nesting
	tomato := &Plant{Id: 1, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")

	log(string(out))
}
