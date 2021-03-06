// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: signature

// Package signature defines types representing interface and method signatures.
package signature

import (
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Embed describes the signature of an embedded interface.
type Embed struct {
	Name    string
	PkgPath string
	Doc     string
}

func (Embed) __VDLReflect(struct {
	Name string `vdl:"signature.Embed"`
}) {
}

func (x Embed) VDLIsZero() bool {
	return x == Embed{}
}

func (x Embed) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if x.PkgPath != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.PkgPath); err != nil {
			return err
		}
	}
	if x.Doc != "" {
		if err := enc.NextFieldValueString(2, vdl.StringType, x.Doc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Embed) VDLRead(dec vdl.Decoder) error {
	*x = Embed{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_1 {
			index = __VDLType_struct_1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.PkgPath = value
			}
		case 2:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Doc = value
			}
		}
	}
}

// Arg describes the signature of a single argument.
type Arg struct {
	Name string
	Doc  string
	Type *vdl.Type // Type of the argument.
}

func (Arg) __VDLReflect(struct {
	Name string `vdl:"signature.Arg"`
}) {
}

func (x Arg) VDLIsZero() bool {
	if x.Name != "" {
		return false
	}
	if x.Doc != "" {
		return false
	}
	if x.Type != nil && x.Type != vdl.AnyType {
		return false
	}
	return true
}

func (x Arg) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if x.Doc != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Doc); err != nil {
			return err
		}
	}
	if x.Type != nil && x.Type != vdl.AnyType {
		if err := enc.NextFieldValueTypeObject(2, x.Type); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Arg) VDLRead(dec vdl.Decoder) error {
	*x = Arg{
		Type: vdl.AnyType,
	}
	if err := dec.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_2 {
			index = __VDLType_struct_2.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Doc = value
			}
		case 2:
			switch value, err := dec.ReadValueTypeObject(); {
			case err != nil:
				return err
			default:
				x.Type = value
			}
		}
	}
}

// Method describes the signature of an interface method.
type Method struct {
	Name      string
	Doc       string
	InArgs    []Arg        // Input arguments
	OutArgs   []Arg        // Output arguments
	InStream  *Arg         // Input stream (optional)
	OutStream *Arg         // Output stream (optional)
	Tags      []*vdl.Value // Method tags
}

func (Method) __VDLReflect(struct {
	Name string `vdl:"signature.Method"`
}) {
}

func (x Method) VDLIsZero() bool {
	if x.Name != "" {
		return false
	}
	if x.Doc != "" {
		return false
	}
	if len(x.InArgs) != 0 {
		return false
	}
	if len(x.OutArgs) != 0 {
		return false
	}
	if x.InStream != nil {
		return false
	}
	if x.OutStream != nil {
		return false
	}
	if len(x.Tags) != 0 {
		return false
	}
	return true
}

func (x Method) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_3); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if x.Doc != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Doc); err != nil {
			return err
		}
	}
	if len(x.InArgs) != 0 {
		if err := enc.NextField(2); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.InArgs); err != nil {
			return err
		}
	}
	if len(x.OutArgs) != 0 {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.OutArgs); err != nil {
			return err
		}
	}
	if x.InStream != nil {
		if err := enc.NextField(4); err != nil {
			return err
		}
		enc.SetNextStartValueIsOptional()
		if err := x.InStream.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.OutStream != nil {
		if err := enc.NextField(5); err != nil {
			return err
		}
		enc.SetNextStartValueIsOptional()
		if err := x.OutStream.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Tags) != 0 {
		if err := enc.NextField(6); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.Tags); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []Arg) error {
	if err := enc.StartValue(__VDLType_list_4); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []*vdl.Value) error {
	if err := enc.StartValue(__VDLType_list_6); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if elem == nil {
			if err := enc.NilValue(vdl.AnyType); err != nil {
				return err
			}
		} else {
			if err := elem.VDLWrite(enc); err != nil {
				return err
			}
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Method) VDLRead(dec vdl.Decoder) error {
	*x = Method{}
	if err := dec.StartValue(__VDLType_struct_3); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_3 {
			index = __VDLType_struct_3.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Doc = value
			}
		case 2:
			if err := __VDLReadAnon_list_1(dec, &x.InArgs); err != nil {
				return err
			}
		case 3:
			if err := __VDLReadAnon_list_1(dec, &x.OutArgs); err != nil {
				return err
			}
		case 4:
			if err := dec.StartValue(__VDLType_optional_5); err != nil {
				return err
			}
			if dec.IsNil() {
				x.InStream = nil
				if err := dec.FinishValue(); err != nil {
					return err
				}
			} else {
				x.InStream = new(Arg)
				dec.IgnoreNextStartValue()
				if err := x.InStream.VDLRead(dec); err != nil {
					return err
				}
			}
		case 5:
			if err := dec.StartValue(__VDLType_optional_5); err != nil {
				return err
			}
			if dec.IsNil() {
				x.OutStream = nil
				if err := dec.FinishValue(); err != nil {
					return err
				}
			} else {
				x.OutStream = new(Arg)
				dec.IgnoreNextStartValue()
				if err := x.OutStream.VDLRead(dec); err != nil {
					return err
				}
			}
		case 6:
			if err := __VDLReadAnon_list_2(dec, &x.Tags); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]Arg) error {
	if err := dec.StartValue(__VDLType_list_4); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]Arg, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem Arg
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]*vdl.Value) error {
	if err := dec.StartValue(__VDLType_list_6); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]*vdl.Value, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem *vdl.Value
			elem = new(vdl.Value)
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// Interface describes the signature of an interface.
type Interface struct {
	Name    string
	PkgPath string
	Doc     string
	Embeds  []Embed  // No special ordering.
	Methods []Method // Ordered by method name.
}

