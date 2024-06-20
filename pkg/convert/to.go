package convert

func ToBool(value any) bool {
	switch val := value.(type) {
	case bool:
		return val
	case uint8:
		return val > 0
	case uint16:
		return val > 0
	case uint32:
		return val > 0
	case uint64:
		return val > 0
	case uint:
		return val > 0
	case int8:
		return val > 0
	case int16:
		return val > 0
	case int32:
		return val > 0
	case int64:
		return val > 0
	case int:
		return val > 0
	case float32:
		return val > 0
	case float64:
		return val > 0
	}
	return false
}

func ToUint8(value any) uint8 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint8(val)
	case uint16:
		return uint8(val)
	case uint32:
		return uint8(val)
	case uint64:
		return uint8(val)
	case uint:
		return uint8(val)
	case int8:
		return uint8(val)
	case int16:
		return uint8(val)
	case int32:
		return uint8(val)
	case int64:
		return uint8(val)
	case int:
		return uint8(val)
	case float32:
		return uint8(val)
	case float64:
		return uint8(val)
	}
	return 0
}

func ToUint16(value any) uint16 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint16(val)
	case uint16:
		return uint16(val)
	case uint32:
		return uint16(val)
	case uint64:
		return uint16(val)
	case uint:
		return uint16(val)
	case int8:
		return uint16(val)
	case int16:
		return uint16(val)
	case int32:
		return uint16(val)
	case int64:
		return uint16(val)
	case int:
		return uint16(val)
	case float32:
		return uint16(val)
	case float64:
		return uint16(val)
	}
	return 0
}

func ToUint32(value any) uint32 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint32(val)
	case uint16:
		return uint32(val)
	case uint32:
		return uint32(val)
	case uint64:
		return uint32(val)
	case uint:
		return uint32(val)
	case int8:
		return uint32(val)
	case int16:
		return uint32(val)
	case int32:
		return uint32(val)
	case int64:
		return uint32(val)
	case int:
		return uint32(val)
	case float32:
		return uint32(val)
	case float64:
		return uint32(val)
	}
	return 0
}

func ToUint64(value any) uint64 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	case uint:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case int:
		return uint64(val)
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	}
	return 0
}

func ToInt8(value any) int8 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int8(val)
	case uint16:
		return int8(val)
	case uint32:
		return int8(val)
	case uint64:
		return int8(val)
	case uint:
		return int8(val)
	case int8:
		return int8(val)
	case int16:
		return int8(val)
	case int32:
		return int8(val)
	case int64:
		return int8(val)
	case int:
		return int8(val)
	case float32:
		return int8(val)
	case float64:
		return int8(val)
	}
	return 0
}

func ToInt16(value any) int16 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int16(val)
	case uint16:
		return int16(val)
	case uint32:
		return int16(val)
	case uint64:
		return int16(val)
	case uint:
		return int16(val)
	case int8:
		return int16(val)
	case int16:
		return int16(val)
	case int32:
		return int16(val)
	case int64:
		return int16(val)
	case int:
		return int16(val)
	case float32:
		return int16(val)
	case float64:
		return int16(val)
	}
	return 0
}

func ToInt32(value any) int32 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int32(val)
	case uint16:
		return int32(val)
	case uint32:
		return int32(val)
	case uint64:
		return int32(val)
	case uint:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int32:
		return int32(val)
	case int64:
		return int32(val)
	case int:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	}
	return 0
}

func ToInt64(value any) int64 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case uint:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return int64(val)
	case int:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	}
	return 0
}

func ToFloat32(value any) float32 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return float32(val)
	case uint16:
		return float32(val)
	case uint32:
		return float32(val)
	case uint64:
		return float32(val)
	case uint:
		return float32(val)
	case int8:
		return float32(val)
	case int16:
		return float32(val)
	case int32:
		return float32(val)
	case int64:
		return float32(val)
	case int:
		return float32(val)
	case float32:
		return val
	case float64:
		return float32(val)
	}
	return 0
}

func ToFloat64(value any) float64 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case uint:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case int:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	}
	return 0
}
