package chayns

import (
	"fmt"
	"net/url"
)

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#get-all-uac-members
func (c *Client) Members(groupId int, filter ...map[string]string) ([]User, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []User `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User"+filterString, nil, get); err != nil {
		return []User{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#get-uac-member
func (c *Client) Member(groupId int, userId int) (User, error) {
	var result struct {
		Data []User `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User/"+fmt.Sprint(userId), nil, get); err != nil {
		return User{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#add-user-to-uac-group
func (c *Client) AddUserUac(groupId int, userId int) (User, error) {
	var result struct {
		Data []User `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User/"+fmt.Sprint(userId), nil, post); err != nil {
		return User{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group-Member#delete-user-from-uac-group
func (c *Client) RemoveUserUac(groupId int, userId int) (bool, error) {
	var result struct {
		Data []struct {
		} `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId)+"/User/"+fmt.Sprint(userId), nil, delete); err != nil {
		return false, err
	}
	return len(result.Data) == 0, nil
}
