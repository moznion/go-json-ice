package serializer

import "strconv"

// AppendSerializedUint appends a serialized uint value to buff.
func AppendSerializedUint(buff []byte, v uint64) []byte {
	return strconv.AppendUint(buff, v, 10)
}
