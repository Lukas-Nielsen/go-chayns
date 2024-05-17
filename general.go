package chayns

type config struct {
	locationID int
	tappID     int
	secret     string
}

type Client struct {
	c config
}

func NewClient(location int) *Client {
	return &Client{
		c: config{
			locationID: location,
		},
	}
}

func (c *Client) SetLocation(location int) *Client {
	c.c.locationID = location
	return c
}

func (c *Client) SetTapp(tapp int) *Client {
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
	UserGroupID int    `json:"userGroupId"`
	CountMember int    `json:"CountMember"`
	TappID      int    `json:"tappId"`
}

type Device struct {
	SysName     string `json:"sysName"`
	SysVersion  string `json:"sysVersion"`
	LastRequest string `json:"lastRequest"`
	DeviceToken string `json:"deviceToken"`
	UDID        string `json:"udid"`
	DeviceID    int    `json:"deviceId"`
	AppVersion  int    `json:"appVersion"`
}
