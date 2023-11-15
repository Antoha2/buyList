package repository

import "fmt"

func (*repositoryImplDB) Create(unit *RepBuyUnit) error {
	fmt.Println("rep ", unit)
	return nil
}
