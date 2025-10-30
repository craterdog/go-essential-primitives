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
	bin "encoding/binary"
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
	reg "regexp"
	sli "slices"
)

// CLASS INTERFACE

// Access Function

func TagClass() TagClassLike {
	return tagClass()
}

// Constructor Methods

func (c *tagClass_) Tag(
	bytes []byte,
) TagLike {
	c.validateSize(uti.ArraySize(bytes))
	var encoded = uti.Base32Encode(bytes)
	return tag_("#" + encoded)
}

func (c *tagClass_) TagWithSize(
	size uint,
) TagLike {
	c.validateSize(size)
	var bytes = uti.RandomBytes(size)
	return c.Tag(bytes)
}

func (c *tagClass_) TagFromSequence(
	sequence Sequential[byte],
) TagLike {
	var bytes = sequence.AsArray()
	c.validateSize(uti.ArraySize(bytes))
	return c.Tag(bytes)
}

func (c *tagClass_) TagFromSource(
	source string,
) TagLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the tag constructor method: %s",
			source,
		)
		panic(message)
	}
	return tag_(source)
}

// Constant Methods

// Function Methods

func (c *tagClass_) Concatenate(
	first TagLike,
	second TagLike,
) TagLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Tag(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v tag_) GetClass() TagClassLike {
	return tagClass()
}

func (v tag_) AsIntrinsic() []byte {
	var base32 = string(v[1:]) // Strip off the leading "#".
	var bytes = uti.Base32Decode(base32)
	return bytes
}

func (v tag_) AsSource() string {
	return string(v)
}

func (v tag_) GetHash() uint64 {
	return bin.BigEndian.Uint64(v.AsIntrinsic())
}

// Attribute Methods

// Accessible[byte] Methods

func (v tag_) GetValue(
	index int,
) byte {
	var bytes = v.AsIntrinsic()
	var size = uti.ArraySize(bytes)
	var goIndex = uti.RelativeToCardinal(index, size)
	return bytes[goIndex]
}

func (v tag_) GetValues(
	first int,
	last int,
) Sequential[byte] {
	var bytes = v.AsIntrinsic()
	var size = uti.ArraySize(bytes)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return tagClass().Tag(bytes[goFirst : goLast+1])
}

func (v tag_) GetIndex(
	value byte,
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

// Searchable[byte] Methods

func (v tag_) ContainsValue(
	value byte,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v tag_) ContainsAny(
	values Sequential[byte],
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

func (v tag_) ContainsAll(
	values Sequential[byte],
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

// Sequential[byte] Methods

func (v tag_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v tag_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v tag_) AsArray() []byte {
	return v.AsIntrinsic()
}

func (v tag_) GetIterator() uti.Ratcheted[byte] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v tag_) String() string {
	return v.AsSource()
}

// Private Methods

func (c *tagClass_) validateSize(
	size uint,
) {
	if size < 8 {
		var message = fmt.Sprintf(
			"A tag must be at least eight bytes long: %v",
			size,
		)
		panic(message)
	}
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	base32_ = base10_ + "|[A-DF-HJ-NP-TV-Z]"
)

// Instance Structure

type tag_ string // This type must support the "comparable" type contraint.

// Class Structure

type tagClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func tagClass() *tagClass_ {
	return tagClassReference_
}

var tagClassReference_ = &tagClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^#((?:" + base32_ + ")+)"),
}
