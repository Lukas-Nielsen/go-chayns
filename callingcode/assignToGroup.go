package callingcode

import "fmt"

// gid = id of group
//
// cid = id of callingcode
func (c *Client) AssignToGroup(gid int, cid string) error {
	var res any
	return c.request(&res, fmt.Sprintf("Group/%d/%s", gid, cid), nil, POST)
}
