package canonical

type StreetMarket struct {
	Id            int64   `validate:"required,min=1" json:"id" bun:",pk,autoincrement"`
	Longitude     float64 `json:"long" validate:"required"`
	Latitude      float64 `json:"lat" validate:"required"`
	SectorCense   int64   `json:"sector" validate:"required"`
	AreaPonderate int64   `json:"area" validate:"required"`
	DistrictCode  int64   `json:"dist_code" validate:"required"`
	District      string  `json:"district" validate:"required"`
	SubTownCode   int64   `json:"subtown_code" validate:"required"`
	SubTown       string  `json:"subtown" validate:"required"`
	Region5       string  `json:"region_5" validate:"required"`
	Region8       string  `json:"region_8" validate:"required"`
	Name          string  `json:"name" validate:"required" bun:"name_alias"`
	Registry      string  `json:"registry" validate:"required"`
	Address       string  `json:"addr" validate:"required" bun:"addr"`
	Number        string  `json:"number" validate:"required" bun:"addr_number"`
	Neighborhood  string  `json:"neighborhood" validate:"required"`
	Reference     string  `json:"reference" validate:"required"`
}

type StreetMarketFilter struct {
	District     string `json:"district"`
	Region5      string `json:"region_5"`
	Name         string `json:"name"`
	Neighborhood string `json:"neighborhood"`
}
