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
			jsonObj := ipfile.AsJson[ipfile.AmazonWebServicesFile](i.DownloadFileName)
			cidrs = jsonObj.Process(cidrs)
		case "google":
			jsonObj := ipfile.AsJson[ipfile.GoogleCloudFile](i.DownloadFileName)
			cidrs = jsonObj.Process(cidrs)
		case "azure":
			jsonObj := ipfile.AsJson[ipfile.AzureFile](i.DownloadFileName)
			cidrs = jsonObj.Process(cidrs)
		case "digitalocean":
			cidrsObj := ipfile.AsCSV(i.DownloadFileName, 0)
			cidrs = cidrsObj.Process(cidrs)
		case "cloudflare":
			cidrsObj := ipfile.AsText(i.DownloadFileName)
			cidrs = cidrsObj.Process(cidrs)
		case "oracle":
			jsonObj := ipfile.AsJson[ipfile.OracleFile](i.DownloadFileName)
			cidrs = jsonObj.Process(cidrs)
		}

		for _, cidr := range cidrs {
			processedCidr, err := ipnet.PrepareCidrforDB(cidr)
			if err != nil {
				fmt.Println("Error: ", err)
			}

			if processedCidr.Iptype == "IPv4" {
				c := models.CidrObject{
					Net:           cidr,
					Start_ip:      processedCidr.NetIPDecimal,
					End_ip:        processedCidr.BcastIPDecimal,
					Url:           i.Url,
					Cloudplatform: i.Cloudplatform,
					Iptype:        processedCidr.Iptype,
				}

				models.AddCidr(db, c)

			}

		}
	}

}
