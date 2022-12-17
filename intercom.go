package chayns

import "fmt"

// struct for location info
type Intercom struct {
	Message             string  `json:"Message"`
	UacIds              []int   `json:"UacIds,omitempty"`
	UserIds             []int   `json:"UserIds,omitempty"`
	ReceiverLocationIds []int   `json:"ReceiverLocationIds,omitempty"`
	ThreadName          string  `json:"ThreadName,omitempty"`
	UseGroupChat        bool    `json:"UseGroupChat"`
	UserAccessToken     string  `json:"UserAccessToken"`
	Images              []Image `json:"Images,omitempty"`
}

type Image struct {
	URL string `json:"url"`
}

// send a intercom message to given recipients or the current site
//
// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (conf *Conf) Intercom(data Intercom) (bool, error) {
	if len(data.Message) == 0 {
		return false, fmt.Errorf("'Message' must not be empty")
	}

	if len(data.UacIds) == 0 && len(data.UserIds) == 0 && len(data.UserAccessToken) == 0 {
		return false, fmt.Errorf("please provide one of the following 'UacID', 'UserID' or 'UserAccessToken'")
	}
	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := conf.basicRequest(&result, "/Intercom", data, "POST"); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
