package data

//DB is used for referencing db type systems.
const DB = "db"

//SERVER is used for referencing server type systems.
const SERVER = "server"

//TOOL is used for referencing tool type systems.
const TOOL = "tool"

//WEB is used for referencing web type systems.
const WEB = "web"

// Definition stores information about a system, used for importing data only.
type Definition struct {
	ID          string
	Name        string
	Type        string
	User        string
	Password    string
	URL         string
	Server      string
	DB          string
	Host        string
	Port        string
	SID         string
	Env         string
	Location    []string
	Responsible []string
	Notes       []string
}

// YamlData stores information about all systems, used for importing data only.
type YamlData struct {
	Defs []Definition
}

//ConfigMeta contains properties which are not part of the system information, for internal use.
type ConfigMeta struct {
	IsValid bool
}

//User contains information about login credentials.
type User struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Config stores data about a system, used for managing and exporting data.
type Config struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	User        string     `json:"user,omitempty"`
	Password    string     `json:"password,omitempty"`
	URL         string     `json:"url"`
	Server      string     `json:"server,omitempty"`
	DB          string     `json:"db,omitempty"`
	Host        string     `json:"host,omitempty"`
	Port        string     `json:"port,omitempty"`
	SID         string     `json:"sid,omitempty"`
	Env         string     `json:"env,omitempty"`
	Location    []string   `json:"location,omitempty"`
	Responsible []string   `json:"responsible,omitempty"`
	Notes       []string   `json:"notes,omitempty"`
	Meta        ConfigMeta `json:"-"`
}
