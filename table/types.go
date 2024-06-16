package table

type Document map[string]interface{}

type Foreign struct {
	ForeignTable string `json:"table,omitempty"`
	ForeignField string `json:"foreign,omitempty"`
	LocalField   string `json:"local,omitempty"`
	As           string `json:"as,omitempty"`
}

type Field struct {
	Name     string `json:"name,omitempty"`
	Label    string `json:"label,omitempty"`
	Type     string `json:"type,omitempty"` //string number date
	Index    bool   `json:"index,omitempty"`
	Required bool   `json:"required,omitempty"`
	Default  any    `json:"default,omitempty"`
}
