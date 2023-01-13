package domain

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   int
	IsPublished bool
	Expiration  string
	Price       float64
}
