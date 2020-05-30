package serializer

import "strconv"

// AppendSerializedFloat appends a serialized float value to buff.
func AppendSerializedFloat(buff []byte, v float64) []byte {
	return strconv.AppendFloat(buff, v, 'e', -1, 64)
}
