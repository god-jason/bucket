package base

type Action struct {
	Batch      bool              `json:"batch,omitempty"` //批量操作
	ProductId  string            `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   string            `json:"device_id,omitempty" bson:"device_id"`
	Action     string            `json:"action"`
	Parameters map[string]string `json:"parameters,omitempty"`
}
