package chayns

import "fmt"

// get data of device
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Device
func (c *Conf) Device(deviceId int) (device, error) {
	var result struct {
		Data []device `json:"data"`
	}
	if err := c.basicRequest(&result, "/Device/"+fmt.Sprint(deviceId), nil, GET); err != nil {
		return device{}, err
	}
	return result.Data[0], nil
}
