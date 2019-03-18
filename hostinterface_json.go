package zabbix

import (
	"fmt"
	"strconv"
)

// jHostInterface is a private map for the Zabbix API HostInterface object.
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/hostinterface/object
type jHostInterface struct {
	InterfaceID string `json:"interfaceid"`
	DNS         string `json:"dns"`
	IP          string `json:"ip"`
	HostID      string `json:"hostid"`
	Main        int    `json:"main,string"`
	Port        string `json:"port"`
	Type        int    `json:"type,string"`
	UseIP       int	   `json:"useip,string"`
	Bulk        int    `json:"bulk,string"`
	Host 		jHosts `json:"hosts"`
}

// HostInterface returns a native Go HostInterface struct mapped from the given JSON HostInterface data.
func (c *jHostInterface) HostInterface() (*HostInterface, error) {
	hostIf := &HostInterface{}
	hostIf.InterfaceID = c.InterfaceID
	hostIf.DNS = c.DNS
	hostIf.IP = c.IP
	hostIf.HostID = c.HostID
	hostIf.Default = c.Main
	hostIf.Port = c.Port
	hostIf.Type = c.Type
	hostIf.ConnectionType = strconv.Itoa(c.UseIP)
	hostIf.Bulk = c.Bulk
	hosts, err := c.Host.Hosts()
	if err != nil {
		return nil, err
	}
	if len(hosts)>0 {
		hostIf.Host = hosts[0]
	}

	return hostIf, nil
}

// jHostInterfaces is a slice of jHostInterface structs.
type jHostInterfaces []jHostInterface

// HostInterfaces returns a native Go slice of Host interfaces mapped from the given JSON HostInterfaces
// data.
func (c jHostInterfaces) HostInterfaces() ([]HostInterface, error) {
	if c != nil {
		hostIfs := make([]HostInterface, len(c))
		for i, jhostif := range c {
			hostIf, err := jhostif.HostInterface()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling HostInterface %d in JSON data: %v", i, err)
			}

			hostIfs[i] = *hostIf
		}

		return hostIfs, nil
	}

	return nil, nil
}
