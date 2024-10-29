package time

import (
	"encoding/json"
	"testing"
)

func TestTime_MarshalJSON(t *testing.T) {
	var abc = struct {
		Tm Time
	}{
		Tm: Now(),
	}

	buf, err := json.Marshal(&abc)
	println(string(buf), err)
}
