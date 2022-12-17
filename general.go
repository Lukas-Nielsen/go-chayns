package chayns

type Conf struct {
	LocationID int
	TappID     int
	Secret     string
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
