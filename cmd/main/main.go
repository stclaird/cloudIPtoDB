package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github/stclaird/cloudIPtoDB/pkg/config"
	"github/stclaird/cloudIPtoDB/pkg/ipfile"
	"github/stclaird/cloudIPtoDB/pkg/ipnet"
	"github/stclaird/cloudIPtoDB/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

var confObj = config.NewConfig()

func setup() *sql.DB {

	os.MkdirAll(confObj.Dbdir, 0755)
	full_path := fmt.Sprintf("%s/%s", confObj.Dbdir, confObj.Dbfile)
	file, err := os.Create(full_path)

	if err != nil {
		log.Println("Os Create Error: ", err)
	}

	file.Close()

	models.DB, _ = sql.Open("sqlite3", full_path)
	if err != nil {
		log.Fatal(err)
	}

	models.SetupDB(models.DB)

	return models.DB
}

func main() {

	db := setup()
	for _, i := range confObj.Ipfiles {

		i.Download()
		var cidrs []string

		switch i.Cloudplatform {
		case "aws":
			jsonObj := ipfile.AsJson[ipfile.GoogleCloudFile](i.DownloadFileName)
			for _, val := range jsonObj.Prefixes {
				exists := ipfile.Str_in_slice(val.Ipv4Prefix, cidrs)
				if exists == false {
					cidrs = append(cidrs, val.Ipv4Prefix)
				}
			}

		case "google":
			jsonObj := ipfile.AsJson[ipfile.GoogleCloudFile](i.DownloadFileName)
			for _, val := range jsonObj.Prefixes {
				exists := ipfile.Str_in_slice(val.Ipv4Prefix, cidrs)
				if exists == false {
					cidrs = append(cidrs, val.Ipv4Prefix)
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
				Start_ip:      processedCidr.NetIPDecimal,
				End_ip:        processedCidr.BcastIPDecimal,
				Url:           i.Url,
				Cloudplatform: i.Cloudplatform,
				Iptype:        "IPv4",
			}

			models.AddCidr(db, c)
		}
	}

}
