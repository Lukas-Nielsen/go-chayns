package callingcode

import (
	"fmt"
	"regexp"
)

type NewCC struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	Type       uint   `json:"type"`
	TypeString string `json:"typeString"`
	SiteID     string `json:"siteId"`
	Name       string `json:"name"`
}

type CC struct {
	c      *Client
	config map[string]any
}

func (c *Client) NewCC(sid string, t ccType, name string) *CC {
	return &CC{
		c: c,
		config: map[string]any{
			"siteId": sid,
			"type":   t,
			"name":   name,
			"appearance": map[string]any{
				"color":       "#005eb8",
				"captionText": "Informieren",
				"useCaption":  true,
			},
		},
	}
}

func (c *CC) Send() (NewCC, error) {
	var resp NewCC
	err := c.c.request(&resp, "", c.config, POST)
	return resp, err
}

// hex color
func (c *CC) SetColor(color string) *CC {
	if !regexp.MustCompile(`^#[0-9a-fA-F]{6}$`).MatchString(color) {
		panic("use a hex code like '#000000'")
	}
	appearance := c.config["appearance"].(map[string]any)

	appearance["color"] = color
	c.config["appearance"] = appearance

	return c
}

func (c *CC) SetIcon(icon icon) *CC {
	appearance := c.config["appearance"].(map[string]any)

	appearance["icon"] = string(icon)
	c.config["appearance"] = appearance

	return c
}

func (c *CC) SetCaptionText(captionText string) *CC {
	appearance := c.config["appearance"].(map[string]any)

	appearance["captionText"] = captionText

	if len(captionText) == 0 {
		appearance["useCaption"] = false
	} else {
		appearance["useCaption"] = true
	}

	c.config["appearance"] = appearance

	return c
}

func (c *CC) SetType(t ccType) *CC {
	c.config["type"] = t

	return c
}

func (c *CC) SetSite(sid string) *CC {
	c.config["siteId"] = sid

	return c
}

func (c *CC) SetName(name string) *CC {
	c.config["name"] = name

	return c
}

func (c *CC) SetUrl(url string) *CC {
	c.config["url"] = url

	return c
}

func (c *CC) SetTappId(tappId uint) *CC {
	c.config["tappId"] = tappId

	return c
}

func (c *CC) SetRequestUrl(requestUrl string) *CC {
	c.config["requestUrl"] = requestUrl

	return c
}

func (c *CC) SetMessageBefore(messageBefore string) *CC {
	c.config["messageBefore"] = messageBefore

	return c
}

func (c *CC) SetMessageFailure(messageFailure string) *CC {
	c.config["messageFailure"] = messageFailure

	return c
}

func (c *CC) SetMessageSuccess(messageSuccess string) *CC {
	c.config["messageSuccess"] = messageSuccess

	return c
}

func (c *CC) SetRequestBody(requestBody any) *CC {
	c.config["requestBody"] = requestBody

	return c
}

func (c *CC) SetDialogIframeUrl(dialogIframeUrl string) *CC {
	c.config["dialogIframeUrl"] = dialogIframeUrl

	return c
}

func (c *CC) SetDialogIframeConfig(dialogIframeConfig any) *CC {
	c.config["dialogIframeConfig"] = dialogIframeConfig

	return c
}

// site id, calling code, key, value
func (c *Client) UpdateCC(sid string, code string, key string, value any) error {
	var resp NewCC
	config := map[string]any{
		"siteId": sid,
		"code":   code,
		key:      value,
	}

	return c.request(&resp, "", config, PATCH)
}

// site id, calling code, color (hex)
func (c *Client) ChangeColor(sid string, code string, color string) error {
	if !regexp.MustCompile(`^#[0-9a-fA-F]{6}$`).MatchString(color) {
		return fmt.Errorf("use a hex code like '#000000'")
	}
	return c.UpdateCC(sid, code, "appearance", map[string]string{"color": color})
}

// site id, calling code, icon
func (c *Client) ChangeIcon(sid string, code string, icon icon) error {
	return c.UpdateCC(sid, code, "appearance", map[string]string{"icon": string(icon)})
}

// site id, calling code, captionText
func (c *Client) ChangeCaptiontext(sid string, code string, captionText string) error {
	config := map[string]any{
		"captionText": captionText,
	}
	if len(captionText) == 0 {
		config["useCaption"] = false
	} else {
		config["useCaption"] = true
	}

	return c.UpdateCC(sid, code, "appearance", config)
}

func (c *Client) DeleteCC(code string) error {
	var res any
	return c.request(&res, code, nil, DELETE)
}
