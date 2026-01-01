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

package module_test

import (
	pri "github.com/craterdog/go-essential-primitives/v8"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

// ELEMENT

func TestUnits(t *tes.T) {
	ass.Equal(t, "Degrees", pri.Degrees.String())
	ass.Equal(t, "Radians", pri.Radians.String())
	ass.Equal(t, "Gradians", pri.Gradians.String())
}

func TestZeroAngles(t *tes.T) {
	var v = pri.Angle(0)
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, pri.AngleClass().Zero())

	v = pri.Angle(2.0 * mat.Pi)
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, pri.AngleClass().Zero())

	v = pri.AngleFromSource("~0")
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, pri.AngleClass().Zero())

	v = pri.AngleFromSource("~τ")
	ass.Equal(t, "~τ", v.AsSource())
	ass.Equal(t, v, pri.AngleClass().Tau())
}

func TestPositiveAngles(t *tes.T) {
	var v = pri.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, pri.AngleClass().Pi())

	v = pri.AngleFromSource("~π")
	ass.Equal(t, "~π", v.AsSource())
	ass.Equal(t, v, pri.AngleClass().Pi())
}

func TestNegativeAngles(t *tes.T) {
	var v = pri.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, pri.AngleClass().Pi())

	v = pri.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var class = pri.AngleClass()
	var v0 = class.Zero()
	var v1 = pri.Angle(mat.Pi * 0.25)
	var v2 = pri.Angle(mat.Pi * 0.5)
	var v3 = pri.Angle(mat.Pi * 0.75)
	var v4 = class.Pi()
	var v5 = pri.Angle(mat.Pi * 1.25)
	var v6 = pri.Angle(mat.Pi * 1.5)
	var v7 = pri.Angle(mat.Pi * 1.75)
	var v8 = class.Tau()

	ass.Equal(t, v4, class.Inverse(v0))
	ass.Equal(t, v5, class.Inverse(v1))
	ass.Equal(t, v6, class.Inverse(v2))
	ass.Equal(t, v7, class.Inverse(v3))
	ass.Equal(t, v0, class.Inverse(v4))
	ass.Equal(t, v4, class.Inverse(v8))

	ass.Equal(t, v1, class.Sum(v0, v1))
	ass.Equal(t, v0, class.Difference(v1, v1))
	ass.Equal(t, v3, class.Sum(v1, v2))
	ass.Equal(t, v1, class.Difference(v3, v2))
	ass.Equal(t, v5, class.Sum(v2, v3))
	ass.Equal(t, v2, class.Difference(v5, v3))
	ass.Equal(t, v7, class.Sum(v3, v4))
	ass.Equal(t, v3, class.Difference(v7, v4))
	ass.Equal(t, v1, class.Sum(v8, v1))
	ass.Equal(t, v0, class.Difference(v8, v8))

	ass.Equal(t, v3, class.Scaled(v1, 3.0))
	ass.Equal(t, v0, class.Scaled(v4, 2.0))
	ass.Equal(t, v4, class.Scaled(v4, -1.0))
	ass.Equal(t, v0, class.Scaled(v8, 1.0))

	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v0)))
	ass.Equal(t, v1, class.ArcCosine(class.Cosine(v1)))
	ass.Equal(t, v2, class.ArcCosine(class.Cosine(v2)))
	ass.Equal(t, v3, class.ArcCosine(class.Cosine(v3)))
	ass.Equal(t, v4, class.ArcCosine(class.Cosine(v4)))
	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v8)))

	ass.Equal(t, v0, class.ArcSine(class.Sine(v0)))
	ass.Equal(t, v1, class.ArcSine(class.Sine(v1)))
	ass.Equal(t, v2, class.ArcSine(class.Sine(v2)))
	ass.Equal(t, v6, class.ArcSine(class.Sine(v6)))
	ass.Equal(t, v7, class.ArcSine(class.Sine(v7)))
	ass.Equal(t, v0, class.ArcSine(class.Sine(v8)))

	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v0), class.Sine(v0)))
	ass.Equal(t, v1, class.ArcTangent(class.Cosine(v1), class.Sine(v1)))
	ass.Equal(t, v2, class.ArcTangent(class.Cosine(v2), class.Sine(v2)))
	ass.Equal(t, v3, class.ArcTangent(class.Cosine(v3), class.Sine(v3)))
	ass.Equal(t, v4, class.ArcTangent(class.Cosine(v4), class.Sine(v4)))
	ass.Equal(t, v5, class.ArcTangent(class.Cosine(v5), class.Sine(v5)))
	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v8), class.Sine(v8)))
}

