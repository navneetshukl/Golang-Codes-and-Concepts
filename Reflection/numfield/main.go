// The NumField() method returns the number of fields
// in a struct and the Field(i int) method returns the
//  reflect.Value of the ith field.

package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

func GetField(inter interface{}) {
	if reflect.ValueOf(inter).Kind() == reflect.Struct {
		v := reflect.ValueOf(inter)

		fmt.Println("Number of fields ", v.NumField())

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main(){
	o:=order{
		orderId: 15,
		customerId: 80,
	}
	GetField(o)
}
