package serializer

import "strconv"

func AppendSerializedFloat(buff []byte, v float64) []byte {
	return strconv.AppendFloat(buff, v, 'e', -1, 64)
}
