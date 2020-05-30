package serializer

var (
	trueBytes  = []byte("true")
	falseBytes = []byte("false")
)

// AppendSerializedBool appends a serialized bool value to buff.
func AppendSerializedBool(buff []byte, b bool) []byte {
	if b {
		return append(buff, trueBytes...)
	}
	return append(buff, falseBytes...)
}
