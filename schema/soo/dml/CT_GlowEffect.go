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
	"fmt"
	"strconv"

	"github.com/codrocker/gooxml"
)

type CT_GlowEffect struct {
	RadAttr   *int64
	ScrgbClr  *CT_ScRgbColor
	SrgbClr   *CT_SRgbColor
	HslClr    *CT_HslColor
	SysClr    *CT_SystemColor
	SchemeClr *CT_SchemeColor
	PrstClr   *CT_PresetColor
}

func NewCT_GlowEffect() *CT_GlowEffect {
	ret := &CT_GlowEffect{}
	return ret
}

func (m *CT_GlowEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rad"},
			Value: fmt.Sprintf("%v", *m.RadAttr)})
	}
	e.EncodeToken(start)
	if m.ScrgbClr != nil {
		sescrgbClr := xml.StartElement{Name: xml.Name{Local: "a:scrgbClr"}}
		e.EncodeElement(m.ScrgbClr, sescrgbClr)
	}
	if m.SrgbClr != nil {
		sesrgbClr := xml.StartElement{Name: xml.Name{Local: "a:srgbClr"}}
		e.EncodeElement(m.SrgbClr, sesrgbClr)
	}
	if m.HslClr != nil {
		sehslClr := xml.StartElement{Name: xml.Name{Local: "a:hslClr"}}
		e.EncodeElement(m.HslClr, sehslClr)
	}
	if m.SysClr != nil {
		sesysClr := xml.StartElement{Name: xml.Name{Local: "a:sysClr"}}
		e.EncodeElement(m.SysClr, sesysClr)
	}
	if m.SchemeClr != nil {
		seschemeClr := xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}}
		e.EncodeElement(m.SchemeClr, seschemeClr)
	}
	if m.PrstClr != nil {
		seprstClr := xml.StartElement{Name: xml.Name{Local: "a:prstClr"}}
		e.EncodeElement(m.PrstClr, seprstClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GlowEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rad" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RadAttr = &parsed
			continue
		}
	}
lCT_GlowEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scrgbClr"}:
				m.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(m.ScrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "srgbClr"}:
				m.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(m.SrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hslClr"}:
				m.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(m.HslClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sysClr"}:
				m.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(m.SysClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "schemeClr"}:
				m.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(m.SchemeClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstClr"}:
				m.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(m.PrstClr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_GlowEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GlowEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GlowEffect and its children
func (m *CT_GlowEffect) Validate() error {
	return m.ValidateWithPath("CT_GlowEffect")
}

// ValidateWithPath validates the CT_GlowEffect and its children, prefixing error messages with path
func (m *CT_GlowEffect) ValidateWithPath(path string) error {
	if m.RadAttr != nil {
		if *m.RadAttr < 0 {
			return fmt.Errorf("%s/m.RadAttr must be >= 0 (have %v)", path, *m.RadAttr)
		}
		if *m.RadAttr > 27273042316900 {
			return fmt.Errorf("%s/m.RadAttr must be <= 27273042316900 (have %v)", path, *m.RadAttr)
		}
	}
	if m.ScrgbClr != nil {
		if err := m.ScrgbClr.ValidateWithPath(path + "/ScrgbClr"); err != nil {
			return err
		}
	}
	if m.SrgbClr != nil {
		if err := m.SrgbClr.ValidateWithPath(path + "/SrgbClr"); err != nil {
			return err
		}
	}
	if m.HslClr != nil {
		if err := m.HslClr.ValidateWithPath(path + "/HslClr"); err != nil {
			return err
		}
	}
	if m.SysClr != nil {
		if err := m.SysClr.ValidateWithPath(path + "/SysClr"); err != nil {
			return err
		}
	}
	if m.SchemeClr != nil {
		if err := m.SchemeClr.ValidateWithPath(path + "/SchemeClr"); err != nil {
			return err
		}
	}
	if m.PrstClr != nil {
		if err := m.PrstClr.ValidateWithPath(path + "/PrstClr"); err != nil {
			return err
		}
	}
	return nil
}
