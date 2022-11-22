package logger

import "reflect"

const (
	maskTag       = "mask"
	sliceByteMask = "X@BQ1"
)

func ToField(key string, val interface{}) (field Field) {
	field = Field{
		Key: key,
		Val: val,
	}
	return
}

func masking(data interface{}) interface{} {
	original := reflect.ValueOf(data)
	altered := reflect.New(original.Type()).Elem()

	switch original.Kind() {
	case reflect.Ptr:
		// check if value is nil
		if !isNil(original) {
			elem := original.Elem()
			switch elem.Kind() {
			case reflect.Struct, reflect.Interface, reflect.Ptr:
				altered.Set(masking(elem.Interface()).(reflect.Value).Addr())
			case reflect.Slice:
				altered = maskSlice(elem)
			case reflect.Map:
				altered = maskMap(elem)
			default:
				altered.Set(elem.Addr())
			}
		}
	case reflect.Slice:
		altered = maskSlice(original)
	case reflect.Map:
		altered = maskMap(original)
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			field := original.Field(i)
			switch field.Kind() {
			case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice, reflect.Ptr:
				if altered.Field(i).CanSet() && !isNil(field) {

					// []byte mostly used for byte file
					if field.Type() == TypeSliceOfBytes {
						if _, ok := original.Type().Field(i).Tag.Lookup(maskTag); ok {
							if !original.Field(i).IsNil() {
								altered.Field(i).SetBytes([]byte(sliceByteMask))
							}
						} else {
							altered.Field(i).Set(original.Field(i))
						}
					} else {
						altered.Field(i).Set(masking(field.Interface()).(reflect.Value))
					}
				}
			default:
				if _, ok := original.Type().Field(i).Tag.Lookup(maskTag); ok {
					var value string
					switch original.Field(i).Interface().(type) {
					case string:
						value = original.Field(i).Interface().(string)
					}
					if len(value) > 0 {
						altered.Field(i).SetString(value)
					} else {
						altered.Field(i).Set(original.Field(i))
					}
				} else {
					if altered.Field(i).CanSet() {
						altered.Field(i).Set(original.Field(i))
					} else {
						switch original.Type() {
						case TypeTime:
							altered.Set(original)
							i += 2
						}
					}
				}
			}
		}
	default:
		altered.Set(original)
	}

	return altered
}

func maskSlice(elem reflect.Value) (altered reflect.Value) {
	altered = reflect.MakeSlice(elem.Type(), elem.Len(), elem.Len())
	for i := 0; i < elem.Len(); i++ {
		value := elem.Index(i)
		switch value.Kind() {
		case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice, reflect.Ptr:
			// check if value is nil
			if !isNil(value) {
				altered.Index(i).Set(masking(value.Interface()).(reflect.Value))
			}
		default:
			altered.Index(i).Set(value)
		}
	}

	return
}

func maskMap(elem reflect.Value) (altered reflect.Value) {
	altered = reflect.MakeMapWithSize(elem.Type(), len(elem.MapKeys()))
	mapRange := elem.MapRange()
	for mapRange.Next() {
		switch mapRange.Value().Kind() {
		case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice, reflect.Ptr:
			// check if value is nil
			if !isNil(mapRange.Value()) {
				altered.SetMapIndex(
					mapRange.Key(),
					masking(mapRange.Value().Interface()).(reflect.Value),
				)
			}
		default:
			altered.SetMapIndex(mapRange.Key(), mapRange.Value())
		}
	}

	return
}

func isNil(elem reflect.Value) bool {
	return elem.Interface() == nil ||
		(reflect.ValueOf(elem.Interface()).Kind() == reflect.Ptr && reflect.ValueOf(elem.Interface()).IsNil())
}
