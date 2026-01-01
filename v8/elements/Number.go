/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	cmp "math/cmplx"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func NumberClass() NumberClassLike {
	return numberClass()
}

// Constructor Methods

// NOTE:
// This constructor creates a new number element from the specified complex
// number.  Each number element is mapped to the Riemann Sphere.
//   - https://en.wikipedia.org/wiki/Riemann_sphere
//
// This mapping removes many of the idiosyncrasies associated with the normal
// complex plane. There is only one value for zero and one value for infinity.
// This simplifies a lot of the mathematical operations:
//
//	z + zero => z
//	z + infinity => infinity
//
//	z - infinity => infinity     {z != infinity}
//	infinity - z => infinity     {z != infinity}
//
//	z * zero => zero             {z != infinity}
//	z * infinity => infinity     {z != zero}
//
//	z / zero => infinity         {z != zero}
//	zero / z => zero             {z != zero}
//
//	z / infinity => zero         {z != infinity}
//	infinity / z => infinity     {z != infinity}
//
//	z ^ zero => one              {by definition}
//	zero ^ z => zero             {z != zero}
//
//	z ^ infinity => zero         {|z| < one}
//	z ^ infinity => one          {|z| = one}
//	z ^ infinity => infinity     {|z| > one}
//	infinity ^ z => infinity     {z != zero}
//
//	log(z, zero) => infinity     {zero < z < infinity}
//	log(zero, z) => zero         {zero < z < infinity}
//
//	log(z, one) => zero          {zero < z}
//	log(one, z) => zero          {zero < z}
//
//	log(z, infinity) => infinity {zero < z < infinity}
//	log(infinity, z) => zero     {zero < z < infinity}
//
// This leaves only the following operations undefined:
//
//	infinity - infinity => undefined
//	zero * infinity => undefined
//	zero / zero => undefined
//	infinity / infinity => undefined
//	log(zero, zero) => undefined
//	log(zero, infinity) => undefined
//	log(infinity, zero) => undefined
//	log(infinity, infinity) => undefined
//
// The resulting number system is easier to use for most applications. For
// numerical analysis the ANSI plus and minus zero values are often used as a
// crutch that leads to misleading convergence information for oscillating
// functions. In the case of numerical analysis it is probably better to track
// the course of the function as it converges than to look at the final value.
func (c *numberClass_) Number(
	complex_ complex128,
) NumberLike {
	return c.normalize(complex_)
}

func (c *numberClass_) NumberFromPolar(
	magnitude float64,
	angle float64,
) NumberLike {
	// Go complex types only use rectangular form so convert to rectangular.
	var complex_ = cmp.Rect(magnitude, angle)
	return c.normalize(complex_)
}

func (c *numberClass_) NumberFromRectangular(
	real_ float64,
	imaginary float64,
) NumberLike {
	var complex_ = complex(real_, imaginary)
	return c.normalize(complex_)
}

func (c *numberClass_) NumberFromInteger(
	integer int,
) NumberLike {
	var complex_ = complex(float64(integer), 0)
	return c.normalize(complex_)
}

func (c *numberClass_) NumberFromFloat(
	float float64,
) NumberLike {
	var complex_ = complex(float, 0)
	return c.normalize(complex_)
}

func (c *numberClass_) NumberFromSource(
	source string,
) NumberLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the number constructor method: %s",
			source,
		)
		panic(message)
	}
	var complex_ = c.complexFromMatches(matches)
	return c.normalize(complex_)
}

// Constant Methods

func (c *numberClass_) Undefined() NumberLike {
	return c.undefined_
}

func (c *numberClass_) Zero() NumberLike {
	return c.zero_
}

func (c *numberClass_) One() NumberLike {
	return c.one_
}

func (c *numberClass_) I() NumberLike {
	return c.i_
}

func (c *numberClass_) E() NumberLike {
	return c.e_
}

func (c *numberClass_) Pi() NumberLike {
	return c.pi_
}

func (c *numberClass_) Phi() NumberLike {
	return c.phi_
}

func (c *numberClass_) Tau() NumberLike {
	return c.tau_
}

