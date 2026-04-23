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
			Name:"STEVE MADDEN LOIRE", 
			Price:25000, 
			Category:"Heels Wanita", 
			Stock:50,
			Description:"Steve Madden LOIRE Women's Shoes Heels- Rose Gold - Rose Gold",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"HELIVA", 
			Price:450000, 
			Category:"Heels Wanita", 
			Stock:50,
			Description:"Heliva Sherly Women High Heels Glitter Jewel Accesories",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Yongki Komaladi", 
			Price:200000, 
			Category:"Heels Wanita", 
			Stock:50,
			Description:"YONGKI KOMALADI ELIZA HEELS WANITA SGL-272158-A20 GOLD",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Trendyol", 
			Price:700000, 
			Category:"Heels Wanita", 
			Stock:50,
			Description:"Baby Blue Ankle Strap Women's Thin Heeled Shoes Takss25To00097",
			ImageURL:"https://picsum.photos/600/400",
		},
		{
			Name:"Cavasso", 
			Price:244.000, 
			Category:"Heels Wanita", 
			Stock:50,
			Description:"Sintesa Suede - Black",
			ImageURL:"https://ibb.co.com/r258zWd4",
		},

	}
	for _, p := range products {
	config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}