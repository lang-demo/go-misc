package main

import "github.com/lang-library/go-global"

func main() {
	global.Echo(global.ExeDir(), "global.ExeDir()")
}
