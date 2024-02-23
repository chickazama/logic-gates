package logicgates

type T interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func AND[n T](a, b n) n {
	return a & b
}