func (c *numberClass_) Minimum() NumberLike {
	return c.minimum_
}

func (c *numberClass_) Maximum() NumberLike {
	return c.maximum_
}

func (c *numberClass_) Infinity() NumberLike {
	return c.infinity_
}

// Function Methods

func (c *numberClass_) Inverse(
	number NumberLike,
) NumberLike {
	if number.HasMagnitude() {
		number = c.normalize(-number.AsIntrinsic())
	}
	return number
}

func (c *numberClass_) Reciprocal(
	number NumberLike,
) NumberLike {
	return c.normalize(1.0 / number.AsIntrinsic())
}

func (c *numberClass_) Conjugate(
	number NumberLike,
) NumberLike {
	return c.normalize(cmp.Conj(number.AsIntrinsic()))
}

func (c *numberClass_) Sum(
	first NumberLike,
	second NumberLike,
) NumberLike {
	switch {
	case !first.IsDefined() || !second.IsDefined():
		return c.undefined_
	case first.IsInfinite() || second.IsInfinite():
		return c.infinity_
	default:
		return c.normalize(first.AsIntrinsic() + second.AsIntrinsic())
	}
}

func (c *numberClass_) Difference(
	first NumberLike,
	second NumberLike,
) NumberLike {
	switch {
	case !first.IsDefined() || !second.IsDefined():
		return c.undefined_
	case first.IsInfinite() && second.IsInfinite():
		return c.undefined_
	case first.IsInfinite() || second.IsInfinite():
		return c.infinity_
	default:
		return c.normalize(first.AsIntrinsic() - second.AsIntrinsic())
	}
}

func (c *numberClass_) Scaled(
	number NumberLike,
	factor float64,
) NumberLike {
	number = c.Product(number, c.normalize(complex(factor, 0)))
	return number
}

func (c *numberClass_) Product(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var number NumberLike
	switch {
	case !first.IsDefined() || !second.IsDefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined_
	case first.IsInfinite() && !second.IsZero():
		// Infinity times anything other than zero is infinite.
		number = c.infinity_
	case second.IsInfinite() && !first.IsZero():
		// Anything other than zero times infinity is infinite.
		number = c.infinity_
	default:
		number = c.normalize(first.AsIntrinsic() * second.AsIntrinsic())
	}
	return number
}

func (c *numberClass_) Quotient(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var number NumberLike
	switch {
	case !first.IsDefined() || !second.IsDefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined_
	case first.IsZero() && second.IsZero():
		// Zero divided by zero is undefined.
		number = c.undefined_
	case first.IsInfinite() && second.IsInfinite():
		// Infinity divided by infinity is undefined.
		number = c.undefined_
	case first.IsZero() && !second.IsZero():
		// Zero divided by anything other than zero is zero.
		number = c.zero_
	case second.IsZero() && !first.IsZero():
		// Anything other than zero divided by zero is infinite.
		number = c.infinity_
	case first.IsInfinite() && !second.IsInfinite():
		// Zero divided by anything other than zero is zero.
		number = c.infinity_
	case second.IsInfinite() && !first.IsInfinite():
		// Anything other than zero divided by zero is infinite.
		number = c.zero_
	default:
		number = c.normalize(first.AsIntrinsic() / second.AsIntrinsic())
	}
	return number
}

func (c *numberClass_) Remainder(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var m1 = cmp.Abs(first.AsIntrinsic())
	var p1 = cmp.Phase(first.AsIntrinsic())
	var m2 = cmp.Abs(second.AsIntrinsic())
	var p2 = cmp.Phase(second.AsIntrinsic())
	var magnitude = c.lockFloat(mat.Remainder(m1, m2))
	var angle = c.lockAngle(p2 - p1)
	var number = c.NumberFromPolar(magnitude, angle)
	return number
}

