// Code generated by "json-ice --type=OmitemptyStruct"; DO NOT EDIT.

package tests

import (
	"bytes"

	"github.com/moznion/go-json-ice/serializer"
)

func (s *OmitemptyStruct) MarshalJSON() ([]byte, error) {
	var err error
	buff := bytes.NewBuffer([]byte("{"))
	if s.EmptyBool != false {
		_, err = buff.WriteString(serializer.SerializePropertyName("empty_bool") + ":" + string(serializer.SerializeBool(s.EmptyBool)) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.EmptyInt != 0 {
		_, err = buff.WriteString(serializer.SerializePropertyName("empty_int") + ":" + string(serializer.SerializeInt(int64(s.EmptyInt))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.EmptyUint != 0 {
		_, err = buff.WriteString(serializer.SerializePropertyName("empty_uint") + ":" + string(serializer.SerializeUint(uint64(s.EmptyUint))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.EmptyFloat != 0 {
		_, err = buff.WriteString(serializer.SerializePropertyName("empty_float") + ":" + string(serializer.SerializeFloat(float64(s.EmptyFloat))) + ",")
		if err != nil {
			return nil, err
		}
	}
	if s.EmptyString != "" {
		_, err = buff.WriteString(serializer.SerializePropertyName("empty_string") + ":" + string(serializer.SerializeString(s.EmptyString)) + ",")
		if err != nil {
			return nil, err
		}
	}
	_, err = buff.WriteString(serializer.SerializePropertyName("not_empty_string") + ":" + string(serializer.SerializeString(s.NotEmptyString)) + ",")
	if err != nil {
		return nil, err
	}
	bs := buff.Bytes()
	if bs[len(bs)-1] == ',' {
		bs[len(bs)-1] = '}'
	} else {
		bs = append(bs, '}')
	}
	return bs, nil
}
