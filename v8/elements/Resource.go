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
	uri "net/url"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func ResourceClass() ResourceClassLike {
	return resourceClass()
}

// Constructor Methods

func (c *resourceClass_) Resource(
	uri string,
) ResourceLike {
	return c.ResourceFromSource("<" + uri + ">")
}

func (c *resourceClass_) ResourceFromSource(
	source string,
) ResourceLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the resource constructor method: %s",
			source,
		)
		panic(message)
	}
	return resource_(matches[1]) // Strip off the angle brackets.
}

func (c *resourceClass_) ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return resource_(url.String())
}

// Constant Methods

func (c *resourceClass_) Undefined() ResourceLike {
	return c.undefined_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v resource_) GetClass() ResourceClassLike {
	return resourceClass()
}

func (v resource_) AsIntrinsic() string {
	return string(v)
}

func (v resource_) AsSource() string {
	return "<" + string(v) + ">"
}

func (v resource_) AsUri() *uri.URL {
	var url, _ = uri.Parse(string(v))
	return url
}

func (v resource_) GetScheme() string {
	var url = v.AsUri()
	return url.Scheme
}

func (v resource_) GetAuthority() string {
	var authority string
	var url = v.AsUri()
	var user = url.User.String()
	var host = url.Host
	if len(user) > 0 {
		authority = user + "@"
	}
	authority += host
	return authority
}

func (v resource_) GetPath() string {
	var url = v.AsUri()
	return url.Path
}

func (v resource_) GetQuery() string {
	var url = v.AsUri()
	return url.RawQuery
}

func (v resource_) GetFragment() string {
	var url = v.AsUri()
	return url.Fragment
}

// Attribute Methods

// PROTECTED INTERFACE

func (v resource_) String() string {
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
	authority_    = "[^/" + control_ + "]+"
	fragment_     = "[^>" + control_ + "]*"
	path_         = "[^\\?#>" + control_ + "]*"
	query_        = "[^#>" + control_ + "]*"
	scheme_       = alpha_ + "(?:" + alphanumeric_ + "|\\+|-|\\.)*"
)

// Instance Structure

type resource_ string

// Class Structure

type resourceClass_ struct {
	// Declare the class constants.
	matcher_   *reg.Regexp
	undefined_ ResourceLike
}

// Class Reference

func resourceClass() *resourceClass_ {
	return resourceClassReference_
}

var resourceClassReference_ = &resourceClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^<((" + scheme_ + "):(?://(" + authority_ + "))?(" + path_ +
			")(?:\\?(" + query_ + "))?(?:#(" + fragment_ + "))?)>",
	),
	undefined_: resource_("<>"),
}
