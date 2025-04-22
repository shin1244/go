package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/beevik/etree"
	"github.com/xuri/excelize/v2"
)

type Person struct {
	ID         string `json:"ID"`
	Name       string `json:"Name"`
	Age        string `json:"Age"`
	Department string `json:"Department"`
}

func main() {
	xlsxToJson("xlsx")
	xmlToJson("xml")
}

func xlsxToJson(extension string) {
	f, err := excelize.OpenFile("sample_data." + extension)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	people := []Person{}

	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		newPerson := Person{
			ID:         row[0],
			Name:       row[1],
			Age:        row[2],
			Department: row[3],
		}
		people = append(people, newPerson)
	}
	makeJsonFile(people, extension)
}

func xmlToJson(extension string) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile("sample_data." + extension)
	if err != nil {
		fmt.Println(err)
		return
	}
	root := doc.SelectElement("employees")
	people := []Person{}

	for _, emp := range root.SelectElements("employee") {

		newPerson := Person{
			ID:         tagText(emp, "ID"),
			Name:       tagText(emp, "Name"),
			Age:        tagText(emp, "Age"),
			Department: tagText(emp, "Department"),
		}
		people = append(people, newPerson)
	}
	makeJsonFile(people, extension)
}

func tagText(elem *etree.Element, str string) string {
	return elem.SelectElement(str).Text()
}

func makeJsonFile(data []Person, extension string) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFile, err := os.Create(extension + "_to_json.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}
