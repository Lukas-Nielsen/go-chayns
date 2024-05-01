package chayns

type location struct {
	AppName           string `json:"appName"`
	Name              string `json:"name"`
	SiteID            string `json:"siteId"`
	LocationID        uint   `json:"locationId"`
	CountMember       uint   `json:"countMember"`
	AndroidAppVersion uint   `json:"androidAppVersion"`
	IosAppVersion     uint   `json:"iosAppVersion"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Location
func (c *Client) Location() (location, error) {
	var result struct {
		Data []location `json:"data"`
	}
	if err := c.basicRequest(&result, "", nil, get); err != nil {
		return location{}, err
	}
	return result.Data[0], nil
}
