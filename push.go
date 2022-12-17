package chayns

import "fmt"

type alert struct {
	Alert string `json:"alert"`
}

// push message to group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Push#push-to-uac-group
func (conf *Conf) PushGroup(groupId int, msg string) (bool, error) {
	var result struct {
		Data []any `json:"data"`
	}

	if err := conf.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/Push", alert{Alert: msg}, "POST"); err != nil {
		return false, err
	}
	return true, nil
}

// push message to user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Push#push-to-user
func (conf *Conf) PushUser(userId int, msg string) (bool, error) {
	var result struct {
		Data []any `json:"data"`
	}

	if err := conf.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Push", alert{Alert: msg}, "POST"); err != nil {
		return false, err
	}
	return true, nil
}
