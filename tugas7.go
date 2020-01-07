package main

import "fmt"
import "runtime"
import "reflect"

func main() {
	runtime.GOMAXPROCS(2)

	go bacatipe1()
	bacatipe2()

	var input string
	fmt.Scanln(&input)
}

func bacatipe1() {
	var umur = 24
	var reflectValue = reflect.ValueOf(umur)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}

func bacatipe2() {
	var nama = "Andika"
	var reflectValue = reflect.ValueOf(nama)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.String {
		fmt.Println("nilai variabel :", reflectValue.String())
	}
}
