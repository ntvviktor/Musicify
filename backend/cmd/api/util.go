package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type MPD struct {
	XMLName                   xml.Name `xml:"MPD"`
	Xmlns                     string   `xml:"xmlns,attr"`
	Profiles                  string   `xml:"profiles,attr"`
	MinBufferTime             string   `xml:"minBufferTime,attr"`
	MediaPresentationDuration string   `xml:"mediaPresentationDuration,attr"`
	Type                      string   `xml:"type,attr"`
	Period                    struct {
		AdaptationSet struct {
			MimeType         string `xml:"mimeType,attr"`
			StartWithSAP     string `xml:"startWithSAP,attr"`
			SegmentAlignment string `xml:"segmentAlignment,attr"`
			Representation   struct {
				ID                        string `xml:"id,attr"`
				Codecs                    string `xml:"codecs,attr"`
				Bandwidth                 string `xml:"bandwidth,attr"`
				AudioSamplingRate         string `xml:"audioSamplingRate,attr"`
				AudioChannelConfiguration struct {
					SchemeIdUri string `xml:"schemeIdUri,attr"`
					Value       string `xml:"value,attr"`
				} `xml:"AudioChannelConfiguration"`
				SegmentList SegmentList `xml:"SegmentList"`
			} `xml:"Representation"`
		} `xml:"AdaptationSet"`
	} `xml:"Period"`
}

type SegmentList struct {
	Timescale      string `xml:"timescale,attr"`
	Duration       string `xml:"duration,attr"`
	Initialization struct {
		SourceURL string `xml:"sourceURL,attr"`
		Range     string `xml:"range,attr"`
	} `xml:"Initialization"`
	SegmentURL []struct {
		MediaRange string `xml:"mediaRange,attr"`
	} `xml:"SegmentURL"`
}

func ParseInXML(dirname string) *SegmentList {
	xmlFile, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	var Metadata MPD
	err = xml.NewDecoder(xmlFile).Decode(&Metadata)
	if err != nil {
		log.Fatal(err)
	}
	SegmentList := Metadata.Period.AdaptationSet.Representation.SegmentList

	return &SegmentList
}
