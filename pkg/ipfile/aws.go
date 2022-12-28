package ipfile

type AmazonWebServicesFile struct {
	IpfileJson
	Prefixes []struct {
		IPPrefix string `json:"ip_prefix"`
	} `json:"prefixes"`
}

func (a *AmazonWebServicesFile) Process(cidrs []string) []string {

	for _, val := range a.Prefixes {
		exists := StrInSlice(val.IPPrefix, cidrs)
		if exists == false {
			cidrs = append(cidrs, val.IPPrefix)
		}
	}

	return cidrs
}
