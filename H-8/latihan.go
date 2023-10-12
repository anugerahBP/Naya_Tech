package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// Konfigurasi koneksi ke database
	dsn := "user=postgres password=@nugerahBP26 dbname=postgres host=db.vbxvyoewbdvxmiiwokba.supabase.co port=5432 sslmode=require"

	// Membuka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Successfully connected to the database")

	// Melakukan migrasi data dari database(opsional, tergantung kebutuhan)
	db.AutoMigrate(&User{})

	// DML: Menambahkan data
	user1 := User{Name: "Ibrahim"}
	db.Create(&user1)

	// DQL: Mengambil data
	var users []User
	db.Find(&users)

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}

	// DML: Memperbarui data
	db.Model(&user1).Update("Name", "Jose")

	// Menghapus data berdasarkan nama
	db.Where("name = ?", "Kim").Delete(&User{})
}
