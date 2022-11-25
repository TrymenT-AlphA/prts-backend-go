package json

// import (
// 	"errors"
// 	"reflect"

// 	"github.com/valyala/fastjson"
// )

// func JSONParse(src *fastjson.Value, dst interface{}) error {

// 	dstValue := reflect.ValueOf(dst).Elem()
// 	if !dstValue.CanAddr() {
// 		return errors.New("[JSONParse] interface can not addr")
// 	}

// 	fieldNum := dstValue.NumField()
// 	for fieldIndex := 0; fieldIndex < fieldNum; fieldIndex++ {
// 		fieldType := dstValue.Field(fieldIndex).Type().String()
// 		if fieldType == "string" {
// 			dstValue.Field(fieldIndex).SetString(string(src.GetStringBytes(dstValue.Field(fieldIndex).Name)))
// 		} else {
// 			return errors.New("[JSONParse] interface type do not support")
// 		}
// 	}
// 	return nil
// }
