
package main
import "fmt"
func main() {

var osnov float64

var rychnoi float64

var dopRychnoi float64

fmt.Print("напишите вес основного богажа в кг" ) 
fmt.Scan(&osnov)
fmt.Print("напишите вес ручного богажа в кг" )
fmt.Scan(&rychnoi)
fmt.Print("напишите вес дополнительного ручного богажа в кг" ) 
fmt.Scan(&dopRychnoi)
var y float64 = osnov + rychnoi + dopRychnoi 
fmt.Println("общий вес богажа в кг:", y)
}