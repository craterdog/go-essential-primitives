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

func SymbolClass() SymbolClassLike {
	return symbolClass()
}

// Constructor Methods

func (c *symbolClass_) Symbol(
	identifier []rune,
) SymbolLike {
	return c.SymbolFromSource("$" + string(identifier))
}

func (c *symbolClass_) SymbolFromSequence(
	sequence Sequential[rune],
) SymbolLike {
	return c.Symbol(sequence.AsArray())
}

func (c *symbolClass_) SymbolFromSource(
	source string,
) SymbolLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the symbol constructor method: %s",
			source,
		)
		panic(message)
	}
	return symbol_(matches[1]) // Strip off the leading "$".
}

// Constant Methods

func (c *symbolClass_) Undefined() SymbolLike {
	return c.undefined_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v symbol_) GetClass() SymbolClassLike {
	return symbolClass()
}

func (v symbol_) AsIntrinsic() []rune {
	return []rune(v)
}

func (v symbol_) AsSource() string {
	return "$" + string(v)
}

// Attribute Methods

// Accessible[rune] Methods

func (v symbol_) GetValue(
	index int,
) rune {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goIndex = uti.RelativeToCardinal(index, size)
	return characters[goIndex]
}

func (v symbol_) GetValues(
	first int,
	last int,
) Sequential[rune] {
	var characters = v.AsIntrinsic()
	var size = uti.ArraySize(characters)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return symbolClass().Symbol(characters[goFirst : goLast+1])
}

func (v symbol_) GetIndex(
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

// Ordered[SymbolLike] Methods

func (v symbol_) IsBefore(
	value SymbolLike,
) bool {
	return sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) < 0
}

// Searchable[rune] Methods

func (v symbol_) ContainsValue(
	value rune,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v symbol_) ContainsAny(
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

func (v symbol_) ContainsAll(
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

func (v symbol_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v symbol_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v symbol_) AsArray() []rune {
	return v.AsIntrinsic()
}

func (v symbol_) GetIterator() uti.Ratcheted[rune] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v symbol_) String() string {
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
	identifier_ = "(?:" + letter_ + ")(?:" + letter_ + "|" + digit_ + "|-)*"
)

// Instance Structure

type symbol_ string

// Class Structure

type symbolClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ SymbolLike
}

// Class Reference

func symbolClass() *symbolClass_ {
	return symbolClassReference_
}

var symbolClassReference_ = &symbolClass_{
	// Initialize the class constants.
	matcher_:   reg.MustCompile("^\\$(" + identifier_ + ")"),
	undefined_: symbol_("$"),
}
