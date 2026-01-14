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

package sequences

import (
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
	reg "regexp"
	sli "slices"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NameClass() NameClassLike {
	return nameClass()
}

// Constructor Methods

func (c *nameClass_) Name(
	segments []string,
) NameLike {
	var source = "/"
	source += sts.Join(segments, "/")
	return name_(source)
}

func (c *nameClass_) NameFromSequence(
	sequence Sequential[string],
) NameLike {
	return c.Name(sequence.AsArray())
}

func (c *nameClass_) NameFromSource(
	source string,
) NameLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the name constructor method: %s",
			source,
		)
		panic(message)
	}
	return name_(source)
}

// Constant Methods

// Function Methods

func (c *nameClass_) Concatenate(
	first NameLike,
	second NameLike,
) NameLike {
	var firstSegments = first.AsIntrinsic()
	var secondSegments = second.AsIntrinsic()
	var allSegments = make(
		[]string,
		len(firstSegments)+len(secondSegments),
	)
	copy(allSegments, firstSegments)
	copy(allSegments[len(firstSegments):], secondSegments)
	return c.Name(allSegments)
}

// INSTANCE INTERFACE

// Principal Methods

func (v name_) GetClass() NameClassLike {
	return nameClass()
}

func (v name_) AsIntrinsic() []string {
	var name = string(v)
	var segments = sts.Split(name, "/") // Extract the segments.
	segments = segments[1:]             // Ignore the empty segment.
	return segments
}

func (v name_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Accessible[string] Methods

func (v name_) GetValue(
	index int,
) string {
	var segments = v.AsIntrinsic()
	var size = uti.ArraySize(segments)
	var goIndex = uti.RelativeToCardinal(index, size)
	return segments[goIndex]
}

func (v name_) GetValues(
	first int,
	last int,
) Sequential[string] {
	var segments = v.AsIntrinsic()
	var size = uti.ArraySize(segments)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return nameClass().Name(segments[goFirst : goLast+1])
}

func (v name_) GetIndex(
	value string,
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

// Ordered[NameLike] Methods

func (v name_) IsBefore(
	value NameLike,
) bool {
	return sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) < 0
}

// Searchable[string] Methods

func (v name_) ContainsValue(
	value string,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v name_) ContainsAny(
	values Sequential[string],
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

func (v name_) ContainsAll(
	values Sequential[string],
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

// Sequential[string] Methods

func (v name_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v name_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v name_) AsArray() []string {
	return v.AsIntrinsic()
}

func (v name_) GetIterator() uti.Ratcheted[string] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v name_) String() string {
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
	digit_  = "\\p{Nd}"
	letter_ = lower_ + "|" + upper_
	lower_  = "\\p{Ll}"
	upper_  = "\\p{Lu}"
)

// Instance Structure

type name_ string // This type must support the "comparable" type contraint.

// Class Structure

type nameClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func nameClass() *nameClass_ {
	return nameClassReference_
}

var nameClassReference_ = &nameClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^(?:/(?:" + letter_ + "|" + digit_ + ")((-|\\.)?(?:" + letter_ + "|" + digit_ + "))+)+",
	),
}
