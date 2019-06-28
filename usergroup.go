package zabbix

import (
//	"fmt"
)

type Usergroups struct {
        UsergroupID string `json:"usrgrpid,omitempty"`
}


// HostUpdateParams struct represents the Zabbix basic parameters for
// updating the host by Zabbix API
/*type HostUpdateParams struct {
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
	fmt.Println(body)
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
*/