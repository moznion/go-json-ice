package serializer

import "strconv"

func SerializeFloat(v float64) []byte {
	return []byte(strconv.FormatFloat(v, 'e', -1, 64))
}
