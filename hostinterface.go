package zabbix

import (
	"fmt"
)

const (
	HostInterfaceTypeAgent = 1
	HostInterfaceTypeSNMP = 2
	HostInterfaceTypeIPMI = 3
	HostInterfaceTypeJMX = 4
)

const (
	HostInterfaceNotDefault = 0
	HostInterfaceDefault = 1
)

const (
	HostInterfaceConnTypeDNS = 0
	HostInterfaceConnTypeIP = 1
)

const (
	HostInterfaceBulkDisabled = 0
	HostInterfaceBulkEnabled = 1
)

// HostInterface represents a Zabbix HostInterface returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/hostinterface/object
type HostInterface struct {
	// ID of the interface.
	InterfaceID string

	// DNS name used by the interface.
	// Can be empty if the connection is made via IP.
	DNS         string

	// IP address used by the interface.
	// Can be empty if the connection is made via DNS.
	IP          string

	// ID of the host the interface belongs to.
	HostID      string

	// Whether the interface is used as default on the host.
	//
	// Default must be HostInterfaceDefault or HostInterfaceNotDefault
	Default        int

	// Port number used by the interface.
	Port        string

	// Interface type.
	//
	// Type must be one of the HostInterfaceType constants.
	Type        int

	// Whether the connection should be made via IP.
	//
	// Type must be one of the HostInterfaceConnType constants.
	ConnectionType       int

	// Whether to use bulk SNMP requests.
	//
	// Type must be one of the HostInterfaceBulk constants.
	Bulk        int

	// Host that uses the interface.
	Host Host

	// TODO
	// Items []Item
}

// HostInterfaceGetParams represent the parameters for a `hostinterface.get` API call.
//
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/hostinterface/get
type HostInterfaceGetParams struct {
	GetParameters

	InterfaceIDs []string `json:"interfaceids,omitempty"`

	HostIDs []string `json:"hostids,omitempty"`

	ItemIDs []string `json:"itemids,omitempty"`

	TriggerIDs []string `json:"itemids,omitempty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`

	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

// GetHostInterfaces queries the Zabbix API for Host interfaces matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetHostInterfaces(params HostInterfaceGetParams) ([]HostInterface, error) {
	hostifs := make([]jHostInterface, 0)
	err := c.Get("hostinterface.get", params, &hostifs)
	if err != nil {
		return nil, err
	}

	if len(hostifs) == 0 {
		return nil, ErrNotFound
	}

	// map JSON HostInterfaces to Go HostInterfaces
	out := make([]HostInterface, len(hostifs))
	for i, jhostif := range hostifs {
		hostif, err := jhostif.HostInterface()
		if err != nil {
			return nil, fmt.Errorf("Error mapping HostInterface %d in response: %v", i, err)
		}

		out[i] = *hostif
	}

	return out, nil
}
