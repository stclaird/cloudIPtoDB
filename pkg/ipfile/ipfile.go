package ipfile

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const downloaddir = "ipfiles/"

type DownloadFile struct {
	Url              string `json:"url"`
	DownloadFileName string `json:"DownloadFileName"`
	Cloudplatform    string `json:"Cloudplatform"`
}

func (i *DownloadFile) Download() (err error) {

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

type IpfileJson struct {
	SyncToken    string `json:"syncToken"`
	CreationTime string `json:"creationTime"`
}

type IpfileCSV struct {
	Prefixes []string
}

type IpfileTXT struct {
	Prefixes []string
}

func MatchIp(pattern string) []string {
	//match ip addresses from string pattern and return slice of ips as string
	re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?:/\d{1,2}|)`)
	result := re.FindAllString(pattern, -1)

	return result
}

func StrInSlice(str string, slice []string) bool {
	//find a string in slice return boolean
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
func AsJson[T any](DownloadFileName string) (fileOut T) {
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

func AsCSV(DownloadFileName string, column int) (ipfile IpfileCSV) {
	// Open a CSV and retrieve CIDR
	var cidrs []string
	csvfile, err := os.Open(downloaddir + DownloadFileName)
	if err != nil {
		log.Println("Error", err)
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		cidrs = append(cidrs, record[column])
	}
	ipfile.Prefixes = cidrs
	return ipfile
}

func AsText(DownloadFileName string) (ipfile IpfileTXT) {
	file, err := os.Open(downloaddir + DownloadFileName)
	if err != nil {
		log.Println("Error", err)
	}
	defer file.Close()

	var cidrs []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		matched := MatchIp(txt)
		for _, cidr := range matched {
			cidrs = append(cidrs, cidr)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ipfile.Prefixes = cidrs

	return ipfile
}
