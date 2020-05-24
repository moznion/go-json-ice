package serializer

var (
	trueBytes  = []byte("true")
	falseBytes = []byte("false")
)

func SerializeBool(b bool) []byte {
	if b {
		return trueBytes
	}
	return falseBytes
}
