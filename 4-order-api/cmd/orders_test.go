package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-adv/4-order-api/internal/auth"
	"go-adv/4-order-api/internal/order"
	"go-adv/4-order-api/pkg/models"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var productMock = &models.Product{
	Name:        "Apple Watch Ultra 2",
	Description: "Apple Watch Ultra 2 GPS + Cellular, 49 мм, корпус из черного титана, ремешок Milanese черного цвета, размер M",
	Images:      []string{"img1.png"},
}

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&models.User{
		PhoneNumber: "89312222225",
		SessionId:   "JlumraZI5lTHAr6xYumbU4q6hvCu_FF6VwlMK47JygY=",
		Code:        "4849",
	})

	db.Create(productMock)
}

func removeData(db *gorm.DB) {
	db.Unscoped().Where("phone_number = ?", "89312222225").Delete(&models.User{})
	db.Unscoped().Where("name = ?", "Apple Watch Ultra 2").Delete(&models.Product{})
}

func TestNewOrder(t *testing.T) {
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.UserVerifyRequest{
		SessionId: "JlumraZI5lTHAr6xYumbU4q6hvCu_FF6VwlMK47JygY=",
		Code:      "4849",
	})

	res, err := http.Post(ts.URL+"/auth/verify", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData auth.VerifyResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token empty")
	}

	var productsSlice []*models.Product
	productsSlice = append(productsSlice, productMock)

	productData, _ := json.Marshal(&order.OrderCreateRequest{
		Products: productsSlice,
	})

	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewReader(productData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", resData.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Fatalf("Expected %d got %d", 201, res.StatusCode)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var orderResData uint
	err = json.Unmarshal(body, &orderResData)
	if err != nil {
		t.Fatal(err)
	}
	if orderResData == 0 {
		t.Fatal("ID empty")
	}

	removeData(db)
}
