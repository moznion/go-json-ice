// Code generated by "json-ice --type=PointerTypes"; DO NOT EDIT.

package tests

import "github.com/moznion/go-json-ice/serializer"

func MarshalPointerTypesAsJSON(s *PointerTypes) ([]byte, error) {
	buff := make([]byte, 1, 500)
	buff[0] = '{'
	if s.BoolValue == nil {
		buff = append(buff, "\"bool_value\":null,"...)
	} else {
		buff = append(buff, "\"bool_value\":"...)
		buff = serializer.AppendSerializedBool(buff, *s.BoolValue)
		buff = append(buff, ',')
	}
	if s.IntValue == nil {
		buff = append(buff, "\"int_value\":null,"...)
	} else {
		buff = append(buff, "\"int_value\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.IntValue))
		buff = append(buff, ',')
	}
	if s.Int8Value == nil {
		buff = append(buff, "\"int8_value\":null,"...)
	} else {
		buff = append(buff, "\"int8_value\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.Int8Value))
		buff = append(buff, ',')
	}
	if s.Int16Value == nil {
		buff = append(buff, "\"int16_value\":null,"...)
	} else {
		buff = append(buff, "\"int16_value\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.Int16Value))
		buff = append(buff, ',')
	}
	if s.Int32Value == nil {
		buff = append(buff, "\"int32_value\":null,"...)
	} else {
		buff = append(buff, "\"int32_value\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.Int32Value))
		buff = append(buff, ',')
	}
	if s.Int64Value == nil {
		buff = append(buff, "\"int64_value\":null,"...)
	} else {
		buff = append(buff, "\"int64_value\":"...)
		buff = serializer.AppendSerializedInt(buff, int64(*s.Int64Value))
		buff = append(buff, ',')
	}
	if s.UintValue == nil {
		buff = append(buff, "\"uint_value\":null,"...)
	} else {
		buff = append(buff, "\"uint_value\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.UintValue))
		buff = append(buff, ',')
	}
	if s.Uint8Value == nil {
		buff = append(buff, "\"uint8_value\":null,"...)
	} else {
		buff = append(buff, "\"uint8_value\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.Uint8Value))
		buff = append(buff, ',')
	}
	if s.Uint16Value == nil {
		buff = append(buff, "\"uint16_value\":null,"...)
	} else {
		buff = append(buff, "\"uint16_value\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.Uint16Value))
		buff = append(buff, ',')
	}
	if s.Uint32Value == nil {
		buff = append(buff, "\"uint32_value\":null,"...)
	} else {
		buff = append(buff, "\"uint32_value\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.Uint32Value))
		buff = append(buff, ',')
	}
	if s.Uint64Value == nil {
		buff = append(buff, "\"uint64_value\":null,"...)
	} else {
		buff = append(buff, "\"uint64_value\":"...)
		buff = serializer.AppendSerializedUint(buff, uint64(*s.Uint64Value))
		buff = append(buff, ',')
	}
	if s.Float32Value == nil {
		buff = append(buff, "\"float32_value\":null,"...)
	} else {
		buff = append(buff, "\"float32_value\":"...)
		buff = serializer.AppendSerializedFloat(buff, float64(*s.Float32Value))
		buff = append(buff, ',')
	}
	if s.Float64Value == nil {
		buff = append(buff, "\"float64_value\":null,"...)
	} else {
		buff = append(buff, "\"float64_value\":"...)
		buff = serializer.AppendSerializedFloat(buff, float64(*s.Float64Value))
		buff = append(buff, ',')
	}
	if s.StringValue == nil {
		buff = append(buff, "\"string_value\":null,"...)
	} else {
		buff = append(buff, "\"string_value\":"...)
		buff = serializer.AppendSerializedString(buff, *s.StringValue)
		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}
