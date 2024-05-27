package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID    uint
    Name  string
    Email string
    Orders []Order
}

type Order struct {
    ID       uint
    UserID   uint
    ProductID uint
    Quantity uint
    User     User
    Product  Product
}

type Product struct {
    ID    uint
    Name  string
    Price float64
}

func ConnectDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=password dbname=testdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func main() {
    db, err := ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    fmt.Println("Connected to the database successfully")

    // Migrasi schema
    if err := db.AutoMigrate(&User{}, &Order{}, &Product{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Contoh data
    db.Create(&User{Name: "John Doe", Email: "john.doe@example.com"})
    db.Create(&Product{Name: "Product A", Price: 100.0})
    db.Create(&Order{UserID: 1, ProductID: 1, Quantity: 2})

    // Membuat context dengan timeout 5 detik
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var results []struct {
        UserName     string
        UserEmail    string
        ProductName  string
        ProductPrice float64
        Quantity     uint
    }

    db.WithContext(ctx).Table("orders").
        Select("users.name as user_name, users.email as user_email, products.name as product_name, products.price as product_price, orders.quantity as quantity").
        Joins("left join users on users.id = orders.user_id").
        Joins("left join products on products.id = orders.product_id").
        Scan(&results)

    for _, result := range results {
        fmt.Printf("User: %s, Email: %s, Product: %s, Price: %.2f, Quantity: %d\n",
            result.UserName, result.UserEmail, result.ProductName, result.ProductPrice, result.Quantity)
    }
}
