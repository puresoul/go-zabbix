package zabbix

import (
    "fmt"
)

type UserCreateParams struct {
	GetParameters
	Alias      string `json:"alias,omitempty"`
	Passwd     string `json:"passwd,omitempty"`
	Usergroup  []Usergroups `json:"usrgrps,omitempty"`
}

type UserResponse struct {
	UserIDs []string `json:"Userids"`
}

// CreateUsers creates a single or multiple new Users.
// Returns a list of ids of created Users.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/User/create
func (c *Session) CreateUser(params UserCreateParams) (UserIds []string, err error) {
	var body UserResponse

	if err := c.Get("user.create", params, &body); err != nil {
		return nil, err
	}
	fmt.Println(body)
	if (body.UserIDs == nil) || (len(body.UserIDs) == 0) {
		return nil, ErrNotFound
	}

	return body.UserIDs, nil
}

// DeleteUsers method allows to delete Users.
// Returns a list of deleted Users ids.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/User/delete
func (c *Session) DeleteUsers(UserIDs ...string) (UserIds []string, err error) {
	var body UserResponse

	if err := c.Get("user.delete", UserIDs, &body); err != nil {
		return nil, err
	}

	if (body.UserIDs == nil) || (len(body.UserIDs) == 0) {
		return nil, ErrNotFound
	}

	return body.UserIDs, nil
}

// UserUpdateParams struct represents the Zabbix basic parameters for
// updating the User by Zabbix API
/*type UserUpdateParams struct {
    User
    Interfaces      UserInterfaces `json:"interfaces,omitempty"`
    Templates       Templates      `json:"templates,omitempty"`
    UnlinkTemplates Templates      `json:"templates,omitempty"`
    // Inventory Inventory `json:"inventory,omitempty"`
}

// GetUsers queries the Zabbix API for Users matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/User/get
func (c *Session) GetUsers(params UserGetParams) ([]User, error) {
	Users := make([]jUser, 0)
	err := c.Get("user.get", params, &Users)
	if err != nil {
		return nil, err
	}

	if len(Users) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]User, len(Users))
	for i, jUser := range Users {
		User, err := jUser.User()
		if err != nil {
			return nil, fmt.Errorf("Error mapping User %d in response: %v", i, err)
		}

		out[i] = *User
	}

	return out, nil
}*/

