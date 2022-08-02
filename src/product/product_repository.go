package product

import "time"

type ProductRepository struct {
	products map[int]Product
}

// Retorno nomeado
func (r ProductRepository) GetById(id int) (p Product) {
	time.Sleep(time.Millisecond * 300)
	p = r.products[id]
	return
}

func (r *ProductRepository) Seed() {
	r.products = make(map[int]Product)
	r.products[1] = Product{Id: 1, Description: "Vassoura elétrica", Brand: "Limpa Limpa LTDA", Price: 39.90}
	r.products[2] = Product{Id: 2, Description: "Panela sem tampa gourmet", Brand: "Tramonfina", Price: 499.99}
	r.products[3] = Product{Id: 3, Description: "Pote de feijão", Brand: "Tapuer", Price: 99.49}
	r.products[4] = Product{Id: 4, Description: "Saco de lixo 50L", Brand: "Dexter & Dexter", Price: 29.90}
	r.products[5] = Product{Id: 5, Description: "Farinha branca 1Kg", Brand: "Trigo Escobar", Price: 14.89}
}

func (r ProductRepository) Clear() {
	for k := range r.products {
		delete(r.products, k)
	}
}
