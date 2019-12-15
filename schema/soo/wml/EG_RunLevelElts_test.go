// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/carmel/gooxml/schema/soo/wml"
)

func TestEG_RunLevelEltsConstructor(t *testing.T) {
	v := wml.NewEG_RunLevelElts()
	if v == nil {
		t.Errorf("wml.NewEG_RunLevelElts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RunLevelElts should validate: %s", err)
	}
}

func TestEG_RunLevelEltsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RunLevelElts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RunLevelElts()
	xml.Unmarshal(buf, v2)
}
