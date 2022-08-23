package canonical

func (f *StreetMarketFilter) Extract() ([]string, []interface{}) {
	whereCol := []string{}
	whereVal := []interface{}{}

	if f.Name != "" {
		whereCol = append(whereCol, "name_alias LIKE ?")
		whereVal = append(whereVal, f.Name+"%")
	}

	if f.District != "" {
		whereCol = append(whereCol, "district LIKE ?")
		whereVal = append(whereVal, f.District+"%")
	}

	if f.Neighborhood != "" {
		whereCol = append(whereCol, "neighborhood LIKE ?")
		whereVal = append(whereVal, f.Neighborhood+"%")
	}

	if f.Region5 != "" {
		whereCol = append(whereCol, "region5 LIKE ?")
		whereVal = append(whereVal, f.Region5+"%")
	}

	return whereCol, whereVal
}
