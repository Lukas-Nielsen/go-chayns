package chayns

import (
	"fmt"
	"net/url"
)

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Gender      string `json:"gender"`
	UserID      int    `json:"userId"`
	PersonID    string `json:"personId"`
	CountGroups int    `json:"countGroups"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-all-users
func (c *Conf) Users(filter ...map[string]string) ([]User, error) {
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
	if err := c.basicRequest(&result, "/User"+filterString, nil, GET); err != nil {
		return []User{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-user
func (c *Conf) User(userId int) (User, error) {
	var result struct {
		Data []User `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId), nil, GET); err != nil {
		return User{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-uac-groups
func (c *Conf) UserUAC(userId int, filter ...map[string]string) ([]UAC, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []UAC `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/UAC"+filterString, nil, GET); err != nil {
		return []UAC{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-all-devices-from-user
func (c *Conf) UserDevices(userId int, filter ...map[string]string) ([]Device, error) {
	filterString := ""
	if len(filter) == 1 && len(filter[0]) > 0 {
		params := url.Values{}
		for key, value := range filter[0] {
			params.Add(key, value)
		}
		filterString = "?" + params.Encode()
	}

	var result struct {
		Data []Device `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Device"+filterString, nil, GET); err != nil {
		return []Device{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-User#get-device-from-user
func (c *Conf) UserDevice(userId int, deviceId int) (Device, error) {
	var result struct {
		Data []Device `json:"data"`
	}
	if err := c.basicRequest(&result, "/User/"+fmt.Sprint(userId)+"/Device/"+fmt.Sprint(deviceId), nil, GET); err != nil {
		return Device{}, err
	}
	return result.Data[0], nil
}
