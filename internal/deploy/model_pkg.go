package deploy

type Pkg struct {
	Service   string            `json:"service" yaml:"service"`
	Version   string            `json:"version" yaml:"version"`
	BuildDate string            `json:"build_date" yaml:"build_date"`
	Pkg       map[string]string `json:"pkg" yaml:"pkg"`
	Validate  map[string]string `json:"validate" yaml:"validate"`
}
