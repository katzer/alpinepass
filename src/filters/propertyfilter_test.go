package filters

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/appPlant/alpinepass/src/data"

func TestFilterNameTrue(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Earth"}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.IsValid,
		"The Config should match the filter.")
}

func TestFilterNameFalse(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Mars"}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.IsValid,
		"The Config should not match the filter.")
}

func TestFilterLocationTrue(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Earth"}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.IsValid,
		"The Config should match the filter.")
}

func TestFilterLocationFalse(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Mars"}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.IsValid,
		"The Config should not match the filter.")
}
