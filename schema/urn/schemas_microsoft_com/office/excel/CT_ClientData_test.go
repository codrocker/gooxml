// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package excel_test

import (
	"encoding/xml"
	"testing"

	"github.com/carmel/gooxml/schema/urn/schemas_microsoft_com/office/excel"
)

func TestCT_ClientDataConstructor(t *testing.T) {
	v := excel.NewCT_ClientData()
	if v == nil {
		t.Errorf("excel.NewCT_ClientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed excel.CT_ClientData should validate: %s", err)
	}
}

func TestCT_ClientDataMarshalUnmarshal(t *testing.T) {
	v := excel.NewCT_ClientData()
	buf, _ := xml.Marshal(v)
	v2 := excel.NewCT_ClientData()
	xml.Unmarshal(buf, v2)
}
