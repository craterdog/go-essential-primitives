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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-essential-primitives/wiki
*/
package module

import (
	ele "github.com/craterdog/go-essential-primitives/v8/elements"
	str "github.com/craterdog/go-essential-primitives/v8/strings"
	uri "net/url"
)

// TYPE ALIASES

// Elements

type (
	Units = ele.Units
)

const (
	Degrees  = ele.Degrees
	Radians  = ele.Radians
	Gradians = ele.Gradians
)

type (
	AngleClassLike       = ele.AngleClassLike
	BooleanClassLike     = ele.BooleanClassLike
	DurationClassLike    = ele.DurationClassLike
	GlyphClassLike       = ele.GlyphClassLike
	MomentClassLike      = ele.MomentClassLike
	NumberClassLike      = ele.NumberClassLike
	PercentageClassLike  = ele.PercentageClassLike
	ProbabilityClassLike = ele.ProbabilityClassLike
	ResourceClassLike    = ele.ResourceClassLike
	SymbolClassLike      = ele.SymbolClassLike
)

type (
	AngleLike       = ele.AngleLike
	BooleanLike     = ele.BooleanLike
	DurationLike    = ele.DurationLike
	GlyphLike       = ele.GlyphLike
	MomentLike      = ele.MomentLike
	NumberLike      = ele.NumberLike
	PercentageLike  = ele.PercentageLike
	ProbabilityLike = ele.ProbabilityLike
	ResourceLike    = ele.ResourceLike
	SymbolLike      = ele.SymbolLike
)

type (
	Continuous = ele.Continuous
	Discrete   = ele.Discrete
	Factored   = ele.Factored
	Polarized  = ele.Polarized
	Temporal   = ele.Temporal
)

// Strings

type (
	Folder = str.Folder
)

type (
	BinaryClassLike    = str.BinaryClassLike
	BytecodeClassLike  = str.BytecodeClassLike
	NameClassLike      = str.NameClassLike
	NarrativeClassLike = str.NarrativeClassLike
	PatternClassLike   = str.PatternClassLike
	QuoteClassLike     = str.QuoteClassLike
	TagClassLike       = str.TagClassLike
	VersionClassLike   = str.VersionClassLike
)

type (
	BinaryLike    = str.BinaryLike
	BytecodeLike  = str.BytecodeLike
	NameLike      = str.NameLike
	NarrativeLike = str.NarrativeLike
	PatternLike   = str.PatternLike
	QuoteLike     = str.QuoteLike
	TagLike       = str.TagLike
	VersionLike   = str.VersionLike
)

type (
	Accessible[V any] = str.Accessible[V]
	Searchable[V any] = str.Searchable[V]
	Sequential[V any] = str.Sequential[V]
	Ordered[V any]    = str.Ordered[V]
)

// CLASS ACCESSORS

// Elements

func AngleClass() AngleClassLike {
	return ele.AngleClass()
}

func Angle(
	radians float64,
) AngleLike {
	return AngleClass().Angle(
		radians,
	)
}

func AngleFromSource(
	source string,
) AngleLike {
	return AngleClass().AngleFromSource(
		source,
	)
}

func BooleanClass() BooleanClassLike {
	return ele.BooleanClass()
}

func Boolean(
	boolean bool,
) BooleanLike {
	return BooleanClass().Boolean(
		boolean,
	)
}

func BooleanFromSource(
	source string,
) BooleanLike {
	return BooleanClass().BooleanFromSource(
		source,
	)
}

func DurationClass() DurationClassLike {
	return ele.DurationClass()
}

func Duration(
	milliseconds int,
) DurationLike {
	return DurationClass().Duration(
		milliseconds,
	)
}

func DurationFromSource(
	source string,
) DurationLike {
	return DurationClass().DurationFromSource(
		source,
	)
}

func GlyphClass() GlyphClassLike {
	return ele.GlyphClass()
}

func Glyph(
	rune_ rune,
) GlyphLike {
	return GlyphClass().Glyph(
		rune_,
	)
}

func GlyphFromInteger(
	integer int,
) GlyphLike {
	return GlyphClass().GlyphFromInteger(
		integer,
	)
}

func GlyphFromSource(
	source string,
) GlyphLike {
	return GlyphClass().GlyphFromSource(
		source,
	)
}

func MomentClass() MomentClassLike {
	return ele.MomentClass()
}

func Moment(
	milliseconds int,
) MomentLike {
	return MomentClass().Moment(
		milliseconds,
	)
}

func MomentFromSource(
	source string,
) MomentLike {
	return MomentClass().MomentFromSource(
		source,
	)
}

func NumberClass() NumberClassLike {
	return ele.NumberClass()
}

func Number(
	complex_ complex128,
) NumberLike {
	return NumberClass().Number(
		complex_,
	)
}

func NumberFromPolar(
	magnitude float64,
	angle float64,
) NumberLike {
	return NumberClass().NumberFromPolar(
		magnitude,
		angle,
	)
}

func NumberFromRectangular(
	real_ float64,
	imaginary float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(
		real_,
		imaginary,
	)
}

func NumberFromInteger(
	integer int,
) NumberLike {
	return NumberClass().NumberFromInteger(
		integer,
	)
}

func NumberFromFloat(
	float float64,
) NumberLike {
	return NumberClass().NumberFromFloat(
		float,
	)
}

