package ipfile

type GoogleCloudFile struct {
	IpfileJson
	Prefixes []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
	} `json:"prefixes"`
}

func (g *GoogleCloudFile) Process(cidrs []string) []string {
	for _, val := range g.Prefixes {
		exists := StrInSlice(val.Ipv4Prefix, cidrs)
		if exists == false {
			cidrs = append(cidrs, val.Ipv4Prefix)
		}
	}

	return cidrs
}
