twofactor - Two-factor authentication library
=============================================

[![Coverage Status](https://img.shields.io/coveralls/zhevron/twofactor.svg)](https://coveralls.io/r/zhevron/twofactor)
[![Build Status](https://travis-ci.org/zhevron/twofactor.svg?branch=master)](https://travis-ci.org/zhevron/twofactor)

**twofactor** is a library for handling two-factor authentication in [Go](https://golang.org/).  
To use the library, you need to install one of the implementations below.

## YAML

[![GoDoc](https://godoc.org/github.com/zhevron/twofactor/totp?status.svg)](https://godoc.org/github.com/zhevron/twofactor/totp)

This implementation reads and writes [YAML](http://www.yaml.org/) files.

```
go get gopkg.in/zhevron/twofactor.v1/totp
```

## License

**twofactor** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
