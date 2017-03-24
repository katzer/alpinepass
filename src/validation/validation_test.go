package validation

import (
	"testing"

	"github.com/appPlant/alpinepass/src/data"
)

func TestValidateMandatoryFieldsFail(t *testing.T) {
	config := data.Config{Location: "Earth"}
	verifyMandatoryFields(config)
}

func TestValidateMandatoryFieldsSuccess(t *testing.T) {
	config := data.Config{
		ID:       "test.mandatory.success",
		Location: "Earth",
	}
	verifyMandatoryFields(config)
}
