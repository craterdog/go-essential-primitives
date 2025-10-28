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

package elements

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v8"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func SymbolClass() SymbolClassLike {
	return symbolClass()
}

// Constructor Methods

func (c *symbolClass_) Symbol(
	identifier string,
) SymbolLike {
	return c.SymbolFromSource("$" + identifier)
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

func (v symbol_) AsIntrinsic() string {
	return string(v)
}

func (v symbol_) AsSource() string {
	return "$" + string(v)
}

// Attribute Methods

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
	digit_      = "\\p{Nd}"
	identifier_ = "(?:" + letter_ + ")(?:" + letter_ + "|" + digit_ + "|-)*"
	letter_     = lower_ + "|" + upper_
	lower_      = "\\p{Ll}"
	upper_      = "\\p{Lu}"
	version_    = "v" + ordinal_ + "(?:\\." + ordinal_ + ")*"
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
