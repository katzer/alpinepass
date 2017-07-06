package filters

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/appPlant/alpinepass/src/data"

func TestFilterNameExactMatchTrue(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Earth", IsValid: true}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.IsValid,
		"The Config should match the filter.")
}

func TestFilterNameExactMatchFalse(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Mars", IsValid: true}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.IsValid,
		"The Config should not match the filter.")
}

func TestFilterLocationTrue(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Earth"}, IsValid: true}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.IsValid,
		"The Config should match the filter.")
}

func TestFilterLocationFalse(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Mars"}, IsValid: true}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.IsValid,
		"The Config should not match the filter.")
}

func TestVerifyFlags(t *testing.T) {
	flags := []string{"location=Earth", "location:Earth"}
	assert.NotPanics(t, func() { validateFlags(flags) })
}

func TestVerifyFlagsError(t *testing.T) {
	flags := []string{"location==Earth", "location::Earth"}
	assert.Panics(t, func() { validateFlags(flags) })
}
