package riskmisc

import "github.com/mitchellh/mapstructure"

func demo() {
	mapstructure.Decode()
}

// DecodeJson 在将 map转换成普通结构体时, map中的键值对的匹配根据结构体定义的json来寻找, 而不是字段名.
func DecodeJson(input interface{}, output interface{}) error {
	c := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  "json",
	}

	decoder, err := mapstructure.NewDecoder(c)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
