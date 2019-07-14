# go-zabbix

Go bindings for the Zabbix API

## Overview

This project provides bindings to interoperate between programs written in Go
language and the Zabbix monitoring API with very ugly hacks to orchestrate 
user/host/meda/templates/actions/groups creation. 

## Getting started with very rough examle

```go
package main

import (
	"crypto/tls"
	"fmt"
	"time"
	"github.com/puresoul/go-zabbix"
)

func main() {
	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix", time.Duration(60) * time.Second)
	if err != nil {
		fmt.Println(err)
	}

	v := zabbix.UsergroupRight{Permission: 0, ID: "2"}

    tmp1, err := session.CreateUsergroup(zabbix.UsergroupCreateParams{Name: "test", Rights: v, UserID: "3"})

	fmt.Println(tmp1,err)

	var t1 []zabbix.Usergroups
	t1 = append(t1, zabbix.Usergroups{UsergroupID: "84"})
	tmp2, err := session.CreateUser(zabbix.UserCreateParams{Alias: "test", Passwd: "password", Usergroup: t1})

	fmt.Println(tmp2,err)

    tmp3, err := session.CreateHostgroup(zabbix.HostgroupCreateParams{Name: "test"})

	fmt.Println(tmp3,err)

	var t4 []zabbix.CreateHostgroup
	var t5 zabbix.Templates

	t2 := zabbix.Template{TemplateID: "18"}
	t3 := zabbix.CreateHostInterface{Type: 1, Main: 1, Useip: 1, IP: "10.0.0.1", Port: "10050", DNS: ""}

	t4 = append(t4, zabbix.CreateHostgroup{GroupID: "1"})
	t5 = append(t5, t2)

	tmp4, err := session.CreateHosts(zabbix.HostCreateParams{Hosts: "test", Hostgroups: t4, Interfaces: t3, Templates: t5})

	fmt.Println(tmp4,err)
}
```

## License

Released under the [GNU GPL License](https://github.com/cavaliercoder/go-zabbix/blob/master/LICENSE)
