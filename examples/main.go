package main

import (
	"fmt"

	"github.com/samber/go-clevercloud-api/clever"
)

func main() {
	client, err := clever.NewClient(&clever.ClientConfig{
		OrgId:      "orga_d33e695e-dc55-4175-9331-7437a323de48",
		AuthToken:  "93c58d8cde8f4f3bb8af5585d07a7cd4",
		AuthSecret: "82443b9f50fb43fe935b3835fce41042",
		Endpoint:   "https://api.clever-cloud.com/v2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := example_addon(client); err != nil {
		fmt.Println(err)
		return
	}
	if err := example_application(client); err != nil {
		fmt.Println(err)
		return
	}
}
