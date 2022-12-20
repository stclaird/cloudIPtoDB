package ipfile

type OracleFile struct {
	IpfileJson
	LastUpdatedTimestamp string `json:"last_updated_timestamp"`
	Regions              []struct {
		Region string `json:"region"`
		Cidrs  []struct {
			Cidr string   `json:"cidr"`
			Tags []string `json:"tags"`
		} `json:"cidrs"`
	} `json:"regions"`
}

func (a *OracleFile) Process(cidrs []string) []string {
	for _, region := range a.Regions {
		for _, cidr := range region.Cidrs {
			exists := Str_in_slice(cidr.Cidr, cidrs)
			if exists == false {
				cidrs = append(cidrs, cidr.Cidr)
			}
		}
	}

	return cidrs
}
