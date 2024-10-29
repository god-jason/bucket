package smart

/**
智能表单

类型
  text:
  password:
  number:
  slider:
  radio:
  rate:
  select:
  tags:
  color:
  checkbox:
  switch:
  textarea:
  date:
  time:
  datetime:
  file:
  image:
  images:
  object:
  list:
  table:
*/

type Field struct {
	Key         string         `json:"key"`
	Label       string         `json:"label"`
	Type        string         `json:"type,omitempty"` //type object array
	Default     any            `json:"default,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
	Tips        string         `json:"tips,omitempty"`
	Pattern     string         `json:"pattern,omitempty"`
	Options     []SelectOption `json:"options,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Min         float64        `json:"min,omitempty"`
	Max         float64        `json:"max,omitempty"`
	Step        float64        `json:"step,omitempty"`

	Disabled bool `json:"disabled,omitempty"`
	Hidden   bool `json:"hidden,omitempty"`

	Array    bool    `json:"array,omitempty"`
	Children []Field `json:"children,omitempty"` //子级？

	Auto []AutoOption `json:"auto,omitempty"`

	Time   bool   `json:"time,omitempty"`
	Upload string `json:"upload,omitempty"` //上传路径
}

type Form []Field

type SelectOption struct {
	Value    any    `json:"value"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled,omitempty"`
}

type AutoOption struct {
	Label string `json:"label"`
	Value any    `json:"value"`
}
