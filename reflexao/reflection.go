package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	percorreValor := func(valor reflect.Value) {
		percorre(valor.Interface(), fn)
	}

	switch valor.Kind() {
	case reflect.Struct:
		for i := 0; i < valor.NumField(); i++ {
			percorreValor(valor.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < valor.Len(); i++ {
			percorreValor(valor.Index(i))
		}
	case reflect.String:
		fn(valor.String())
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorre(valor.MapIndex(chave).Interface(), fn)
		}
	}
}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}
