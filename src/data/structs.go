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
	ID          string `json:"id,"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Environment string `json:"environment"`
	User        string `json:"user"`
	Password    string `json:"password,omitempty"`
	Host        string `json:"host,omitempty"`
	IsValid     bool   `json:"-"`
}
