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
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func AngleClass() AngleClassLike {
	return angleClass()
}

// Constructor Methods

func (c *angleClass_) Angle(
	radians float64,
) AngleLike {
	return c.angleFromFloat(radians)
}

func (c *angleClass_) AngleFromSource(
	source string,
) AngleLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the angle constructor method: %s",
			source,
		)
		panic(message)
	}
	var match = matches[1] // Strip off the leading '~' character.
	switch match {
	case "pi", "π":
		return c.pi_
	case "tau", "τ":
		return c.tau_
	default:
		var float, _ = stc.ParseFloat(match, 64)
		return c.angleFromFloat(float)
	}
}

// Constant Methods

func (c *angleClass_) Undefined() AngleLike {
	return c.undefined_
}

func (c *angleClass_) Zero() AngleLike {
	return c.zero_
}

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Function Methods

func (c *angleClass_) Inverse(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angle.AsFloat() - angleClass().Pi().AsFloat())
	return result_
}

func (c *angleClass_) Sum(
	first AngleLike,
	second AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(first.AsFloat() + second.AsFloat())
	return result_
}

func (c *angleClass_) Difference(
	first AngleLike,
	second AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(first.AsFloat() - second.AsFloat())
	return result_
}

func (c *angleClass_) Scaled(
	angle AngleLike,
	factor float64,
) AngleLike {
	var result_ = c.angleFromFloat(angle.AsFloat() * factor)
	return result_
}

func (c *angleClass_) Complement(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angleClass().Pi().AsFloat()/2.0 - angle.AsFloat())
	return result_
}

func (c *angleClass_) Supplement(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angleClass().Pi().AsFloat() - angle.AsFloat())
	return result_
}

func (c *angleClass_) Conjugate(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(-angle.AsFloat())
	return result_
}

func (c *angleClass_) Cosine(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 1.0
	case mat.Pi * 0.25:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		result_ = 0.0
	case mat.Pi * 0.75:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi:
		result_ = -1.0
	case mat.Pi * 1.25:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		result_ = 0.0
	case mat.Pi * 1.75:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		result_ = 1.0
	default:
		result_ = mat.Cos(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcCosine(
	x float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Acos(x))
	return result_
}

func (c *angleClass_) Sine(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 0.0
	case mat.Pi * 0.25:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		result_ = 1.0
	case mat.Pi * 0.75:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi:
		result_ = 0.0
	case mat.Pi * 1.25:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		result_ = -1.0
	case mat.Pi * 1.75:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		result_ = 0.0
	default:
		result_ = mat.Sin(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcSine(
	y float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Asin(y))
	return result_
}

func (c *angleClass_) Tangent(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 0.0
	case mat.Pi * 0.25:
		result_ = 1.0
	case mat.Pi * 0.5:
		result_ = mat.Inf(1)
	case mat.Pi * 0.75:
		result_ = -1.0
	case mat.Pi:
		result_ = 0.0
	case mat.Pi * 1.25:
		result_ = 1.0
	case mat.Pi * 1.5:
		result_ = mat.Inf(1)
	case mat.Pi * 1.75:
		result_ = -1.0
	case mat.Pi * 2.0:
		result_ = 0.0
	default:
		result_ = mat.Tan(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcTangent(
	x float64,
	y float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Atan2(y, x))
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v angle_) GetClass() AngleClassLike {
	return angleClass()
}

func (v angle_) AsIntrinsic() float64 {
	return float64(v)
}

func (v angle_) AsUnits(
	units Units,
) float64 {
	var result_ float64
	var radians = v.AsIntrinsic()
	var pi = angleClass().pi_.AsIntrinsic()
	var tau = angleClass().tau_.AsIntrinsic()
	switch units {
	case Degrees:
		result_ = 360.0 * radians / tau
	case Radians:
		result_ = radians
	case Gradians:
		result_ = 200.0 / pi
	}
	return result_
}

func (v angle_) AsParts() (
	x float64,
	y float64,
) {
	var complex_ = v.AsIntrinsic()
	x = mat.Cos(complex_)
	y = mat.Sin(complex_)
	return
}

// Attribute Methods

// Continuous Methods

func (v angle_) AsSource() string {
	var result_ = angleClass().sourceFromAngle(v)
	return result_
}

func (v angle_) AsFloat() float64 {
	return float64(v)
}

func (v angle_) HasMagnitude() bool {
	return !v.IsZero()
}

func (v angle_) IsInfinite() bool {
	return false
}

func (v angle_) IsDefined() bool {
	return v != angleClass().undefined_
}

func (v angle_) IsMinimum() bool {
	return v == angleClass().zero_
}

func (v angle_) IsZero() bool {
	return v == angleClass().zero_ || v == angleClass().tau_
}

func (v angle_) IsMaximum() bool {
	return v == angleClass().tau_
}

// PROTECTED INTERFACE

func (v Units) String() string {
	var source string
	switch v {
	case Degrees:
		source = "Degrees"
	case Radians:
		source = "Radians"
	case Gradians:
		source = "Gradians"
	}
	return source
}

func (v angle_) String() string {
	return v.AsSource()
}

// Private Methods

func (c *angleClass_) angleFromFloat(float float64) angle_ {
	float = c.normalizeValue(float)
	float = c.lockAngle(float)
	return angle_(float)
}

func (c *angleClass_) lockAngle(value float64) float64 {
	var pi = angleClass().Pi().AsIntrinsic()
	var value32 = float32(value)
	switch {
	case mat.Abs(value) <= 1.2246467991473515e-16:
		value = 0
	case value32 == float32(0.5*pi):
		value = 0.5 * pi
	case value32 == float32(pi):
		value = pi
	case value32 == float32(1.5*pi):
		value = 1.5 * pi
	}
	return value
}

func (c *angleClass_) normalizeValue(value float64) float64 {
	var tau = angleClass().Tau().AsIntrinsic()
	if value < -tau || value >= tau {
		// Normalize the value to the range [-τ..τ).
		value = mat.Remainder(value, tau)
	}
	if value < 0.0 {
		// Normalize the value to the range [0..τ).
		value = value + tau
	}
	return value
}

func (c *angleClass_) sourceFromAngle(angle angle_) string {
	var source string
	switch angle {
	case c.pi_:
		source = "~π"
	case c.tau_:
		source = "~τ"
	default:
		source = "~" + stc.FormatFloat(float64(angle), 'G', 15, 64)
	}
	return source
}

// Instance Structure

type angle_ float64

// Class Structure

type angleClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ AngleLike
	zero_      AngleLike
	pi_        AngleLike
	tau_       AngleLike
}

// Class Reference

func angleClass() *angleClass_ {
	return angleClassReference_
}

var angleClassReference_ = &angleClass_{
	// Initialize the class constants.
	matcher_:   reg.MustCompile("^~(0|" + amplitude_ + ")"),
	undefined_: angle_(mat.NaN()),
	zero_:      angle_(0.0),
	pi_:        angle_(mat.Pi),
	tau_:       angle_(2.0 * mat.Pi),
}
