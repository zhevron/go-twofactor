package twofactor

import (
	"bytes"
	"encoding/base32"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TwoFactorSuite struct{}

var _ = Suite(&TwoFactorSuite{})

func (s *TwoFactorSuite) TestInt64ToBytes(c *C) {
	ints := []int64{10, 200, 3000, 40000, 500000}
	intBytes := [][]byte{
		[]byte{0, 0, 0, 0, 0, 0, 0, 10},
		[]byte{0, 0, 0, 0, 0, 0, 0, 200},
		[]byte{0, 0, 0, 0, 0, 0, 11, 184},
		[]byte{0, 0, 0, 0, 0, 0, 156, 64},
		[]byte{0, 0, 0, 0, 0, 7, 161, 32},
	}
	for i, n := range ints {
		c.Check(bytes.Compare(Int64ToBytes(n), intBytes[i]), Equals, 0)
	}
}

func (s *TwoFactorSuite) TestNewSecret(c *C) {
	secret, _ := NewSecret(0)
	c.Check(len(secret), Equals, 16)
}

func (s *TwoFactorSuite) TestSecretBytes(c *C) {
	str := "thisisasecret"
	strb, _ := base32.StdEncoding.DecodeString(str)
	secret := Secret(str)
	bytes, _ := secret.Bytes()
	c.Check(len(bytes), Equals, len(strb))
	for i, b := range bytes {
		c.Check(b, Equals, strb[i])
	}
}

func (s *TwoFactorSuite) TestSecretString(c *C) {
	str := "thisisasecret"
	secret := Secret(str)
	c.Check(secret.String(), Equals, str)
}

func (s *TwoFactorSuite) BenchmarkNewSecret(c *C) {
	for i := 0; i < c.N; i++ {
		NewSecret(0)
	}
}
