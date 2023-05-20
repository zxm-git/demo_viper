package main

import (
	"fmt"
	"viper/vip"
)

func main() {
	vip.Init()
	fmt.Println(vip.C)
}
