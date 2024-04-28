package callingcode

const (
	BASE_URL = "https://webapi.tobit.com/ChaynsCodeAdministration/v1.0/Chaynscode/"
	UA       = "github.com/Lukas-Nielsen/go-chayns/callingcode"

	POST   method = "POST"
	GET    method = "GET"
	PUT    method = "PUT"
	DELETE method = "DELETE"
	PATCH  method = "PATCH"

	TOBITLOGO         icon = "tobitlogo"
	SPACEINVADER      icon = "spaceinvader"
	PACMNAGHOST       icon = "pacmnaghost"
	BUBBLE_1          icon = "bubble1"
	BUBBLE_2          icon = "bubble2"
	EXCLAMATIONMARK_1 icon = "exclamationmark1"
	EXCLAMATIONMARK_2 icon = "exclamationmark2"
	HEART             icon = "heart"
	LOCATION_1        icon = "location1"
	LOCATION_2        icon = "location2"
	PENCIL            icon = "pencil"
	PRICETAG          icon = "pricetag"
	QUESTIONMARK_1    icon = "questionmark1"
	QUESTIONMARK_2    icon = "questionmark2"
	SMILEY_1          icon = "smiley1"
	SMILEY_2          icon = "smiley2"
	SWITCH            icon = "switch"
	THUMBSUP          icon = "thumbsup"
	THUMBSDOWN        icon = "thumbsdown"
	TABLENUMBER       icon = "tablenumber"

	SVG ccFormat = "SVG"
	PNG ccFormat = "PNG"
	JPG ccFormat = "JPG"
	PDF ccFormat = "PDF"

	CC_SITE          ccType = 1
	CC_PAGE          ccType = 2
	CC_URL           ccType = 3
	CC_POST          ccType = 4
	CC_DIALOG_IFRAME ccType = 11

	APP_WITHOUT     ccApp = 0
	APP_RECOMMENDED ccApp = 1
	APP_REQUIRED    ccApp = 2
)
