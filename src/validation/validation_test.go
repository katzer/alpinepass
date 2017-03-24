package validation

import (
	"testing"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/stretchr/testify/assert"
)

func TestValidateMandatoryFieldsFail(t *testing.T) {
	config := data.Config{
		Location: "Earth",
	}

	assert.NotEqual(t, "", verifyMandatoryFields(config),
		"The validation should fail.")
}

func TestValidateMandatoryFieldsSuccess(t *testing.T) {
	config := data.Config{
		ID:       "test.mandatory.success",
		Env:      "int",
		Location: "Earth",
	}

	assert.Equal(t, "", verifyMandatoryFields(config),
		"The validation should not fail.")
}