func TestFalseBooleans(t *tes.T) {
	ass.False(t, pri.BooleanClass().False().AsIntrinsic())
	var v = pri.Boolean(false)
	ass.False(t, v.AsIntrinsic())
	v = pri.BooleanFromSource("false")
	ass.Equal(t, "false", v.AsSource())
	ass.Equal(t, v, pri.BooleanClass().False())
}

func TestTrueBooleans(t *tes.T) {
	ass.True(t, pri.BooleanClass().True().AsIntrinsic())
	var v = pri.Boolean(true)
	ass.True(t, v.AsIntrinsic())
	v = pri.BooleanFromSource("true")
	ass.Equal(t, "true", v.AsSource())
	ass.Equal(t, v, pri.BooleanClass().True())
}

func TestBooleansLibrary(t *tes.T) {
	var T = pri.Boolean(true)
	var F = pri.Boolean(false)
	var class = pri.BooleanClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Ior(class.San(T, T), class.San(T, T))
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Ior(class.San(T, F), class.San(F, T))
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Ior(class.San(F, T), class.San(T, F))
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Ior(class.San(F, F), class.San(F, F))
	ass.Equal(t, xor, class.Xor(F, F))
}

var DurationClass = pri.DurationClass()

var zero uint = 0
var one uint = 1

func TestZeroDurations(t *tes.T) {
	var v = pri.Duration(0)
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, zero, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsMilliseconds())
	ass.Equal(t, 0.0, v.AsSeconds())
	ass.Equal(t, 0.0, v.AsMinutes())
	ass.Equal(t, 0.0, v.AsHours())
	ass.Equal(t, 0.0, v.AsDays())
	ass.Equal(t, 0.0, v.AsWeeks())
	ass.Equal(t, 0.0, v.AsMonths())
	ass.Equal(t, 0.0, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, zero, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

func TestStringDurations(t *tes.T) {
	var duration = pri.DurationFromSource("~P1Y2M3DT4H5M6S")
	ass.Equal(t, "~P1Y2M3DT4H5M6S", duration.AsSource())
	duration = pri.DurationFromSource("~P0W")
	ass.Equal(t, "~P0W", duration.AsSource())
}

func TestDurations(t *tes.T) {
	var v = pri.Duration(60000)
	ass.Equal(t, "~PT1M", v.AsSource())
	ass.Equal(t, 60000, v.AsInteger())
	ass.Equal(t, uint(60000), v.AsIntrinsic())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, one, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

var GlyphClass = pri.GlyphClass()

func TestGlyphs(t *tes.T) {
	var v = pri.GlyphFromSource("'''")
	ass.Equal(t, "'''", v.AsSource())

	v = pri.Glyph('a')
	ass.Equal(t, "'a'", v.AsSource())

	v = pri.Glyph('"')
	ass.Equal(t, `'"'`, v.AsSource())

	v = pri.Glyph('😊')
	ass.Equal(t, "'😊'", v.AsSource())

	v = pri.Glyph('界')
	ass.Equal(t, "'界'", v.AsSource())

	v = pri.Glyph('\'')
	ass.Equal(t, "'''", v.AsSource())

	v = pri.Glyph('\\')
	ass.Equal(t, "'\\'", v.AsSource())

	v = pri.Glyph('\n')
	ass.Equal(t, "'\n'", v.AsSource())

	v = pri.Glyph('\t')
	ass.Equal(t, "'\t'", v.AsSource())
}

var MomentClass = pri.MomentClass()

func TestIntegerMoments(t *tes.T) {
	var v = pri.Moment(1238589296789)
	ass.False(t, v.IsNegative())
	ass.Equal(t, 1238589296789, v.AsIntrinsic())
	ass.Equal(t, 1238589296789, v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391204, v.AsDays())
	ass.Equal(t, 2047.9320383416004, v.AsWeeks())
	ass.Equal(t, 470.9919881193849, v.AsMonths())
	ass.Equal(t, 39.24933234328208, v.AsYears())
	ass.Equal(t, uint(789), v.GetMilliseconds())
	ass.Equal(t, uint(56), v.GetSeconds())
	ass.Equal(t, uint(34), v.GetMinutes())
	ass.Equal(t, uint(12), v.GetHours())
	ass.Equal(t, uint(1), v.GetDays())
	ass.Equal(t, uint(14), v.GetWeeks())
	ass.Equal(t, uint(4), v.GetMonths())
	ass.Equal(t, uint(2009), v.GetYears())
}

func TestStringMoments(t *tes.T) {
	var sources = []string{
		"<1-01-01T01:01:01.001>",
		"<1-01-01T01:01:01>",
		"<1-01-01T01:01>",
		"<1-01-01T01>",
		"<1-01-01>",
	}
	for _, source := range sources {
		var moment = pri.MomentFromSource(source)
		ass.Equal(t, source, moment.AsSource())
	}
	var v = pri.MomentFromSource("<-1-02-03T04:05:06.700>")
	ass.True(t, v.IsNegative())
	ass.Equal(t, "<-1-02-03T04:05:06.700>", v.AsSource())
	v = pri.Moment(-62167219200000)
	ass.Equal(t, "<0-01-01>", v.AsSource())
}

func TestMomentsLibrary(t *tes.T) {
	var before = pri.MomentClass().Now()
	var duration = pri.Duration(12345)
	var after = pri.Moment(before.AsInteger() + duration.AsInteger())
	var class = pri.MomentClass()

	ass.Equal(t, duration, class.Duration(before, after))
	ass.Equal(t, duration, class.Duration(after, before))
	ass.Equal(t, after, class.Later(before, duration))
	ass.Equal(t, before, class.Earlier(after, duration))
}

func TestZero(t *tes.T) {
	var v = pri.Number(0 + 0i)
	ass.Equal(t, 0+0i, v.AsIntrinsic())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "0", v.AsSource())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, v, pri.NumberClass().Zero())
}

