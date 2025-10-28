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
	uti "github.com/craterdog/go-essential-utilities/v8"
	mat "math"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func PercentageClass() PercentageClassLike {
	return percentageClass()
}

// Constructor Methods

func (c *percentageClass_) Percentage(
	float float64,
) PercentageLike {
	return percentage_(float / 100.0)
}

func (c *percentageClass_) PercentageFromInteger(
	integer int,
) PercentageLike {
	var float = float64(integer)
	return percentage_(float / 100.0)
}

func (c *percentageClass_) PercentageFromSource(
	source string,
) PercentageLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the percentage constructor method: %s",
			source,
		)
		panic(message)
	}
	var float, _ = stc.ParseFloat(matches[1], 64) // Strip off the '%' suffix.
	return percentage_(float / 100.0)
}

// Constant Methods

func (c *percentageClass_) Undefined() PercentageLike {
	return c.undefined_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v percentage_) GetClass() PercentageClassLike {
	return percentageClass()
}

func (v percentage_) AsIntrinsic() float64 {
	return float64(v)
}

// Attribute Methods

// Continuous Methods

func (v percentage_) AsSource() string {
	return numberClass().sourceFromFloat(float64(v)*100.0) + "%"
}

func (v percentage_) AsFloat() float64 {
	return float64(v * 100.0)
}

func (v percentage_) HasMagnitude() bool {
	return v.IsDefined() && !(v.IsZero() || v.IsInfinite())
}

func (v percentage_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

func (v percentage_) IsDefined() bool {
	return !mat.IsNaN(float64(v))
}

func (v percentage_) IsMinimum() bool {
	return v == -mat.MaxFloat64
}

func (v percentage_) IsZero() bool {
	return v == 0
}

func (v percentage_) IsMaximum() bool {
	return v == mat.MaxFloat64
}

// Polarized Methods

func (v percentage_) IsNegative() bool {
	return v < 0
}

// PROTECTED INTERFACE

func (v percentage_) String() string {
	return v.AsSource()
}

// Private Methods

// Instance Structure

type percentage_ float64

// Class Structure

type percentageClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ PercentageLike
}

// Class Reference

func percentageClass() *percentageClass_ {
	return percentageClassReference_
}

var percentageClassReference_ = &percentageClass_{
	// Initialize the class constants.
	matcher_:   reg.MustCompile("^(" + real_ + ")%"),
	undefined_: percentage_(mat.NaN()),
}
