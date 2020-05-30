package serializer

import "strconv"

// AppendSerializedString appends a serialized string value to buff.
func AppendSerializedString(buff []byte, v string) []byte {
	return strconv.AppendQuote(buff, v)
}
