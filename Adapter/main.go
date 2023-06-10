package main

import (
	"fmt"
)

type fetchInterface interface {
	marshal() fetchInterface
}

type xml struct {
	xmlData string
}

func (d *xml) marshal() fetchInterface {
	return &xml{
		xmlData: d.xmlData,
	}
}

type json struct {
	jsonData string
}

func (d *json) marshal() fetchInterface {
	return &json{
		jsonData: d.jsonData,
	}
}

type adapter struct {
	json *json
	xml  *xml
}

func (a *adapter) convert() *xml {

	jsonData := a.json.jsonData

	convertJsonToXml := jsonData + " converted to xml" // some conversion logic

	return &xml{
		xmlData: convertJsonToXml,
	}
}

func main() {
	jsonData := &json{
		jsonData: "json data",
	}

	adapter := &adapter{
		json: jsonData,
	}

	convertedData := adapter.convert()

	fmt.Println(convertedData.marshal())
}
