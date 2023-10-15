package chayns

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type permission struct {
	Data []string `json:"permissions"`
}

type AccessTokenServer struct {
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

type AccessTokenLocal struct {
	Jti         string    `json:"jti,omitempty"`
	Sub         string    `json:"sub,omitempty"`
	Type        int       `json:"type,omitempty"`
	Exp         time.Time `json:"exp,omitempty"`
	Iat         time.Time `json:"iat,omitempty"`
	LocationID  int       `json:"LocationID,omitempty"`
	SiteID      string    `json:"SiteID,omitempty"`
	IsAdmin     bool      `json:"IsAdmin,omitempty"`
	TobitUserID int       `json:"TobitUserID,omitempty"`
	PersonID    string    `json:"PersonID,omitempty"`
	FirstName   string    `json:"FirstName,omitempty"`
	LastName    string    `json:"LastName,omitempty"`
	Prov        int       `json:"prov,omitempty"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Authorization
func (c *Client) NewPageAccessToken(pem ...string) (string, error) {
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
func (c *Client) ValidateAccessToken(token string, uac ...int) (AccessTokenServer, error) {
	path := "/AccessToken"
	if len(uac) > 0 {
		path += "?RequiredUacGroups=" + strings.Join(strings.Split(fmt.Sprint(uac), " "), ",")
	}

	var result struct {
		Data []AccessTokenServer `json:"data"`
	}
	if err := c.bearerRequest(token, &result, path, nil, get); err != nil {
		return AccessTokenServer{}, err
	}
	return result.Data[0], nil
}

func ValidateAccessTokenAlt(token string) (bool, error) {
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

func InspectAccesstoken(token string) (AccessTokenLocal, error) {
	tokenSplit := strings.Split(token, ".")
	var tokenData AccessTokenLocal
	tokenDataStr, _ := base64.StdEncoding.DecodeString(tokenSplit[1])

	if err := json.Unmarshal(tokenDataStr, &tokenData); err != nil {
		return AccessTokenLocal{}, err
	}

	return tokenData, nil
}
