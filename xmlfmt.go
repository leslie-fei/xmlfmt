package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"strings"
)

type element struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	//Content []byte     `xml:",innerxml"`
	Childs []element `xml:",any,omitempty"`
}

func FormatXML(content []byte, prefix, indent string, replaces []string) ([]byte, error) {
	var el element
	err := xml.Unmarshal(content, &el)
	if err != nil {
		return nil, err
	}

	rs, err := xml.MarshalIndent(&el, prefix, indent)
	if err != nil {
		return nil, err
	}

	for _, rep := range replaces {
		rr := strings.SplitN(rep, "=", 2)
		if len(rr) != 2 {
			return nil, errors.New("replace must split `=` result length != 2")
		}
		rs = bytes.ReplaceAll(rs, []byte(rr[0]), []byte(rr[1]))
	}
	return rs, nil
}
