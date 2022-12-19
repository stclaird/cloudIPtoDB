package ipfile

type AzureFile struct {
	ChangeNumber int    `json:"changeNumber"`
	Cloud        string `json:"cloud"`
	Values       []struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		Properties struct {
			ChangeNumber    int      `json:"changeNumber"`
			Region          string   `json:"region"`
			RegionID        int      `json:"regionId"`
			Platform        string   `json:"platform"`
			SystemService   string   `json:"systemService"`
			AddressPrefixes []string `json:"addressPrefixes"`
			NetworkFeatures []string `json:"networkFeatures"`
		} `json:"properties"`
	} `json:"values"`
}

func (a *AzureFile) Process(cidrs []string) []string {
	//	fmt.Println(a.Values)
	for _, val := range a.Values {
		for _, cidr := range val.Properties.AddressPrefixes {
			exists := Str_in_slice(cidr, cidrs)
			if exists == false {
				cidrs = append(cidrs, cidr)
			}
		}
	}

	return cidrs
}
