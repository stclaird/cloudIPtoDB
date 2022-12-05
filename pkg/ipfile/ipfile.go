package ipfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const downloaddir = "ipfiles/"

type GoogleCloudFile struct {
	SyncToken    string `json:"syncToken"`
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}

type AmazonWebServicesFile struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix string `json:"ip_prefix"`
	} `json:"prefixes"`
}

type Ipfile struct {
	Url              string `json:"url"`
	DownloadFileName string `json:"DownloadFileName"`
	Cloudplatform    string `json:"Cloudplatform"`
}

func (i *Ipfile) Download() (err error) {

	full_path := fmt.Sprintf("%s/%s", downloaddir, i.DownloadFileName)
	log.Printf("Downloading %s to %s", i.Url, full_path)
	//Download the IP Address file
	// Create the file
	fileOut, err := os.Create(full_path)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	resp, err := http.Get(i.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Errorf("bad status: %s", resp.Status))
	}

	// Write the body to file
	_, err = io.Copy(fileOut, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func GoogleAsJson(DownloadFileName string) (fileOut GoogleCloudFile) {
	// Open downloaded file and return as json
	jsonFile, err := os.Open(downloaddir + DownloadFileName)
	if err != nil {
		log.Println("Error", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &fileOut)

	return fileOut
}

func AmazonAsJson(DownloadFileName string) (fileOut AmazonWebServicesFile) {
	// Open downloaded file and return as json

	jsonFile, err := os.Open(downloaddir + DownloadFileName)
	if err != nil {
		log.Println("Error", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &fileOut)

	return fileOut
}

func Str_in_slice(str string, slice []string) bool {
	//find a string in slice return boolean
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
