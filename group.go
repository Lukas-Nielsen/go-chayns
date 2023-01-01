package chayns

import (
	"fmt"
	"net/url"
)

type groupData struct {
	ShowName string `json:"showName"`
	Name     string `json:"name"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#get-all-uac-groups
func (c *Conf) Groups(filter ...map[string]string) ([]uac, error) {
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
	if err := c.basicRequest(&result, "/UAC"+filterString, nil, GET); err != nil {
		return []uac{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#get-uac-group
func (c *Conf) Group(groupId int) (uac, error) {
	var result struct {
		Data []uac `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), nil, GET); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#create-uac-group
func (c *Conf) NewGroup(name string, showName string) (uac, error) {
	if len(name) == 0 {
		return uac{}, fmt.Errorf("'name' must not be empty")
	}
	if len(showName) == 0 {
		return uac{}, fmt.Errorf("'showName' must not be empty")
	}

	var result struct {
		Data []uac `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC", groupData{ShowName: showName, Name: name}, POST); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#modify-uac-group
func (c *Conf) ModifyGroup(groupId int, name string, showName string) (uac, error) {
	var result struct {
		Data []uac `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), groupData{ShowName: showName, Name: name}, PATCH); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#delete-uac-group
func (c *Conf) DeleteGroup(groupId int) (bool, error) {
	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), nil, DELETE); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