func (c *numberClass_) Power(
	base NumberLike,
	exponent NumberLike,
) NumberLike {
	var number NumberLike
	switch {
	case !base.IsDefined() || !exponent.IsDefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined_
	case exponent.IsZero():
		// Anything to the zero power is 1 by definition.
		number = c.one_
	case base.IsZero():
		// Zero to any power other than zero is still zero.
		number = c.zero_
	case base.IsInfinite():
		// Infinity to any power other than zero is infinite.
		number = c.infinity_
	case exponent.IsInfinite():
		var magnitude = base.GetMagnitude()
		switch {
		case magnitude < 1:
			// A magnitude less than one to an infinite power is zero.
			number = c.zero_
		case magnitude == 1:
			// A magnitude equal to one to an infinite power is one.
			number = c.one_
		case magnitude > 1:
			// A magnitude greater than one to an infinite power is infinite.
			number = c.infinity_
		default:
			panic(fmt.Sprintf("An impossible magnitude was encountered: %v", magnitude))
		}
	default:
		number = c.normalize(cmp.Pow(base.AsIntrinsic(), exponent.AsIntrinsic()))
	}
	return number
}

func (c *numberClass_) Logarithm(
	base NumberLike,
	number NumberLike,
) NumberLike {
	// logB(z) => ln(z) / ln(b)
	var lnB = cmp.Log(base.AsIntrinsic())
	var lnZ = cmp.Log(number.AsIntrinsic())
	var logB = lnZ / lnB
	return c.normalize(logB)
}

// INSTANCE INTERFACE

// Principal Methods

func (v number_) GetClass() NumberClassLike {
	return numberClass()
}

func (v number_) AsIntrinsic() complex128 {
	return complex128(v)
}

func (v number_) AsRectangular() string {
	var source string
	switch {
	case v.IsZero():
		source = "0"
	case v.IsMinimum():
		source = "-∞"
	case v.IsMaximum():
		source = "+∞"
	case v.IsInfinite():
		source = "∞"
	case !v.IsDefined():
		source = "undefined"
	default:
		var realPart = v.GetReal()
		var imagPart = v.GetImaginary()
		switch {
		case imagPart == 0:
			source = numberClass().sourceFromFloat(realPart)
		case realPart == 0:
			source = numberClass().sourceFromFloat(imagPart) + "i"
		default:
			source = numberClass().sourceFromFloat(realPart)
			if imagPart > 0 {
				source += "+"
			}
			source += numberClass().sourceFromFloat(imagPart) + "i"
		}
	}
	return source
}

func (v number_) AsPolar() string {
	var source string
	switch {
	case v.IsZero():
		source = "0"
	case v.IsMinimum():
		source = "-∞"
	case v.IsMaximum():
		source = "+∞"
	case v.IsInfinite():
		source = "∞"
	case !v.IsDefined():
		source = "undefined"
	default:
		var magnitude = v.GetMagnitude()
		var angle = v.GetAngle()
		source = numberClass().sourceFromFloat(magnitude)
		if angle > 0 {
			source += "e^~" + numberClass().sourceFromFloat(angle) + "i"
		}
	}
	return source
}

func (v number_) GetReal() float64 {
	return real(v)
}

func (v number_) GetImaginary() float64 {
	return imag(v)
}

func (v number_) GetMagnitude() float64 {
	return numberClass().lockFloat(cmp.Abs(complex128(v)))
}

func (v number_) GetAngle() float64 {
	return cmp.Phase(complex128(v))
}

// Attribute Methods

// Continuous Methods

func (v number_) AsSource() string {
	return v.AsRectangular()
}

func (v number_) AsFloat() float64 {
	return real(v)
}

func (v number_) HasMagnitude() bool {
	return v.IsDefined() && !v.IsZero() && !v.IsMinimum() &&
		!v.IsMaximum() && !v.IsInfinite()
}

func (v number_) IsInfinite() bool {
	return v == numberClass().infinity_
}

func (v number_) IsDefined() bool {
	return !(mat.IsNaN(real(v)) || mat.IsNaN(imag(v)))
}

func (v number_) IsMinimum() bool {
	return v == numberClass().minimum_
}

func (v number_) IsZero() bool {
	return v == numberClass().zero_
}

func (v number_) IsMaximum() bool {
	return v == numberClass().maximum_
}

// Polarized Methods

func (v number_) IsNegative() bool {
	return real(v) < 0
}

// PROTECTED INTERFACE

func (v number_) String() string {
	return v.AsSource()
}

