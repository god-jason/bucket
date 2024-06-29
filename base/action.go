package base

type Action struct {
	ProductId  string         `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   string         `json:"device_id,omitempty" bson:"device_id"`
	Name       string         `json:"action"`
	Parameters map[string]any `json:"parameters,omitempty"`
}
