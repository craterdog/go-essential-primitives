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
)

// CLASS INTERFACE

// Access Function

func IdentifierClass() IdentifierClassLike {
	return identifierClass()
}

// Constructor Methods

func (c *identifierClass_) Identifier(
	characters []rune,
) IdentifierLike {
	return c.IdentifierFromSource(string(characters))
}

func (c *identifierClass_) IdentifierFromSequence(
	sequence Sequential[rune],
) IdentifierLike {
	return c.Identifier(sequence.AsArray())
}

func (c *identifierClass_) IdentifierFromSource(
	source string,
) IdentifierLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the identifier constructor method: %s",
			source,
		)
		panic(message)
	}
	return identifier_(source)
}

// Constant Methods

func (c *identifierClass_) Undefined() IdentifierLike {
	return c.undefined_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v identifier_) GetClass() IdentifierClassLike {
	return identifierClass()
}

func (v identifier_) AsIntrinsic() []rune {
	return []rune(v)
}

func (v identifier_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Accessible[rune] Methods

func (v identifier_) GetValue(
	index int,
) rune {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goIndex = uti.RelativeToCardinal(index, size)
	return characters[goIndex]
}

func (v identifier_) GetValues(
	first int,
	last int,
) Sequential[rune] {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return identifierClass().Identifier(characters[goFirst : goLast+1])
}

func (v identifier_) GetIndex(
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

// Ordered[IdentifierLike] Methods

func (v identifier_) IsBefore(
	value IdentifierLike,
) bool {
	return sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) < 0
}

// Searchable[rune] Methods

func (v identifier_) ContainsValue(
	value rune,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v identifier_) ContainsAny(
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

func (v identifier_) ContainsAll(
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

func (v identifier_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v identifier_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v identifier_) AsArray() []rune {
	return v.AsIntrinsic()
}

func (v identifier_) GetIterator() uti.Ratcheted[rune] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v identifier_) String() string {
	return v.AsSource()
}

// Private Methods

// Instance Structure

type identifier_ string

// Class Structure

type identifierClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ IdentifierLike
}

// Class Reference

func identifierClass() *identifierClass_ {
	return identifierClassReference_
}

var identifierClassReference_ = &identifierClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^((?:" + letter_ + ")((-)?(?:" + letter_ + "|" + digit_ + "))*)",
	),
	undefined_: identifier_(""),
}
