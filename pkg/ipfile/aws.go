package ipfile

type AmazonWebServicesFile struct {
	Ipfile
	Prefixes []struct {
		IPPrefix string `json:"ip_prefix"`
	} `json:"prefixes"`
}
