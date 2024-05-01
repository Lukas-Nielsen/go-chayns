# go-chayns

package to wrap the api of [chanys®](https://chayns.net/), you can find the documentation of the api at the repository of TobitSoftware [GitHub](https://github.com/TobitSoftware/chayns-backend)

## used packages

- [Resty](https://github.com/go-resty/resty)

## getting started

```sh
go get github.com/Lukas-Nielsen/go-chayns
```

```go
import chayns "github.com/Lukas-Nielsen/go-chayns"
```

## usage

### conf

for all actions you need a `locationId` and for the most actions you need a `tappId` and a `secret`

```go
// conf
c := chayns.NewConf(<locationId>)

c.
    SetLocation(<locationId>).
    SetTapp(<tappId>).
    SetSecret("<secret>")
```

### filter

Reference at [GitHub-Wiki @ TobitSoftware](https://github.com/TobitSoftware/chayns-backend/wiki/Parameters,-Fields-&-Filters#fields-and-filters)

some functions accept a filter parameter this has to be of type `map[string]string`, eg:

```go
filter := map[string]string{
    "gender": "female",
}
```

### functions

#### accesstoken

```go
// generate page accesstoken
// string, error
result, err := c.NewPageAccessToken(..."<permission>")

// validate an accesstoken
// accessToken (information about the token), error
result, err := c.ValidateAccessToken("<your token>", ...<groupId>)
```

#### user

```go
// get all users
// []user, error
result, err := c.Users(...filter)

// get user by id
// user, error
result, err := c.User(<userId>)

// get groups of user by id
// []uac, error
result, err := c.UserUAC(<userId>, ...filter)

// get devices of user by id
// []device, error
result, err := c.UserDevices(<userId>, ...filter)

// get device by id of user by id
// []device, error
result, err := c.UserDevice(<userId>, <deviceId>)
```

#### group

```go
// get all groups
// []uac, error
result, err := c.Groups(...filter)

// get group by id
// uac, error
result, err := c.Group(<goupId>)

// create a new group
// uac, error
result, err := c.NewGroup("<name>", "<showName>")

// modify a group
// uac, error
result, err := c.ModifyGroup(<groupId>, "<name>", "<showName>")

// delete a group
// bool, error
result, err := c.DeleteGroup(<groupId>)
```

#### group member

```go
// get all members of group by id
// []user, error
result, err := c.Members(<groupId>, ...filter)

// get member by id of group by id
// user, error
result, err := c.Member(<goupId>, <userId>)

// add user to group
// user, error
result, err := c.AddUserUac(<groupId>, <userId>)

// remove user from group
// bool, error
result, err := c.RemoveUserUac(<groupId>, <userId>)
```

#### device

```go
// get device by id
// device, error
result, err := c.Device(<deviceId>)
```

#### location

```go
// get information about the location
// location, error
result, err := c.Location()
```

#### push

```go
// send push notification to group
// bool, error
result, err := c.PushGroup(<groupId>, "<message>")

// send push notification to user
// bool, error
result, err := c.PushUser(<userId>, "<message>")
```

#### uintercom

```go
// uintercom
i := c.NewIntercom("<message>")

i.
    AddUser(...<userId>).
    AddGroup(...<groupId>).
    AddLocation(...<locationId>).
    AddImage(..."<url>").
    SetAccessToken("<token>").
    SetThreadName("<name>").
    SetGroupChat(boolean)

// send the message
// bool, error
result, err := i.Send()
```
