package ipfile

type GoogleCloudFile struct {
	IpfileJson
	GoogleCloudFilePrefix
}

type GoogleCloudFilePrefix struct {
	Prefixes []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}

func (g *GoogleCloudFilePrefix) Process(cidrs []string) []string {
	for _, val := range g.Prefixes {
		exists := Str_in_slice(val.Ipv4Prefix, cidrs)
		if exists == false {
			cidrs = append(cidrs, val.Ipv4Prefix)
		}
	}

	return cidrs
}
