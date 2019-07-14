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

type TemplateGetParams struct {
	GetParameters

	// GroupIDs filters search results to hosts that are members of the given
	// Group IDs.
	Output []string `json:"output,omitempty"`

	// ApplicationIDs filters search results to hosts that have items in the
	// given Application IDs.
	Filter Filter `json:"filter,omitempty"`
}

type Filter struct {
	Host []string `json:"host,omitempty"`
}

func (c *Session) GetTemplates(params TemplateGetParams) ([]Template, error) {
	templates := make([]jTemplate, 0)
	err := c.Get("template.get", params, &templates)
	if err != nil {
		return nil, err
	}

	if len(templates) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Template, len(templates))
	for i, jtemp := range templates {
		host, _ := jtemp.Template()
		out[i] = *host
	}

	return out, nil
}
