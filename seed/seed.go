package main

import (
	"log"

	"github.com/ginaa07/ReginaSafarina-gin-firebase-backend/config"
	"github.com/ginaa07/ReginaSafarina-gin-firebase-backend/models"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
			Name:"Nasi Goreng Spesial", 
			Price:25000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Nasi goreng dengan telur dan ayam",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Es Jeruk", 
			Price:5000, 
			Category:"Minuman", 
			Stock:50,
			Description:"Minuman Sachetan",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Mie Gacoan", 
			Price:20000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Mie Pedas No 1 di Indonesia",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Martabak Manis", 
			Price:15000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Martabak manis dengan isi cokelat dan krim",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Nasi Padang", 
			Price:25000, 
			Category:"Makanan", 
			Stock:50,
			Description:"Enak parah cuyy",
			ImageURL:"https://picsum.photos/600/400",
		},

	}
	for _, p := range products {
	config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}