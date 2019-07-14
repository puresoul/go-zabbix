package zabbix

import (
)

type MediaCreateParams struct {
	GetParameters
	User []Users `json:"users,omitempty"`
	Media     Medias       `json:"medias,omitempty"`
}

type Users struct {
UserID  string `json:"userid,omitempty"`
}


type Medias struct {
MediaTypeID  string `json:"mediatypeid"`
SendTo  string `json:"sendto"`
Active  int `json:"active"`
Severity  int `json:"severity"`
Period  string `json:"period"`
}


type MediaResponse struct {
	MediaIDs []string `json:"mediaids"`
}


func (c *Session) CreateMedia(params MediaCreateParams) ([]string, error) {
	var body MediaResponse

	if err := c.Get("user.addmedia", params, &body); err != nil {
		return []string{""}, err
	}

	if (body.MediaIDs == nil) || (len(body.MediaIDs) == 0) {
		return []string{""}, ErrNotFound
	}

	return body.MediaIDs, nil
}


func (c *Session) DeleteMedia(MediaIDs ...string) ([]string, error) {
	var body MediaResponse

	if err := c.Get("user.deletemedia", MediaIDs, &body); err != nil {
		return []string{""}, err
	}

	if (body.MediaIDs == nil) || (len(body.MediaIDs) == 0) {
		return []string{""}, ErrNotFound
	}

	return body.MediaIDs, nil
}

