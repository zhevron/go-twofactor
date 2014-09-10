package hotp

import (
	"testing"

	"github.com/zhevron/go2fa"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type HOTPSuite struct {
	secret go2fa.Secret
	hotp   *HOTP
}

var _ = Suite(&HOTPSuite{})

func (s *HOTPSuite) SetUpTest(c *C) {
	s.secret, _ = go2fa.NewSecret(0)
	s.hotp = NewHOTP(s.secret)
}
