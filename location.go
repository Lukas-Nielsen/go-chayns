package chayns

type location struct {
	AppName           string `json:"appName"`
	Name              string `json:"name"`
	SiteID            string `json:"siteId"`
	LocationID        int    `json:"locationId"`
	CountMember       int    `json:"countMember"`
	AndroidAppVersion int    `json:"androidAppVersion"`
	IosAppVersion     int    `json:"iosAppVersion"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Location
func (c *Conf) Location() (location, error) {
	var result struct {
		Data []location `json:"data"`
	}
	if err := c.basicRequest(&result, "", nil, get); err != nil {
		return location{}, err
	}
	return result.Data[0], nil
}
