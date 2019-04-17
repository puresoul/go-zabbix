package zabbix

// Template structure for mapping template object fields with JSON
type Template struct {
	TemplateID  string `json:"templateid,omitempty"`
	Name        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"name,omitempty"`
	Macros      string `json:"macros,omitempty"`
	Groups      string `json:"groups,omitempty"`
}

// Templates type for number of template returned from Zabbix API
type Templates []Template
