package zabbix

import (
	"fmt"
)

type Usergroups struct {
        UsergroupID string `json:"usrgrpid,omitempty"`
}

type UsergroupResponse struct {
        UsergroupIDs string `json:"usrgrpids,omitempty"`
}

type UsergroupCreateParams struct {
	GetParameters
	Name string `json:"name,omitempty"`
    Rights UsergroupRight `json:"rights,omitempty"`
    UserID string `json:"userids,omitempty"`
}

type UsergroupRight struct {
	Permission int `json:"permission"`
	ID string `json:"id,omitempty"`
}

func (c *Session) CreateUsergroup(params UsergroupCreateParams) (UsertgroupIDs string, err error) {
	var body UsergroupResponse

	if err := c.Get("usergroup.create", params, &body); err != nil {
		return "", err
	}
	fmt.Println(body)
	if (len(body.UsergroupIDs) == 0) {
		return "", ErrNotFound
	}

	return body.UsergroupIDs, nil
}


// HostUpdateParams struct represents the Zabbix basic parameters for
// updating the host by Zabbix API
/*type UsergroupUpdateParams struct {
}