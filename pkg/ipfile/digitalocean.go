package ipfile

func (i *IpfileCSV) Process(cidrs []string) []string {
	for _, val := range i.Prefixes {
		exists := Str_in_slice(val, cidrs)
		if exists == false {
			cidrs = append(cidrs, val)
		}
	}

	return cidrs
}
