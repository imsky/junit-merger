package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type JUnitReport struct {
	Name string `xml:"name,attr"`
}

func main() {
	//outputFileName := flag.String("o", "merged.xml", "merged report filename")
	flag.Parse()
	files := flag.Args()
	//todo: walk directories recursively
	for _, fileName := range files {
		xmlFile, err := os.Open(fileName)
		if err != nil {
			//todo: use logging library
			fmt.Println("Error opening file:", err)
			return
		}

		var q JUnitReport
		data, _ := ioutil.ReadFile(fileName)
		xml.Unmarshal(data, &q)

		fmt.Println(q)
		xmlFile.Close()
	}
}
