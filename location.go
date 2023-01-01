package chayns

// struct for location info
type location struct {
	AppName           string `json:"appName"`
	Name              string `json:"name"`
	SiteID            string `json:"siteId"`
	LocationID        int    `json:"locationId"`
	CountMember       int    `json:"countMember"`
	AndroidAppVersion int    `json:"androidAppVersion"`
	IosAppVersion     int    `json:"iosAppVersion"`
}

// getting the location info
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Location
func (c *Conf) Location() (location, error) {
	var result struct {
		Data []location `json:"data"`
	}
	if err := c.basicRequest(&result, "", nil, GET); err != nil {
		return location{}, err
	}
	return result.Data[0], nil
}
