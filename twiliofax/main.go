package main

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

func main() {

	restCall := resty.R()
	restCall.SetBasicAuth("myuser", "mypass")
	restCall.SetFormData(map[string]string{
		"From": "+15017122661",
		"To":  "+15558675310",
		"MediaUrl": "https://www.twilio.com/docs/documents/25/justthefaxmaam.pdf"
	})
	response, err := restCall.Post("https://fax.twilio.com/v1/Faxes")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(">>> %v", string(response.Body()))

}
