// Code generated by "json-ice --type=BasicTypes"; DO NOT EDIT.

package benchmark

import "github.com/moznion/go-json-ice/serializer"

func MarshalBasicTypesAsJSON(s *BasicTypes) ([]byte, error) {
	buff := make([]byte, 1, 500)
	buff[0] = '{'
	buff = append(buff, "\"bool_value\":"...)
	buff = serializer.AppendSerializedBool(buff, s.BoolValue)
	buff = append(buff, ',')
	buff = append(buff, "\"int_value\":"...)
	buff = serializer.AppendSerializedInt(buff, int64(s.IntValue))
	buff = append(buff, ',')
	buff = append(buff, "\"int8_value\":"...)
	buff = serializer.AppendSerializedInt(buff, int64(s.Int8Value))
	buff = append(buff, ',')
	buff = append(buff, "\"int16_value\":"...)
	buff = serializer.AppendSerializedInt(buff, int64(s.Int16Value))
	buff = append(buff, ',')
	buff = append(buff, "\"int32_value\":"...)
	buff = serializer.AppendSerializedInt(buff, int64(s.Int32Value))
	buff = append(buff, ',')
	buff = append(buff, "\"int64_value\":"...)
	buff = serializer.AppendSerializedInt(buff, int64(s.Int64Value))
	buff = append(buff, ',')
	buff = append(buff, "\"uint_value\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.UintValue))
	buff = append(buff, ',')
	buff = append(buff, "\"uint8_value\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Uint8Value))
	buff = append(buff, ',')
	buff = append(buff, "\"uint16_value\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Uint16Value))
	buff = append(buff, ',')
	buff = append(buff, "\"uint32_value\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Uint32Value))
	buff = append(buff, ',')
	buff = append(buff, "\"uint64_value\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Uint64Value))
	buff = append(buff, ',')
	buff = append(buff, "\"float32_value\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.Float32Value))
	buff = append(buff, ',')
	buff = append(buff, "\"float64_value\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.Float64Value))
	buff = append(buff, ',')
	buff = append(buff, "\"string_value\":"...)
	buff = serializer.AppendSerializedString(buff, s.StringValue)
	buff = append(buff, ',')
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}