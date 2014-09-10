// Copyright (C) 2014 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package hotp contains an RFC4226 implementation of HMAC-based one time passwords.
package hotp

import "github.com/zhevron/go2fa"

// HOTP implements HMAC-based one time passwords as specified in RFC4226 (http://tools.ietf.org/html/rfc4226).
type HOTP struct {
	secret   go2fa.Secret
	length   int
	duration int
}

// NewHOTP makes a new HOTP object for generating OTP codes from a given Secret.
func NewHOTP(s go2fa.Secret) *HOTP {
	return &HOTP{
		secret: s,
	}
}
