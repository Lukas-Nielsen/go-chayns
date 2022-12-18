package chayns

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c *conf) basicRequest(result any, path string, data any, method string) error {
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

	url := baseUrl + version + "/" + fmt.Sprint(c.locationID) + path

	client := resty.New()

	req := client.R().
		SetHeader("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(fmt.Sprint(c.tappID)+":"+c.secret))).
		SetResult(&result).
		SetError(&respError)

	var err error
	err = nil
	switch method {
	case "POST":
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Post(url)

	case "PATCH":
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Patch(url)

	case "GET":
		_, err = req.
			Get(url)

	case "DELETE":
		_, err = req.
			Delete(url)
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}

func (c *conf) bearerRequest(token string, result any, path string, data any, method string) error {
	if c.locationID == 0 {
		return fmt.Errorf("'locationID' must not be empty")
	}
	var respError respError

	url := baseUrl + version + "/" + fmt.Sprint(c.locationID) + path

	client := resty.New()

	req := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetResult(&result).
		SetError(&respError)

	var err error
	err = nil
	switch method {
	case "POST":
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Post(url)

	case "GET":
		_, err = req.
			Get(url)
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}
