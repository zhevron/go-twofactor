go2fa - Two-factor authentication for Google Go
===============================================

[![Coverage Status](https://img.shields.io/coveralls/zhevron/go2fa.svg)](https://coveralls.io/r/zhevron/go2fa)
[![Build Status](https://travis-ci.org/zhevron/go2fa.svg?branch=master)](https://travis-ci.org/zhevron/go2fa)
[![GoDoc](https://godoc.org/github.com/zhevron/go2fa?status.svg)](https://godoc.org/github.com/zhevron/go2fa)

go2fa is a Go library for handling two-factor authentication.

You can install the library using Go:

```
go get gopkg.in/zhevron/go2fa.v1
```

You will also need to install one of the implementations.

Time-based One Time Password (TOTP) [![GoDoc](https://godoc.org/github.com/zhevron/go2fa/totp?status.svg)](https://godoc.org/github.com/zhevron/go2fa/totp)
```
go get gopkg.in/zhevron/go2fa.v1/totp
```

go2fa is licensed under the [MIT license](http://opensource.org/licenses/MIT).
