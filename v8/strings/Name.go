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
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NameClass() NameClassLike {
	return nameClass()
}

// Constructor Methods

func (c *nameClass_) Name(
	folders []Folder,
) NameLike {
	var source string
	for _, folder := range folders {
		source += "/" + string(folder)
	}
	return name_(source)
}

func (c *nameClass_) NameFromSequence(
	sequence Sequential[Folder],
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
	var firstFolders = first.AsIntrinsic()
	var secondFolders = second.AsIntrinsic()
	var allFolders = make(
		[]Folder,
		len(firstFolders)+len(secondFolders),
	)
	copy(allFolders, firstFolders)
	copy(allFolders[len(firstFolders):], secondFolders)
	return c.Name(allFolders)
}

// INSTANCE INTERFACE

// Principal Methods

func (v name_) GetClass() NameClassLike {
	return nameClass()
}

func (v name_) AsIntrinsic() []Folder {
	var name = string(v)
	var strings = sts.Split(name, "/")[1:] // Extract the folders.
	var folders = make([]Folder, len(strings))
	for index, folder := range strings {
		folders[index] = Folder(folder)
	}
	return folders
}

func (v name_) AsSource() string {
	return string(v)
}

// Attribute Methods

// Spectral Methods

func (v name_) IsBefore(
	value NameLike,
) bool {
	return sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) < 0
}

// Searchable[Folder] Methods

func (v name_) ContainsValue(
	value Folder,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v name_) ContainsAny(
	values Sequential[Folder],
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
	values Sequential[Folder],
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

// Sequential[Folder] Methods

func (v name_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v name_) GetSize() uint {
	return uti.ArraySize(v.AsIntrinsic())
}

func (v name_) AsArray() []Folder {
	return v.AsIntrinsic()
}

func (v name_) GetIterator() uti.Ratcheted[Folder] {
	return uti.Iterator(v.AsIntrinsic())
}

// Accessible[Folder] Methods

func (v name_) GetValue(
	index int,
) Folder {
	var folders = v.AsIntrinsic()
	var size = uti.ArraySize(folders)
	var goIndex = uti.RelativeToCardinal(index, size)
	return folders[goIndex]
}

func (v name_) GetValues(
	first int,
	last int,
) Sequential[Folder] {
	var folders = v.AsIntrinsic()
	var size = uti.ArraySize(folders)
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return nameClass().Name(folders[goFirst : goLast+1])
}

func (v name_) GetIndex(
	value Folder,
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
		"^(?:/(?:" + letter_ + "|" + digit_ + "|-)+" + ")+",
	),
}
