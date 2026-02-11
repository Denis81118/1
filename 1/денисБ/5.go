package main

import ("fmt"; "strings")
func validateUser(booboo []string) {
	if len(booboo) < 3  {
		fmt.Println("недостатончно данных.")
	}
var name string
var age int64
var email string

if age <=17 {
	fmt.Println("ошибка: мало лет")
return
}

if len(name) >= 50 {
	fmt.Println("ошибка: длинное имя")
return
}

if !strings.Contains(email, "@") {
	fmt.Println("ошибка:неправильный email")
return
}
fmt.Println("все норм")
}
func main(){
	validateUser([]string{"amir", "17", "amir@gmail.com"})
}