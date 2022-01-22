package main

import "fmt"

func main(){
    var nilai32 int32 = 11010101
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	// hati hati dalam melakukan konversi karena masing masing tipe data memiliki batasan

	// int 16 akan mengeluarkan output 53 karena nilai input konversi melwati batas int16

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16,"\n\n")


	// markicob konversi terhapad string
	// markicob kita ambil nilai pertama dari string

	var name = "Nden"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}