// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package schemaLibrary

import "github.com/codrocker/gooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_Schema", NewCT_Schema)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_SchemaLibrary", NewCT_SchemaLibrary)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "schemaLibrary", NewSchemaLibrary)
}
