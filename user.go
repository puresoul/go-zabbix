package zabbix

import (
	"fmt"
)

type UserCreateParams struct {
	GetParameters
	Alias     string       `json:"alias,omitempty"`
	Passwd    string       `json:"passwd,omitempty"`
	Usergroup []Usergroups `json:"usrgrps,omitempty"`
}

type UserResponse struct {
	UserIDs []string `json:"userids"`
}

// CreateUsers creates a single or multiple new Users.
// Returns a list of ids of created Users.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/User/create
func (c *Session) CreateUser(params UserCreateParams) ([]string, error) {
	var body UserResponse

	if err := c.Get("user.create", params, &body); err != nil {
		return []string{""}, err
	}
	fmt.Println(body)
	if (body.UserIDs == nil) || (len(body.UserIDs) == 0) {
		return []string{""}, ErrNotFound
	}

	return body.UserIDs, nil
}

// DeleteUsers method allows to delete Users.
// Returns a list of deleted Users ids.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/User/delete
func (c *Session) DeleteUsers(UserIDs ...string) ([]string, error) {
	var body UserResponse

	if err := c.Get("user.delete", UserIDs, &body); err != nil {
		return []string{""}, err
	}

	if (body.UserIDs == nil) || (len(body.UserIDs) == 0) {
		return []string{""}, ErrNotFound
	}

	return body.UserIDs, nil
}

// UserUpdateParams struct represents the Zabbix basic parameters for
// updating the User by Zabbix API
/*
type UserUpdateParams struct {
}
*/
