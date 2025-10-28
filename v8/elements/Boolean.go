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
	uti "github.com/craterdog/go-essential-utilities/v8"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func BooleanClass() BooleanClassLike {
	return booleanClass()
}

// Constructor Methods

func (c *booleanClass_) Boolean(
	boolean bool,
) BooleanLike {
	return boolean_(boolean)
}

func (c *booleanClass_) BooleanFromSource(
	source string,
) BooleanLike {
	// Our booleans are more restrictive than the Go strconv package.
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the boolean constructor method: %s",
			source,
		)
		panic(message)
	}
	var boolean, _ = stc.ParseBool(matches[0])
	return boolean_(boolean)
}

// Constant Methods

func (c *booleanClass_) False() BooleanLike {
	return c.false_
}

func (c *booleanClass_) True() BooleanLike {
	return c.true_
}

// Function Methods

func (c *booleanClass_) Not(
	boolean BooleanLike,
) BooleanLike {
	var result_ = boolean_(!boolean.AsIntrinsic())
	return result_
}

func (c *booleanClass_) And(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsIntrinsic() && second.AsIntrinsic())
	return result_
}

func (c *booleanClass_) San(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsIntrinsic() && !second.AsIntrinsic())
	return result_
}

func (c *booleanClass_) Ior(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsIntrinsic() || second.AsIntrinsic())
	return result_
}

func (c *booleanClass_) Xor(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = c.Ior(c.San(first, second), c.San(second, first))
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v boolean_) GetClass() BooleanClassLike {
	return booleanClass()
}

func (v boolean_) AsIntrinsic() bool {
	return bool(v)
}

// Attribute Methods

// Discrete Methods

func (v boolean_) AsSource() string {
	var result_ = stc.FormatBool(bool(v))
	return result_
}

func (v boolean_) AsInteger() int {
	var result_ int
	if v {
		result_ = 1
	}
	return result_
}

func (v boolean_) IsDefined() bool {
	return true
}

func (v boolean_) IsMinimum() bool {
	return !v.AsIntrinsic()
}

func (v boolean_) IsZero() bool {
	return !v.AsIntrinsic()
}

func (v boolean_) IsMaximum() bool {
	return v.AsIntrinsic()
}

// PROTECTED INTERFACE

func (v boolean_) String() string {
	return v.AsSource()
}

// Private Methods

// Instance Structure

type boolean_ bool

// Class Structure

type booleanClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
	false_   BooleanLike
	true_    BooleanLike
}

// Class Reference

func booleanClass() *booleanClass_ {
	return booleanClassReference_
}

var booleanClassReference_ = &booleanClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^false|true"),
	false_:   boolean_(false),
	true_:    boolean_(true),
}
