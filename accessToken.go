package chayns

import (
	"fmt"
	"strings"
	"time"
)

type permission struct {
	Data []string `json:"permissions"`
}

type AccessToken struct {
	LocationID     int       `json:"locationId"`
	DeveloperID    int       `json:"developerId"`
	TappID         int       `json:"tappId"`
	Permissions    []string  `json:"permissions"`
	UserID         int       `json:"userId"`
	FacebookUserID string    `json:"facevookUserId"`
	PersonID       string    `json:"personId"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Expires        time.Time `json:"expires"`
	TokenType      struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"tokenType"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Authorization
func (c *Conf) NewPageAccessToken(pem ...string) (string, error) {
	var result struct {
		Data []struct {
			PageAccessToken string `json:"pageAccessToken"`
		} `json:"data,omitempty"`
	}
	if err := c.basicRequest(&result, "/AccessToken", permission{Data: pem}, POST); err != nil {
		return "", err
	}
	return result.Data[0].PageAccessToken, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-AccessToken#read-accesstoken
func (c *Conf) ValidateAccessToken(token string, uac ...int) (AccessToken, error) {
	path := "/AccessToken"
	if len(uac) > 0 {
		path += "?RequiredUacGroups=" + strings.Join(strings.Split(fmt.Sprint(uac), " "), ",")
	}

	var result struct {
		Data []AccessToken `json:"data"`
	}
	if err := c.bearerRequest(token, &result, path, nil, GET); err != nil {
		return AccessToken{}, err
	}
	return result.Data[0], nil
}
