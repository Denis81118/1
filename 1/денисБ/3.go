package main
import "fmt"
type Order struct {
    id          int
    items       []int
    total       float64
    address     string
    isCompleted bool
}

func main() {
    addOrder(1, []int{101, 102}, 1500.50, "главная", false)
    addOrder(2, []int{201}, 999.99, "центр", true)
	 fmt.Println("ЗАказы:")
     
    for id, order := range orders {
        fmt.Println("номер:", id)
        fmt.Println("Товары:", order.items)
        fmt.Println("Сумма:", order.total)
        fmt.Println("Адрес:", order.address)
        fmt.Println("Выполнен:", order.isCompleted)
       
}
}
var orders = make(map[int]Order)

func addOrder(id int, items []int, total float64, address string, isCompleted bool) {
    orders[id] = Order{id, items, total, address, isCompleted}
}