func TestInfinity(t *tes.T) {
	var v = pri.Number(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.AsIntrinsic())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "∞", v.AsSource())
	ass.Equal(t, mat.Inf(1), v.AsFloat())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
	ass.Equal(t, v, pri.NumberClass().Infinity())
}

func TestUndefined(t *tes.T) {
	var v = pri.Number(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.AsIntrinsic()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsFloat()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositivePureReals(t *tes.T) {
	var v = pri.Number(0.25)
	ass.Equal(t, 0.25+0i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	var integer = 5
	v = pri.NumberFromInteger(integer)
	ass.Equal(t, 5.0, v.AsFloat())
	var float = 5.0
	v = pri.NumberFromFloat(float)
	ass.Equal(t, 5.0, v.AsFloat())
	v = pri.NumberFromSource("1.23456789E+100")
	ass.Equal(t, "1.23456789E+100", v.AsSource())
	v = pri.NumberFromSource("1.23456789E-10")
	ass.Equal(t, "1.23456789E-10", v.AsSource())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = pri.Number(0.25i)
	ass.Equal(t, 0+0.25i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = pri.Number(-0.75)
	ass.Equal(t, -0.75+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = pri.Number(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumberFromPolar(t *tes.T) {
	var v = pri.NumberFromPolar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetAngle())

	v = pri.NumberFromSource("5e^~1i")
	ass.Equal(t, 5.0, v.GetMagnitude())
	ass.Equal(t, 1.0, v.GetAngle())
	ass.Equal(t, "5e^~1i", v.AsPolar())
}

func TestNumberFromSource(t *tes.T) {
	var v = pri.NumberFromSource("1e^~πi")
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, "-1", v.AsSource())
	ass.Equal(t, "1e^~πi", v.AsPolar())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetAngle())

	v = pri.NumberFromSource("-1.2-3.4i")
	ass.Equal(t, "-1.2-3.4i", v.AsSource())
	ass.Equal(t, -1.2, v.GetReal())
	ass.Equal(t, -3.4, v.GetImaginary())

	v = pri.NumberFromSource("-π+τi")
	ass.Equal(t, "-π+τi", v.AsSource())
	ass.Equal(t, -3.141592653589793, v.GetReal())
	ass.Equal(t, 6.283185307179586, v.GetImaginary())

	v = pri.NumberFromSource("undefined")
	ass.Equal(t, "undefined", v.AsSource())
	ass.False(t, v.IsDefined())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("+infinity")
	ass.Equal(t, "+∞", v.AsSource())
	ass.True(t, v.IsMaximum())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("infinity")
	ass.Equal(t, "∞", v.AsSource())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("-infinity")
	ass.Equal(t, "-∞", v.AsSource())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("∞")
	ass.Equal(t, "∞", v.AsSource())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("-∞")
	ass.Equal(t, "-∞", v.AsSource())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = pri.NumberFromSource("+1")
	ass.Equal(t, "1", v.AsSource())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = pri.NumberFromSource("1")
	ass.Equal(t, "1", v.AsSource())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = pri.NumberFromSource("-π")
	ass.Equal(t, "-π", v.AsSource())
	ass.Equal(t, -mat.Pi, v.GetReal())
	ass.Equal(t, mat.Pi, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.True(t, v.IsNegative())

	v = pri.NumberFromSource("+1i")
	ass.Equal(t, "1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = pri.NumberFromSource("1i")
	ass.Equal(t, "1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = pri.NumberFromSource("-1i")
	ass.Equal(t, "-1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, -mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = pri.NumberFromSource("-1.2345678E+90")
	ass.Equal(t, "-1.2345678E+90", v.AsSource())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.2345678e+90, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())

	v = pri.NumberFromSource("-1.2345678E+90i")
	ass.Equal(t, "-1.2345678E+90i", v.AsSource())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.2345678e+90, v.GetImaginary())

	v = pri.NumberFromSource("1.2345678E+90e^~1.2345678E-90i")
	ass.Equal(t, "1.2345678E+90e^~1.2345678E-90i", v.AsPolar())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 1.2345678e+90, v.GetMagnitude())
	ass.Equal(t, 1.2345678e-90, v.GetAngle())
}

func TestNumberLibrary(t *tes.T) {
	var class = pri.NumberClass()
	var zero = class.Zero()
	var i = class.I()
	var minusi = pri.Number(-1i)
	var half = pri.Number(0.5)
	var minushalf = pri.Number(-0.5)
	var one = class.One()
	var minusone = pri.Number(-1)
	var two = pri.Number(2.0)
	var minustwo = pri.Number(-2.0)
	var infinity = class.Infinity()
	var undefined = class.Undefined()

	//	-z
	ass.Equal(t, zero, class.Inverse(zero))
	ass.Equal(t, minushalf, class.Inverse(half))
	ass.Equal(t, minusone, class.Inverse(one))
	ass.Equal(t, minusi, class.Inverse(i))
	ass.Equal(t, infinity, class.Inverse(infinity))
	ass.False(t, class.Inverse(undefined).IsDefined())

	//	z + zero => z
	ass.Equal(t, minusi, class.Sum(minusi, zero))
	ass.Equal(t, minusone, class.Sum(minusone, zero))
	ass.Equal(t, zero, class.Sum(zero, zero))
	ass.Equal(t, one, class.Sum(one, zero))
	ass.Equal(t, i, class.Sum(i, zero))
	ass.Equal(t, infinity, class.Sum(infinity, zero))
	ass.False(t, class.Sum(undefined, zero).IsDefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, class.Sum(minusi, infinity))
	ass.Equal(t, infinity, class.Sum(minusone, infinity))
	ass.Equal(t, infinity, class.Sum(zero, infinity))
	ass.Equal(t, infinity, class.Sum(one, infinity))
	ass.Equal(t, infinity, class.Sum(i, infinity))
	ass.Equal(t, infinity, class.Sum(infinity, infinity))
	ass.False(t, class.Sum(undefined, infinity).IsDefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(minusi, infinity))
	ass.Equal(t, infinity, class.Difference(minusone, infinity))
	ass.Equal(t, infinity, class.Difference(zero, infinity))
	ass.Equal(t, infinity, class.Difference(one, infinity))
	ass.Equal(t, infinity, class.Difference(i, infinity))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, infinity).IsDefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(infinity, minusi))
	ass.Equal(t, infinity, class.Difference(infinity, minusone))
	ass.Equal(t, infinity, class.Difference(infinity, zero))
	ass.Equal(t, infinity, class.Difference(infinity, one))
	ass.Equal(t, infinity, class.Difference(infinity, i))
	ass.False(t, class.Difference(infinity, undefined).IsDefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, class.Difference(minusi, minusi))
	ass.Equal(t, zero, class.Difference(minusone, minusone))
	ass.Equal(t, zero, class.Difference(zero, zero))
	ass.Equal(t, zero, class.Difference(one, one))
	ass.Equal(t, zero, class.Difference(i, i))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, undefined).IsDefined())

	//	z * r
	ass.Equal(t, minusi, class.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, class.Scaled(minusone, 0.5))
	ass.Equal(t, zero, class.Scaled(zero, 5.0))
	ass.Equal(t, half, class.Scaled(one, 0.5))
	ass.Equal(t, i, class.Scaled(i, 1.0))
	ass.Equal(t, infinity, class.Scaled(infinity, 5.0))
	ass.False(t, class.Scaled(undefined, 5.0).IsDefined())

	//	/z
	ass.Equal(t, infinity, class.Reciprocal(zero))
	ass.Equal(t, two, class.Reciprocal(half))
	ass.Equal(t, one, class.Reciprocal(one))
	ass.Equal(t, minushalf, class.Reciprocal(minustwo))
	ass.Equal(t, minusi, class.Reciprocal(i))
	ass.Equal(t, zero, class.Reciprocal(infinity))
	ass.False(t, class.Reciprocal(undefined).IsDefined())

	//	*z
	ass.Equal(t, zero, class.Conjugate(zero))
	ass.Equal(t, one, class.Conjugate(one))
	ass.Equal(t, minusi, class.Conjugate(i))
	ass.Equal(t, i, class.Conjugate(minusi))
	ass.False(t, class.Conjugate(undefined).IsDefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, class.Product(zero, zero))
	ass.Equal(t, zero, class.Product(one, zero))
	ass.Equal(t, zero, class.Product(i, zero))
	ass.False(t, class.Product(infinity, zero).IsDefined())
	ass.False(t, class.Product(undefined, zero).IsDefined())

	//	z * one => z
	ass.Equal(t, zero, class.Product(zero, one))
	ass.Equal(t, one, class.Product(one, one))
	ass.Equal(t, i, class.Product(i, one))
	ass.Equal(t, infinity, class.Product(infinity, one))
	ass.False(t, class.Product(undefined, one).IsDefined())

	//	z * infinity => infinity  {z != zero}
	ass.False(t, class.Product(zero, infinity).IsDefined())
	ass.Equal(t, infinity, class.Product(one, infinity))
	ass.Equal(t, infinity, class.Product(i, infinity))
	ass.Equal(t, infinity, class.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.False(t, class.Quotient(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Quotient(zero, one))
	ass.Equal(t, zero, class.Quotient(zero, i))
	ass.Equal(t, zero, class.Quotient(zero, infinity))
	ass.False(t, class.Quotient(zero, undefined).IsDefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, class.Quotient(one, zero))
	ass.Equal(t, infinity, class.Quotient(i, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.False(t, class.Quotient(undefined, zero).IsDefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, class.Quotient(one, infinity))
	ass.Equal(t, zero, class.Quotient(i, infinity))
	ass.False(t, class.Quotient(infinity, infinity).IsDefined())
	ass.False(t, class.Quotient(undefined, infinity).IsDefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, one))
	ass.Equal(t, infinity, class.Quotient(infinity, i))
	ass.False(t, class.Quotient(infinity, undefined).IsDefined())

	//	y / z
	ass.Equal(t, one, class.Quotient(one, one))
	ass.Equal(t, one, class.Quotient(i, i))
	ass.Equal(t, i, class.Quotient(i, one))
	ass.Equal(t, two, class.Quotient(one, half))
	ass.Equal(t, one, class.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, class.Power(minusi, zero))
	ass.Equal(t, one, class.Power(minusone, zero))
	ass.Equal(t, one, class.Power(zero, zero))
	ass.Equal(t, one, class.Power(one, zero))
	ass.Equal(t, one, class.Power(i, zero))
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.False(t, class.Power(undefined, zero).IsDefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, class.Power(zero, one))
	ass.Equal(t, zero, class.Power(zero, i))
	ass.Equal(t, zero, class.Power(zero, infinity))
	ass.False(t, class.Power(zero, undefined).IsDefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, class.Power(minustwo, infinity))
	ass.Equal(t, one, class.Power(minusi, infinity))
	ass.Equal(t, one, class.Power(minusone, infinity))
	ass.Equal(t, zero, class.Power(minushalf, infinity))
	ass.Equal(t, zero, class.Power(half, infinity))
	ass.Equal(t, one, class.Power(one, infinity))
	ass.Equal(t, one, class.Power(i, infinity))
	ass.Equal(t, infinity, class.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.Equal(t, infinity, class.Power(infinity, one))
	ass.Equal(t, infinity, class.Power(infinity, i))
	ass.Equal(t, infinity, class.Power(infinity, infinity))
	ass.False(t, class.Power(infinity, undefined).IsDefined())

	//	one ^ z => one
	ass.Equal(t, one, class.Power(one, one))
	ass.Equal(t, one, class.Power(one, i))
	ass.Equal(t, one, class.Power(one, minusone))
	ass.Equal(t, one, class.Power(one, minusi))

	//	log(zero, z) => zero
	ass.False(t, class.Logarithm(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(zero, i))
	ass.Equal(t, zero, class.Logarithm(zero, one))
	ass.False(t, class.Logarithm(zero, infinity).IsDefined())
	ass.False(t, class.Logarithm(zero, undefined).IsDefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, class.Logarithm(one, zero))
	ass.False(t, class.Logarithm(one, one).IsDefined())
	ass.Equal(t, infinity, class.Logarithm(one, infinity))
	ass.False(t, class.Logarithm(one, undefined).IsDefined())

	//	log(infinity, z) => zero
	ass.False(t, class.Logarithm(infinity, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(infinity, one))
	ass.False(t, class.Logarithm(infinity, infinity).IsDefined())
	ass.False(t, class.Logarithm(infinity, undefined).IsDefined())
}

func TestZeroPercentages(t *tes.T) {
	var v = pri.Percentage(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestPositivePercentages(t *tes.T) {
	var v = pri.Percentage(25)
	ass.Equal(t, 0.25, v.AsIntrinsic())
	ass.Equal(t, 25.0, v.AsFloat())
}

func TestNegativePercentages(t *tes.T) {
	var v = pri.Percentage(-75)
	ass.Equal(t, -0.75, v.AsIntrinsic())
	ass.Equal(t, -75.0, v.AsFloat())
}

func TestStringPercentages(t *tes.T) {
	var v = pri.PercentageFromSource("1.7%")
	//ass.Equal(t, -1.0, v.AsIntrinsic())
	//ass.Equal(t, -100.0, v.AsFloat())
	ass.Equal(t, "1.7%", v.AsSource())
}

func TestBooleanProbabilities(t *tes.T) {
	var v1 = pri.ProbabilityFromBoolean(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = pri.ProbabilityFromBoolean(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = pri.Probability(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = pri.Probability(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbability(t *tes.T) {
	pri.ProbabilityClass().Random()
}

func TestStringProbabilities(t *tes.T) {
	var v = pri.ProbabilityFromSource("p0")
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "p0", v.AsSource())

	v = pri.ProbabilityFromSource("p0.5")
	ass.Equal(t, 0.5, v.AsIntrinsic())
	ass.Equal(t, 0.5, v.AsFloat())
	ass.Equal(t, "p0.5", v.AsSource())

	v = pri.ProbabilityFromSource("p1")
	ass.Equal(t, 1.0, v.AsIntrinsic())
	ass.Equal(t, 1.0, v.AsFloat())
	ass.Equal(t, "p1", v.AsSource())
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = pri.Probability(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = pri.Probability(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = pri.Probability(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = pri.Probability(0.75)
	var F = pri.Probability(0.25)
	var class = pri.ProbabilityClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Probability(class.San(T, T).AsFloat() + class.San(T, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Probability(class.San(T, F).AsFloat() + class.San(F, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Probability(class.San(F, T).AsFloat() + class.San(T, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Probability(class.San(F, F).AsFloat() + class.San(F, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, F))
}

func TestResource(t *tes.T) {
	var v = pri.Resource("https://craterdog.com/About.html")
	ass.Equal(t, "https://craterdog.com/About.html", v.AsIntrinsic())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = pri.ResourceFromSource("<https://craterdog.com/About.html>")
	ass.Equal(t, "<https://craterdog.com/About.html>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = pri.ResourceFromSource("<mailto:craterdog@google.com>")
	ass.Equal(t, "<mailto:craterdog@google.com>", v.AsSource())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = pri.ResourceFromSource("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = pri.ResourceFromSource("<https://craterdog.com/#Home>")
	ass.Equal(t, "<https://craterdog.com/#Home>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = pri.ResourceFromSource("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestSymbol(t *tes.T) {
	var foobar = []rune("foo-bar")
	var v = pri.Symbol(foobar)
	ass.Equal(t, foobar, v.AsIntrinsic())
}

func TestSymbolFromSource(t *tes.T) {
	var foobar = "$foo-bar"
	var v = pri.SymbolFromSource(foobar)
	ass.Equal(t, foobar, v.AsSource())
}

// STRING

func TestEmptyBinary(t *tes.T) {
	var binary = `'><'`
	var v = pri.BinaryFromSource(binary)
	ass.Equal(t, binary, v.AsSource())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestBinary(t *tes.T) {
	var binary = `'>
    abcd1234
<'`
	var v = pri.BinaryFromSource(binary)
	ass.Equal(t, binary, v.AsSource())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), pri.Binary(v.AsArray()).AsArray())
}

func TestBinaryLibrary(t *tes.T) {
	var b1 = `'>
    abcd
<'`
	var b2 = `'>
    12345678
<'`
	var b3 = `'>
    abcd12345678
<'`
	var v1 = pri.BinaryFromSource(b1)
	var v2 = pri.BinaryFromSource(b2)
	var class = pri.BinaryClass()
	ass.Equal(t, b3, class.Concatenate(v1, v2).AsSource())

	v1 = pri.Binary([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = pri.Binary([]byte{0x03, 0x00, 0x01, 0x02})
	var not = pri.Binary([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = pri.Binary([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = pri.Binary([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = pri.Binary([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = pri.Binary([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = pri.Binary([]byte{0x03, 0x00, 0x01, 0x00, 0x00})

	ass.Equal(t, not, class.Not(v1))
	ass.Equal(t, and, class.And(v1, v2))
	ass.Equal(t, sans, class.San(v1, v2))
	ass.Equal(t, or, class.Ior(v1, v2))
	ass.Equal(t, xor, class.Xor(v1, v2))
	ass.Equal(t, sans2, class.San(v2, v1))
}

func TestBytecode(t *tes.T) {
	var bytecode = `'>
    :abcd:1234
<'`
	var v = pri.BytecodeFromSource(bytecode)
	ass.Equal(t, bytecode, v.AsSource())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), pri.Bytecode(v.AsArray()).AsArray())
}

func TestName(t *tes.T) {
	var v1 = pri.NameFromSource("/bali-nebula/types/abstractions/5String")
	ass.Equal(t, "/bali-nebula/types/abstractions/5String", v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, int(v1.GetSize()))
	ass.Equal(t, pri.Folder("bali-nebula"), v1.GetValue(1))
	ass.Equal(t, pri.Folder("5String"), v1.GetValue(-1))
	var v2 = pri.Name(v1.AsArray())
	ass.Equal(t, v1.AsSource(), v2.AsSource())
	var v3 = pri.NameFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, 1, v1.GetIndex("bali-nebula"))
	ass.Equal(t, "/bali-nebula/types", v3.AsSource())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = pri.NameFromSource("/bali-nebula/types/abstractions")
	var v2 = pri.NameFromSource("/String")
	ass.Equal(
		t,
		"/bali-nebula/types/abstractions/String",
		pri.NameClass().Concatenate(v1, v2).AsSource(),
	)
}

const n0 = `"><"`

const n1 = `">
    abcd本
<"`

const n2 = `">
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

const n3 = `">
    abcd本
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

func TestEmptyNarrative(t *tes.T) {
	var v0 = pri.NarrativeFromSource(n0)
	ass.Equal(t, n0, v0.AsSource())
	ass.True(t, v0.IsEmpty())
	ass.Equal(t, 0, int(v0.GetSize()))
	ass.Equal(t, 0, len(v0.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v1 = pri.NarrativeFromSource(n1)
	ass.Equal(t, n1, v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 1, int(v1.GetSize()))

	var v3 = pri.NarrativeFromSource(n3)
	ass.Equal(t, n3, v3.AsSource())
	ass.False(t, v3.IsEmpty())
	ass.Equal(t, 5, int(v3.GetSize()))

	ass.Equal(t, n3, pri.Narrative(v3.AsArray()).AsSource())
	ass.Equal(t, 5, len(v3.AsArray()))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = pri.NarrativeFromSource(n1)
	var v2 = pri.NarrativeFromSource(n2)
	var v3 = pri.NarrativeClass().Concatenate(v1, v2)
	ass.Equal(t, v1.GetValue(1), v3.GetValue(1))
	ass.Equal(t, v2.GetValue(-1), v3.GetValue(-1))
	ass.Equal(t, n3, v3.AsSource())
}

func TestNonePattern(t *tes.T) {
	var v = pri.PatternClass().None()
	ass.Equal(t, `none`, v.AsSource())

	v = pri.PatternFromSource(`none`)
	ass.Equal(t, `none`, v.AsSource())
	ass.Equal(t, v, pri.PatternClass().None())

	var text = ""
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *tes.T) {
	var v = pri.PatternClass().Any()
	ass.Equal(t, `any`, v.AsSource())

	v = pri.PatternFromSource(`any`)
	ass.Equal(t, `any`, v.AsSource())
	ass.Equal(t, v, pri.PatternClass().Any())

	var text = ""
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *tes.T) {
	var v = pri.PatternFromSource(`"c(.+t)"?`)
	ass.Equal(t, `"c(.+t)"?`, v.AsSource())

	var text = "ct"
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}

func TestEmptyQuote(t *tes.T) {
	var v = pri.Quote([]rune{})
	ass.Equal(t, []rune{}, v.AsIntrinsic())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestQuote(t *tes.T) {
	var v = pri.QuoteFromSource(`"abcd本1234"`)
	ass.Equal(t, `"abcd本1234"`, v.AsSource())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, int(v.GetSize()))
	ass.Equal(t, 'a', rune(v.GetValue(1)))
	ass.Equal(t, '4', rune(v.GetValue(-1)))
	ass.Equal(t, `"d本1"`, pri.QuoteFromSequence(v.GetValues(4, 6)).AsSource())
	ass.Equal(t, 8, v.GetIndex('3'))
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = pri.QuoteFromSource(`"abcd本"`)
	var v2 = pri.QuoteFromSource(`"1234"`)
	ass.Equal(t, `"abcd本1234"`, pri.QuoteClass().Concatenate(v1, v2).AsSource())
}

func TestStringTags(t *tes.T) {
	var size uint
	for size = 8; size < 33; size++ {
		var t1 = pri.TagWithSize(size)
		ass.Equal(t, len(t1.AsSource()), 1+int(mat.Ceil(float64(size)*8.0/5.0)))
		var s1 = t1.AsSource()
		var t2 = pri.TagFromSource(s1)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsSource()
		ass.Equal(t, s1, s2)
		ass.Equal(t, t1.AsArray(), t2.AsArray())
	}
}

func TestVersion(t *tes.T) {
	var v1 = pri.VersionFromSource("v1.2.3")
	ass.Equal(t, "v1.2.3", v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, int(v1.GetSize()))
	ass.Equal(t, uint(1), v1.GetValue(1))
	ass.Equal(t, uint(3), v1.GetValue(-1))
	var v3 = pri.VersionFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, 2, v1.GetIndex(2))
	ass.Equal(t, "v1.2", v3.AsSource())
}

func TestVersionsLibrary(t *tes.T) {
	var v1 = pri.Version([]uint{1})
	var v2 = pri.Version([]uint{2, 3})
	var class = pri.VersionClass()

	var v3 = class.Concatenate(v1, v2)
	ass.Equal(t, []uint{1, 2, 3}, v3.AsIntrinsic())
	ass.False(t, class.IsValidNextVersion(v1, v1))
	ass.Equal(t, "v2", class.GetNextVersion(v1, 0).AsSource())
	ass.Equal(t, "v2", class.GetNextVersion(v1, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 3).AsSource())

	ass.False(t, class.IsValidNextVersion(v2, v2))
	ass.Equal(t, "v3", class.GetNextVersion(v2, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 0).AsSource())
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "v2.3.1", class.GetNextVersion(v2, 3).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 3), v2))

	ass.False(t, class.IsValidNextVersion(v3, v3))
	ass.Equal(t, "v2", class.GetNextVersion(v3, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "v1.3", class.GetNextVersion(v3, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 0).AsSource())
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 3).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "v1.2.3.1", class.GetNextVersion(v3, 4).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 4)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 4), v3))
}
