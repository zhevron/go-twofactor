go2fa - Two-factor authentication for Google Go
===============================================

[![Coverage Status](https://img.shields.io/coveralls/golinux-io/go2fa.svg)](https://coveralls.io/r/golinux-io/go2fa)
[![Build Status](https://travis-ci.org/golinux-io/go2fa.svg?branch=master)](https://travis-ci.org/golinux-io/go2fa)
[![GoDoc](https://godoc.org/github.com/golinux-io/go2fa?status.svg)](https://godoc.org/github.com/golinux-io/go2fa)

go2fa is a Go library for handling two-factor authentication.

You can install the library using Go:

```
go get gopkg.in/golinux-io/go2fa.v1
```

You will also need to install one of the implementations.

Time-based One Time Password (TOTP) [![GoDoc](https://godoc.org/github.com/golinux-io/go2fa/totp?status.svg)](https://godoc.org/github.com/golinux-io/go2fa/totp)
```
go get gopkg.in/golinux-io/go2fa.v1/totp
```

HMAC-based One Time Password (HOTP) [![GoDoc](https://godoc.org/github.com/zhevron/go2fa/hotp?status.svg)](https://godoc.org/github.com/zhevron/go2fa/hotp)
```
go get gopkg.in/zhevron/go2fa.v1/hotp
```

go2fa is licensed under the [MIT license](http://opensource.org/licenses/MIT).
