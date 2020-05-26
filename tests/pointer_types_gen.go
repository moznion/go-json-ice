// Code generated by "json-ice --type=PointerTypes"; DO NOT EDIT.

package tests

import (
	"bytes"

	"github.com/moznion/go-json-ice/serializer"
)

func (s *PointerTypes) MarshalJSON() ([]byte, error) {
	var err error
	buff := bytes.NewBuffer([]byte("{"))
	if s.BoolValue == nil {
		_, err = buff.WriteString("\"bool_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"bool_value\"" + ":" + string(serializer.SerializeBool(*s.BoolValue)) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.IntValue == nil {
		_, err = buff.WriteString("\"int_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"int_value\"" + ":" + string(serializer.SerializeInt(int64(*s.IntValue))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Int8Value == nil {
		_, err = buff.WriteString("\"int8_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"int8_value\"" + ":" + string(serializer.SerializeInt(int64(*s.Int8Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Int16Value == nil {
		_, err = buff.WriteString("\"int16_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"int16_value\"" + ":" + string(serializer.SerializeInt(int64(*s.Int16Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Int32Value == nil {
		_, err = buff.WriteString("\"int32_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"int32_value\"" + ":" + string(serializer.SerializeInt(int64(*s.Int32Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Int64Value == nil {
		_, err = buff.WriteString("\"int64_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"int64_value\"" + ":" + string(serializer.SerializeInt(int64(*s.Int64Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.UintValue == nil {
		_, err = buff.WriteString("\"uint_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"uint_value\"" + ":" + string(serializer.SerializeUint(uint64(*s.UintValue))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Uint8Value == nil {
		_, err = buff.WriteString("\"uint8_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"uint8_value\"" + ":" + string(serializer.SerializeUint(uint64(*s.Uint8Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Uint16Value == nil {
		_, err = buff.WriteString("\"uint16_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"uint16_value\"" + ":" + string(serializer.SerializeUint(uint64(*s.Uint16Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Uint32Value == nil {
		_, err = buff.WriteString("\"uint32_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"uint32_value\"" + ":" + string(serializer.SerializeUint(uint64(*s.Uint32Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Uint64Value == nil {
		_, err = buff.WriteString("\"uint64_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"uint64_value\"" + ":" + string(serializer.SerializeUint(uint64(*s.Uint64Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Float32Value == nil {
		_, err = buff.WriteString("\"float32_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"float32_value\"" + ":" + string(serializer.SerializeFloat(float64(*s.Float32Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.Float64Value == nil {
		_, err = buff.WriteString("\"float64_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"float64_value\"" + ":" + string(serializer.SerializeFloat(float64(*s.Float64Value))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.StringValue == nil {
		_, err = buff.WriteString("\"string_value\"" + ":" + string(serializer.SerializeNull()) + ",")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = buff.WriteString("\"string_value\"" + ":" + string(serializer.SerializeString(*s.StringValue)) + ",")
		if err != nil {
			return nil, err
		}
	}
	bs := buff.Bytes()
	if bs[len(bs)-1] == ',' {
		bs[len(bs)-1] = '}'
	} else {
		bs = append(bs, '}')
	}
	return bs, nil
}
