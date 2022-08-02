package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mrbrunelli/go-concurrency/src/product"
)

func getProduct(id int, repo *product.ProductRepository, ch chan product.Product) {
	fmt.Printf("Buscando produto %v\n", id)

	product := repo.GetById(id)
	ch <- product

}

func main() {
	start := time.Now()

	repo := product.ProductRepository{}
	repo.Seed()

	ids := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	ch := make(chan product.Product)

	for _, id := range ids {
		// Preciso avisar a main thread que estou criando +1 goroutine
		wg.Add(1)

		go func(id int) {
			// Avisar a main thread que essa goroutine terminou
			defer wg.Done()
			getProduct(id, &repo, ch)
		}(id)
	}

	// Esperar todas goroutines concluírem e então fechar o channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Enquanto houver mensagens no channel, irá consumir
	for product := range ch {
		fmt.Println(product)
	}

	fmt.Println(time.Since(start).Seconds(), "segundos")
}
