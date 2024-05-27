package main
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   uint  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255)"`
}

type Product struct {
	ID   uint  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
	Price int `gorm:"type:int"`
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Product{})
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=ada-onb0001 password=Kmzway87a@ dbname=belajar_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}


func CreateUser(db *gorm.DB, user User) error {
	return db.Create(&user).Error
}

// func UpdateUserEmail(db *gorm.DB, id int, newEmail string) error {
// 	var user User
// 	if err := db.First(&user, id).Error; err != nil {
// 		return err
// 	}
// 	user.Email = newEmail
// 	log.Println(user)
// 	return db.Save(&user).Error
// }

// func DeleteUser(db *gorm.DB, id int) error {
// 	return db.Delete(&User{}, id).Error
// }

// func FindUserById(db *gorm.DB, id int) (User, error) {
// 	var user User
// 	if err := db.First(&user, id).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func FindUserAll(db *gorm.DB) ([]User, error) {
// 	var users []User
// 	if err := db.Find(&users).Error; err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func FindUserByEmail(db *gorm.DB, email string) (User, error) {
// 	var user User
// 	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func FindUserByName(db *gorm.DB, name string) (User, error) {
// 	var user User
// 	if err := db.Where("name LIKE ?", "%" + name + "%").First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func UpdateNameByEmail(db *gorm.DB, email string, newName string) error {
// 	var user User
// 	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
// 		return err
// 	}
// 	user.Name = newName
// 	return db.Save(&user).Error
// }

func DeleteUserByEmail(db *gorm.DB, email string) error {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return db.Delete(&user).Error
}


func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database succesfully", db)
	if err := MigrateDB(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Migrated table succesfully")

	// user := User{
	// 	Name: "Bracmatya Doe",
	// 	Email: "bracmatya@mail.com",
	// }
	// if err := CreateUser(db, user); err != nil {
	// 	log.Fatalf("Failed to create user: %v", err)
	// }

	// if err := UpdateUserEmail(db, 1, "mail@mail.com"); err != nil {
	// 	log.Fatalf("Failed to update user: %v", err)
	// }

	// if err := DeleteUser(db, 2); err != nil {
	// 	log.Fatalf("Failed to delete user: %v", err)
	// }

	// user, err := FindUserById(db,1)
	// if err != nil {
	// 	log.Fatalf("Failed to find user: %v", err)
	// }
	// log.Println("User found: ", user)

	// users, err := FindUserAll(db)
	// if err != nil {
	// 	log.Fatalf("Failed to find user: %v", err)
	// }
	// log.Println("User found: %+v\n", users)

	// dataUser, err := FindUserByEmail(db, "anggiet@mail.com")
	// if err != nil {
	// 	log.Fatalf("Failed to find user: %v", err)
	// }

	// dataUser , err := FindUserByName(db, "Anggiet")
	// if err != nil {
	// 	log.Fatalf("Failed to find user: %v", err)
	// }

	if err := DeleteUserByEmail(db, "mail@mail.com"); err != nil {

		log.Fatalf("Failed to find user: %v", err)
	}

	log.Println("User updated succesfully")


	// log.Println("User deleted succesfully")
}