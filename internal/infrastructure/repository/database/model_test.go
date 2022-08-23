package database

import (
	"testing"

	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/stretchr/testify/assert"
)

func TestResetModel(t *testing.T) {
	// t.Skip()
	config.InitializeTest()

	dbStruct := DatabaseConnect()

	err := ResetAllModel(dbStruct)

	assert.Nil(t, err)
}
