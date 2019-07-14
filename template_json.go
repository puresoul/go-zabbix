package zabbix

import (
	"fmt"
)

// jTemplate is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/object
type jTemplate struct {
ProxyHostID	string	`json:"proxy_hostid",omitempty`
Hostname	string	`json:"host",omitempty`
Hosts       []string `json:"hosts"`
Groups      string   `json:"groupids"`
Description string   `json:"description,omitempty"`
Macros      string   `json:"-"`
Status	string	`json:"status",omitempty`
DisableUntil	string	`json:"disable_until",omitempty`
Error	string	`json:"error",omitempty`
Available	string	`json:"available",omitempty`
ErrorsFrom	string	`json:"errors_from",omitempty`
Lastaccess	string	`json:"lastaccess",omitempty`
Ipmi_authtype	string	`json:"ipmi_authtype",omitempty`
Ipmi_privilege	string	`json:"ipmi_privilege",omitempty`
Ipmi_username	string	`json:"ipmi_username",omitempty`
Ipmi_password	string	`json:"ipmi_password",omitempty`
Ipmi_disable_until	string	`json:"ipmi_disable_until",omitempty`
Ipmi_available	string	`json:"ipmi_available",omitempty`
Snmp_disable_until	string	`json:"snmp_disable_until",omitempty`
Snmp_available	string	`json:"snmp_available",omitempty`
MaintenanceID	string	`json:"maintenanceid",omitempty`
Maintenance_status	string	`json:"maintenance_status",omitempty`
Maintenance_type	string	`json:"maintenance_type",omitempty`
Maintenance_from	string	`json:"maintenance_from",omitempty`
Ipmi_errors_from	string	`json:"ipmi_errors_from",omitempty`
Snmp_errors_from	string	`json:"snmp_errors_from",omitempty`
Ipmi_error	string	`json:"ipmi_error",omitempty`
Snmp_error	string	`json:"snmp_error",omitempty`
Jmx_disable_until	string	`json:"jmx_disable_until",omitempty`
Jmx_available	string	`json:"jmx_available",omitempty`
Jmx_errors_from	string	`json:"jmx_errors_from",omitempty`
Jmx_error	string	`json:"jmx_error",omitempty`
DisplayName	string	`json:"name",omitempty`
Flags	string	`json:"flags",omitempty`
TemplateID	string	`json:"templateid",omitempty`
TLS_connect	string	`json:"tls_connect",omitempty`
TLS_accept	string	`json:"tls_accept",omitempty`
TLS_issuer	string	`json:"tls_issuer",omitempty`
TLS_subject	string	`json:"tls_subject",omitempty`
TLS_psk_identity	string	`json:"tls_psk_identity",omitempty`
TLS_psk	string	`json:"tls_psk",omitempty`
}

// Host returns a native Go Host struct mapped from the given JSON Host data.
func (c *jTemplate) Template() (*Template, error) {
	//var err error

	template := &Template{}
	template.TemplateID = c.TemplateID
	template.Name = c.Hostname
	template.DisplayName = c.DisplayName
	template.Description = c.Description
	template.Macros = c.Macros
	template.Groups = c.Groups
	/*
		host.Source, err = strconv.Atoi(c.Flags)
		if err != nil {
			return nil, fmt.Errorf("Error parsing Host Flags: %v", err)
		}
	*/
	return template, nil
}

// jHosts is a slice of jHost structs.
type jTemplates []jTemplate

// Hosts returns a native Go slice of Templates mapped from the given JSON Temlates
// data.
func (c jTemplates) Tempaltes() ([]Template, error) {
	if c != nil {
		templates := make([]Template, len(c))
		for i, jTemplate := range c {
			template, err := jTemplate.Template()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Host %d in JSON data: %v", i, err)
			}

			templates[i] = *template
		}

		return templates, nil
	}

	return nil, nil
}
