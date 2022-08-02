package main

import (
	"fmt"
	"sync"

	"github.com/mrbrunelli/go-concurrency/src/product"
)

func getProduct(id int, repo *product.ProductRepository, ch chan product.Product) {
	fmt.Printf("Buscando produto %v\n", id)

	product := repo.GetById(id)
	ch <- product

}

func main() {
	repo := product.ProductRepository{}
	repo.Seed()

	ids := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	ch := make(chan product.Product)

	for _, id := range ids {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			getProduct(id, &repo, ch)
		}(id)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for product := range ch {
		fmt.Printf("O produto %s da marca %s está saindo por apenas %v\n", product.Description, product.Brand, product.Price)
	}

}
