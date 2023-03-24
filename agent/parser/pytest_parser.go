package parser

import (
	"encoding/xml"
	"io/ioutil"
	"time"
)

type Testsuites struct {
	XMLName   xml.Name    `xml:"testsuites"`
	Testsuite []Testsuite `xml:"testsuite"`
}

type Testsuite struct {
	Name      string     `xml:"name,attr"`
	Errors    int        `xml:"errors,attr"`
	Failures  int        `xml:"failures,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Tests     int        `xml:"tests,attr"`
	Time      float64    `xml:"time,attr"`
	Timestamp time.Time  `xml:"timestamp,attr"`
	Hostname  string     `xml:"hostname,attr"`
	TestCases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Classname string  `xml:"classname,attr"`
	Name      string  `xml:"name,attr"`
	Time      float64 `xml:"time,attr"`
	Failure   *struct {
		Message string `xml:"message,attr"`
		Text    string `xml:",chardata"`
	} `xml:"failure"`
}

// type Failure struct {
// 	Message string `xml:"message,attr"`
// 	Text    string `xml:",chardata"`
// }

func Parse(filepath string) Testsuites {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var testsuites Testsuites
	if err := xml.Unmarshal(data, &testsuites); err != nil {
		panic(err)
	}

	return testsuites
}
