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
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func BinaryClass() BinaryClassLike {
	return binaryClass()
}

// Constructor Methods

func (c *binaryClass_) Binary(
	bytes []byte,
) BinaryLike {
	var encoded = uti.Base64Encode(bytes)
	var length = len(encoded)
	var source = "'>"
	if length > 0 {
		source += "\n"
		var width = 60
		var indentation = "    "
		var index int
		for index = 0; index+width < length; index += width {
			source += indentation + encoded[index:index+width] + "\n"
		}
		source += indentation + encoded[index:] + "\n"
	}
	source += "<'"
	return binary_(source)
}

func (c *binaryClass_) BinaryFromSequence(
	sequence Sequential[byte],
) BinaryLike {
	return c.Binary(sequence.AsArray())
}

func (c *binaryClass_) BinaryFromSource(
	source string,
) BinaryLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the binary constructor method: %s",
			source,
		)
		panic(message)
	}
	return binary_(source)
}

// Constant Methods

// Function Methods

func (c *binaryClass_) Not(
	binary BinaryLike,
) BinaryLike {
	var bytes = binary.AsIntrinsic()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return c.Binary(bytes)
}

func (c *binaryClass_) And(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsIntrinsic()
	var secondBytes = second.AsIntrinsic()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] & secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) San(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsIntrinsic()
	var secondBytes = second.AsIntrinsic()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] &^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Ior(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsIntrinsic()
	var secondBytes = second.AsIntrinsic()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] | secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Xor(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsIntrinsic()
	var secondBytes = second.AsIntrinsic()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] ^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Concatenate(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var firstBytes = first.AsIntrinsic()
	var secondBytes = second.AsIntrinsic()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Binary(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v binary_) GetClass() BinaryClassLike {
	return binaryClass()
}

func (v binary_) AsIntrinsic() []byte {
	var binary = string(v)
	var base64 = binary[2 : len(v)-2]         // Strip off the delimiters.
	base64 = sts.ReplaceAll(base64, " ", "")  // Remove all spaces.
	base64 = sts.ReplaceAll(base64, "\n", "") // Remove all newlines.
	var bytes = uti.Base64Decode(base64)
	return bytes
}

func (v binary_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Sequential[byte] Methods

func (v binary_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v binary_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v binary_) AsArray() []byte {
	return v.AsIntrinsic()
}

func (v binary_) GetIterator() uti.Ratcheted[byte] {
	return uti.Iterator(v.AsIntrinsic())
}

// PROTECTED INTERFACE

func (v binary_) String() string {
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
	alpha_        = "[A-Za-z]"
	alphanumeric_ = alpha_ + "|" + base10_
	base10_       = "[0-9]"
	base64_       = alphanumeric_ + "|[\\+/]"
	space_        = " "
)

// Instance Structure

type binary_ string // This type must support the "comparable" type contraint.

// Class Structure

type binaryClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func binaryClass() *binaryClass_ {
	return binaryClassReference_
}

var binaryClassReference_ = &binaryClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'>(" + eol_ + "((?:" + space_ + ")*(?:" + base64_ + "){2,60}" +
			eol_ + ")+(?:" + space_ + ")*)?<'",
	),
}
