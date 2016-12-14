package data

// Definition stores information about a system, used for importing data.
type Definition struct {
	Title    string
	Type     string
	Env      string
	Location string
	User     string
	Password string
	URL      string
	Notes    string
	Tags     []string
}

// YamlData stores information about all systems, used for importing data.
type YamlData struct {
	Defs []Definition
}
