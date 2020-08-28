package aws

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// IPPrefix define ip_prefix struct in AWS
type IPPrefix struct {
	gorm.Model         `json:"-"`
	IPPrefix           string `json:"ip_prefix,omitempty"`
	Region             string `json:"region,omitempty"`
	Service            string `json:"service,omitempty"`
	NetworkBorderGroup string `json:"network_border_group,omitempty"`
}

// IPv6Prefix defines ipv6_prefix struct in AWS
type IPv6Prefix struct {
	gorm.Model         `json:"-"`
	Ipv6Prefix         string `json:"ipv6_prefix,omitempty"`
	Region             string `json:"region,omitempty"`
	Service            string `json:"service,omitempty"`
	NetworkBorderGroup string `json:"network_border_group,omitempty"`
}

// IPRange defines ip range struct in AWS
type IPRange struct {
	gorm.Model   `json:"-"`
	SyncToken    string       `json:"syncToken,omitempty"`
	CreateDate   string       `json:"createDate,omitempty"`
	Prefixes     []IPPrefix   `json:"prefixes,omitempty" gorm:"-"`
	Ipv6Prefixes []IPv6Prefix `json:"ipv6_prefixes,omitempty" gorm:"-"`
}

// GetAWSIPRange returns the struct for AWS ip ranges as defined above
func GetAWSIPRange() IPRange {

	// URL to retrieve AWS IPs world wide
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"

	log.Info("Read json file from AWS")
	res, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
	}

	defer res.Body.Close()

	var rangeData IPRange

	err = json.NewDecoder(res.Body).Decode(&rangeData)
	if err != nil {
		log.Error(err.Error())
	}

	log.Infof("Read timestamp for ipramges %v", rangeData.CreateDate)
	log.Infof("Read %v entries for prefixes", len(rangeData.Prefixes))
	return rangeData
}

// fill DB with ranges
func fillDB(db *gorm.DB) {
	rangeData := GetAWSIPRange()

	db.Create(&IPRange{SyncToken: rangeData.SyncToken, CreateDate: rangeData.CreateDate})

}

// PrepareDB three tables : timestamp, ipprefix, ipv6_prefix
func PrepareDB(db *gorm.DB) {
	db.AutoMigrate(&IPRange{})
	db.AutoMigrate(&IPPrefix{})
	db.AutoMigrate(&IPv6Prefix{})

	fillDB(db)
}
