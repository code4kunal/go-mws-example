package utils

import (
	"encoding/xml"
	"time"
)
type Date time.Time

type MarshalerAttr interface {
	MarshalXMLAttr(name xml.Name) (xml.Attr, error)
}


func (d Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	dateString := time.Time(d).Format("2006-01-02")
	attr := xml.Attr {
		name,
		dateString,
	}

	return attr, nil
}