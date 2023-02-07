package main

import (
	smartPhone "github.com/fandyaditya/golang-learn-module/smartphone"
)

func main() {
	sm := smartPhone.Samsung{ModelName: "Galaxy S21"}
	xo := smartPhone.Xiaomi{ModelName: "Redmi Note 3000"}
	sm.Nfc()
	sm.Infrared()
	xo.Nfc()
	xo.Bluetooth()
}
