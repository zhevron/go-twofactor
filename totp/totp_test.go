package totp

import (
	"testing"

	"github.com/zhevron/twofactor"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TOTPSuite struct {
	secret twofactor.Secret
	totp   *TOTP
}

var _ = Suite(&TOTPSuite{})

func (s *TOTPSuite) SetUpTest(c *C) {
	s.secret, _ = twofactor.NewSecret(0)
	s.totp = NewTOTP(s.secret, 0, 0)
}

func (s *TOTPSuite) TestNewTOTP(c *C) {
	c.Check(s.totp.secret, Equals, s.secret)
	c.Check(s.totp.length, Equals, 6)
	c.Check(s.totp.duration, Equals, 30)
}

func (s *TOTPSuite) TestTOTPCode(c *C) {
	code, _ := s.totp.Code()
	c.Check(code, Not(Equals), 0)
}

func (s *TOTPSuite) TestTOTPSecret(c *C) {
	c.Check(s.totp.Secret().String(), Equals, s.secret.String())
}

func (s *TOTPSuite) TestValidate(c *C) {
	code, _ := s.totp.Code()
	prevCode, _ := s.totp.ForPeriodOffset(-1)
	c.Check(s.totp.Validate(code, 1), Equals, true)
	c.Check(s.totp.Validate(prevCode, 1), Equals, true)
	c.Check(s.totp.Validate(123456, 1), Equals, false)
}

func (s *TOTPSuite) TestTOTPForPeriod(c *C) {
	code, _ := s.totp.ForPeriod(1)
	c.Check(code, Not(Equals), 0)
}

func (s *TOTPSuite) TestTOTPForPeriodOffset(c *C) {
	code, _ := s.totp.ForPeriodOffset(-1)
	c.Check(code, Not(Equals), 0)
}

func (s *TOTPSuite) TestTOTPCurrentPeriod(c *C) {
	c.Check(s.totp.currentPeriod(), Not(Equals), 0)
}

func (s *TOTPSuite) TestTOTPGenerateCode(c *C) {
	totp := NewTOTP("ONSWG4TFOQ======", 0, 0)
	periods := []int64{1000, 2000, 3000, 4000, 5000}
	codes := []int32{562944, 414684, 722632, 500486, 376835}
	for i, n := range periods {
		code, _ := totp.generateCode(n)
		c.Check(code, Equals, codes[i])
	}
}

func (s *TOTPSuite) BenchmarkGenerateCode(c *C) {
	for i := 0; i < c.N; i++ {
		s.totp.generateCode(int64(i))
	}
}
