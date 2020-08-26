package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

// IPPrefix define ip_prefix struct in AWS
type IPPrefix struct {
	IPPrefix           string `json:"ip_prefix,omitempty"`
	Region             string `json:"region,omitempty"`
	Service            string `json:"service,omitempty"`
	NetworkBorderGroup string `json:"network_border_group,omitempty"`
}

// IPv6Prefix defines ipv6_prefix struct in AWS
type IPv6Prefix struct {
	Ipv6Prefix         string `json:"ipv6_prefix,omitempty"`
	Region             string `json:"region,omitempty"`
	Service            string `json:"service,omitempty"`
	NetworkBorderGroup string `json:"network_border_group,omitempty"`
}

// IPRange defines ip range struct in AWS
type IPRange struct {
	SyncToken    string       `json:"syncToken,omitempty"`
	CreateDate   string       `json:"createDate,omitempty"`
	Prefixes     []IPPrefix   `json:"prefixes,omitempty"`
	Ipv6Prefixes []IPv6Prefix `json:"ipv6_prefixes,omitempty"`
}

// getAWSIPRange returns the struct for AWS ip ranges as defined above
func getAWSIPRange() IPRange {

	// URL to retrieve AWS IPs world wide
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	var rangeData IPRange

	err = json.NewDecoder(res.Body).Decode(&rangeData)
	if err != nil {
		panic(err.Error())
	}
	return rangeData
}

func main() {

	// region and service, which are interesting
	myRegionPtr := flag.String("region", "eu-central-1", "region code from AWS. Default = eu-central-1")
	myServicePtr := flag.String("service", "EC2", "service code from AWS. Default = ECS")

	// read the commanline args
	flag.Parse()

	res := getAWSIPRange()

	fmt.Printf("IP Range creation date: %v\n", res.CreateDate)
	fmt.Printf("Range for %v, %v\n-----\n", *myRegionPtr, *myServicePtr)

	for i := range res.Prefixes {
		prefix := res.Prefixes[i]
		if prefix.Region == *myRegionPtr && prefix.Service == *myServicePtr {
			fmt.Printf("%v: %v \n", i, prefix.IPPrefix)
		}

	}
}
