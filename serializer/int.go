package serializer

import "strconv"

func SerializeInt(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
