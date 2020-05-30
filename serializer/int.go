package serializer

import "strconv"

func AppendSerializedInt(buff []byte, v int64) []byte {
	return strconv.AppendInt(buff, v, 10)
}
