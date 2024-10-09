package common

import (
	"encoding/xml"
)

type StyledText struct {
	Text  string
	Style string
}

type OcrPage struct {
	XMLName xml.Name  `xml:"div"`
	Areas   []OcrArea `xml:"div"`
}

type OcrArea struct {
	XMLName xml.Name `xml:"div"`
	Blocks  []OcrPar `xml:"p"`
	BBox    string   `xml:"bbox,attr"`
}

type OcrPar struct {
	XMLName xml.Name  `xml:"p"`
	Lines   []OcrLine `xml:"span"`
}

type OcrLine struct {
	XMLName xml.Name  `xml:"span"`
	Words   []OcrWord `xml:"span"`
}

type OcrWord struct {
	XMLName xml.Name `xml:"span"`
	Text    string   `xml:",chardata"`
}
