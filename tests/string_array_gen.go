// Code generated by "json-ice --type=[]string --toplevel-array"; DO NOT EDIT.

package tests

import "github.com/moznion/go-json-ice/serializer"

func MarshalStringArrayAsJSON(st []string) ([]byte, error) {
	buff := make([]byte, 0, 20)
	buff = append(buff, '[')
	for _, v := range st {
		buff = serializer.AppendSerializedString(buff, v)
		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = ']'
	} else {
		buff = append(buff, ']')
	}

	return buff, nil
}
