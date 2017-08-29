package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	type User struct {
		Name string
		Age  int
	}
	t := User{"zzp", 18}
	content, e := xml.Marshal(t)
	fmt.Println(string(content), e)

	m := make(map[string]interface{})
	xml.Unmarshal(content, &m)
	fmt.Println(m)

	decoder := xml.NewDecoder(bytes.NewReader(content))
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		fmt.Println("print:", t)
		switch t.(type) {
		case xml.StartElement:
			e := t.(xml.StartElement)
			fmt.Println(e.Name.Local)
		case xml.EndElement:
		case xml.CharData:
			e := t.(xml.CharData)
			fmt.Println(string(e))
		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		}
	}
}
