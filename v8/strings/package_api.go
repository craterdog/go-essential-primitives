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

/*
Package "strings" provides a framework of aspects and class definitions for a
rich set of primitive data types that can be iterated over.  All primitive types
are immutable and—for better performance—are implemented as extensions to
existing Go primitive types.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-essential-primitives/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package strings

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
	reg "regexp"
)

// TYPE DECLARATIONS

/*
Folder is a constrained type representing a string of the form:
['0'..'9' 'A'..'Z' 'a'..'z'] (('-' | '.')? ['0'..'9' 'A'..'Z' 'a'..'z'])+
*/
type Folder string

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
BinaryClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
binary-like concrete class.
*/
type BinaryClassLike interface {
	// Constructor Methods
	Binary(
		bytes []byte,
	) BinaryLike
	BinaryFromSequence(
		sequence Sequential[byte],
	) BinaryLike
	BinaryFromSource(
		source string,
	) BinaryLike

	// Function Methods
	Not(
		binary BinaryLike,
	) BinaryLike
	And(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	San(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Ior(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Xor(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Concatenate(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
}

/*
BytecodeClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete bytecode-like class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		instructions []uint16,
	) BytecodeLike
	BytecodeFromSequence(
		sequence Sequential[uint16],
	) BytecodeLike
	BytecodeFromSource(
		source string,
	) BytecodeLike
}

/*
NameClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
name-like concrete class.
*/
type NameClassLike interface {
	// Constructor Methods
	Name(
		folders []Folder,
	) NameLike
	NameFromSequence(
		sequence Sequential[Folder],
	) NameLike
	NameFromSource(
		source string,
	) NameLike

	// Function Methods
	Concatenate(
		first NameLike,
		second NameLike,
	) NameLike
}

/*
NarrativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
narrative-like concrete class.
*/
type NarrativeClassLike interface {
	// Constructor Methods
	Narrative(
		lines []string,
	) NarrativeLike
	NarrativeFromSequence(
		sequence Sequential[string],
	) NarrativeLike
	NarrativeFromSource(
		source string,
	) NarrativeLike

	// Function Methods
	Concatenate(
		first NarrativeLike,
		second NarrativeLike,
	) NarrativeLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
pattern-like concrete class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Pattern(
		characters []rune,
	) PatternLike
	PatternFromSequence(
		sequence Sequential[rune],
	) PatternLike
	PatternFromSource(
		source string,
	) PatternLike

	// Constant Methods
	None() PatternLike
	Any() PatternLike

	// Function Methods
	Concatenate(
		first PatternLike,
		second PatternLike,
	) PatternLike
}

/*
QuoteClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
quote-like concrete class.
*/
type QuoteClassLike interface {
	// Constructor Methods
	Quote(
		characters []rune,
	) QuoteLike
	QuoteFromSequence(
		sequence Sequential[rune],
	) QuoteLike
	QuoteFromSource(
		source string,
	) QuoteLike

	// Function Methods
	Concatenate(
		first QuoteLike,
		second QuoteLike,
	) QuoteLike
}

/*
TagClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
tag-like concrete class.
*/
type TagClassLike interface {
	// Constructor Methods
	Tag(
		bytes []byte,
	) TagLike
	TagWithSize(
		size uint,
	) TagLike
	TagFromSequence(
		sequence Sequential[byte],
	) TagLike
	TagFromSource(
		source string,
	) TagLike

	// Function Methods
	Concatenate(
		first TagLike,
		second TagLike,
	) TagLike
}

/*
VersionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
version-like concrete class.
*/
type VersionClassLike interface {
	// Constructor Methods
	Version(
		ordinals []uint,
	) VersionLike
	VersionFromSequence(
		sequence Sequential[uint],
	) VersionLike
	VersionFromSource(
		source string,
	) VersionLike

	// Function Methods
	IsValidNextVersion(
		current VersionLike,
		next VersionLike,
	) bool
	GetNextVersion(
		current VersionLike,
		level uint,
	) VersionLike
	Concatenate(
		first VersionLike,
		second VersionLike,
	) VersionLike
}

// INSTANCE DECLARATIONS

/*
BinaryLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete binary-like class.
*/
type BinaryLike interface {
	// Principal Methods
	GetClass() BinaryClassLike
	AsIntrinsic() []byte
	AsSource() string

	// Aspect Interfaces
	Sequential[byte]
}

/*
BytecodeLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	AsIntrinsic() []uint16
	AsSource() string

	// Aspect Interfaces
	Sequential[uint16]
}

/*
NameLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete name-like class.
*/
type NameLike interface {
	// Principal Methods
	GetClass() NameClassLike
	AsIntrinsic() []Folder
	AsSource() string

	// Aspect Interfaces
	Accessible[Folder]
	Searchable[Folder]
	Sequential[Folder]
	Ordered[NameLike]
}

/*
NarrativeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete narrative-like class.
*/
type NarrativeLike interface {
	// Principal Methods
	GetClass() NarrativeClassLike
	AsIntrinsic() []string
	AsSource() string

	// Aspect Interfaces
	Accessible[string]
	Searchable[string]
	Sequential[string]
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a pattern-like elemental class.
*/
type PatternLike interface {
	// Principal Methods
	GetClass() PatternClassLike
	AsIntrinsic() []rune
	AsSource() string
	AsRegexp() *reg.Regexp
	MatchesText(
		text string,
	) bool
	GetMatches(
		text string,
	) []string

	// Aspect Interfaces
	Accessible[rune]
	Searchable[rune]
	Sequential[rune]
}

/*
QuoteLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete quote-like class.
*/
type QuoteLike interface {
	// Principal Methods
	GetClass() QuoteClassLike
	AsIntrinsic() []rune
	AsSource() string

	// Aspect Interfaces
	Accessible[rune]
	Searchable[rune]
	Sequential[rune]
	Ordered[QuoteLike]
}

/*
TagLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete tag-like class.
*/
type TagLike interface {
	// Principal Methods
	GetClass() TagClassLike
	AsIntrinsic() []byte
	AsSource() string
	GetHash() uint64

	// Aspect Interfaces
	Accessible[byte]
	Searchable[byte]
	Sequential[byte]
}

/*
VersionLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete version-like class.
*/
type VersionLike interface {
	// Principal Methods
	GetClass() VersionClassLike
	AsIntrinsic() []uint
	AsSource() string

	// Aspect Interfaces
	Accessible[uint]
	Searchable[uint]
	Sequential[uint]
	Ordered[VersionLike]
}

// ASPECT DECLARATIONS

/*
Accessible[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of an accessible concrete
class.

An accessible class maintains values that can be accessed using indices. The
indices of an accessible sequence are ORDINAL rather than ZERO based—which
never really made sense except for pointer offsets. What is the "zeroth
value" anyway? It's the "first value", right?  So we start fresh...

This approach allows for positive indices starting at the beginning of the
sequence, and negative indices starting at the end of the sequence as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]
	   -N        -(N-1)      -(N-2)          -1

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.
*/
type Accessible[V any] interface {
	GetValue(
		index int,
	) V
	GetValues(
		first int,
		last int,
	) Sequential[V]
	GetIndex(
		value V,
	) int
}

/*
Searchable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a searchable concrete
class.
*/
type Searchable[V any] interface {
	ContainsValue(
		value V,
	) bool
	ContainsAny(
		values Sequential[V],
	) bool
	ContainsAll(
		values Sequential[V],
	) bool
}

/*
Sequential[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sequential concrete
class.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() uint
	AsArray() []V
	GetIterator() uti.IteratorLike[V]
}

/*
Ordered[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of an ordered concrete class.
*/
type Ordered[V any] interface {
	AsSource() string
	IsBefore(
		value V,
	) bool
}
