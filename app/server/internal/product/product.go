package product

type Product struct {
	Id    uint64  `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float32 `json:"price" db:"price"`
	Rank  int     `json:"rank" db:"rank"`
}

// Products can be used to inject abilities to Product slices like sorting, filtering, etc.
type Products []Product
