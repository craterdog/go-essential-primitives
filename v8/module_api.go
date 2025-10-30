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
	seq "github.com/craterdog/go-essential-primitives/v8/sequences"
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
)

type (
	Continuous = ele.Continuous
	Discrete   = ele.Discrete
	Factored   = ele.Factored
	Polarized  = ele.Polarized
	Temporal   = ele.Temporal
)

// Sequences

type (
	Folder = seq.Folder
)

type (
	BinaryClassLike    = seq.BinaryClassLike
	BytecodeClassLike  = seq.BytecodeClassLike
	NameClassLike      = seq.NameClassLike
	NarrativeClassLike = seq.NarrativeClassLike
	PatternClassLike   = seq.PatternClassLike
	QuoteClassLike     = seq.QuoteClassLike
	SymbolClassLike    = seq.SymbolClassLike
	TagClassLike       = seq.TagClassLike
	VersionClassLike   = seq.VersionClassLike
)

type (
	BinaryLike    = seq.BinaryLike
	BytecodeLike  = seq.BytecodeLike
	NameLike      = seq.NameLike
	NarrativeLike = seq.NarrativeLike
	PatternLike   = seq.PatternLike
	QuoteLike     = seq.QuoteLike
	SymbolLike    = seq.SymbolLike
	TagLike       = seq.TagLike
	VersionLike   = seq.VersionLike
)

type (
	Accessible[V any] = seq.Accessible[V]
	Searchable[V any] = seq.Searchable[V]
	Sequential[V any] = seq.Sequential[V]
	Ordered[V any]    = seq.Ordered[V]
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
	milliseconds uint,
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

// Sequences

func BinaryClass() BinaryClassLike {
	return seq.BinaryClass()
}

func Binary(
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence seq.Sequential[byte],
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
	return seq.BytecodeClass()
}

func Bytecode(
	instructions []uint16,
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromSequence(
	sequence seq.Sequential[uint16],
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
	return seq.NameClass()
}

func Name(
	folders []seq.Folder,
) NameLike {
	return NameClass().Name(
		folders,
	)
}

func NameFromSequence(
	sequence seq.Sequential[seq.Folder],
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
	return seq.NarrativeClass()
}

func Narrative(
	lines []string,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence seq.Sequential[string],
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
	return seq.PatternClass()
}

func Pattern(
	characters []rune,
) PatternLike {
	return PatternClass().Pattern(
		characters,
	)
}

func PatternFromSequence(
	sequence seq.Sequential[rune],
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
	return seq.QuoteClass()
}

func Quote(
	characters []rune,
) QuoteLike {
	return QuoteClass().Quote(
		characters,
	)
}

func QuoteFromSequence(
	sequence seq.Sequential[rune],
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

func SymbolClass() SymbolClassLike {
	return seq.SymbolClass()
}

func Symbol(
	identifier []rune,
) SymbolLike {
	return SymbolClass().Symbol(
		identifier,
	)
}

func SymbolFromSequence(
	sequence seq.Sequential[rune],
) SymbolLike {
	return SymbolClass().SymbolFromSequence(
		sequence,
	)
}

func SymbolFromSource(
	source string,
) SymbolLike {
	return SymbolClass().SymbolFromSource(
		source,
	)
}

func TagClass() TagClassLike {
	return seq.TagClass()
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
	sequence seq.Sequential[byte],
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
	return seq.VersionClass()
}

func Version(
	ordinals []uint,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence seq.Sequential[uint],
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
