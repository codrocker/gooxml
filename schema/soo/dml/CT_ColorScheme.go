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

	"github.com/carmel/gooxml"
)

type CT_ColorScheme struct {
	NameAttr string
	Dk1      *CT_Color
	Lt1      *CT_Color
	Dk2      *CT_Color
	Lt2      *CT_Color
	Accent1  *CT_Color
	Accent2  *CT_Color
	Accent3  *CT_Color
	Accent4  *CT_Color
	Accent5  *CT_Color
	Accent6  *CT_Color
	Hlink    *CT_Color
	FolHlink *CT_Color
	ExtLst   *CT_OfficeArtExtensionList
}

func NewCT_ColorScheme() *CT_ColorScheme {
	ret := &CT_ColorScheme{}
	ret.Dk1 = NewCT_Color()
	ret.Lt1 = NewCT_Color()
	ret.Dk2 = NewCT_Color()
	ret.Lt2 = NewCT_Color()
	ret.Accent1 = NewCT_Color()
	ret.Accent2 = NewCT_Color()
	ret.Accent3 = NewCT_Color()
	ret.Accent4 = NewCT_Color()
	ret.Accent5 = NewCT_Color()
	ret.Accent6 = NewCT_Color()
	ret.Hlink = NewCT_Color()
	ret.FolHlink = NewCT_Color()
	return ret
}

func (m *CT_ColorScheme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	sedk1 := xml.StartElement{Name: xml.Name{Local: "a:dk1"}}
	e.EncodeElement(m.Dk1, sedk1)
	selt1 := xml.StartElement{Name: xml.Name{Local: "a:lt1"}}
	e.EncodeElement(m.Lt1, selt1)
	sedk2 := xml.StartElement{Name: xml.Name{Local: "a:dk2"}}
	e.EncodeElement(m.Dk2, sedk2)
	selt2 := xml.StartElement{Name: xml.Name{Local: "a:lt2"}}
	e.EncodeElement(m.Lt2, selt2)
	seaccent1 := xml.StartElement{Name: xml.Name{Local: "a:accent1"}}
	e.EncodeElement(m.Accent1, seaccent1)
	seaccent2 := xml.StartElement{Name: xml.Name{Local: "a:accent2"}}
	e.EncodeElement(m.Accent2, seaccent2)
	seaccent3 := xml.StartElement{Name: xml.Name{Local: "a:accent3"}}
	e.EncodeElement(m.Accent3, seaccent3)
	seaccent4 := xml.StartElement{Name: xml.Name{Local: "a:accent4"}}
	e.EncodeElement(m.Accent4, seaccent4)
	seaccent5 := xml.StartElement{Name: xml.Name{Local: "a:accent5"}}
	e.EncodeElement(m.Accent5, seaccent5)
	seaccent6 := xml.StartElement{Name: xml.Name{Local: "a:accent6"}}
	e.EncodeElement(m.Accent6, seaccent6)
	sehlink := xml.StartElement{Name: xml.Name{Local: "a:hlink"}}
	e.EncodeElement(m.Hlink, sehlink)
	sefolHlink := xml.StartElement{Name: xml.Name{Local: "a:folHlink"}}
	e.EncodeElement(m.FolHlink, sefolHlink)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorScheme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Dk1 = NewCT_Color()
	m.Lt1 = NewCT_Color()
	m.Dk2 = NewCT_Color()
	m.Lt2 = NewCT_Color()
	m.Accent1 = NewCT_Color()
	m.Accent2 = NewCT_Color()
	m.Accent3 = NewCT_Color()
	m.Accent4 = NewCT_Color()
	m.Accent5 = NewCT_Color()
	m.Accent6 = NewCT_Color()
	m.Hlink = NewCT_Color()
	m.FolHlink = NewCT_Color()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
lCT_ColorScheme:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "dk1"}:
				if err := d.DecodeElement(m.Dk1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lt1"}:
				if err := d.DecodeElement(m.Lt1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "dk2"}:
				if err := d.DecodeElement(m.Dk2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lt2"}:
				if err := d.DecodeElement(m.Lt2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent1"}:
				if err := d.DecodeElement(m.Accent1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent2"}:
				if err := d.DecodeElement(m.Accent2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent3"}:
				if err := d.DecodeElement(m.Accent3, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent4"}:
				if err := d.DecodeElement(m.Accent4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent5"}:
				if err := d.DecodeElement(m.Accent5, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "accent6"}:
				if err := d.DecodeElement(m.Accent6, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hlink"}:
				if err := d.DecodeElement(m.Hlink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "folHlink"}:
				if err := d.DecodeElement(m.FolHlink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ColorScheme %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorScheme
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorScheme and its children
func (m *CT_ColorScheme) Validate() error {
	return m.ValidateWithPath("CT_ColorScheme")
}

// ValidateWithPath validates the CT_ColorScheme and its children, prefixing error messages with path
func (m *CT_ColorScheme) ValidateWithPath(path string) error {
	if err := m.Dk1.ValidateWithPath(path + "/Dk1"); err != nil {
		return err
	}
	if err := m.Lt1.ValidateWithPath(path + "/Lt1"); err != nil {
		return err
	}
	if err := m.Dk2.ValidateWithPath(path + "/Dk2"); err != nil {
		return err
	}
	if err := m.Lt2.ValidateWithPath(path + "/Lt2"); err != nil {
		return err
	}
	if err := m.Accent1.ValidateWithPath(path + "/Accent1"); err != nil {
		return err
	}
	if err := m.Accent2.ValidateWithPath(path + "/Accent2"); err != nil {
		return err
	}
	if err := m.Accent3.ValidateWithPath(path + "/Accent3"); err != nil {
		return err
	}
	if err := m.Accent4.ValidateWithPath(path + "/Accent4"); err != nil {
		return err
	}
	if err := m.Accent5.ValidateWithPath(path + "/Accent5"); err != nil {
		return err
	}
	if err := m.Accent6.ValidateWithPath(path + "/Accent6"); err != nil {
		return err
	}
	if err := m.Hlink.ValidateWithPath(path + "/Hlink"); err != nil {
		return err
	}
	if err := m.FolHlink.ValidateWithPath(path + "/FolHlink"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
