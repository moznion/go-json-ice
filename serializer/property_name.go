package serializer

import "strconv"

func SerializePropertyName(propertyName string) string {
	return strconv.Quote(propertyName)
}
