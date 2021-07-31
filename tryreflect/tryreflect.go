package tryreflect

import (
	"fmt"
	"reflect"
)

type Example struct {
	String  string `key:"string" default:"test" required:"true"`
	Number  int    `key:"number" default:"12345" required:"true"`
	Array   [5]int
	Slice   []int
	Maps    map[string]string
	Channel chan int
	Pointer *string
}

func sprintType(k reflect.Kind) string {
	switch k {
	case reflect.String:
		return fmt.Sprint("string")
	case reflect.Int:
		return fmt.Sprint("int")
	case reflect.Array:
		return fmt.Sprint("array")
	case reflect.Slice:
		return fmt.Sprint("slice")
	case reflect.Map:
		return fmt.Sprint("map")
	case reflect.Chan:
		return fmt.Sprint("chan")
	case reflect.Struct:
		return fmt.Sprint("struct")
	case reflect.Ptr:
		return fmt.Sprint("pointer")
	case reflect.Invalid:
		return fmt.Sprint("invalid")
	default:
		return fmt.Sprint("Kind value is", k)
	}
}

func printlnType(input interface{}) {
	fmt.Println(sprintType(TryReflectType(input)))
}

func TryReflectType(input interface{}) reflect.Kind {
	return reflect.TypeOf(input).Kind()
}

func TryReflectTypeElem(input interface{}) reflect.Kind {
	//Only
	t := reflect.TypeOf(input)
	tk := t.Kind()
	if tk != reflect.Ptr &&
		tk != reflect.Array &&
		tk != reflect.Slice &&
		tk != reflect.Chan &&
		tk != reflect.Map {
		return reflect.Invalid
	}
	return t.Elem().Kind()
}

func printlnTypeElem(input interface{}) {
	fmt.Println(sprintType(TryReflectTypeElem(input)))
}

func TryReflectValue(input interface{}) {
	fmt.Println(reflect.ValueOf(input).Interface())
}

func TryTraverseStruct(input interface{}) {
	kind := TryReflectType(input)
	if kind == reflect.Struct {
		values := reflect.ValueOf(input)
		types := reflect.TypeOf(input)
		count := values.NumField()
		//count := types.NumField()  // just ok too.
		for i := 0; i < count; i++ {
			fmt.Printf("%d element is %v, type is %v, value is %v.\n",
				i, types.Field(i).Name, sprintType(values.Field(i).Kind()), values.Field(i).Interface())
		}
	}
}

func TryTraverseStructPointer(input interface{}) {
	kind := TryReflectType(input)
	if kind == reflect.Ptr {
		transType := reflect.TypeOf(input).Elem()
		//transType := reflect.PtrTo(reflect.TypeOf(input)) // if input type is Ptr (&Struct), this kind is reflect.Ptr
		if transType.Kind() == reflect.Struct {
			values := reflect.ValueOf(input).Elem()
			types := transType
			count := types.NumField()
			for i := 0; i < count; i++ {
				fmt.Printf("%d element is %v, type is %v, value is %v.\n",
					i, types.Field(i).Name, sprintType(values.Field(i).Kind()), values.Field(i).Interface())
			}
		} else {
			fmt.Println(sprintType(transType.Kind()))
		}
	}
}

// TryTraverseStructModifyElement
// if the element type is int, then assign the element value 5555
// if string, then "I am changed."
// if slice int, then []{5,5,5,5,5}
// if map[string]string, then append map{"key","value"}
func TryTraverseStructModifyElement(input interface{}) {
	k := reflect.TypeOf(input).Kind()
	// type must be Ptr and elem type must be Struct
	if k == reflect.Ptr {
		if ke := reflect.TypeOf(input).Elem().Kind(); ke == reflect.Struct {
			values := reflect.ValueOf(input).Elem()
			count := values.NumField()
			for i := 0; i < count; i++ {
				f := values.Field(i)
				t := f.Kind()
				switch t {
				case reflect.Int:
					f.SetInt(5555)
				case reflect.String:
					f.SetString("I am changed.")
				}
			}
		}
	}
}

func TryGetElementFromStructByName(eleName string, input interface{}) interface{} {
	var values reflect.Value
	var types reflect.Type
	k := reflect.TypeOf(input).Kind()
	if k == reflect.Struct {
		values = reflect.ValueOf(input)
		types = reflect.TypeOf(input)
	} else if k == reflect.Ptr {
		if ke := reflect.TypeOf(input).Elem().Kind(); ke == reflect.Struct {
			values = reflect.ValueOf(input).Elem()
			types = reflect.TypeOf(input).Elem()
		} else {
			return nil
		}
	} else {
		return nil
	}

	count := values.NumField()
	for i := 0; i < count; i++ {
		if eleName == types.Field(i).Name {
			return values.Field(i).Interface()
		}
	}

	return nil
}

func TryTraverseSlice(input interface{}) {
	var values reflect.Value
	var can bool
	k := reflect.TypeOf(input).Kind()
	if k == reflect.Slice {
		values = reflect.ValueOf(input)
		can = true
	}
	if k == reflect.Ptr {
		ke := reflect.TypeOf(input).Elem().Kind()
		if ke == reflect.Slice {
			values = reflect.ValueOf(input).Elem()
			can = true
		}
	}
	if can {
		//values.Index()
		//values.Len()
		//values.Cap()
		count := values.Len()
		for i := 0; i < count; i++ {
			fmt.Println(values.Index(i).Interface())
		}
	}
}

func TryTraverseTag(input interface{}) {
	types := reflect.TypeOf(input)
	count := 0
	if types.Kind() == reflect.Struct {
		count = types.NumField()
	} else if types.Kind() == reflect.Ptr && types.Elem().Kind() == reflect.Struct {
		count = types.Elem().NumField()
		types = types.Elem()
	} else {
		return
	}

	for i := 0; i < count; i++ {
		fmt.Printf("The %d tag:=======\n", i)
		fieldTag := types.Field(i).Tag
		// if get no key, return key is ""
		// Get calls Lookup essentially
		key := fieldTag.Get("key")
		fmt.Println("    key is", key)
		if def, ok := fieldTag.Lookup("default"); ok {
			fmt.Println("    default is", def)
		}
		if req, ok := fieldTag.Lookup("required"); ok {
			fmt.Println("    required is", req)
		}
	}
}
