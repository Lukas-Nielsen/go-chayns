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
func (c *Client) Groups(filter ...map[string]string) ([]UAC, error) {
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
	if err := c.basicRequest(&result, "/UAC"+filterString, nil, get); err != nil {
		return []UAC{}, err
	}
	return result.Data, nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#get-uac-group
func (c *Client) Group(groupId int) (UAC, error) {
	var result struct {
		Data []UAC `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), nil, get); err != nil {
		return UAC{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#create-uac-group
func (c *Client) NewGroup(name string, showName string) (UAC, error) {
	if len(name) == 0 {
		return UAC{}, fmt.Errorf("'name' must not be empty")
	}
	if len(showName) == 0 {
		return UAC{}, fmt.Errorf("'showName' must not be empty")
	}

	var result struct {
		Data []UAC `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC", groupData{ShowName: showName, Name: name}, post); err != nil {
		return UAC{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#modify-uac-group
func (c *Client) ModifyGroup(groupId int, name string, showName string) (UAC, error) {
	var result struct {
		Data []UAC `json:"data"`
	}
	if err := c.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), groupData{ShowName: showName, Name: name}, patch); err != nil {
		return UAC{}, err
	}
	return result.Data[0], nil
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#delete-uac-group
func (c *Client) DeleteGroup(groupId int) (bool, error) {
	if err := c.basicRequest(nil, "/UAC/"+fmt.Sprint(groupId), nil, delete); err != nil {
		return false, err
	}
	return true, nil
}
