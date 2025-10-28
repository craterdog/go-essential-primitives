/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package elements

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v8"
	mat "math"
	reg "regexp"
	uni "unicode"
	utf "unicode/utf8"
)

// CLASS INTERFACE

// Access Function

func GlyphClass() GlyphClassLike {
	return glyphClass()
}

// Constructor Methods

func (c *glyphClass_) Glyph(
	rune_ rune,
) GlyphLike {
	return glyph_(rune_)
}

func (c *glyphClass_) GlyphFromInteger(
	integer int,
) GlyphLike {
	return glyph_(integer)
}

func (c *glyphClass_) GlyphFromSource(
	source string,
) GlyphLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the glyph constructor method: %s",
			source,
		)
		panic(message)
	}
	var rune_, _ = utf.DecodeRuneInString(matches[1]) // Strip off the single quotes.
	return glyph_(rune_)
}

// Constant Methods

func (c *glyphClass_) Undefined() GlyphLike {
	return c.undefined_
}

// Function Methods

func (c *glyphClass_) ToLowercase(glyph GlyphLike) GlyphLike {
	var rune_ = glyph.AsIntrinsic()
	rune_ = uni.ToLower(rune_)
	return glyph_(rune_)
}

func (c *glyphClass_) ToUppercase(glyph GlyphLike) GlyphLike {
	var rune_ = glyph.AsIntrinsic()
	rune_ = uni.ToUpper(rune_)
	return glyph_(rune_)
}

// INSTANCE INTERFACE

// Principal Methods

func (v glyph_) GetClass() GlyphClassLike {
	return glyphClass()
}

func (v glyph_) AsIntrinsic() rune {
	return rune(v)
}

// Attribute Methods

// Discrete Methods

func (v glyph_) AsSource() string {
	return "'" + string([]rune{rune(v)}) + "'"
}

func (v glyph_) AsInteger() int {
	return int(v)
}

func (v glyph_) IsDefined() bool {
	return v >= 0
}

func (v glyph_) IsMinimum() bool {
	return v == 0
}

func (v glyph_) IsZero() bool {
	return v == 0
}

func (v glyph_) IsMaximum() bool {
	return v == mat.MaxInt32
}

// PROTECTED INTERFACE

func (v glyph_) String() string {
	return v.AsSource()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	base16_  = base10_ + "|[a-f]"
	control_ = "\\p{Cc}"
	escape_  = "\\\\" + unicode_ + "|[abfnrtv\\\\]"
	unicode_ = "u(?:" + base16_ + "){4}|U(?:" + base16_ + "){8}"
)

// Instance Structure

type glyph_ rune

// Class Structure

type glyphClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ GlyphLike
}

// Class Reference

func glyphClass() *glyphClass_ {
	return glyphClassReference_
}

var glyphClassReference_ = &glyphClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'((?:" + escape_ + ")|[^" + control_ + "])'",
	),
	undefined_: glyph_(-1),
}
