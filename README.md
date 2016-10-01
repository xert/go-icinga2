# go-icinga2

go-icinga2 is a Go client library for accessing the [Icinga2 API](http://docs.icinga.org/icinga2/latest/doc/module/icinga2/chapter/icinga2-api).

[![License](https://img.shields.io/badge/license-BSD-red.svg)](./LICENSE)
[![Build Status](https://travis-ci.org/xert/go-icinga2.svg?branch=master)](https://travis-ci.org/xert/go-icinga2)
[![Go Report Card](https://goreportcard.com/badge/github.com/xert/go-icinga2)](https://goreportcard.com/report/github.com/xert/go-icinga2)
[![codecov](https://codecov.io/gh/xert/go-icinga2/branch/master/graph/badge.svg)](https://codecov.io/gh/xert/go-icinga2)
[![GoDoc](https://godoc.org/github.com/xert/go-icinga2/icinga?status.svg)](https://godoc.org/github.com/xert/go-icinga2/icinga)

## Status ##

In very early stage of development. We'd love to accept your pull requests and contributions to this project.

## Usage ##

Not very usable at the moment.

```go
transport := &icinga.BasicAuthTransport{
	Username: "user",
	Password: "pass",

	Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
}

icinga := icinga.NewClient("https://localhost:5665/", transport.Client())

host, resp, err := icinga.Hosts.Get("hostname")
```

## Licence ##

This library is distributed under the New BSD License found in the [LICENSE](./LICENSE) file.
