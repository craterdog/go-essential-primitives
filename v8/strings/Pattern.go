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

package strings

import (
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
	reg "regexp"
	sli "slices"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func PatternClass() PatternClassLike {
	return patternClass()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	characters []rune,
) PatternLike {
	var pattern = string(characters)
	reg.MustCompile(pattern)
	var source string
	switch pattern {
	case `^none$`:
		source = `none`
	case `.*`:
		source = `any`
	default:
		source = stc.Quote(pattern) + "?"
	}
	return pattern_(source)
}

func (c *patternClass_) PatternFromSequence(
	sequence Sequential[rune],
) PatternLike {
	return c.Pattern(sequence.AsArray())
}

func (c *patternClass_) PatternFromSource(
	source string,
) PatternLike {

	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the pattern constructor method: %s",
			source,
		)
		panic(message)
	}
	return pattern_(source)
}

// Constant Methods

func (c *patternClass_) None() PatternLike {
	return c.none_
}

func (c *patternClass_) Any() PatternLike {
	return c.any_
}

// Function Methods

func (c *patternClass_) Concatenate(
	first PatternLike,
	second PatternLike,
) PatternLike {
	return c.Pattern(uti.CombineArrays(first.AsIntrinsic(), second.AsIntrinsic()))
}

// INSTANCE INTERFACE

// Principal Methods

func (v pattern_) GetClass() PatternClassLike {
	return patternClass()
}

func (v pattern_) AsIntrinsic() []rune {
	var pattern = string(v)
	switch pattern {
	case "none":
		return []rune(`^none$`)
	case "any":
		return []rune(`.*`)
	default:
		pattern = pattern[:len(pattern)-1]     // Strip off the trailing "?".
		var unquoted, _ = stc.Unquote(pattern) // Strip off the double quotes.
		return []rune(unquoted)
	}
}

func (v pattern_) AsSource() string {
	return string(v)
}

func (v pattern_) AsRegexp() *reg.Regexp {
	var regexp = string(v.AsIntrinsic())
	return reg.MustCompile(regexp)
}

func (v pattern_) MatchesText(
	text string,
) bool {
	var regexp = string(v.AsIntrinsic())
	var matcher = reg.MustCompile(regexp)
	return matcher.MatchString(text)
}

func (v pattern_) GetMatches(
	text string,
) []string {
	var regexp = string(v.AsIntrinsic())
	var matcher = reg.MustCompile(regexp)
	return matcher.FindStringSubmatch(text)
}

// Attribute Methods

// Searchable[rune] Methods

func (v pattern_) ContainsValue(
	value rune,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v pattern_) ContainsAny(
	values Sequential[rune],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This set contains at least one of the values.
			return true
		}
	}
	// This set does not contain any of the values.
	return false
}

func (v pattern_) ContainsAll(
	values Sequential[rune],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This set is missing at least one of the values.
			return false
		}
	}
	// This set does contains all of the values.
	return true
}

// Sequential[rune] Methods

func (v pattern_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v pattern_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v pattern_) AsArray() []rune {
	return v.AsIntrinsic()
}

func (v pattern_) GetIterator() uti.IteratorLike[rune] {
	return uti.Iterator(v.AsIntrinsic())
}

// Accessible[rune] Methods

func (v pattern_) GetValue(
	index int,
) rune {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goIndex = uti.RelativeToCardinal(index, size)
	return characters[goIndex]
}

func (v pattern_) GetValues(
	first int,
	last int,
) Sequential[rune] {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return patternClass().Pattern(characters[goFirst : goLast+1])
}

func (v pattern_) GetIndex(
	value rune,
) int {
	var index int
	var iterator = v.GetIterator()
	for iterator.HasNext() {
		index++
		var candidate = iterator.GetNext()
		if candidate == value {
			// Found the value.
			return index
		}
	}
	// The value was not found.
	return 0
}

// PROTECTED INTERFACE

func (v pattern_) String() string {
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
	regex_ = "\"((?:" + character_ + ")+)\"\\\\?"
)

// Instance Structure

type pattern_ string // This type must support the "comparable" type contraint.

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
	none_    PatternLike
	any_     PatternLike
}

// Class Reference

func patternClass() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^" + regex_ + "|any|none"),
	none_:    pattern_(`none`),
	any_:     pattern_(`any`),
}
