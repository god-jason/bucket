package table

type Foreign struct {
	Table string `json:"table,omitempty"`
	Field string `json:"field,omitempty"`
	Let   string `json:"let,omitempty"` //引用变量
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
	Updated  bool     `json:"updated,omitempty"` //更新时间
	Children []*Field `json:"children,omitempty"`
}

type TimeSeries struct {
	TimeField          string `json:"timeField,omitempty"`
	MetaField          string `json:"metaField,omitempty"`
	Granularity        string `json:"granularity,omitempty"` //seconds, minutes, hours
	ExpireAfterSeconds int64  `json:"expireAfterSeconds,omitempty"`
}
