package main

import (
	"gogogo/code"
	"gogogo/eat"
	"gogogo/sleep"
)

func main() {
	isAlive := true
	for isAlive {
		eat.Eat("🍚", 2)
		code.Code()
		sleep.Sleep(7)
	}
}
