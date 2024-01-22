package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

func GetTypeAndValue(val interface{}) {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	k := t.Kind()

	fmt.Println("Type of the Given interface is ", t)
	fmt.Println("Value of the given interface is ", v)
	fmt.Println("Kind of the given interface is ", k)
}

func main() {
	o := order{
		orderId:    10,
		customerId: 20,
	}

	GetTypeAndValue(o)
}
