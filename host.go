package zabbix

import (
	"fmt"
)

const (
	// HostSourceDefault indicates that a Host was created in the normal way.
	HostSourceDefault = 0

	// HostSourceDiscovery indicates that a Host was created by Host discovery.
	HostSourceDiscovery = 4
)

const (
	HostStatusMonitored   = "0"
	HostStatusUnmonitored = "1"
)

const (
	HostEncryptionDisabled = 1
	HostEncryptionPSK      = 2
	HostEncryptionCert     = 4
)

// Host represents a Zabbix Host returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/2.2/manual/config/hosts
type Host struct {
	// HostID is the unique ID of the Host.
	HostID string `json:"hostid,omitempty"`

	// Hostname is the technical name of the Host.
	Hostname string `json:"host,omitempty"`

	// DisplayName is the visible name of the Host.
	DisplayName string `json:"name,omitempty"`

	// Source is the origin of the Host and must be one of the HostSource
	// constants.
	Source int `json:"flags,string,omitempty"`

	// Macros contains all Host Macros assigned to the Host.
	Macros HostMacros `json:"macros,omitempty"`

	// Groups contains all Host Groups assigned to the Host.
	Groups Hostgroups `json:"groups,omitempty"`

	// Description of the host.
	Description string `json:"description,omitempty"`

	// Status and function of the host.
	//
	// Status must be one of the HostStatus constants.
	Status string `json:"status,omitempty"`

	// ProxyID is ID of the proxy that is used to monitor the host.
	ProxyID string `json:"proxy_hostid,omitempty"`

	// Connections to host.
	//
	// TLSConnect must be one of the HostEncryption constants.
	TLSConnect int `json:"tls_connect,string,omitempty"`

	// Connections from host.
	//
	// TLSAccept must be one of the HostEncryption constants.
	TLSAccept int `json:"tls_accept,string,omitempty"`

	// Certificate issuer.
	TLSIssuer string `json:"tls_issuer,omitempty"`

	// Certificate subject.
	TLSSubject string `json:"tls_subject,omitempty"`

	// PSK identity. Required if either TLSConnect or TLSAccept has PSK enabled.
	TLSPSKIdentity string `json:"tls_psk_identity,omitempty"`

	// The preshared key. Required if either TLSConnect or TLSAccept has PSK enabled.
	TLSPSK string `json:"tls_psk,omitempty"`
}

type Hosts []Host

// HostGetParams represent the parameters for a `host.get` API call.
//
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/get#parameters
type HostGetParams struct {
	GetParameters

	// GroupIDs filters search results to hosts that are members of the given
	// Group IDs.
	GroupIDs []string `json:"groupids,omitempty"`

	// ApplicationIDs filters search results to hosts that have items in the
	// given Application IDs.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// DiscoveredServiceIDs filters search results to hosts that are related to
	// the given discovered service IDs.
	DiscoveredServiceIDs []string `json:"dserviceids,omitempty"`

	// GraphIDs filters search results to hosts that have the given graph IDs.
	GraphIDs []string `json:"graphids,omitempty"`

	// HostIDs filters search results to hosts that matched the given Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// WebCheckIDs filters search results to hosts with the given Web Check IDs.
	WebCheckIDs []string `json:"httptestids,omitempty"`

	// InterfaceIDs filters search results to hosts that use the given Interface
	// IDs.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// ItemIDs filters search results to hosts with the given Item IDs.
	ItemIDs []string `json:"itemids,omitempty"`

	// MaintenanceIDs filters search results to hosts that are affected by the
	// given Maintenance IDs
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`

	// MonitoredOnly filters search results to return only monitored hosts.
	MonitoredOnly bool `json:"monitored_hosts,omitempty"`

	// ProxyOnly filters search results to hosts which are Zabbix proxies.
	ProxiesOnly bool `json:"proxy_host,omitempty"`

	// ProxyIDs filters search results to hosts monitored by the given Proxy
	// IDs.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// IncludeTemplates extends search results to include Templates.
	IncludeTemplates bool `json:"templated_hosts,omitempty"`

	// SelectGroups causes the Host Groups that each Host belongs to to be
	// attached in the search results.
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`

	// SelectApplications causes the Applications from each Host to be attached
	// in the search results.
	SelectApplications SelectQuery `json:"selectApplications,omitempty"`

	// SelectDiscoveries causes the Low-Level Discoveries from each Host to be
	// attached in the search results.
	SelectDiscoveries SelectQuery `json:"selectDiscoveries,omitempty"`

	// SelectDiscoveryRule causes the Low-Level Discovery Rule that created each
	// Host to be attached in the search results.
	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`

	// SelectGraphs causes the Graphs from each Host to be attached in the
	// search results.
	SelectGraphs SelectQuery `json:"selectGraphs,omitempty"`

	SelectHostDiscovery SelectQuery `json:"selectHostDiscovery,omitempty"`

	SelectWebScenarios SelectQuery `json:"selectHttpTests,omitempty"`

	SelectInterfaces SelectQuery `json:"selectInterfaces,omitempty"`

	SelectInventory SelectQuery `json:"selectInventory,omitempty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`

	SelectMacros SelectQuery `json:"selectMacros,omitempty"`

	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectScreens         SelectQuery `json:"selectScreens,omitempty"`
	SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"`
}

