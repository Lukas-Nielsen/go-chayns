package chayns

import "fmt"

type intercom struct {
	*Conf
	Message             string  `json:"Message"`
	UacIds              []int   `json:"UacIds,omitempty"`
	UserIds             []int   `json:"UserIds,omitempty"`
	ReceiverLocationIds []int   `json:"ReceiverLocationIds,omitempty"`
	ThreadName          string  `json:"ThreadName,omitempty"`
	UseGroupChat        bool    `json:"UseGroupChat"`
	UserAccessToken     string  `json:"UserAccessToken,omitempty"`
	Images              []image `json:"Images,omitempty"`
}

type image struct {
	URL string `json:"url"`
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (c *Conf) NewIntercom(msg string) *intercom {
	return &intercom{
		Message:      msg,
		UseGroupChat: false,
	}
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) AddGroup(groupId ...int) *intercom {
	i.UacIds = append(i.UacIds, groupId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) AddUser(userId ...int) *intercom {
	i.UserIds = append(i.UserIds, userId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) AddLocation(locationId ...int) *intercom {
	i.ReceiverLocationIds = append(i.ReceiverLocationIds, locationId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) AddImage(url ...string) *intercom {
	var images []image
	if len(url) > 0 {
		for _, entry := range url {
			images = append(images, image{URL: entry})
		}
		i.Images = append(i.Images, images...)
	}
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) SetThreadName(name string) *intercom {
	i.ThreadName = name
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) SetGroupChat(group bool) *intercom {
	i.UseGroupChat = group
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) SetAccessToken(token string) *intercom {
	i.UserAccessToken = token
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *intercom) Send() (bool, error) {
	if len(i.Message) == 0 {
		return false, fmt.Errorf("'Message' must not be empty")
	}

	if len(i.UacIds) == 0 && len(i.UserIds) == 0 && len(i.UserAccessToken) == 0 {
		return false, fmt.Errorf("please provide one of the following 'UacID', 'UserID' or 'UserAccessToken'")
	}

	var result struct {
		Data []struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := i.basicRequest(&result, "/Intercom", i, POST); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
