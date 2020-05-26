package serializer

var (
	trueBytes  = []byte("true")
	falseBytes = []byte("false")
)

func AppendSerializedBool(buff []byte, b bool) []byte {
	if b {
		return append(buff, trueBytes...)
	}
	return append(buff, falseBytes...)
}