type HostCreateParams struct {
	Host
	Interfaces HostInterfaces `json:"interfaces"`
	Templates  Templates      `json:"templates,omitempty"`
	// Inventory Inventory `json:"inventory,omitempty"`
}

type HostUpdateParams struct {
	Host
	Interfaces      HostInterfaces `json:"interfaces,omitempty"`
	Templates       Templates      `json:"templates,omitempty"`
	UnlinkTemplates Templates      `json:"templates,omitempty"`
	// Inventory Inventory `json:"inventory,omitempty"`
}

// HostResponse represent host action response body
type HostResponse struct {
	HostIDs []string `json:"hostids"`
}

// GetHosts queries the Zabbix API for Hosts matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/host/get
func (c *Session) GetHosts(params HostGetParams) ([]Host, error) {
	hosts := make([]jHost, 0)
	err := c.Get("host.get", params, &hosts)
	if err != nil {
		return nil, err
	}

	if len(hosts) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Host, len(hosts))
	for i, jhost := range hosts {
		host, err := jhost.Host()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Host %d in response: %v", i, err)
		}

		out[i] = *host
	}

	return out, nil
}

// CreateHosts creates a single or multiple new hosts.
// Returns a list of ids of created hosts.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/host/create
func (c *Session) CreateHosts(params ...HostCreateParams) (hostIds []string, err error) {
	var body HostResponse

	if err := c.Get("host.create", params, &body); err != nil {
		return nil, err
	}

	if (body.HostIDs == nil) || (len(body.HostIDs) == 0) {
		return nil, ErrNotFound
	}

	return body.HostIDs, nil
}

// DeleteHosts method allows to delete hosts.
// Returns a list of deleted hosts ids.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/host/delete
func (c *Session) DeleteHosts(hostIDs ...string) (hostIds []string, err error) {
	var body HostResponse

	if err := c.Get("host.delete", hostIDs, &body); err != nil {
		return nil, err
	}

	if (body.HostIDs == nil) || (len(body.HostIDs) == 0) {
		return nil, ErrNotFound
	}

	return body.HostIDs, nil
}

// UpdateHosts method allows to update hosts.
// Returns a list of updated hosts ids.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/host/update
func (c *Session) UpdateHosts(params ...HostUpdateParams) (hostIds []string, err error) {
	var body HostResponse

	if err := c.Get("host.update", params, &body); err != nil {
		return nil, err
	}

	if (body.HostIDs == nil) || (len(body.HostIDs) == 0) {
		return nil, ErrNotFound
	}

	return body.HostIDs, nil
}
