package main

import (
	"fmt"

	"github.com/mrbrunelli/go-concurrency/src/product"
)

func getProduct(id int, repo *product.ProductRepository, ch chan string) {
	product := repo.GetById(id)
	ch <- fmt.Sprintf("Produto %s da marca %s está saindo por apenas R$ %v", product.Description, product.Brand, product.Price)
}

func main() {
	repo := product.ProductRepository{}
	repo.Seed()

	ids := [5]int{1, 2, 3, 4, 5}

	ch := make(chan string)

	for _, id := range ids {
		fmt.Printf("Buscando produto %v\n", id)
		go getProduct(id, &repo, ch)
	}

	for v := range ch {
		fmt.Println(v)
	}

}