func NumberFromSource(
	source string,
) NumberLike {
	return NumberClass().NumberFromSource(
		source,
	)
}

func PercentageClass() PercentageClassLike {
	return ele.PercentageClass()
}

func Percentage(
	float float64,
) PercentageLike {
	return PercentageClass().Percentage(
		float,
	)
}

func PercentageFromInteger(
	integer int,
) PercentageLike {
	return PercentageClass().PercentageFromInteger(
		integer,
	)
}

func PercentageFromSource(
	source string,
) PercentageLike {
	return PercentageClass().PercentageFromSource(
		source,
	)
}

func ProbabilityClass() ProbabilityClassLike {
	return ele.ProbabilityClass()
}

func Probability(
	float float64,
) ProbabilityLike {
	return ProbabilityClass().Probability(
		float,
	)
}

func ProbabilityFromBoolean(
	boolean bool,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromBoolean(
		boolean,
	)
}

func ProbabilityFromSource(
	source string,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromSource(
		source,
	)
}

func ResourceClass() ResourceClassLike {
	return ele.ResourceClass()
}

func Resource(
	uri string,
) ResourceLike {
	return ResourceClass().Resource(
		uri,
	)
}

func ResourceFromSource(
	source string,
) ResourceLike {
	return ResourceClass().ResourceFromSource(
		source,
	)
}

func ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return ResourceClass().ResourceFromUri(
		url,
	)
}

func SymbolClass() SymbolClassLike {
	return ele.SymbolClass()
}

func Symbol(
	identifier string,
) SymbolLike {
	return SymbolClass().Symbol(
		identifier,
	)
}

func SymbolFromSource(
	source string,
) SymbolLike {
	return SymbolClass().SymbolFromSource(
		source,
	)
}

// Strings

func BinaryClass() BinaryClassLike {
	return str.BinaryClass()
}

func Binary(
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence str.Sequential[byte],
) BinaryLike {
	return BinaryClass().BinaryFromSequence(
		sequence,
	)
}

func BinaryFromSource(
	source string,
) BinaryLike {
	return BinaryClass().BinaryFromSource(
		source,
	)
}

func BytecodeClass() BytecodeClassLike {
	return str.BytecodeClass()
}

func Bytecode(
	instructions []uint16,
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromSequence(
	sequence str.Sequential[uint16],
) BytecodeLike {
	return BytecodeClass().BytecodeFromSequence(
		sequence,
	)
}

func BytecodeFromSource(
	source string,
) BytecodeLike {
	return BytecodeClass().BytecodeFromSource(
		source,
	)
}

func NameClass() NameClassLike {
	return str.NameClass()
}

func Name(
	folders []str.Folder,
) NameLike {
	return NameClass().Name(
		folders,
	)
}

func NameFromSequence(
	sequence str.Sequential[str.Folder],
) NameLike {
	return NameClass().NameFromSequence(
		sequence,
	)
}

func NameFromSource(
	source string,
) NameLike {
	return NameClass().NameFromSource(
		source,
	)
}

func NarrativeClass() NarrativeClassLike {
	return str.NarrativeClass()
}

func Narrative(
	lines []string,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence str.Sequential[string],
) NarrativeLike {
	return NarrativeClass().NarrativeFromSequence(
		sequence,
	)
}

func NarrativeFromSource(
	source string,
) NarrativeLike {
	return NarrativeClass().NarrativeFromSource(
		source,
	)
}

func PatternClass() PatternClassLike {
	return str.PatternClass()
}

func Pattern(
	characters []rune,
) PatternLike {
	return PatternClass().Pattern(
		characters,
	)
}

func PatternFromSequence(
	sequence str.Sequential[rune],
) PatternLike {
	return PatternClass().PatternFromSequence(
		sequence,
	)
}

func PatternFromSource(
	source string,
) PatternLike {
	return PatternClass().PatternFromSource(
		source,
	)
}

func QuoteClass() QuoteClassLike {
	return str.QuoteClass()
}

func Quote(
	characters []rune,
) QuoteLike {
	return QuoteClass().Quote(
		characters,
	)
}

func QuoteFromSequence(
	sequence str.Sequential[rune],
) QuoteLike {
	return QuoteClass().QuoteFromSequence(
		sequence,
	)
}

func QuoteFromSource(
	source string,
) QuoteLike {
	return QuoteClass().QuoteFromSource(
		source,
	)
}

func TagClass() TagClassLike {
	return str.TagClass()
}

func Tag(
	bytes []byte,
) TagLike {
	return TagClass().Tag(
		bytes,
	)
}

func TagWithSize(
	size uint,
) TagLike {
	return TagClass().TagWithSize(
		size,
	)
}

func TagFromSequence(
	sequence str.Sequential[byte],
) TagLike {
	return TagClass().TagFromSequence(
		sequence,
	)
}

func TagFromSource(
	source string,
) TagLike {
	return TagClass().TagFromSource(
		source,
	)
}

func VersionClass() VersionClassLike {
	return str.VersionClass()
}

func Version(
	ordinals []uint,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence str.Sequential[uint],
) VersionLike {
	return VersionClass().VersionFromSequence(
		sequence,
	)
}

func VersionFromSource(
	source string,
) VersionLike {
	return VersionClass().VersionFromSource(
		source,
	)
}

// GLOBAL FUNCTIONS
