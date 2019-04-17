package zabbix

import (
	"os"
	"testing"
)

func TestCreateHosts(t *testing.T) {
	session := GetTestSession(t)
	os.Setenv("ZBX_DEBUG", "1")

	wantname := "testgo"
	wantdispname := "GO_host"
	wantgroupid := Hostgroup{GroupID: "2"}

	listhostgroup := []Hostgroup{wantgroupid}

	wanthostinterface := HostInterface{
		Type:    1,
		Default: 1,
		IP:      "192.168.33.11",
		DNS:     "",
		Port:    "10050",
	}
	listhostinterfaces := []HostInterface{wanthostinterface}

	params := HostCreateParams{
		Host: Host{
			Hostname:    wantname,
			DisplayName: wantdispname,
			Groups:      listhostgroup,
		},
		Interfaces: listhostinterfaces,
	}
	hostid, err := session.CreateHosts(params)
	t.Error(hostid, err)
}
