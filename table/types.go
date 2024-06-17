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
	Created  bool     `json:"created,omitempty"` //创建时间
}

type TimeSeries struct {
	TimeField          string `json:"timeField,omitempty"`
	MetaField          string `json:"metaField,omitempty"`
	Granularity        string `json:"granularity,omitempty"` //seconds, minutes, hours
	ExpireAfterSeconds int64  `json:"expireAfterSeconds,omitempty"`
}
