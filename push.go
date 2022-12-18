package chayns

import "fmt"

type alert struct {
	Alert string `json:"alert"`
}

// push message to group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Push#push-to-uac-group
func (c *conf) PushGroup(groupId int, msg string) (bool, error) {
	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/Push", alert{Alert: msg}, "POST"); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}

// push message to user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Push#push-to-user
func (c *conf) PushUser(userId int, msg string) (bool, error) {
	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Push", alert{Alert: msg}, "POST"); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
