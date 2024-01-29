/*
 GoLang DLL example. Goals: load golang dll into fpc/lazarus, and load golang
 dll into another go executable.

 The syntax
   //export SomeFunc
 needs to be used above the function you want to export

 To compile this program run:
   go build -buildmode=c-archive exportgo.go
 then compile goDLL.c that exports the functions for GCC to link, and run:
   gcc -shared -pthread -o goDLL.dll goDLL.c exportgo.a -lWinMM -lntdll -lWS2_32

*/

package main

/*
#cgo LDFLAGS: -L.
#include "goDLL.h"
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/lang-library/go-global"
	"github.com/lang-library/go-winlib"
	"golang.org/x/sys/windows"
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

func (u User) Add2(args any) any {
	var _result any
	_result = nil
	if _args, _ok := args.([]any); _ok {
		_result = _args[0].(float64) + _args[1].(float64)
	}
	return _result
}

var (
	funcTable map[string]func(any) any = make(map[string]func(any) any)
	user                               = User{
		NickName: "user1",
		Age:      30,
	}
)

func init() {
	funcTable["add2"] = add2_impl
}

// https://stackoverflow.com/questions/71587996/cannot-use-type-assertion-on-type-parameter-value
func isInt[T any](x T) (ok bool) {

	_, ok = any(x).(int) // convert, then assert
	return
}

//export Call
func Call(_namePtr, _jsonPtr uintptr) uintptr {
	_name := winlib.UTF8AddrToString(_namePtr)
	global.Echo(_name, "_name")
	if funcTable[_name] == nil {
		return winlib.StringToUTF8Addr("null")
	}
	/*
		global.Echo(ptv.NumMethod(), "tv.NumMethod()")
		if _method, _found = tv.MethodByName(_name); !_found {
			global.Echo(_method, "_method")
			global.Echo(_found, "_found")
			return winlib.StringToUTF8Addr("null")
		}
	*/
	global.Echo("method found!")
	global.Echo(_jsonPtr, "_jsonPtr")
	_json := winlib.UTF8AddrToString(_jsonPtr)
	_input := global.FromJson(_json)
	global.Print(_input, "_input")
	_answer := funcTable["add2"](_input)
	global.Print(_answer, "_answer")
	_output := global.ToPrettyJson(_answer)
	global.Echo(_output, "_output")
	return winlib.StringToUTF8Addr(_output)
}

//export __ThisPath__
func __ThisPath__() uintptr {
	s := winlib.GetModuleFileName((windows.Handle)(unsafe.Pointer(&C.__ImageBase)))
	return winlib.StringToUTF8Addr(s)
}

//export add2
func add2(_json uintptr) uintptr {
	global.Echo(_json, "_json")
	_input := global.FromJson(winlib.UTF8AddrToString(_json))
	global.Print(_input, "_input")
	_map := _input.(map[string]any)
	global.Echo(_map, "_map")
	_args := _map["args"]
	global.Print(_args, "_args")
	//_answer := add2_impl(_args)
	_answer := funcTable["add2"](_args)
	global.Print(_answer, "_answer")
	_map["result"] = _answer
	global.Echo(_map, "_map(2)")
	_output := global.ToPrettyJson(_map)
	global.Echo(_output, "_output")
	return winlib.StringToUTF8Addr(_output)
}

func add2_impl(_args any) any {
	global.Print(_args, "_args")
	_argsAsArray := _args.([]any)
	global.Print(_argsAsArray[0], "[0]")
	global.Print(_argsAsArray[1], "[1]")
	_answer := _argsAsArray[0].(float64) + _argsAsArray[1].(float64)
	global.Print(_answer, "_answer")
	return []any{_answer}
}

//export GetIntFromDLL
func GetIntFromDLL() int32 {
	return 42
}

//export GetStringFromDLL
func GetStringFromDLL() string {
	return "<string from golang: helloハロー©>"
}

//export PrintHello
func PrintHello(name string) {
	fmt.Printf("From DLL: Hello, %s!\n", name)
}

//export PrintBye
func PrintBye() {
	fmt.Println("From DLL: Bye!")
}

//export GetStringPtrFromDLL
func GetStringPtrFromDLL() uintptr {
	msg := "<string from golang: helloハロー©>"
	return winlib.StringToWideAddr(msg)
}

//export PrintHelloPtr
func PrintHelloPtr(name uintptr) {
	nameStr := winlib.WideAddrToString(name)
	fmt.Printf("From PrintHelloPtr: Hello, %s!\n", nameStr)
}

//export GetStringPtrFromDLL2
func GetStringPtrFromDLL2() uintptr {
	msg := "<string from golang: helloハロー©>"
	return winlib.StringToUTF8Addr(msg)
}

//export PrintHelloPtr2
func PrintHelloPtr2(name uintptr) {
	nameStr := winlib.UTF8AddrToString(name)
	fmt.Printf("From PrintHelloPtr2: Hello, %s!\n", nameStr)
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