// Private Methods

// This private function returns the complex number associated with the
// specified regular expression matches.
func (c *numberClass_) complexFromMatches(matches []string) complex128 {
	var complex_ complex128
	switch {
	case len(matches[1]) > 0:
		// This is a complex number in polar form.
		var magnitude = c.floatFromSource(matches[2])
		var angle = c.floatFromSource(matches[3])
		complex_ = cmp.Rect(magnitude, angle)
	case len(matches[4]) > 0:
		// This is a complex number in rectangular form.
		var realPart = c.floatFromSource(matches[5])
		var imaginaryPart = c.floatFromSource(matches[6])
		complex_ = complex(realPart, imaginaryPart)
	case len(matches[7]) > 0:
		// This is a pure imaginary number.
		var realPart = 0.0
		var imaginaryPart = c.floatFromSource(matches[7][:len(matches[7])-1])
		complex_ = complex(realPart, imaginaryPart)
	case len(matches[8]) > 0:
		// This is a pure real number.
		switch matches[8] {
		case "∞", "infinity":
			complex_ = complex(mat.Inf(0), mat.Inf(0))
		default:
			var realPart = c.floatFromSource(matches[8])
			var imaginaryPart = 0.0
			complex_ = complex(realPart, imaginaryPart)
		}
	}
	return complex_
}

// This private function returns the floating point value for the specified
// string.
func (c *numberClass_) floatFromSource(source string) float64 {
	var float float64
	switch source {
	case "+e", "e":
		float = mat.E
	case "-e":
		float = -mat.E
	case "+pi", "+π", "pi", "π":
		float = mat.Pi
	case "-pi", "-π":
		float = -mat.Pi
	case "+phi", "+φ", "phi", "φ":
		float = mat.Phi
	case "-phi", "-φ":
		float = -mat.Phi
	case "+tau", "+τ", "tau", "τ":
		float = mat.Pi * 2.0
	case "-tau", "-τ":
		float = -mat.Pi * 2.0
	case "+infinity", "+∞", "infinity", "∞":
		float = mat.Inf(1)
	case "-infinity", "-∞":
		float = mat.Inf(-1)
	case "undefined":
		float = mat.NaN()
	default:
		float, _ = stc.ParseFloat(source, 64)
	}
	return float
}

// NOTE:
// This private function uses the single precision floating point range to lock
// a double precision floating point value onto 0, 1, -1, or ∞ when the
// floating point value falls outside the single precision range for these
// values. Otherwise, the floating point value is returned unchanged.
func (c *numberClass_) lockFloat(float float64) float64 {
	var limited = float32(float)
	switch {
	case mat.Abs(float) <= 1.2246467991473515e-16:
		// We must handle round-off errors for some trigonometric functions.
		float = 0
	case limited == -1:
		float = -1
	case limited == 1:
		float = 1
	case mat.IsInf(float, 0):
		float = mat.Inf(1)
	}
	return float
}

// NOTE:
// This private function uses the single precision floating point range to lock
// a double precision floating point angle onto 0, π/2, π, or 3π/2 when the
// angle falls outside of the single precision range for these values.
// Otherwise, the angle is returned unchanged.
func (c *numberClass_) lockAngle(angle float64) float64 {
	var angle32 = float32(angle)
	switch {
	case mat.Abs(angle) <= 1.2246467991473515e-16:
		// We must handle round-off errors for some trigonometric functions.
		angle = 0
	case angle32 == float32(0.5*mat.Pi):
		angle = 0.5 * mat.Pi
	case angle32 == float32(mat.Pi):
		angle = mat.Pi
	case angle32 == float32(1.5*mat.Pi):
		angle = 1.5 * mat.Pi
	}
	return angle
}

