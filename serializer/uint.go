package serializer

import "strconv"

func SerializeUint(v uint64) []byte {
	return []byte(strconv.FormatUint(v, 10))
}
