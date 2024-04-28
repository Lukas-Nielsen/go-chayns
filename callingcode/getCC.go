package callingcode

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetCC(t ccFormat, code string, color string, text string, icon icon, tablenumber string) ([]byte, error) {
	resp, err := resty.New().R().SetQueryParams(map[string]string{
		"value":       "https://chayns.cc/" + code,
		"color":       color,
		"text":        text,
		"icon":        string(icon),
		"tablenumber": tablenumber,
	}).Get("https://cube.tobit.cloud/qr-code-generator/v1.0/" + string(t))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("an error occured")
	}

	return resp.Body(), nil
}
