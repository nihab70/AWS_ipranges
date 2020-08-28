package main

import (
	"flag"
	"fmt"

	"github.com/nihab70/cloudorama/aux/envconfig"
	"github.com/nihab70/cloudorama/cloudkit/aws"
)

func main() {

	//set up logger ...
	envconfig.InitLogging()

	//init DB
	initDB()

	// region and service, which are interesting
	myRegionPtr := flag.String("region", "eu-central-1", "region code from AWS. Default = eu-central-1")
	myServicePtr := flag.String("service", "EC2", "service code from AWS. Default = ECS")

	// read the commanline args
	flag.Parse()

	res := aws.GetAWSIPRange()

	fmt.Printf("IP Range creation date: %v\n", res.CreateDate)
	fmt.Printf("Range for %v, %v\n-----\n", *myRegionPtr, *myServicePtr)

	for i, prefix := range res.Prefixes {
		if prefix.Region == *myRegionPtr && prefix.Service == *myServicePtr {
			fmt.Printf("%v: %v \n", i, prefix.IPPrefix)
		}

	}
}
