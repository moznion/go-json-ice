package serializer

import "strconv"

func AppendSerializedString(buff []byte, v string) []byte {
	return strconv.AppendQuote(buff, v)
}
