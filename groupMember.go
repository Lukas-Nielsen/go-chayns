package chayns

import (
	"fmt"
	"net/url"
)

// get all members of a group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#get-all-uac-members
func (c *conf) Members(groupId int, filter ...map[string]string) ([]user, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []user `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User"+filterString, nil, GET); err != nil {
		return []user{}, err
	}
	return result.Data, nil
}

// get data of given user in given group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#get-uac-member
func (c *conf) Member(groupId int, userId int) (user, error) {
	var result struct {
		Data []user `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User/"+fmt.Sprint(userId), nil, GET); err != nil {
		return user{}, err
	}
	return result.Data[0], nil
}

// add user to group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#add-user-to-uac-group
func (c *conf) AddUserUac(groupId int, userId int) (user, error) {
	var result struct {
		Data []user `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC", nil, POST); err != nil {
		return user{}, err
	}
	return result.Data[0], nil
}

// remove user from group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#delete-user-from-uac-group
func (c *conf) RemoveUserUac(groupId int, userId int) (bool, error) {
	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User/"+fmt.Sprint(userId), nil, DELETE); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
