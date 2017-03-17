package data

// Definition stores information about a system, used for importing data.
type Definition struct {
	Id   string
	Name string
	Type string
	User string
	URL  string

	Env      string
	Location string
	Password string
	Notes    string
	Tags     []string

	Title string //obsolete
}

// YamlData stores information about all systems, used for importing data.
type YamlData struct {
	Defs []Definition
}

// Config stores data about a system, used for exporting data.
type Config struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password,omitempty"`
	URL      string `json:"url"`
	Location string `json:"location"`
	Env      string `json:"env"`
	Host     string `json:"host,omitempty"`
	IsValid  bool   `json:"-"`
}
