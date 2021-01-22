package zabbix

import ()

type MediaCreateParams struct {
	GetParameters
	UserID string   `json:"userid,omitempty"`
	Media  []Medias `json:"user_medias,omitempty"`
}

type Medias struct {
	MediaTypeID string `json:"mediatypeid"`
	SendTo      string `json:"sendto"`
	Active      int    `json:"active"`
	Severity    int    `json:"severity"`
	Period      string `json:"period"`
}

type MediaResponse struct {
	MediaIDs []string `json:"mediaids"`
}

func (c *Session) UpdateMedia(params MediaCreateParams) (error) {
	var body MediaResponse

	if err := c.Get("user.update", params, &body); err != nil {
		return err
	}

	return nil
}

