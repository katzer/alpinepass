package validation

import (
	"testing"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/stretchr/testify/assert"
)

func TestValidateCommonFieldsFail(t *testing.T) {
	config := data.Config{
		Location: "Earth",
	}

	assert.NotEqual(t, "", verifyCommonFields(config),
		"The validation should fail.")
}

func TestValidateCommonFieldsSuccess(t *testing.T) {
	config := data.Config{
		ID:   "test.mandatory.success",
		Env:  "int",
		Type: "db",
		User: "someUser",
		URL:  "url.com",
	}

	assert.Equal(t, "", verifyCommonFields(config),
		"The validation should not fail.")
}

func TestValidateTypeSuccess(t *testing.T) {

}

func TestValidateTypeFail(t *testing.T) {

}

func TestDbFieldsSuccess(t *testing.T) {

}

func TestDbFieldsFail(t *testing.T) {

}

func TestWebFieldsSuccess(t *testing.T) {

}

func TestWebFieldsFail(t *testing.T) {

}
