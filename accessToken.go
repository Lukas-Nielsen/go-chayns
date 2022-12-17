package chayns

import (
	"fmt"
	"strings"
	"time"
)

type permission struct {
	Data []string `json:"permissions"`
}

// struct for access token info
type accessToken struct {
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

// getting the page acces token based on the given permission
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Authorization
func (conf *Conf) NewPageAccessToken(pem ...string) (string, error) {
	var result struct {
		Data []struct {
			PageAccessToken string `json:"pageAccessToken"`
		} `json:"data,omitempty"`
	}
	if err := conf.basicRequest(&result, "/AccessToken", permission{Data: pem}, "POST"); err != nil {
		return "", err
	}
	return result.Data[0].PageAccessToken, nil
}

// validate the given access token (user or page)
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-AccessToken#read-accesstoken
func (conf *Conf) ValidateAccessToken(token string, uac ...int) (accessToken, error) {
	var result struct {
		Data []accessToken `json:"data"`
	}

	path := "/AccessToken"
	if len(uac) > 0 {
		path = "?RequiredUacGroups=" + strings.Join(strings.Split(fmt.Sprint(uac), " "), ",")
	}
	if err := conf.bearerRequest(token, &result, path, nil, "GET"); err != nil {
		return accessToken{}, err
	}
	return result.Data[0], nil
}
