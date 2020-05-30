package serializer

import "strconv"

// AppendSerializedInt appends a serialized int value to buff.
func AppendSerializedInt(buff []byte, v int64) []byte {
	return strconv.AppendInt(buff, v, 10)
}
