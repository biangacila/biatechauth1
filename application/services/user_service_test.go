package services

import (
	"fmt"
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"github.com/biangacila/luvungula-go/global"
	"testing"
)

func initProducts(t *testing.T) []aggregates.ProductAggregate {
	beer, err := aggregates.NewProduct("Beer", "Healthy", 50)
	if err != nil {
		t.Fatal(err)
	}
	peanut, err := aggregates.NewProduct("Peanut", "Snacks", 20)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := aggregates.NewProduct("Wine", "masty drink", 20)
	if err != nil {
		t.Fatal(err)
	}
	return []aggregates.ProductAggregate{beer, peanut, wine}
}

func TestUserService(t *testing.T) {
	products := initProducts(t)
	uo, err := NewUserService(
		WithMemoryUserRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	user, err := aggregates.NewUser("biatechauth01", "UC1001", "User 1",
		"Sur 1", "user1@bia.com", "0729139504", "admin", "admin1", "987654321")
	if err != nil {
		t.Error(err)
	}
	err = uo.users.Add(user)
	if err != nil {
		t.Error(err)
	}

	_id, _err := uo.users.Get("UC100")
	if _err != nil {
		t.Error(err)
	}
	global.DisplayObject("users in memory", _id)
	var order = []string{
		products[0].AggregateID(),
	}
	useId, err := user.GetUserID()
	fmt.Println(">>User id: ", useId)
	if err != nil {
		t.Error(err)
	}
	_, err = uo.CreateOrder(useId, order)
	if err != nil {
		t.Error(err)
	}
}
