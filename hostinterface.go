package zabbix

import (
	"fmt"
)

const (
	// HostInterfaceTypeAgent represents the host interface flag for Agent type
	HostInterfaceTypeAgent = 1
	// HostInterfaceTypeSNMP represents the host interface flag for SNMP type
	HostInterfaceTypeSNMP = 2
	// HostInterfaceTypeIPMI represents the host interface flag for IPMI type
	HostInterfaceTypeIPMI = 3
	// HostInterfaceTypeJMX represents the host interface flag for JMX type
	HostInterfaceTypeJMX = 4
)

const (
	// HostInterfaceNotDefault represents the host interface flag when it not delfaut
	HostInterfaceNotDefault = 0
	// HostInterfaceDefault represents the host interface flag when if delfaut
	HostInterfaceDefault = 1
)

const (
	// HostInterfaceConnTypeDNS represents the host interface type flag when
	// DNS name is default for performing the connection
	HostInterfaceConnTypeDNS = "0"
	// HostInterfaceConnTypeIP represents the host interface type flag when
	// IP address is default for performing the connection
	HostInterfaceConnTypeIP = "1"
)

const (
	// HostInterfaceBulkDisabled represents SNMP host interface parameter for
	// Disable state of SNMP Bulk polling
	HostInterfaceBulkDisabled = 0
	// HostInterfaceBulkEnabled represents SNMP host interface parameter for
	// Enable state of SNMP Bulk polling
	HostInterfaceBulkEnabled = 1
)

// HostInterface represents a Zabbix HostInterface returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/hostinterface/object
type HostInterface struct {
	// ID of the interface.
	InterfaceID string `json:"interfaceid,omitempty"`

	// DNS name used by the interface.
	// Can be empty if the connection is made via IP.
	DNS string `json:"dns,omitempty"`

	// IP address used by the interface.
	// Can be empty if the connection is made via DNS.
	IP string `json:"ip,omitempty"`

	// ID of the host the interface belongs to.
	HostID string `json:"hostid,omitempty"`

	// Whether the interface is used as default on the host.
	//
	// Default must be HostInterfaceDefault or HostInterfaceNotDefault
	Default int `json:"main,string,omitempty"`

	// Port number used by the interface.
	Port string `json:"port,omitempty"`

	// Interface type.
	//
	// Type must be one of the HostInterfaceType constants.
	Type int `json:"type,string,omitempty"`

	// Whether the connection should be made via IP.
	//
	// Type must be one of the HostInterfaceConnType constants.
	ConnectionType string `json:"useip,omitempty"`

	// Whether to use bulk SNMP requests.
	//
	// Type must be one of the HostInterfaceBulk constants.
	Bulk int `json:"bulk,string,omitempty"`

	// Host that uses the interface.
	Host Host `json:"hosts,omitempty"`

	// TODO
	// Items []Item `json:"items,omitempty"`
}

// HostInterfaces represents the slice of Host interfaces returned from Zabbix API
type HostInterfaces []HostInterface

// HostInterfaceGetParams represent the parameters for a `hostinterface.get` API call.
//
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/hostinterface/get
type HostInterfaceGetParams struct {
	GetParameters

	InterfaceIDs []string `json:"interfaceids,omitempty"`

	HostIDs []string `json:"hostids,omitempty"`

	ItemIDs []string `json:"itemids,omitempty"`

	TriggerIDs []string `json:"triggerids,omitempty"`

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
