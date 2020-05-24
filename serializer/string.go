package serializer

import "strconv"

func SerializeString(v string) []byte {
	return []byte(strconv.Quote(v))
}
