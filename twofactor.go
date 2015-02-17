// Copyright (C) 2014-2015 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package twofactor contains support functions and types for two-factor authentication.
//
// This package provides shared structs and functions for all implementations.
//
// To generate a new secret for 2FA, use the NewSecret method:
//		import "gopkg.in/zhevron/twofactor.v1"
//		secret := twofactor.NewSecret(0)
package twofactor

import (
	"crypto/rand"
	"encoding/base32"
)

// Secret wraps a string with functions to generate a cryptographically secure secret.
type Secret string

// NewSecret generates a cryptograpically secure string of the given length.
// If the length provided is 0 or less, it will default to a length of 10.
func NewSecret(length int) (Secret, error) {
	if length <= 0 {
		length = 10
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	return Secret(base32.StdEncoding.EncodeToString(b)), err
}

// Bytes decodes the secret into a byte slice.
func (s Secret) Bytes() ([]byte, error) {
	return base32.StdEncoding.DecodeString(s.String())
}

// String returns the secret as a string.
func (s Secret) String() string {
	return string(s)
}

// Int64ToBytes converts a 64-bit integer to a 8-byte slice.
func Int64ToBytes(n int64) []byte {
	b := make([]byte, 8)
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n = n >> 8
	}
	return b
}
