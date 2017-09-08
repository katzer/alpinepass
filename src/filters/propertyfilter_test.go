package filters

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/appPlant/alpinepass/src/data"

func TestFilterNameExactMatchTrue(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Earth", Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.Meta.IsValid,
		"The Config should match the filter.")
}

func TestFilterNameExactMatchFalse(t *testing.T) {
	filter := "name=Earth"
	config := data.Config{Name: "Mars", Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.Meta.IsValid,
		"The Config should not match the filter.")
}

//TODO
func TestFilterNameContainsMatchTrue(t *testing.T) {
	filter := "name:Ear"
	config := data.Config{Name: "Earth", Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.Meta.IsValid,
		"The Config should match the filter.")
}

//TODO
func TestFilterNameContainsMatchFalse(t *testing.T) {
	filter := "name=Ear"
	config := data.Config{Name: "Mars", Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.Meta.IsValid,
		"The Config should not match the filter.")
}

func TestFilterLocationTrue(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Earth"}, Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.Meta.IsValid,
		"The Config should match the filter.")
}

func TestFilterLocationFalse(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: []string{"Solar System", "Mars"}, Meta: data.ConfigMeta{IsValid: true}}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.Meta.IsValid,
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
