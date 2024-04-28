# CallingCodes

## Types

### Overview

there are the folowing types of CallingCodes

- site (1) | opens mainpage of site
- page (2) | opens a given page of site
- url (3) | opens given url
- post (4)| performs a post request with optional parameter
- dialog iframe (11) | opens a given url in a dialog

### config

all require the following attributes

- siteId string
- name string
- type number

and these ones are optional

- appearance
  - icon hex color like #000000
  - color string
  - captionText string
- redirectMode 0|1|2

some have additional attributes

| CC type       | attribute          | type   | required | ref                                                                                     |
| ------------- | ------------------ | ------ | -------- | --------------------------------------------------------------------------------------- |
| page          | tappId             | number | true     |                                                                                         |
| url           | url                | string | true     |                                                                                         |
| post          | requestUrl         | string | true     |                                                                                         |
|               | messageFailure     | string | true     |                                                                                         |
|               | messageBefore      | string | false    |                                                                                         |
|               | messageSuccess     | string | false    |                                                                                         |
|               | requestBody        | json   | false    |                                                                                         |
| iframe dialog | dialogIframeUrl    | string | true     |                                                                                         |
|               | dialogIframeConfig | json   | false    | [Github tobit.software](https://github.com/TobitSoftware/chayns-js/wiki/Dialogs#iframe) |
