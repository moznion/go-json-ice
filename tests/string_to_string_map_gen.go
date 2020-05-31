// Code generated by "json-ice --type=map[string]string --toplevel-map"; DO NOT EDIT.

package tests

import "github.com/moznion/go-json-ice/serializer"

func MarshalStringToStringMapAsJSON(mt map[string]string) ([]byte, error) {
	buff := make([]byte, 0, 20)
	buff = append(buff, '{')
	for mapKey, mapValue := range mt {
		buff = serializer.AppendSerializedString(buff, mapKey)
		buff = append(buff, ':')
		buff = serializer.AppendSerializedString(buff, mapValue)
		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}

	return buff, nil
}