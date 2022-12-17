# go-chayns

package to wrap the api of [chanys®](https://chayns.net/), you can find the documentation of the api at the repository of TobitSoftware [GitHub](https://github.com/TobitSoftware/chayns-backend)

## used packages

-   [Resty](https://github.com/go-resty/resty)

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
conf := chanys.Conf{
    LocationId: <your locationId>,
    TappId: <your tappId>,
    Secret: <your secret>,
}
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
token, err := conf.NewPageAccessToken(..."<permission>")

// validate an accesstoken
// accessToken (information about the token), error
tokenStruct, err := conf.ValidateAccessToken("<your token>", ...<groupId>)
```

#### user

```go
// get all users
// []user, error
users, err := conf.Users(...filter)

// get user by id
// user, error
user, err := conf.User(<userId>)

// get groups of user by id
// []uac, error
uacs, err := conf.UserUAC(<userId>, ...filter)

// get devices of user by id
// []device, error
devices, err := conf.UserDevices(<userId>, ...filter)

// get device by id of user by id
// []device, error
device, err := conf.UserDevice(<userId>, <deviceId>)
```

#### group

```go
// get all groups
// []uac, error
uacs, err := conf.Groups(...filter)

// get group by id
// uac, error
uac, err := conf.Group(<goupId>)

// create a new group
// uac, error
uac, err := conf.NewGroup("<name>", "<showName>")

// modify a group
// uac, error
uac, err := conf.ModifyGroup(<groupId>, "<name>", "<showName>")

// delete a group
// bool, error
bool, err := conf.DeleteGroup(<groupId>)
```

#### group member

```go
// get all members of group by id
// []user, error
users, err := conf.Members(<groupId>, ...filter)

// get member by id of group by id
// user, error
user, err := conf.Member(<goupId>, <userId>)

// add user to group
// user, error
user, err := conf.AddUserUac(<groupId>, <userId>)

// remove user from group
// bool, error
bool, err := conf.RemoveUserUac(<groupId>, <userId>)
```

#### device

```go
// get device by id
// device, error
device, err := conf.Device(<deviceId>)
```

#### location

```go
// get information about the location
// location, error
location, err := conf.Location()
```

#### push

```go
// send push notification to group
// bool, error
bool, err := conf.PushGroup(<groupId>, "<message>")

// send push notification to user
// bool, error
bool, err := conf.PushUser(<userId>, "<message>")
```

#### intercom

```go
// send a intercom
// bool, error
bool, err := conf.Intercom(data Intercom)
```
