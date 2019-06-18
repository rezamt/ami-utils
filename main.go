package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/rezamt/ami-utils/resources"
)

func main() {

	// result, err := resources.ListAMIs("amzn-ami-minimal-hvm-*")
	result, err := resources.ListAMIs("rez-*")

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	for _, im := range result.Images {
		fmt.Println(*im.CreationDate, " - ", *im.Name)
	}

}