func (Interface) __VDLReflect(struct {
	Name string `vdl:"signature.Interface"`
}) {
}

func (x Interface) VDLIsZero() bool {
	if x.Name != "" {
		return false
	}
	if x.PkgPath != "" {
		return false
	}
	if x.Doc != "" {
		return false
	}
	if len(x.Embeds) != 0 {
		return false
	}
	if len(x.Methods) != 0 {
		return false
	}
	return true
}

func (x Interface) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_7); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if x.PkgPath != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.PkgPath); err != nil {
			return err
		}
	}
	if x.Doc != "" {
		if err := enc.NextFieldValueString(2, vdl.StringType, x.Doc); err != nil {
			return err
		}
	}
	if len(x.Embeds) != 0 {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_3(enc, x.Embeds); err != nil {
			return err
		}
	}
	if len(x.Methods) != 0 {
		if err := enc.NextField(4); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_4(enc, x.Methods); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_3(enc vdl.Encoder, x []Embed) error {
	if err := enc.StartValue(__VDLType_list_8); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_4(enc vdl.Encoder, x []Method) error {
	if err := enc.StartValue(__VDLType_list_9); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Interface) VDLRead(dec vdl.Decoder) error {
	*x = Interface{}
	if err := dec.StartValue(__VDLType_struct_7); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_7 {
			index = __VDLType_struct_7.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.PkgPath = value
			}
		case 2:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Doc = value
			}
		case 3:
			if err := __VDLReadAnon_list_3(dec, &x.Embeds); err != nil {
				return err
			}
		case 4:
			if err := __VDLReadAnon_list_4(dec, &x.Methods); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_3(dec vdl.Decoder, x *[]Embed) error {
	if err := dec.StartValue(__VDLType_list_8); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]Embed, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem Embed
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

func __VDLReadAnon_list_4(dec vdl.Decoder, x *[]Method) error {
	if err := dec.StartValue(__VDLType_list_9); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]Method, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem Method
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1   *vdl.Type
	__VDLType_struct_2   *vdl.Type
	__VDLType_struct_3   *vdl.Type
	__VDLType_list_4     *vdl.Type
	__VDLType_optional_5 *vdl.Type
	__VDLType_list_6     *vdl.Type
	__VDLType_struct_7   *vdl.Type
	__VDLType_list_8     *vdl.Type
	__VDLType_list_9     *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Embed)(nil))
	vdl.Register((*Arg)(nil))
	vdl.Register((*Method)(nil))
	vdl.Register((*Interface)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Embed)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*Arg)(nil)).Elem()
	__VDLType_struct_3 = vdl.TypeOf((*Method)(nil)).Elem()
	__VDLType_list_4 = vdl.TypeOf((*[]Arg)(nil))
	__VDLType_optional_5 = vdl.TypeOf((*Arg)(nil))
	__VDLType_list_6 = vdl.TypeOf((*[]*vdl.Value)(nil))
	__VDLType_struct_7 = vdl.TypeOf((*Interface)(nil)).Elem()
	__VDLType_list_8 = vdl.TypeOf((*[]Embed)(nil))
	__VDLType_list_9 = vdl.TypeOf((*[]Method)(nil))

	return struct{}{}
}
