// Copyright (C) 2014-2015 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package totp contains an RFC6238 implementation of time-based one time passwords.
//
// To be able to generate codes, you need to create a TOTP object using a secret:
//		import (
//			"gopkg.in/zhevron/go2fa.v1"
//			"gopkg.in/zhevron/go2fa.v1/totp"
//		)
//		secret := go2fa.NewSecret(0)
//		otp := totp.NewTOTP(secret, 0, 0)
//
// After creating the TOTP object, you can easily generate the current code by calling the Code method:
//		code, err := otp.Code()
//		if err != nil {
//			// Failed to generate the code
//		}
//
// A convenience method is also provided for checking if a code is valid using time period offset:
//		code := 123456
//		if otp.Validate(code, 0) {
//			// Code is valid
//		} else {
//			// Code is not valid
//		}
package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"math"
	"time"

	"github.com/zhevron/go2fa"
)

// TOTP implements time-based one time passwords as specified in RFC6238 (http://tools.ietf.org/html/rfc6238).
type TOTP struct {
	secret   go2fa.Secret
	length   int
	duration int
}

// NewTOTP makes a new TOTP object for generating OTP codes from a given Secret.
// If length is 0 or less, it will default to a length of 6.
// If duration is 0 or less, it will default to a duration of 30 seconds.
func NewTOTP(s go2fa.Secret, length int, duration int) *TOTP {
	if length <= 0 {
		length = 6
	}
	if duration <= 0 {
		duration = 30
	}
	return &TOTP{
		secret:   s,
		length:   length,
		duration: duration,
	}
}

// Code returns the current OTP code.
func (t TOTP) Code() (int32, error) {
	return t.ForPeriod(t.currentPeriod())
}

// Secret returns the secret for this TOTP object.
func (t TOTP) Secret() go2fa.Secret {
	return t.secret
}

// Validate checks the the provided code is valid.
func (t TOTP) Validate(code int32, offset int) bool {
	var codes []int32
	for i := -offset; i <= offset; i++ {
		if c, err := t.ForPeriodOffset(int64(i)); err == nil {
			codes = append(codes, c)
		}
	}
	for _, c := range codes {
		if c == int32(code) {
			return true
		}
	}
	return false
}

// ForPeriod returns the code for the given period.
func (t TOTP) ForPeriod(n int64) (int32, error) {
	return t.generateCode(n)
}

// ForPeriodOffset returns the code from the time period of current+offset.
func (t TOTP) ForPeriodOffset(n int64) (int32, error) {
	return t.ForPeriod(t.currentPeriod() + n)
}

// currentPeriod returns the current time period.
func (t TOTP) currentPeriod() int64 {
	return time.Now().Unix() / int64(t.duration)
}

// generateCode returns the calculated code for the given time period.
func (t TOTP) generateCode(n int64) (int32, error) {
	var code int32
	b, err := t.secret.Bytes()
	if err == nil {
		h := hmac.New(sha1.New, b)
		h.Write(go2fa.Int64ToBytes(n))
		hash := h.Sum(nil)
		offset := int((len(hash) - 1) & 0xf)
		binary := int32(hash[offset]&0x7f)<<24 |
			int32(hash[offset+1]&0xff)<<16 |
			int32(hash[offset+2]&0xff)<<8 |
			int32(hash[offset+3]&0xff)
		code = int32(int64(binary) % int64(math.Pow10(t.length)))
	}
	return code, err
}
