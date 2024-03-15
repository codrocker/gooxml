// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/codrocker/gooxml"
)

type CT_Worksheet struct {
	// Worksheet Properties
	SheetPr *CT_SheetPr
	// Worksheet Dimensions
	Dimension *CT_SheetDimension
	// Sheet Views
	SheetViews *CT_SheetViews
	// Sheet Format Properties
	SheetFormatPr *CT_SheetFormatPr
	// Column Information
	Cols []*CT_Cols
	// Sheet Data
	SheetData *CT_SheetData
	// Sheet Calculation Properties
	SheetCalcPr *CT_SheetCalcPr
	// Sheet Protection
	SheetProtection *CT_SheetProtection
	// Protected Ranges
	ProtectedRanges *CT_ProtectedRanges
	// Scenarios
	Scenarios *CT_Scenarios
	// AutoFilter
	AutoFilter *CT_AutoFilter
	// Sort State
	SortState *CT_SortState
	// Data Consolidate
	DataConsolidate *CT_DataConsolidate
	// Custom Sheet Views
	CustomSheetViews *CT_CustomSheetViews
	// Merge Cells
	MergeCells *CT_MergeCells
	// Phonetic Properties
	PhoneticPr *CT_PhoneticPr
	// Conditional Formatting
	ConditionalFormatting []*CT_ConditionalFormatting
	// Data Validations
	DataValidations *CT_DataValidations
	// Hyperlinks
	Hyperlinks *CT_Hyperlinks
	// Print Options
	PrintOptions *CT_PrintOptions
	// Page Margins
	PageMargins *CT_PageMargins
	// Page Setup Settings
	PageSetup *CT_PageSetup
	// Header and Footer Settings
	HeaderFooter *CT_HeaderFooter
	// Horizontal Page Breaks
	RowBreaks *CT_PageBreak
	// Vertical Page Breaks
	ColBreaks *CT_PageBreak
	// Custom Properties
	CustomProperties *CT_CustomProperties
	// Cell Watch Items
	CellWatches *CT_CellWatches
	// Ignored Errors
	IgnoredErrors *CT_IgnoredErrors
	// Smart Tags
	SmartTags *CT_SmartTags
	// Drawing
	Drawing *CT_Drawing
	// Legacy Drawing
	LegacyDrawing *CT_LegacyDrawing
	// Legacy Drawing Header Footer
	LegacyDrawingHF *CT_LegacyDrawing
	DrawingHF       *CT_DrawingHF
	// Background Image
	Picture    *CT_SheetBackgroundPicture
	OleObjects *CT_OleObjects
	// Embedded Controls
	Controls *CT_Controls
	// Web Publishing Items
	WebPublishItems *CT_WebPublishItems
	// Table Parts
	TableParts *CT_TableParts
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Worksheet() *CT_Worksheet {
	ret := &CT_Worksheet{}
	ret.SheetData = NewCT_SheetData()
	return ret
}

func (m *CT_Worksheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SheetPr != nil {
		sesheetPr := xml.StartElement{Name: xml.Name{Local: "ma:sheetPr"}}
		e.EncodeElement(m.SheetPr, sesheetPr)
	}
	if m.Dimension != nil {
		sedimension := xml.StartElement{Name: xml.Name{Local: "ma:dimension"}}
		e.EncodeElement(m.Dimension, sedimension)
	}
	if m.SheetViews != nil {
		sesheetViews := xml.StartElement{Name: xml.Name{Local: "ma:sheetViews"}}
		e.EncodeElement(m.SheetViews, sesheetViews)
	}
	if m.SheetFormatPr != nil {
		sesheetFormatPr := xml.StartElement{Name: xml.Name{Local: "ma:sheetFormatPr"}}
		e.EncodeElement(m.SheetFormatPr, sesheetFormatPr)
	}
	if m.Cols != nil {
		secols := xml.StartElement{Name: xml.Name{Local: "ma:cols"}}
		for _, c := range m.Cols {
			e.EncodeElement(c, secols)
		}
	}
	sesheetData := xml.StartElement{Name: xml.Name{Local: "ma:sheetData"}}
	e.EncodeElement(m.SheetData, sesheetData)
	if m.SheetCalcPr != nil {
		sesheetCalcPr := xml.StartElement{Name: xml.Name{Local: "ma:sheetCalcPr"}}
		e.EncodeElement(m.SheetCalcPr, sesheetCalcPr)
	}
	if m.SheetProtection != nil {
		sesheetProtection := xml.StartElement{Name: xml.Name{Local: "ma:sheetProtection"}}
		e.EncodeElement(m.SheetProtection, sesheetProtection)
	}
	if m.ProtectedRanges != nil {
		seprotectedRanges := xml.StartElement{Name: xml.Name{Local: "ma:protectedRanges"}}
		e.EncodeElement(m.ProtectedRanges, seprotectedRanges)
	}
	if m.Scenarios != nil {
		sescenarios := xml.StartElement{Name: xml.Name{Local: "ma:scenarios"}}
		e.EncodeElement(m.Scenarios, sescenarios)
	}
	if m.AutoFilter != nil {
		seautoFilter := xml.StartElement{Name: xml.Name{Local: "ma:autoFilter"}}
		e.EncodeElement(m.AutoFilter, seautoFilter)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "ma:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	if m.DataConsolidate != nil {
		sedataConsolidate := xml.StartElement{Name: xml.Name{Local: "ma:dataConsolidate"}}
		e.EncodeElement(m.DataConsolidate, sedataConsolidate)
	}
	if m.CustomSheetViews != nil {
		secustomSheetViews := xml.StartElement{Name: xml.Name{Local: "ma:customSheetViews"}}
		e.EncodeElement(m.CustomSheetViews, secustomSheetViews)
	}
	if m.MergeCells != nil {
		semergeCells := xml.StartElement{Name: xml.Name{Local: "ma:mergeCells"}}
		e.EncodeElement(m.MergeCells, semergeCells)
	}
	if m.PhoneticPr != nil {
		sephoneticPr := xml.StartElement{Name: xml.Name{Local: "ma:phoneticPr"}}
		e.EncodeElement(m.PhoneticPr, sephoneticPr)
	}
	if m.ConditionalFormatting != nil {
		seconditionalFormatting := xml.StartElement{Name: xml.Name{Local: "ma:conditionalFormatting"}}
		for _, c := range m.ConditionalFormatting {
			e.EncodeElement(c, seconditionalFormatting)
		}
	}
	if m.DataValidations != nil {
		sedataValidations := xml.StartElement{Name: xml.Name{Local: "ma:dataValidations"}}
		e.EncodeElement(m.DataValidations, sedataValidations)
	}
	if m.Hyperlinks != nil {
		sehyperlinks := xml.StartElement{Name: xml.Name{Local: "ma:hyperlinks"}}
		e.EncodeElement(m.Hyperlinks, sehyperlinks)
	}
	if m.PrintOptions != nil {
		seprintOptions := xml.StartElement{Name: xml.Name{Local: "ma:printOptions"}}
		e.EncodeElement(m.PrintOptions, seprintOptions)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "ma:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "ma:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "ma:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	if m.RowBreaks != nil {
		serowBreaks := xml.StartElement{Name: xml.Name{Local: "ma:rowBreaks"}}
		e.EncodeElement(m.RowBreaks, serowBreaks)
	}
	if m.ColBreaks != nil {
		secolBreaks := xml.StartElement{Name: xml.Name{Local: "ma:colBreaks"}}
		e.EncodeElement(m.ColBreaks, secolBreaks)
	}
	if m.CustomProperties != nil {
		secustomProperties := xml.StartElement{Name: xml.Name{Local: "ma:customProperties"}}
		e.EncodeElement(m.CustomProperties, secustomProperties)
	}
	if m.CellWatches != nil {
		secellWatches := xml.StartElement{Name: xml.Name{Local: "ma:cellWatches"}}
		e.EncodeElement(m.CellWatches, secellWatches)
	}
	if m.IgnoredErrors != nil {
		seignoredErrors := xml.StartElement{Name: xml.Name{Local: "ma:ignoredErrors"}}
		e.EncodeElement(m.IgnoredErrors, seignoredErrors)
	}
	if m.SmartTags != nil {
		sesmartTags := xml.StartElement{Name: xml.Name{Local: "ma:smartTags"}}
		e.EncodeElement(m.SmartTags, sesmartTags)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "ma:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	if m.LegacyDrawing != nil {
		selegacyDrawing := xml.StartElement{Name: xml.Name{Local: "ma:legacyDrawing"}}
		e.EncodeElement(m.LegacyDrawing, selegacyDrawing)
	}
	if m.LegacyDrawingHF != nil {
		selegacyDrawingHF := xml.StartElement{Name: xml.Name{Local: "ma:legacyDrawingHF"}}
		e.EncodeElement(m.LegacyDrawingHF, selegacyDrawingHF)
	}
	if m.DrawingHF != nil {
		sedrawingHF := xml.StartElement{Name: xml.Name{Local: "ma:drawingHF"}}
		e.EncodeElement(m.DrawingHF, sedrawingHF)
	}
	if m.Picture != nil {
		sepicture := xml.StartElement{Name: xml.Name{Local: "ma:picture"}}
		e.EncodeElement(m.Picture, sepicture)
	}
	if m.OleObjects != nil {
		seoleObjects := xml.StartElement{Name: xml.Name{Local: "ma:oleObjects"}}
		e.EncodeElement(m.OleObjects, seoleObjects)
	}
	if m.Controls != nil {
		secontrols := xml.StartElement{Name: xml.Name{Local: "ma:controls"}}
		e.EncodeElement(m.Controls, secontrols)
	}
	if m.WebPublishItems != nil {
		sewebPublishItems := xml.StartElement{Name: xml.Name{Local: "ma:webPublishItems"}}
		e.EncodeElement(m.WebPublishItems, sewebPublishItems)
	}
	if m.TableParts != nil {
		setableParts := xml.StartElement{Name: xml.Name{Local: "ma:tableParts"}}
		e.EncodeElement(m.TableParts, setableParts)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Worksheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SheetData = NewCT_SheetData()
lCT_Worksheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetPr"}:
				m.SheetPr = NewCT_SheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dimension"}:
				m.Dimension = NewCT_SheetDimension()
				if err := d.DecodeElement(m.Dimension, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetViews"}:
				m.SheetViews = NewCT_SheetViews()
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetFormatPr"}:
				m.SheetFormatPr = NewCT_SheetFormatPr()
				if err := d.DecodeElement(m.SheetFormatPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cols"}:
				tmp := NewCT_Cols()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cols = append(m.Cols, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetData"}:
				if err := d.DecodeElement(m.SheetData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetCalcPr"}:
				m.SheetCalcPr = NewCT_SheetCalcPr()
				if err := d.DecodeElement(m.SheetCalcPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetProtection"}:
				m.SheetProtection = NewCT_SheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "protectedRanges"}:
				m.ProtectedRanges = NewCT_ProtectedRanges()
				if err := d.DecodeElement(m.ProtectedRanges, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "scenarios"}:
				m.Scenarios = NewCT_Scenarios()
				if err := d.DecodeElement(m.Scenarios, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "autoFilter"}:
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sortState"}:
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dataConsolidate"}:
				m.DataConsolidate = NewCT_DataConsolidate()
				if err := d.DecodeElement(m.DataConsolidate, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customSheetViews"}:
				m.CustomSheetViews = NewCT_CustomSheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "mergeCells"}:
				m.MergeCells = NewCT_MergeCells()
				if err := d.DecodeElement(m.MergeCells, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "phoneticPr"}:
				m.PhoneticPr = NewCT_PhoneticPr()
				if err := d.DecodeElement(m.PhoneticPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "conditionalFormatting"}:
				tmp := NewCT_ConditionalFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConditionalFormatting = append(m.ConditionalFormatting, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dataValidations"}:
				m.DataValidations = NewCT_DataValidations()
				if err := d.DecodeElement(m.DataValidations, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "hyperlinks"}:
				m.Hyperlinks = NewCT_Hyperlinks()
				if err := d.DecodeElement(m.Hyperlinks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "printOptions"}:
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageMargins"}:
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageSetup"}:
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "headerFooter"}:
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rowBreaks"}:
				m.RowBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.RowBreaks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "colBreaks"}:
				m.ColBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.ColBreaks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customProperties"}:
				m.CustomProperties = NewCT_CustomProperties()
				if err := d.DecodeElement(m.CustomProperties, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellWatches"}:
				m.CellWatches = NewCT_CellWatches()
				if err := d.DecodeElement(m.CellWatches, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ignoredErrors"}:
				m.IgnoredErrors = NewCT_IgnoredErrors()
				if err := d.DecodeElement(m.IgnoredErrors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTags"}:
				m.SmartTags = NewCT_SmartTags()
				if err := d.DecodeElement(m.SmartTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawing"}:
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawingHF"}:
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawingHF"}:
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "picture"}:
				m.Picture = NewCT_SheetBackgroundPicture()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleObjects"}:
				m.OleObjects = NewCT_OleObjects()
				if err := d.DecodeElement(m.OleObjects, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "controls"}:
				m.Controls = NewCT_Controls()
				if err := d.DecodeElement(m.Controls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishItems"}:
				m.WebPublishItems = NewCT_WebPublishItems()
				if err := d.DecodeElement(m.WebPublishItems, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tableParts"}:
				m.TableParts = NewCT_TableParts()
				if err := d.DecodeElement(m.TableParts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Worksheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Worksheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Worksheet and its children
func (m *CT_Worksheet) Validate() error {
	return m.ValidateWithPath("CT_Worksheet")
}

// ValidateWithPath validates the CT_Worksheet and its children, prefixing error messages with path
func (m *CT_Worksheet) ValidateWithPath(path string) error {
	if m.SheetPr != nil {
		if err := m.SheetPr.ValidateWithPath(path + "/SheetPr"); err != nil {
			return err
		}
	}
	if m.Dimension != nil {
		if err := m.Dimension.ValidateWithPath(path + "/Dimension"); err != nil {
			return err
		}
	}
	if m.SheetViews != nil {
		if err := m.SheetViews.ValidateWithPath(path + "/SheetViews"); err != nil {
			return err
		}
	}
	if m.SheetFormatPr != nil {
		if err := m.SheetFormatPr.ValidateWithPath(path + "/SheetFormatPr"); err != nil {
			return err
		}
	}
	for i, v := range m.Cols {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cols[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.SheetData.ValidateWithPath(path + "/SheetData"); err != nil {
		return err
	}
	if m.SheetCalcPr != nil {
		if err := m.SheetCalcPr.ValidateWithPath(path + "/SheetCalcPr"); err != nil {
			return err
		}
	}
	if m.SheetProtection != nil {
		if err := m.SheetProtection.ValidateWithPath(path + "/SheetProtection"); err != nil {
			return err
		}
	}
	if m.ProtectedRanges != nil {
		if err := m.ProtectedRanges.ValidateWithPath(path + "/ProtectedRanges"); err != nil {
			return err
		}
	}
	if m.Scenarios != nil {
		if err := m.Scenarios.ValidateWithPath(path + "/Scenarios"); err != nil {
			return err
		}
	}
	if m.AutoFilter != nil {
		if err := m.AutoFilter.ValidateWithPath(path + "/AutoFilter"); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
			return err
		}
	}
	if m.DataConsolidate != nil {
		if err := m.DataConsolidate.ValidateWithPath(path + "/DataConsolidate"); err != nil {
			return err
		}
	}
	if m.CustomSheetViews != nil {
		if err := m.CustomSheetViews.ValidateWithPath(path + "/CustomSheetViews"); err != nil {
			return err
		}
	}
	if m.MergeCells != nil {
		if err := m.MergeCells.ValidateWithPath(path + "/MergeCells"); err != nil {
			return err
		}
	}
	if m.PhoneticPr != nil {
		if err := m.PhoneticPr.ValidateWithPath(path + "/PhoneticPr"); err != nil {
			return err
		}
	}
	for i, v := range m.ConditionalFormatting {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConditionalFormatting[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DataValidations != nil {
		if err := m.DataValidations.ValidateWithPath(path + "/DataValidations"); err != nil {
			return err
		}
	}
	if m.Hyperlinks != nil {
		if err := m.Hyperlinks.ValidateWithPath(path + "/Hyperlinks"); err != nil {
			return err
		}
	}
	if m.PrintOptions != nil {
		if err := m.PrintOptions.ValidateWithPath(path + "/PrintOptions"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if m.RowBreaks != nil {
		if err := m.RowBreaks.ValidateWithPath(path + "/RowBreaks"); err != nil {
			return err
		}
	}
	if m.ColBreaks != nil {
		if err := m.ColBreaks.ValidateWithPath(path + "/ColBreaks"); err != nil {
			return err
		}
	}
	if m.CustomProperties != nil {
		if err := m.CustomProperties.ValidateWithPath(path + "/CustomProperties"); err != nil {
			return err
		}
	}
	if m.CellWatches != nil {
		if err := m.CellWatches.ValidateWithPath(path + "/CellWatches"); err != nil {
			return err
		}
	}
	if m.IgnoredErrors != nil {
		if err := m.IgnoredErrors.ValidateWithPath(path + "/IgnoredErrors"); err != nil {
			return err
		}
	}
	if m.SmartTags != nil {
		if err := m.SmartTags.ValidateWithPath(path + "/SmartTags"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawing != nil {
		if err := m.LegacyDrawing.ValidateWithPath(path + "/LegacyDrawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawingHF != nil {
		if err := m.LegacyDrawingHF.ValidateWithPath(path + "/LegacyDrawingHF"); err != nil {
			return err
		}
	}
	if m.DrawingHF != nil {
		if err := m.DrawingHF.ValidateWithPath(path + "/DrawingHF"); err != nil {
			return err
		}
	}
	if m.Picture != nil {
		if err := m.Picture.ValidateWithPath(path + "/Picture"); err != nil {
			return err
		}
	}
	if m.OleObjects != nil {
		if err := m.OleObjects.ValidateWithPath(path + "/OleObjects"); err != nil {
			return err
		}
	}
	if m.Controls != nil {
		if err := m.Controls.ValidateWithPath(path + "/Controls"); err != nil {
			return err
		}
	}
	if m.WebPublishItems != nil {
		if err := m.WebPublishItems.ValidateWithPath(path + "/WebPublishItems"); err != nil {
			return err
		}
	}
	if m.TableParts != nil {
		if err := m.TableParts.ValidateWithPath(path + "/TableParts"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
