// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/codrocker/gooxml"
)

type EG_TextGeometry struct {
	CustGeom   *CT_CustomGeometry2D
	PrstTxWarp *CT_PresetTextShape
}

func NewEG_TextGeometry() *EG_TextGeometry {
	ret := &EG_TextGeometry{}
	return ret
}

func (m *EG_TextGeometry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:EG_TextGeometry"
	if m.CustGeom != nil {
		secustGeom := xml.StartElement{Name: xml.Name{Local: "a:custGeom"}}
		e.EncodeElement(m.CustGeom, secustGeom)
	}
	if m.PrstTxWarp != nil {
		seprstTxWarp := xml.StartElement{Name: xml.Name{Local: "a:prstTxWarp"}}
		e.EncodeElement(m.PrstTxWarp, seprstTxWarp)
	}
	return nil
}

func (m *EG_TextGeometry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextGeometry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custGeom"}:
				m.CustGeom = NewCT_CustomGeometry2D()
				if err := d.DecodeElement(m.CustGeom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstTxWarp"}:
				m.PrstTxWarp = NewCT_PresetTextShape()
				if err := d.DecodeElement(m.PrstTxWarp, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_TextGeometry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextGeometry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextGeometry and its children
func (m *EG_TextGeometry) Validate() error {
	return m.ValidateWithPath("EG_TextGeometry")
}

// ValidateWithPath validates the EG_TextGeometry and its children, prefixing error messages with path
func (m *EG_TextGeometry) ValidateWithPath(path string) error {
	if m.CustGeom != nil {
		if err := m.CustGeom.ValidateWithPath(path + "/CustGeom"); err != nil {
			return err
		}
	}
	if m.PrstTxWarp != nil {
		if err := m.PrstTxWarp.ValidateWithPath(path + "/PrstTxWarp"); err != nil {
			return err
		}
	}
	return nil
}
