package main

import (
	"fmt"
	"reflect"
)

type User struct {
	NickName string
	Age      int32
}

func (u User) GetNickName() string {
	return u.NickName
}

func (u User) GetAge() int32 {
	return u.Age
}

func main() {
	user := User{
		NickName: "user1",
		Age:      30,
	}

	fmt.Println("=== Type ===")
	tv := reflect.TypeOf(user)
	tm, _ := tv.MethodByName("GetAge") // メソッドが存在しない場合は第二戻り値がfalseを返す
	fmt.Println("Name:", tm.Name, "Type:", tm.Type)
	for i := 0; i < tv.NumMethod(); i++ {
		t := tv.Method(i)
		fmt.Println(i+1, "Name:", t.Name, "Type:", t.Type)
	}

	fmt.Println()

	fmt.Println("=== Value ===")
	rv := reflect.ValueOf(user)
	rm := rv.MethodByName("GetAge")
	fmt.Println("Kind:", rm.Kind(), "Type:", rm.Type())
	for i := 0; i < rv.NumMethod(); i++ {
		v := rv.Method(i)
		fmt.Println(i+1, "Kind:", v.Kind(), "Type:", v.Type())
	}
}
