package table

type Document map[string]interface{}

type Foreign struct {
	Table string `json:"table,omitempty"`
	Field string `json:"field,omitempty"`
	As    string `json:"as,omitempty"`
}

type Field struct {
	Name     string   `json:"name,omitempty"`
	Label    string   `json:"label,omitempty"`
	Type     string   `json:"type,omitempty"` //string number date
	Index    bool     `json:"index,omitempty"`
	Required bool     `json:"required,omitempty"`
	Default  any      `json:"default,omitempty"`
	Foreign  *Foreign `json:"foreign,omitempty"`
}
