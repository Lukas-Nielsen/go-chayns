package chayns

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (conf *Conf) basicRequest(result any, path string, data any, method string) error {
	if conf.LocationID == 0 {
		return fmt.Errorf("'LocationID' must not be empty")
	}
	if conf.TappID == 0 {
		return fmt.Errorf("'TappID' must not be empty")
	}
	if len(conf.Secret) == 0 {
		return fmt.Errorf("'Secret' must not be empty")
	}
	var respError respError
	client := resty.New()

	req := client.R().
		SetHeader("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(fmt.Sprint(conf.TappID)+":"+conf.Secret))).
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
			Post(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)

	case "PATCH":
		if data != nil {
			req = req.
				SetBody(data)
		}
		_, err = req.
			Patch(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)

	case "GET":
		_, err = req.
			Get(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)

	case "DELETE":
		_, err = req.
			Delete(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}

func (conf *Conf) bearerRequest(token string, result any, path string, data any, method string) error {
	if conf.LocationID == 0 {
		return fmt.Errorf("'LocationID' must not be empty")
	}
	var respError respError
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
			Post(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)

	case "GET":
		_, err = req.
			Get(baseUrl + version + "/" + fmt.Sprint(conf.LocationID) + path)
	}
	if err != nil {
		return err
	}

	if len(respError.Message) != 0 {
		return fmt.Errorf(respError.Message)
	}
	return err
}
