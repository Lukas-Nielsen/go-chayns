package chayns

import (
	"fmt"
	"net/url"
)

type groupData struct {
	ShowName string `json:"showName"`
	Name     string `json:"name"`
}

// get data all groups
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#get-all-uac-groups
func (conf *Conf) Groups(filter ...map[string]string) ([]uac, error) {
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
	if err := conf.basicRequest(&result, "/UAC"+filterString, nil, "GET"); err != nil {
		return []uac{}, err
	}
	return result.Data, nil
}

// get data of given group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#get-uac-group
func (conf *Conf) Group(groupId int) (uac, error) {
	var result struct {
		Data []uac `json:"data"`
	}
	if err := conf.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), nil, "GET"); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// create a new group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#create-uac-group
func (conf *Conf) NewGroup(name string, showName string) (uac, error) {
	if len(name) == 0 {
		return uac{}, fmt.Errorf("'name' must not be empty")
	}
	if len(showName) == 0 {
		return uac{}, fmt.Errorf("'showName' must not be empty")
	}

	var result struct {
		Data []uac `json:"data"`
	}
	if err := conf.basicRequest(&result, "/UAC", groupData{ShowName: showName, Name: name}, "POST"); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// modify a group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#modify-uac-group
func (conf *Conf) ModifyGroup(groupId int, name string, showName string) (uac, error) {
	var result struct {
		Data []uac `json:"data"`
	}
	if err := conf.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), groupData{ShowName: showName, Name: name}, "PATCH"); err != nil {
		return uac{}, err
	}
	return result.Data[0], nil
}

// delete a group
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Group#delete-uac-group
func (conf *Conf) DeleteGroup(groupId int) (bool, error) {
	var result struct {
		Data []any `json:"data"`
	}
	if err := conf.basicRequest(&result, "/UAC/"+fmt.Sprint(groupId), nil, "DELETE"); err != nil {
		return false, err
	}
	return true, nil
}
