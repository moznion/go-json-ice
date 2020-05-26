package serializer

import "strconv"

func AppendSerializedUint(buff []byte, v uint64) []byte {
	return strconv.AppendUint(buff, v, 10)
}
