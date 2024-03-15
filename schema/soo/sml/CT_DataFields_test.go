// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/codrocker/gooxml/schema/soo/sml"
)

func TestCT_DataFieldsConstructor(t *testing.T) {
	v := sml.NewCT_DataFields()
	if v == nil {
		t.Errorf("sml.NewCT_DataFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataFields should validate: %s", err)
	}
}

func TestCT_DataFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataFields()
	xml.Unmarshal(buf, v2)
}
