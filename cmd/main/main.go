package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github/stclaird/cloudIPtoDB/pkg/config"
	"github/stclaird/cloudIPtoDB/pkg/ipnet"
	"github/stclaird/cloudIPtoDB/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

var confObj = config.NewConfig()

func setup() *sql.DB{

    os.MkdirAll(confObj.Dbdir, 0755)
	full_path := fmt.Sprintf("%s/%s", confObj.Dbdir, confObj.Dbfile)
   	file, err := os.Create(full_path)

	if err != nil {
		log.Println("Os Create Error: ", err)
	}

	file.Close()

	models.DB, _ = sql.Open("sqlite3", full_path )
	if err != nil {
		log.Fatal(err)
	}

	models.SetupDB(models.DB)

	return models.DB
}

func main() {

	db := setup()

	for _, ipfile := range confObj.Ipfiles {

		ipfile.Download()

		var jsonObj interface{}
		var cidrs []string

		if ipfile.Cloudplatform == "aws" {
			jsonObj = ipfile.amazonAsJson(ipfile.DownloadFileName)
			json := jsonObj.(AmazonWebServicesFile)
			fmt.Printf("Found %v Cidrs from %s\n", len(json.Prefixes), ipfile.Cloudplatform)
			for _, val := range json.Prefixes {
				exists := ipfile.Str_in_slice(val.IPPrefix, cidrs)
				if exists == false {
					cidrs = append(cidrs, val.IPPrefix)
				}
			}
		}

		if ipfile.Cloudplatform == "google" {
			jsonObj = ipfile.googleAsJson(ipfile.DownloadFileName)
			json := jsonObj.(GoogleCloudFile)
			fmt.Printf("Found %v Cidrs from %s\n", len(json.Prefixes), ipfile.Cloudplatform)
			for _, val := range json.Prefixes {
				var cidr string
				if len(val.Ipv4Prefix) > 0 {
					cidr = val.Ipv4Prefix
					exists := ipfile.Str_in_slice(cidr, cidrs)

					if exists == false {
						cidrs = append(cidrs, cidr)
					}
				}

			}
		}

		for _, cidr := range cidrs {
			processedCidr, err := ipnet.ProcessCidr(cidr)

			if err != nil {
				fmt.Println("Error: ", err)
			}

			c := models.CidrObject{
				Net:           cidr,
				Start_ip:      processedCidr.netIPDecimal,
				End_ip:        processedCidr.bcastIPDecimal,
				Url:           ipfile.Url,
				Cloudplatform: ipfile.Cloudplatform,
				Iptype:        "IPv4",
			}

			models.AddCidr(db, c)
		}

	}

}
