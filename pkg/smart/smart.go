package smart

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

	Children []Field `json:"children,omitempty"` //子级？
}

type Form []Field

type SelectOption struct {
	Value    any    `json:"value"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled,omitempty"`
}
