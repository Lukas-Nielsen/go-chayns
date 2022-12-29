package chayns

import (
	"fmt"
	"net/url"
)

type user struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Gender      string `json:"gender"`
	UserID      int    `json:"userId"`
	PersonID    string `json:"personId"`
	CountGroups int    `json:"countGroups"`
}

// get data of all users
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-all-users
func (c *conf) Users(filter ...map[string]string) ([]user, error) {
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
	if err := c.basicRequest(&result, "/User"+filterString, nil, GET); err != nil {
		return []user{}, err
	}
	return result.Data, nil
}

// get data of given user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-user
func (c *conf) User(userId int) (user, error) {
	var result struct {
		Data []user `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId), nil, GET); err != nil {
		return user{}, err
	}
	return result.Data[0], nil
}

// get uac groups of given user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-uac-groups
func (c *conf) UserUAC(userId int, filter ...map[string]string) ([]uac, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []uac `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/UAC"+filterString, nil, GET); err != nil {
		return []uac{}, err
	}
	return result.Data, nil
}

// get devices of given user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-all-devices-from-user
func (c *conf) UserDevices(userId int, filter ...map[string]string) ([]device, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []device `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Device"+filterString, nil, GET); err != nil {
		return []device{}, err
	}
	return result.Data, nil
}

// get given device of given user
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-device-from-user
func (c *conf) UserDevice(userId int, deviceId int) (device, error) {
	var result struct {
		Data []device `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Device/"+fmt.Sprint(deviceId), nil, GET); err != nil {
		return device{}, err
	}
	return result.Data[0], nil
}
