package chayns

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
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
	if err := c.basicRequest(&result, "/AccessToken", permission{Data: pem}, post); err != nil {
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
	if err := c.bearerRequest(token, &result, path, nil, get); err != nil {
		return AccessToken{}, err
	}
	return result.Data[0], nil
}

func (c *Conf) ValidateAccessTokenAlt(token string) (bool, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("User-Agent", useragent).
		Head("https://auth.tobit.com/v2/token/validate")

	if resp.StatusCode() != 200 {
		return false, err
	}
	return true, nil
}