func (c *numberClass_) normalize(complex_ complex128) NumberLike {
	var number NumberLike
	switch {
	case cmp.Abs(complex_) == 0:
		// Normalize any versions of zero.
		number = c.zero_
	case imag(complex_) == 0 && mat.IsInf(real(complex_), -1):
		// Normalize any real negative infinity.
		number = c.minimum_
	case imag(complex_) == 0 && mat.IsInf(real(complex_), 1):
		// Normalize any real positive infinity.
		number = c.maximum_
	case cmp.IsInf(complex_):
		// Normalize any complex infinities.
		number = c.infinity_
	case cmp.IsNaN(complex_):
		// Normalize any complex NaN's.
		number = c.undefined_
	default:
		// Lock onto 0, -1, 1, -i, i, and ∞ if necessary.
		var realPart = c.lockFloat(real(complex_))
		var imaginaryPart = c.lockFloat(imag(complex_))
		number = number_(complex(realPart, imaginaryPart))
	}
	return number
}

func (c *numberClass_) sourceFromFloat(float float64) string {
	var float63 = mat.Float64frombits(mat.Float64bits(float) &^ 1)
	var e = mat.Float64frombits(mat.Float64bits(mat.E) &^ 1)
	var pi = mat.Float64frombits(mat.Float64bits(mat.Pi) &^ 1)
	var tau = mat.Float64frombits(mat.Float64bits(2.0*mat.Pi) &^ 1)
	var phi = mat.Float64frombits(mat.Float64bits(mat.Phi) &^ 1)
	var source string
	switch {
	case float63 == e:
		source = "e"
	case float63 == -e:
		source = "-e"
	case float63 == pi:
		source = "π"
	case float63 == -pi:
		source = "-π"
	case float63 == phi:
		source = "φ"
	case float63 == -phi:
		source = "-φ"
	case float63 == tau:
		source = "τ"
	case float63 == -tau:
		source = "-τ"
	case mat.IsInf(float, 1):
		source = "∞"
	case mat.IsInf(float, -1):
		source = "-∞"
	case mat.IsNaN(float):
		source = "undefined"
	default:
		source = stc.FormatFloat(float63, 'G', 15, 64)
	}
	return source
}

// This constructor creates a new number from the specified polar values.
// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	amplitude_      = "(?:0" + fraction_ + "|" + ordinal_ + "(?:" + fraction_ + ")?(?:" + exponent_ + ")?|" + transcendental_ + ")"
	base10_         = "[0-9]"
	exponent_       = "E(?:" + sign_ + ")?" + ordinal_
	float_          = "(?:" + sign_ + ")?" + amplitude_
	fraction_       = "\\.(?:" + base10_ + ")+"
	imaginary_      = float_ + "i"
	infinity_       = "(?:" + sign_ + ")?(?:infinity|∞)"
	ordinal_        = "[1-9](?:" + base10_ + ")*"
	polar_          = "(" + amplitude_ + ")e\\^~(" + amplitude_ + ")i"
	real_           = float_ + "|0|" + infinity_ + "|" + undefined_
	rectangular_    = "((?:" + sign_ + ")?" + amplitude_ + ")((?:" + sign_ + ")" + amplitude_ + ")i"
	sign_           = "\\+|-"
	transcendental_ = "e|pi|π|tau|τ|phi|φ"
	undefined_      = "undefined"
)

// Instance Structure

type number_ complex128

// Class Structure

type numberClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ NumberLike
	zero_      NumberLike
	one_       NumberLike
	i_         NumberLike
	e_         NumberLike
	pi_        NumberLike
	tau_       NumberLike
	phi_       NumberLike
	minimum_   NumberLike
	maximum_   NumberLike
	infinity_  NumberLike
}

// Class Reference

func numberClass() *numberClass_ {
	return numberClassReference_
}

var numberClassReference_ = &numberClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^(" + polar_ + ")|(" + rectangular_ + ")|(" + imaginary_ + ")|(" + real_ + ")",
	),
	undefined_: number_(complex(mat.NaN(), mat.NaN())),
	zero_:      number_(0),
	one_:       number_(1.0),
	i_:         number_(1.0i),
	e_:         number_(mat.E),
	pi_:        number_(mat.Pi),
	tau_:       number_(2.0 * mat.Pi),
	phi_:       number_(mat.Phi),
	minimum_:   number_(complex(mat.Inf(-1), 0)),
	maximum_:   number_(complex(mat.Inf(1), 0)),
	infinity_:  number_(complex(mat.Inf(0), mat.Inf(0))),
}
