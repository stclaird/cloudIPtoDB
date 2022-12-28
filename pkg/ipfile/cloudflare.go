package ipfile

func (i *IpfileTXT) Process(cidrs []string) []string {
	for _, val := range i.Prefixes {
		exists := StrInSlice(val, cidrs)
		if exists == false {
			cidrs = append(cidrs, val)
		}
	}

	return cidrs
}
