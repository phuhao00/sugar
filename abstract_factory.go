package sugar

type Product interface {
}

type AbstractFactory interface {
	Create() Product
}
