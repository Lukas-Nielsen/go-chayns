package chayns

import "fmt"

type Intercom struct {
	client              *Client
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
func (c *Client) NewIntercom(msg string) *Intercom {
	return &Intercom{
		client:         c,
		Message:      msg,
		UseGroupChat: false,
	}
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) AddGroup(groupId ...int) *Intercom {
	i.UacIds = append(i.UacIds, groupId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) AddUser(userId ...int) *Intercom {
	i.UserIds = append(i.UserIds, userId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) AddLocation(locationId ...int) *Intercom {
	i.ReceiverLocationIds = append(i.ReceiverLocationIds, locationId...)
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) AddImage(url ...string) *Intercom {
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
func (i *Intercom) SetThreadName(name string) *Intercom {
	i.ThreadName = name
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) SetGroupChat(group bool) *Intercom {
	i.UseGroupChat = group
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) SetAccessToken(token string) *Intercom {
	i.UserAccessToken = token
	return i
}

// https://github.com/TobitSoftware/chayns-backend/wiki/Reference-Intercom
func (i *Intercom) Send() (bool, error) {
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
	if err := i.client.basicRequest(&result, "/Intercom", i, post); err != nil {
		return false, err
	}
	return result.Data[0].Success, nil
}
