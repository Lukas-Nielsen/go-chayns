package chayns

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	post      string = "POST"
	get       string = "GET"
	patch     string = "PATCH"
	delete    string = "DELETE"
	useragent string = "github.com/Lukas-Nielsen/go-chayns"
)

func (c *Conf) getUrl(path string) string {
	return fmt.Sprint(c.locationID) + path
}

func (c *Conf) basicRequest(result any, path string, data any, method string) error {
	if c.locationID == 0 {
		return fmt.Errorf("'locationID' must not be empty")
	}
	if c.tappID == 0 {
		return fmt.Errorf("'tappID' must not be empty")
	}
	if len(c.secret) == 0 {
		return fmt.Errorf("'secret' must not be empty")
	}
	var respError respError

	client := resty.New().
		SetBaseURL(BASE_URL + VERSION + "/")

	req := client.R().
		SetHeader("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(fmt.Sprint(c.tappID)+":"+c.secret))).
		SetHeader("User-Agent", useragent).
		SetResult(&result).
		SetError(&respError)

	var err error
	err = nil
	switch method {
	case post:
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Post(c.getUrl(path))

	case patch:
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Patch(c.getUrl(path))

	case get:
		_, err = req.
			Get(c.getUrl(path))

	case delete:
		_, err = req.
			Delete(c.getUrl(path))
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}

func (c *Conf) bearerRequest(token string, result any, path string, data any, method string) error {
	if c.locationID == 0 {
		return fmt.Errorf("'locationID' must not be empty")
	}
	var respError respError

	client := resty.New().
		SetBaseURL(BASE_URL + VERSION + "/")

	req := client.R().
		SetAuthToken(token).
		SetHeader("User-Agent", useragent).
		SetResult(&result).
		SetError(&respError)

	var err error
	err = nil
	switch method {
	case post:
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Post(c.getUrl(path))

	case get:
		_, err = req.
			Get(c.getUrl(path))
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}
