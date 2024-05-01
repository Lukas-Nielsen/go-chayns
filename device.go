package chayns

import "fmt"

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Device
func (c *Client) Device(deviceId uint) (Device, error) {
	var result struct {
		Data []Device `json:"data"`
	}
	if err := c.basicRequest(&result, "/Device/"+fmt.Sprint(deviceId), nil, get); err != nil {
		return Device{}, err
	}
	return result.Data[0], nil
}
