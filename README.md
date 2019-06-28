# go-zabbix

Go bindings for the Zabbix API

[![go report card](https://goreportcard.com/badge/github.com/cavaliercoder/go-zabbix "go report card")](https://goreportcard.com/report/github.com/cavaliercoder/go-zabbix)
[![cover.run](https://cover.run/go/github.com/cavaliercoder/go-zabbix.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fcavaliercoder%2Fgo-zabbix)
[![GPL license](https://img.shields.io/badge/license-GPL-brightgreen.svg)](https://opensource.org/licenses/gpl-license)
[![GoDoc](https://godoc.org/github.com/cavaliercoder/go-zabbix?status.svg)](https://godoc.org/github.com/cavaliercoder/go-zabbix)

## Overview

This project provides bindings to interoperate between programs written in Go
language and the Zabbix monitoring API.

A number of Zabbix API bindings already exist for Go with varying levels of
maturity. This project aims to provide an alternative implementation which is
stable, fast, and allows for loose typing (using types such as `interface{}` or
`map[string]interface{}`) as well as strong types (such as `Host` or `Event`).

The package aims to have comprehensive coverage of Zabbix API methods from v1.8
through to v3.0 without introducing limitations to the native API methods.

## Getting started

```go
package main

import (
	"crypto/tls"
	"fmt"
	"time"
	"github.com/cavaliercoder/go-zabbix"
)

func main() {
	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix", time.Duration(60) * time.Second)
	if err != nil {
		fmt.Println(err)
	}

	var grp []zabbix.Usergroups
	grp = append(grp, zabbix.Usergroups{UsergroupID: "84"})
	tmp, err := session.CreateUser(zabbix.UserCreateParams{Alias: "testicek", Passwd: "password", Usergroup: grp})

	fmt.Println(tmp,err)
}
```

## License

Released under the [GNU GPL License](https://github.com/cavaliercoder/go-zabbix/blob/master/LICENSE)
