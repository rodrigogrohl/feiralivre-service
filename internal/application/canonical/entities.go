package canonical

type StreetMarket struct {
	Id            int64   `json:"id"`
	Longitude     float64 `json:"long"`
	Latitude      float64 `json:"lat"`
	SectorCense   int64   `json:"sector"`
	AreaPonderate int64   `json:"area"`
	DistrictCode  int64   `json:"dist_code"`
	District      string  `json:"district"`
	SubTownCode   int64   `json:"subtown_code"`
	SubTown       string  `json:"subtown"`
	Region5       string  `json:"region_5"`
	Region8       string  `json:"region_8"`
	Name          string  `json:"name"`
	Registry      string  `json:"registry"`
	Address       string  `json:"addr"`
	Number        string  `json:"number"`
	Neighborhood  string  `json:"neighborhood"`
	Reference     string  `json:"reference"`
}

type StreetMarketFilter struct {
	District     string `json:"district"`
	Region5      string `json:"region_5"`
	Name         string `json:"name"`
	Neighborhood string `json:"neighborhood"`
}
