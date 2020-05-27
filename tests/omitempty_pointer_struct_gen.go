// Code generated by "json-ice --type=OmitemptyPointerStruct"; DO NOT EDIT.

package tests

import "github.com/moznion/go-json-ice/serializer"

func MarshalOmitemptyPointerStructAsJSON(s *OmitemptyPointerStruct) ([]byte, error) {
	buff := make([]byte, 1, 500)
	buff[0] = '{'
	if s.EmptyBool != nil && *s.EmptyBool != false {
		buff = append(buff, "\"empty_bool\":"...)
		buff = serializer.AppendSerializedBool(buff, *s.EmptyBool)
		buff = append(buff, ',')
	}
	if s.EmptyInt != nil && *s.EmptyInt != 0 {
		buff = append(buff, "\"empty_int\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.EmptyInt))
		buff = append(buff, ',')
	}
	if s.EmptyUint != nil && *s.EmptyUint != 0 {
		buff = append(buff, "\"empty_uint\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.EmptyUint))
		buff = append(buff, ',')
	}
	if s.EmptyFloat != nil && *s.EmptyFloat != 0 {
		buff = append(buff, "\"empty_float\":"...)
		buff = serializer.AppendSerializedFloat(buff, float64(*s.EmptyFloat))
		buff = append(buff, ',')
	}
	if s.EmptyString != nil && *s.EmptyString != "" {
		buff = append(buff, "\"empty_string\":"...)
		buff = serializer.AppendSerializedString(buff, *s.EmptyString)
		buff = append(buff, ',')
	}
	if s.EmptySlice != nil && len(s.EmptySlice) > 0 {
		buff = append(buff, "\"empty_slice\":"...)
		buff = append(buff, '[')
		for _, v := range s.EmptySlice {
			buff = serializer.AppendSerializedString(buff, v)
			buff = append(buff, ',')
		}
		if buff[len(buff)-1] == ',' {
			buff[len(buff)-1] = ']'
		} else {
			buff = append(buff, ']')
		}
		buff = append(buff, ',')
	}
	if s.NotEmptyString == nil {
		buff = append(buff, "\"not_empty_string\":null,"...)
	} else {
		buff = append(buff, "\"not_empty_string\":"...)
		buff = serializer.AppendSerializedString(buff, *s.NotEmptyString)
		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}
