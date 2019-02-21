package models

type Components struct {
	Components []*Component `json:"components" yaml:"components"`
}

type Component struct {
	Name         string   `json:"component" yaml:"component"`
	Requirements []string `json:"requires" yaml:"requires"`
}
