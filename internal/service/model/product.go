package model

import "github.com/maxwelbm/pod_example/internal/repository"

type Product struct {
	Id           int
	Name         string
	Quantity     int
	Code_value   string
	Is_published bool
	Expiration   string
	Price        float64
}

type ServiceProducts struct {
	Storage repository.RepositoryDB
}

func (s *ServiceProducts) Create(product Product) (Product, error) {
	product, err := s.Storage.Create(product)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func NewServiceProducts(storage repository.RepositoryDB) *ServiceProducts {
	return &ServiceProducts{
		Storage: storage,
	}
}

// func FillProductList(db *map[int]*handlers.Product) {
// 	file, err := os.Open("products.json")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}

// 	defer file.Close()

// 	var products []handlers.Product
// 	if err := json.NewDecoder(file).Decode(&products); err != nil {
// 		fmt.Println("Error decoding JSON:", err)
// 		return
// 	}

// 	for _, product := range products {
// 		(*db)[product.Id] = &product
// 	}
// }
