package main

import "fmt"

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
	Quantity   int
}

//Criando a função getTotal, no qual seu retonar será um float64 e adicionei a mesma
// na struct Order, ou seja essa função e um metodo da struct
func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

// Função setPrice, a mesma recebe um price como parametro e atribui o price ao
// price da struct, essa função é um metodo da struct
func (o *Order) setPrice(price float64) {
	fmt.Println(o.Price)
	o.Price = price
	fmt.Println(o.Price)
}

// Função construtora, a mesma retorna um ponteiro para uma nova instacia do order
func NewOrder() *Order {
	return &Order{}
}

func main() {
	order := NewOrder()
	order.Quantity = 10
	order.Price = 15.0

	//Criando uma order
	// order := Order{
	// 	Id:         "123",
	// 	Price:      10.0,
	// 	Tax:        5.0,
	// 	FinalPrice: 25.0,
	// 	Quantity:   4,
	//}

	//imprimindo os dados da order criada
	//fmt.Println(order.Id, order.Price, order.Tax, order.FinalPrice)

	//Imprimindo o valor total
	//fmt.Println(order.getTotal())

	//Adicionando um novo valor ao price pela função do Order
	//order.setPrice(50.25)
	//fmt.Println(order.getTotal())
	// valor será de 40, mesmo adicionando o 50.25, quando adiciono o ponteiro no
	// (O *Order) do setPrice o valor fica 201

	// na linguem GO não temos um construtor, para inicializar uma struct
	// então é utilizado uma função para inicializar
}
