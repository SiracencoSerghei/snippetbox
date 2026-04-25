package interfaces

import (
	"fmt"
	"reflect"
)

// Inspector — маленький debug tool для інтерфейсів
type Inspector struct{}

// NewInspector створює інстанс
func NewInspector() *Inspector {
	return &Inspector{}
}

// Inspect приймає будь-який interface{} і розбирає його
func (i *Inspector) Inspect(value any) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	fmt.Println("──────────── INTERFACE INSPECTOR ────────────")

	// NIL check
	if !v.IsValid() {
		fmt.Println("value: <nil>")
		return
	}

	fmt.Printf("Type: %T\n", value)
	fmt.Printf("Concrete type: %s\n", t.String())

	// Якщо це pointer — розіменовуємо
	if v.Kind() == reflect.Ptr {
	if v.IsNil() {
		fmt.Println("nil pointer")
		return
	}
	v = v.Elem()
	t = t.Elem()
}

	fmt.Printf("Kind: %s\n", v.Kind())

	// Struct fields dump
	if v.Kind() == reflect.Struct {
		fmt.Println("Fields:")

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			val := v.Field(i)

			var value any
			if val.CanInterface() {
				value = val.Interface()
			} else {
				value = "<unexported>"
			}

			fmt.Printf("  - %s (%s) = %v\n",
				field.Name,
				field.Type,
				value,
)
		}
	} else {
		fmt.Printf("Value: %v\n", v.Interface())
	}

	fmt.Println("─────────────────────────────────────────────")
}