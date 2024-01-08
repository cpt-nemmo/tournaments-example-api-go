package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person represents a <person> node in the XML
type Person struct {
	XMLName   xml.Name   `xml:"Tournaments"`
	DataItems []dataItem `xml:"DataItems"`
}

// Skill represents a <skill> node in the XML
type dataItem struct {
	XMLName    xml.Name `xml:"Tournament"`
	Id         string   `xml:"id,attr"`
	Level      string   `xml:"level,attr"`
	Name       string   `xml:"name,attr"`
	NameEn     string   `xml:"nameEn,attr"`
	SeasonPart string   `xml:"seasonPart,attr"`
	Season     string   `xml:"season,attr"`
	StartDate  string   `xml:"startDate,attr"`
	EndDate    string   `xml:"endDate,attr"`
	GameId     string   `xml:"endid,attr"`
}

var tournaments []byte

func main() {
	const (
		// A generic XML header suitable for use with the output of Marshal.
		// This is not automatically added to any output of this package,
		// it is provided as a convenience.
		Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	)

	players := Person{
		DataItems: []dataItem{
			{
				Id: "160", Level: "khl",
				Name:       "Открытый чемпионат России по хоккею - Чемпионат Континентальной хоккейной лиги",
				NameEn:     "Open Russian Ice Hockey Championship - Kontinental Hockey League Championship",
				SeasonPart: "regular",
				Season:     "2008-2009",
				StartDate:  "2008-09-02",
				EndDate:    "2009-02-26",
				GameId:     "true",
			},
			{
				Id: "165", Level: "khl",
				Name:       "Открытый чемпионат России по хоккею - Чемпионат Континентальной хоккейной лиги",
				NameEn:     "Open Russian Ice Hockey Championship - Kontinental Hockey League Championship",
				SeasonPart: "playoff",
				Season:     "2008-2009",
				StartDate:  "2009-03-01",
				EndDate:    "2009-04-12",
				GameId:     "true",
			},
			{
				Id: "167", Level: "khl",
				Name:       "Чемпионат Континентальной хоккейной лиги - Открытый чемпионат России по хоккею",
				NameEn:     "Kontinental Hockey League Championship - Open Russian Ice Hockey Championship",
				SeasonPart: "regular",
				Season:     "2009-2010",
				StartDate:  "2009-09-10",
				EndDate:    "2010-03-07",
				GameId:     "true",
			},
			{
				Id: "168", Level: "khl",
				Name:       "Чемпионат Континентальной хоккейной лиги - Открытый чемпионат России по хоккею",
				NameEn:     "Kontinental Hockey League Championship - Open Russian Ice Hockey Championship",
				SeasonPart: "playoff",
				Season:     "2009-2010",
				StartDate:  "2010-03-10",
				EndDate:    "2010-04-27",
				GameId:     "true",
			},
		},
	}

	out := new(bytes.Buffer)

	enc := xml.NewEncoder(out)
	enc.Indent("", "  ")
	if err := enc.Encode(players); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	tournaments = bytes.ReplaceAll(out.Bytes(), []byte("></Tournament>"), []byte("/>"))
	// os.Stdout.Write(b)
	// fmt.Printf("%s\n", xml.Header+string(b))
	// tournaments = string(b)
	r := mux.NewRouter()
	r.HandleFunc("/tournaments", getTournaments).Methods("GET")
	fmt.Println("Serve on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getTournaments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	// Write
	w.Write(tournaments)
}
