package zabbix

type Template struct {
	TemplateID  string `json:"templateid,omitempty"`
	Name        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"name,omitempty"`
}

type Templates []Template
