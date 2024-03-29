package config

import "fmt"

// RootConfig is used to tell the server what files are to
// be served in the root directory.
type RootConfig struct {
	Path         string   `json:"path"` // "/"
	Directory    string   `json:"dir"`
	Files        []string `json:"files"`
	AliasDefault []string `json:"alias"`
	Allow        string   `json:"allow"`   // comma-separated list of additional allowed files.
	Default      string   `json:"default"` // Default document. default=index.html
}

func (r *RootConfig) info() {
	println("> Root")
	println(fmt.Sprintf("--> Path         = %s", r.Path))
	println(fmt.Sprintf("--> Directory    = %s", r.Directory))
	println(fmt.Sprintf("--> Files        = %s", r.Files))
	println(fmt.Sprintf("--> AliasDefault = %s", r.AliasDefault))
	println(fmt.Sprintf("--> Default      = %s", r.Default))
}
