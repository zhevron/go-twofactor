twofactor - Two-factor authentication library
=============================================

[![wercker status](https://app.wercker.com/status/fd9c8ba528a1797f2e9cd5bbc4a1dc29/s/master "wercker status")](https://app.wercker.com/project/bykey/fd9c8ba528a1797f2e9cd5bbc4a1dc29)
[![Coverage Status](https://coveralls.io/repos/zhevron/go-twofactor/badge.svg?branch=HEAD)](https://coveralls.io/r/zhevron/go-twofactor)

**go-twofactor** is a library for handling two-factor authentication in [Go](https://golang.org/).  
To use the library, you need to install one of the implementations below.

## TOTP

[![GoDoc](https://godoc.org/github.com/zhevron/go-twofactor/totp?status.svg)](https://godoc.org/github.com/zhevron/go-twofactor/totp)

This implementation generates and verifies codes using the time-based one time password (TOTP) algorithm specified in [RFC6238](https://tools.ietf.org/html/rfc6238).

```
go get gopkg.in/zhevron/go-twofactor.v1/totp
```

## License

**go-twofactor** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
