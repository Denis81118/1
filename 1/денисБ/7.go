package main

import "fmt"

type employee struct {
id int64
name string
position string
salary float64
}
func main(){
	employees := []employee{
		{1, "amir", "разработчик", 100000},
		{1, "amir2", "уборщик", 102000},
		{1, "amir3", "тестировщик", 110000},
	}
	total, avg := fff(employees)
	fmt.Println("общий фонд", total)
	fmt.Println("средняя зп", avg)
}
func fff(employees []employee) (float64, float64) {

 total := 0.0
 for _, e := range employees {
	total += e.salary
 } 
 avg := total / float64(len(employees))
 return total, avg
}
