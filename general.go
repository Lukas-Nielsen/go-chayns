package chayns

type config struct {
	locationID uint
	tappID     uint
	secret     string
}

type Client struct {
	c config
}

func NewClient(location uint) *Client {
	return &Client{
		c: config{
			locationID: location,
		},
	}
}

func (c *Client) SetLocation(location uint) *Client {
	c.c.locationID = location
	return c
}

func (c *Client) SetTapp(tapp uint) *Client {
	c.c.tappID = tapp
	return c
}

func (c *Client) SetSecret(secret string) *Client {
	c.c.secret = secret
	return c
}

type respError struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"errorGuid,omitempty"`
}

type UAC struct {
	Name        string `json:"name"`
	ShowName    string `json:"showName"`
	UserGroupID uint   `json:"userGroupId"`
	CountMember uint   `json:"CountMember"`
	TappID      uint   `json:"tappId"`
}

type Device struct {
	SysName     string `json:"sysName"`
	SysVersion  string `json:"sysVersion"`
	LastRequest string `json:"lastRequest"`
	DeviceToken string `json:"deviceToken"`
	UDID        string `json:"udid"`
	DeviceID    uint   `json:"deviceId"`
	AppVersion  uint   `json:"appVersion"`
}
