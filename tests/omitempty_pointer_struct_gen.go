// Code generated by "json-ice --type=OmitemptyPointerStruct"; DO NOT EDIT.

package tests

import "strconv"

func MarshalOmitemptyPointerStructAsJSON(s *OmitemptyPointerStruct) ([]byte, error) {
	buff := make([]byte, 1, 500)
	buff[0] = '{'
	if s.EmptyBool != nil && *s.EmptyBool != false {
		buff = append(buff, "\"empty_bool\":"...)
		if *s.EmptyBool {
			buff = append(buff, "true"...)
		} else {
			buff = append(buff, "false"...)
		}

		buff = append(buff, ',')
	}
	if s.EmptyInt != nil && *s.EmptyInt != 0 {
		buff = append(buff, "\"empty_int\":"...)
		buff = strconv.AppendInt(buff, int64(*s.EmptyInt), 10)
		buff = append(buff, ',')
	}
	if s.EmptyUint != nil && *s.EmptyUint != 0 {
		buff = append(buff, "\"empty_uint\":"...)
		buff = strconv.AppendUint(buff, uint64(*s.EmptyUint), 10)
		buff = append(buff, ',')
	}
	if s.EmptyFloat != nil && *s.EmptyFloat != 0 {
		buff = append(buff, "\"empty_float\":"...)
		buff = strconv.AppendFloat(buff, float64(*s.EmptyFloat), 'e', -1, 64)
		buff = append(buff, ',')
	}
	if s.EmptyString != nil && *s.EmptyString != "" {
		buff = append(buff, "\"empty_string\":"...)
		buff = strconv.AppendQuote(buff, *s.EmptyString)
		buff = append(buff, ',')
	}
	if s.EmptySlice != nil && len(s.EmptySlice) > 0 {
		buff = append(buff, "\"empty_slice\":"...)
		buff = append(buff, '[')
		for _, v := range s.EmptySlice {
			buff = strconv.AppendQuote(buff, v)
			buff = append(buff, ',')
		}
		if buff[len(buff)-1] == ',' {
			buff[len(buff)-1] = ']'
		} else {
			buff = append(buff, ']')
		}

		buff = append(buff, ',')
	}
	if s.EmptyMap != nil && len(s.EmptyMap) > 0 {
		buff = append(buff, "\"empty_map\":"...)
		buff = append(buff, '{')
		for mapKey, mapValue := range s.EmptyMap {
			buff = strconv.AppendQuote(buff, mapKey)
			buff = append(buff, ':')
			buff = strconv.AppendQuote(buff, mapValue)
			buff = append(buff, ',')
		}
		if buff[len(buff)-1] == ',' {
			buff[len(buff)-1] = '}'
		} else {
			buff = append(buff, '}')
		}

		buff = append(buff, ',')
	}
	if s.NotEmptyString == nil {
		buff = append(buff, "\"not_empty_string\":null,"...)
	} else {
		buff = append(buff, "\"not_empty_string\":"...)
		buff = strconv.AppendQuote(buff, *s.NotEmptyString)
		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}
