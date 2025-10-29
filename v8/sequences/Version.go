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

package sequences

import (
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
	reg "regexp"
	sli "slices"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func VersionClass() VersionClassLike {
	return versionClass()
}

// Constructor Methods

func (c *versionClass_) Version(
	ordinals []uint,
) VersionLike {
	var index = 0
	var source = "v" + stc.Itoa(int(ordinals[index]))
	for index++; index < len(ordinals); index++ {
		source += "." + stc.Itoa(int(ordinals[index]))
	}
	return version_(source)
}

func (c *versionClass_) VersionFromSequence(
	sequence Sequential[uint],
) VersionLike {
	return c.Version(sequence.AsArray())
}

func (c *versionClass_) VersionFromSource(
	source string,
) VersionLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the version constructor method: %s",
			source,
		)
		panic(message)
	}
	return version_(source)
}

// Constant Methods

// Function Methods

func (c *versionClass_) IsValidNextVersion(
	current VersionLike,
	next VersionLike,
) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var currentIterator = uti.Iterator(current.AsArray())
	var nextIterator = uti.Iterator(next.AsArray())
	for currentIterator.HasNext() && nextIterator.HasNext() {
		var currentLevel = currentIterator.GetNext()
		var nextLevel = nextIterator.GetNext()
		if currentLevel == nextLevel {
			// So far the level values are the same.
			continue
		}
		// The last level for the next version must be one more.
		return !nextIterator.HasNext() && nextLevel == currentLevel+1
	}
	// The last level for the next version must be one.
	return nextIterator.HasNext() && nextIterator.GetNext() == 1
}

func (c *versionClass_) GetNextVersion(
	current VersionLike,
	level uint,
) VersionLike {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = uti.ArraySize(ordinals)
	switch {
	case level == 0:
		level = size // Normalize the level to the current size.
	case level < size:
		// The next version will require fewer levels.
		ordinals = ordinals[:level]
	case level > size:
		// The next version will require another level.
		size++
		level = size // Normalize the level to the new size.
		ordinals = append(ordinals, 0)
	}

	// Increment the specified version level.
	var index = level - 1 // Convert to zero based indexing.
	ordinals[index]++

	var version = c.Version(ordinals)
	return version
}

func (c *versionClass_) Concatenate(
	first VersionLike,
	second VersionLike,
) VersionLike {
	var firstOrdinals = first.AsArray()
	var secondOrdinals = second.AsArray()
	var allOrdinals = make(
		[]uint,
		len(firstOrdinals)+len(secondOrdinals),
	)
	copy(allOrdinals, firstOrdinals)
	copy(allOrdinals[len(firstOrdinals):], secondOrdinals)
	return c.Version(allOrdinals)
}

// INSTANCE INTERFACE

// Principal Methods

func (v version_) GetClass() VersionClassLike {
	return versionClass()
}

func (v version_) AsIntrinsic() []uint {
	var version = string(v[1:]) // Strip off the leading "v".
	var levels = sts.Split(version, ".")
	var ordinals = make([]uint, len(levels))
	for index, level := range levels {
		var ordinal, _ = stc.ParseUint(level, 10, 64)
		ordinals[index] = uint(ordinal)
	}
	return ordinals
}

func (v version_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Spectral Methods

func (v version_) IsBefore(
	value VersionLike,
) bool {
	return sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) < 0
}

// Searchable[uint] Methods

func (v version_) ContainsValue(
	value uint,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v version_) ContainsAny(
	values Sequential[uint],
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

func (v version_) ContainsAll(
	values Sequential[uint],
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

// Sequential[uint] Methods

func (v version_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v version_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v version_) AsArray() []uint {
	return v.AsIntrinsic()
}

func (v version_) GetIterator() uti.Ratcheted[uint] {
	return uti.Iterator(v.AsIntrinsic())
}

// Accessible[uint] Methods

func (v version_) GetValue(
	index int,
) uint {
	var ordinals = v.AsIntrinsic()
	var size = uti.ArraySize(ordinals)
	var goIndex = uti.RelativeToCardinal(index, size)
	return ordinals[goIndex]
}

func (v version_) GetValues(
	first int,
	last int,
) Sequential[uint] {
	var ordinals = v.AsIntrinsic()
	var size = uti.ArraySize(ordinals)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return versionClass().Version(ordinals[goFirst : goLast+1])
}

func (v version_) GetIndex(
	value uint,
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

func (v version_) String() string {
	return v.AsSource()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each version to lessen the chance of a version collision with other private Go
// class constants in this package.
const (
	ordinal_ = "[1-9](?:" + base10_ + ")*"
)

// Instance Structure

type version_ string // This type must support the "comparable" type contraint.

// Class Structure

type versionClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func versionClass() *versionClass_ {
	return versionClassReference_
}

var versionClassReference_ = &versionClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^v(" + ordinal_ + "(?:\\." + ordinal_ + ")*)",
	),
}
