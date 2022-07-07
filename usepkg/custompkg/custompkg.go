package custompkg

import (
	"fmt"
	"LEARNGOLANG/usepkg/exinit"
)
type Student struct {
	Name string
	Age int
	score int
}
var Ratio int
var ttt int
const PI = 3.14
const pI2 = 3.1415
func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD()
}
func printCustom2() {
	fmt.Println("this is custom package!2");
}