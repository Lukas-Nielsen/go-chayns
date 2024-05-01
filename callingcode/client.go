package callingcode

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type config struct {
	token string
	debug bool
}

type Client struct {
	c config
}

// token = chayns user token
func NewClient(token string, debug bool) *Client {
	return &Client{
		c: config{
			token: token,
			debug: debug,
		},
	}
}

// token = chayns user token
func (c *Client) SetToken(token string) *Client {
	c.c.token = token
	return c
}

func (c *Client) request(result any, path string, data any, method method) error {
	if len(c.c.token) == 0 {
		return fmt.Errorf("'token' must not be empty")
	}
	var resp *resty.Response

	client := resty.New().
		SetBaseURL(BASE_URL)

	req := client.R().
		SetDebug(c.c.debug).
		SetAuthToken(c.c.token).
		SetHeader("User-Agent", UA).
		SetResult(&result)

	var err error
	err = nil

	switch method {
	case POST:
		if data != nil {
			req = req.
				SetBody(data)
		}
		resp, err = req.
			Post(path)

	case PATCH:
		if data != nil {
			req = req.
				SetBody(data)
		}
		resp, err = req.
			Patch(path)

	case GET:
		resp, err = req.
			Get(path)

	case DELETE:
		resp, err = req.
			Delete(path)
	}
	if err != nil {
		return err
	}

	if resp.IsError() {
		if len(resp.String()) == 0 {
			return fmt.Errorf(resp.Status())
		} else {
			return fmt.Errorf(resp.String())
		}
	}
	return err
}
