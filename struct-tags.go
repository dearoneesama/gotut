package main

import (
	"fmt"
	"reflect"
	"strings"
)

type getFlatMapPairType struct {
	first, second string
}

func getFlatMap(t reflect.Type) (res []getFlatMapPairType) {
	type pairType struct {
		field reflect.StructField
		seen bool
	}

	if t.NumField() == 0 {
		return
	}
	stack := make([]pairType, t.NumField())
	for i := t.NumField() - 1; i >= 0; i-- {
		stack[t.NumField()-i-1] = pairType { t.Field(i), false }
	}
	var prefixes []string

	for len(stack) != 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.seen {
			prefixes = prefixes[:len(prefixes)-1]
			continue
		}
		stack = append(stack, pairType{ pair.field, true })
		field := pair.field
		fieldTag := field.Tag.Get("bson")

		switch fieldType := field.Type.Kind(); fieldType {
		case reflect.Struct:
			for i := field.Type.NumField() - 1; i >= 0; i-- {
				stack = append(stack, pairType{ field.Type.Field(i), false })
			}
			prefixes = append(prefixes, fieldTag + ".")
		case reflect.Array, reflect.Slice:
			prefixes = append(prefixes, fieldTag + "[]")
			elemType := field.Type.Elem()
			// flatten nd arrays
			for elemType.Kind() == reflect.Array || elemType.Kind() == reflect.Slice {
				prefixes[len(prefixes)-1] = prefixes[len(prefixes)-1] + "[]"
				elemType = elemType.Elem()
			}
			if elemType.Kind() == reflect.Struct {
				stack = append(stack, pairType{ reflect.StructField {
					Name: "",
					Type: elemType,
				}, false }) // virtual
			} else {
				res = append(res, getFlatMapPairType{ strings.Join(prefixes, ""), elemType.String() })
			}
		case reflect.Map:
			panic("not supported")
		default:
			prefixes = append(prefixes, fieldTag)
			res = append(res, getFlatMapPairType{ strings.Join(prefixes, ""), field.Type.String() })
		}
	}
	return
}

type testStruct struct {
	FirstName string    `bson:"firstname"`
	LastName string     `bson:"lastname"`
	Address struct {
		Street string      `bson:"street"`
		City string        `bson:"city"`
		Zip int64          `bson:"zip"`
	}                   `bson:"address"`
	anime [] struct {
		Name string        `bson:"name"`
		Yolo struct {
			Life [][]string        `bson:"life"`
			Nest struct {
				Name string          `bson:"name"`
				Deep3D [][][] struct {
					int                  `bson:"int"`
				}                    `bson:"deep"`
			}                     `bson:"nest"`
		}                  `bson:"yolo"`
	}                   `bson:"anime"`
}

func main() {
	var test testStruct
	res := getFlatMap(reflect.TypeOf(test))
	for _, v := range res {
		fmt.Printf("%v: %v\n", v.first, v.second)
	}
	/*
firstname: string
lastname: string
address.street: string
address.city: string
address.zip: int64
anime[].name: string
anime[].yolo.life[][]: string
anime[].yolo.nest.name: string
anime[].yolo.nest.deep[][][].int: int
	*/
}
