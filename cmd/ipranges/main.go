package main

import (
	"flag"
	"fmt"

	aws "github.com/nihab70/AWS_ipranges/cloudkit"
)

func main() {

	// region and service, which are interesting
	myRegionPtr := flag.String("region", "eu-central-1", "region code from AWS. Default = eu-central-1")
	myServicePtr := flag.String("service", "EC2", "service code from AWS. Default = ECS")

	// read the commanline args
	flag.Parse()

	res := aws.GetAWSIPRange()

	fmt.Printf("IP Range creation date: %v\n", res.CreateDate)
	fmt.Printf("Range for %v, %v\n-----\n", *myRegionPtr, *myServicePtr)

	for i := range res.Prefixes {
		prefix := res.Prefixes[i]
		if prefix.Region == *myRegionPtr && prefix.Service == *myServicePtr {
			fmt.Printf("%v: %v \n", i, prefix.IPPrefix)
		}

	}
}
