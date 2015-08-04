twofactor - Two-factor authentication library
=============================================

[![Coverage Status](https://img.shields.io/coveralls/zhevron/twofactor.svg)](https://coveralls.io/r/zhevron/twofactor)
[![Build Status](https://travis-ci.org/zhevron/twofactor.svg?branch=master)](https://travis-ci.org/zhevron/twofactor)

**twofactor** is a library for handling two-factor authentication in [Go](https://golang.org/).  
To use the library, you need to install one of the implementations below.

## TOTP

[![GoDoc](https://godoc.org/github.com/zhevron/twofactor/totp?status.svg)](https://godoc.org/github.com/zhevron/twofactor/totp)

This implementation generates and verifies codes using the time-based one time password (TOTP) algorithm specified in [RFC6238](https://tools.ietf.org/html/rfc6238).

```
go get gopkg.in/zhevron/go-twofactor.v1/totp
```

## License

**twofactor** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
