package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type Aggregator struct {
	//Table  string        //默认 bucket.aggregate
	//Period time.Duration //1h
	Type string `json:"type,omitempty"` //inc sum count avg last first max min
	As   string `json:"as,omitempty"`
}

type Property struct {
	Name        string        `json:"name,omitempty"`
	Label       string        `json:"label,omitempty"`
	Type        string        `json:"type,omitempty"` //bool string number array object
	Default     any           `json:"default,omitempty"`
	Writable    bool          `json:"writable,omitempty"`
	Historical  bool          `json:"historical,omitempty"`
	Aggregators []*Aggregator `json:"aggregators,omitempty"`

	//Children *Property
}

type Product struct {
	Id         primitive.ObjectID `json:"_id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Type       string             `json:"type,omitempty"` //泛类型，比如：电表，水表
	Properties []*Property        `json:"properties,omitempty"`

	properties map[string]*Property
}

func (p *Product) GetProperty(k string) *Property {
	return p.properties[k]
}

func (p *Product) Open() error {
	//创建索引
	p.properties = make(map[string]*Property)
	for _, a := range p.Properties {
		p.properties[a.Name] = a
	}

	return nil
}
