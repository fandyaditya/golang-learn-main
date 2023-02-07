package main

import (
	smartPhone "github.com/fandyaditya/golang-learn-module/smartphone"
)

func main() {
	sm := smartPhone.Samsung{ModelName: "Galaxy S21"}
	xo := smartPhone.Xiaomi{ModelName: "Redmi Note 3000"}
	sm.Nfc()
	//sm.Infrared() //this method call is obsolete in new version module
	xo.Nfc()
	xo.Bluetooth()
}
