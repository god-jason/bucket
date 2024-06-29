package product

// Aggregator 聚合器
type Aggregator struct {
	//Table  string        //默认 bucket.aggregate
	//Period time.Duration //1h
	Type string `json:"type,omitempty"` //inc sum count avg last first max min
	As   string `json:"as,omitempty"`
}

// Property 属性
type Property struct {
	Name        string        `json:"name,omitempty"`        //变量名称
	Label       string        `json:"label,omitempty"`       //显示名称
	Unit        string        `json:"unit,omitempty"`        //单位
	Type        string        `json:"type,omitempty"`        //bool string number array object
	Default     any           `json:"default,omitempty"`     //默认值
	Writable    bool          `json:"writable,omitempty"`    //是否可写
	Historical  bool          `json:"historical,omitempty"`  //是否保存历史
	Aggregators []*Aggregator `json:"aggregators,omitempty"` //聚合计算

	//Children *Property
}

// Alarm 报警器
type Alarm struct {
	Title       string
	Level       string
	Type        string
	Message     string
	Template    string
	Delay       int //延迟时间s
	RepeatDelay int //再次提醒间隔s
	RepeatTotal int //总提醒次数
}

type Product struct {
	Id         string      `json:"_id,omitempty"`
	Name       string      `json:"name,omitempty"` //名称
	Type       string      `json:"type,omitempty"` //泛类型，比如：电表，水表
	Properties []*Property `json:"properties,omitempty"`

	properties map[string]*Property
}

func (p *Product) GetProperty(k string) *Property {
	if p.properties == nil {
		//创建索引
		p.properties = make(map[string]*Property)
		for _, a := range p.Properties {
			p.properties[a.Name] = a
		}
	}
	return p.properties[k]
}
