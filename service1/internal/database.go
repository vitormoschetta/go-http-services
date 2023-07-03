package internal

type Repository struct {
	storage []Product
}

func NewRepository() *Repository {
	return &Repository{
		storage: []Product{},
	}
}

func (r *Repository) FindAll() (products []Product, err error) {
	return r.storage, nil
}

func (r *Repository) FindById(id string) (product Product, err error) {
	for _, product := range r.storage {
		if product.ID == id {
			return product, nil
		}
	}
	return
}

func (r *Repository) Save(product Product) (err error) {
	r.storage = append(r.storage, product)
	return nil
}

func (r *Repository) Update(product Product) (err error) {
	for i, p := range r.storage {
		if p.ID == product.ID {
			r.storage[i] = product
			return nil
		}
	}
	return
}

func (r *Repository) Delete(id string) (err error) {
	for i, p := range r.storage {
		if p.ID == id {
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
			return nil
		}
	}
	return
}
