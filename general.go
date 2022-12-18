package chayns

type conf struct {
	locationID int
	tappID     int
	secret     string
}

// create new conf
func NewConf(location int) *conf {
	return &conf{
		locationID: location,
	}
}

// set locationId of conf
func (c *conf) SetLocation(location int) *conf {
	c.locationID = location
	return c
}

// set tappId of conf
func (c *conf) SetTapp(tapp int) *conf {
	c.tappID = tapp
	return c
}

// set secret of conf
func (c *conf) SetSecret(secret string) *conf {
	c.secret = secret
	return c
}

// struct for the error messages
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
