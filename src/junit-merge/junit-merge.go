package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

type JUnitReport struct {
	XMLName   xml.Name
	XML       string       `xml:",innerxml"`
	Name      string       `xml:"name,attr"`
	Time      float64      `xml:"time,attr"`
	Tests     uint64       `xml:"tests,attr"`
	Failures  uint64       `xml:"failures,attr"`
	XMLBuffer bytes.Buffer `xml:"-"`
}

func main() {
	//outputFileName := flag.String("o", "merged.xml", "merged report filename")
	flag.Parse()
	files := flag.Args()
	//todo: walk directories recursively

	var mergedReport JUnitReport
	startedReading := false

	for _, fileName := range files {
		var report JUnitReport
		in, _ := ioutil.ReadFile(fileName)
		xml.Unmarshal(in, &report)

		if report.XMLName.Local == "testsuite" {
			panic(errors.New("Reports with a root <testsuite> are not supported"))
		}

		if startedReading && report.Name != mergedReport.Name {
			panic(errors.New("All reports must have the same <testsuites> name"))
		}

		startedReading = true
		mergedReport.XMLName = xml.Name{Local: "testsuites"}
		mergedReport.Name = report.Name
		mergedReport.Time += report.Time
		mergedReport.Tests += report.Tests
		mergedReport.Failures += report.Failures
		mergedReport.XMLBuffer.WriteString(report.XML)
	}

	mergedReport.XML = mergedReport.XMLBuffer.String()
	out, _ := xml.MarshalIndent(&mergedReport, "", "  ")
	fmt.Println(string(out))
}
