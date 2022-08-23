package rest

import (
	"encoding/json"

	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
)

type KeyRequest struct {
	Id int64 `json:"id" binding:"required"`
}

func ToJSON(any interface{}) map[string]interface{} {
	var smMap map[string]interface{}
	content, _ := json.Marshal(any)
	json.Unmarshal(content, &smMap)
	return smMap
}

func ListToJSON(myList []*canonical.StreetMarket) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range myList {
		result = append(result, ToJSON(item))
	}
	return result
}
