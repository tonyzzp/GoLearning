package main

import (
	"fmt"
	"reflect"
)

type _User struct {
	Name string `key:"c_name" def:""`
	age  int
}

func main() {
	user := _User{}
	t := reflect.TypeOf(_User{})
	fmt.Println(t)
	fmt.Println("path:", t.PkgPath())
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	fmt.Println(t.NumField())
	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f)
		tag := f.Tag
		fmt.Println(tag.Get("key"))
		fmt.Println(tag.Get("def"))
	}
	fmt.Println(user)
	v := reflect.ValueOf(&user.Name)
	v.Elem().SetString("zzp")
	fmt.Println(user.Name)
}
