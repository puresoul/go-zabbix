package zabbix

import (
	"fmt"
)

// jTemplate is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/host/object
type jTemplate struct {
	TemplateID  string   `json:"hostid"`
	Hostname    string   `json:"host"`
	Description string   `json:"description,omitempty"`
	DisplayName string   `json:"name"`
	Macros      string   `json:"-"`
	Hosts       []string `json:"hosts"`
	Groups      string   `json:"groupids"`
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
