package callingcode

import "fmt"

type group struct {
	Name string `json:"showName"`
	SID  string `json:"siteId"`
}

type Group struct {
	ID       int    `json:"id,omitempty"`
	SiteID   string `json:"siteId,omitempty"`
	ShowName string `json:"showName,omitempty"`
}

// name = display nameof group
//
// sid = chayns site id
func (c *Client) NewGroup(name string, sid string) (Group, error) {
	var res Group
	err := c.request(&res, "Group", group{Name: name, SID: sid}, POST)
	return res, err
}

// gid = id of group
//
// name = display nameof group
//
// sid = chayns site id
func (c *Client) RenameGroup(gid int, name string, sid string) (Group, error) {
	var res Group
	err := c.request(&res, fmt.Sprintf("Group/%d", gid), group{Name: name, SID: sid}, PATCH)
	return res, err
}

// gid = id of group
func (c *Client) DeleteGroup(gid int, deleteCodes bool) error {
	var res Group
	return c.request(&res, fmt.Sprintf("Group/%d?deleteCodes=%t", gid, deleteCodes), nil, DELETE)
}
