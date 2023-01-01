package chayns

type Conf struct {
	locationID int
	tappID     int
	secret     string
}

func NewConf(location int) *Conf {
	return &Conf{
		locationID: location,
	}
}

func (c *Conf) SetLocation(location int) *Conf {
	c.locationID = location
	return c
}

func (c *Conf) SetTapp(tapp int) *Conf {
	c.tappID = tapp
	return c
}

func (c *Conf) SetSecret(secret string) *Conf {
	c.secret = secret
	return c
}

type respError struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"errorGuid,omitempty"`
}

type uac struct {
	Name        string `json:"name"`
	ShowName    string `json:"showName"`
	UserGroupID int    `json:"userGroupId"`
	CountMember int    `json:"CountMember"`
	TappID      int    `json:"tappId"`
}

type device struct {
	SysName     string `json:"sysName"`
	SysVersion  string `json:"sysVersion"`
	LastRequest string `json:"lastRequest"`
	DeviceToken string `json:"deviceToken"`
	UDID        string `json:"udid"`
	DeviceID    int    `json:"deviceId"`
	AppVersion  int    `json:"appVersion"`
}
