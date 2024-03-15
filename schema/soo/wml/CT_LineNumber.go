// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/codrocker/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_LineNumber struct {
	// Line Number Increments to Display
	CountByAttr *int64
	// Line Numbering Starting Value
	StartAttr *int64
	// Distance Between Text and Line Numbering
	DistanceAttr *sharedTypes.ST_TwipsMeasure
	// Line Numbering Restart Setting
	RestartAttr ST_LineNumberRestart
}

func NewCT_LineNumber() *CT_LineNumber {
	ret := &CT_LineNumber{}
	return ret
}

func (m *CT_LineNumber) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:countBy"},
			Value: fmt.Sprintf("%v", *m.CountByAttr)})
	}
	if m.StartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:start"},
			Value: fmt.Sprintf("%v", *m.StartAttr)})
	}
	if m.DistanceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:distance"},
			Value: fmt.Sprintf("%v", *m.DistanceAttr)})
	}
	if m.RestartAttr != ST_LineNumberRestartUnset {
		attr, err := m.RestartAttr.MarshalXMLAttr(xml.Name{Local: "w:restart"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "countBy" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CountByAttr = &parsed
			continue
		}
		if attr.Name.Local == "start" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.StartAttr = &parsed
			continue
		}
		if attr.Name.Local == "distance" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.DistanceAttr = &parsed
			continue
		}
		if attr.Name.Local == "restart" {
			m.RestartAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LineNumber: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LineNumber and its children
func (m *CT_LineNumber) Validate() error {
	return m.ValidateWithPath("CT_LineNumber")
}

// ValidateWithPath validates the CT_LineNumber and its children, prefixing error messages with path
func (m *CT_LineNumber) ValidateWithPath(path string) error {
	if m.DistanceAttr != nil {
		if err := m.DistanceAttr.ValidateWithPath(path + "/DistanceAttr"); err != nil {
			return err
		}
	}
	if err := m.RestartAttr.ValidateWithPath(path + "/RestartAttr"); err != nil {
		return err
	}
	return nil
}
