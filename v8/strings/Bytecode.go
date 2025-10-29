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
	mat "math"
	reg "regexp"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func BytecodeClass() BytecodeClassLike {
	return bytecodeClass()
}

// Constructor Methods

func (c *bytecodeClass_) Bytecode(
	instructions []uint16,
) BytecodeLike {
	var source = "'>"
	var newline = "\n    "
	for index := 0; index < len(instructions); index++ {
		var instruction = instructions[index]
		if mat.Mod(float64(index), 12.0) == 0 {
			source += newline
		}
		source += fmt.Sprintf(":%04x", instruction)
	}
	source += "\n<'"
	return bytecode_(source)
}

func (c *bytecodeClass_) BytecodeFromSequence(
	sequence Sequential[uint16],
) BytecodeLike {
	return c.Bytecode(sequence.AsArray())
}

func (c *bytecodeClass_) BytecodeFromSource(
	source string,
) BytecodeLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the bytecode constructor method: %s",
			source,
		)
		panic(message)
	}
	return bytecode_(source)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v bytecode_) GetClass() BytecodeClassLike {
	return bytecodeClass()
}

func (v bytecode_) AsIntrinsic() []uint16 {
	var base16 = string(v)
	base16 = base16[2 : len(base16)-2]        // Strip off the delimiters.
	base16 = sts.ReplaceAll(base16, "\n", "") // Remove all newlines.
	base16 = sts.ReplaceAll(base16, " ", "")  // Remove all spaces.
	var strings = sts.Split(base16, ":")[1:]  // Extract the instructions.
	var instructions = make([]uint16, len(strings))
	for index, hex := range strings {
		var integer, _ = stc.ParseUint(hex, 16, 16)
		var instruction = uint16(integer)
		instructions[index] = instruction
	}
	return instructions
}

func (v bytecode_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Sequential[uint16] Methods

func (v bytecode_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v bytecode_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v bytecode_) AsArray() []uint16 {
	return v.AsIntrinsic()
}

func (v bytecode_) GetIterator() uti.Ratcheted[uint16] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v bytecode_) String() string {
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
	instruction_ = ":(?:" + base16_ + "){4}"
)

// Instance Structure

type bytecode_ string // This type must support the "comparable" type contraint.

// Class Structure

type bytecodeClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func bytecodeClass() *bytecodeClass_ {
	return bytecodeClassReference_
}

var bytecodeClassReference_ = &bytecodeClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'>" + eol_ + "((?:" + space_ + ")*(?:" + instruction_ + "){1,12}" +
			eol_ + ")+(?:" + space_ + ")*<'",
	),
}
