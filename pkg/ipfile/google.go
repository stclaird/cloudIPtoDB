package ipfile

type GoogleCloudFile struct {
	Ipfile
	GoogleCloudFilePrefix
}

type GoogleCloudFilePrefix struct {
	Prefixes []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}
