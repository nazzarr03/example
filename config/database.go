package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nazzarr03/example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ( // Db değişkeni global olarak tanımlanıyor ki diğer dosyalardan da erişilebilsin.
	Db *gorm.DB // Diğer dosyalardan erişilebilmesi için baş harfi büyük olmalı.
) // Diğer dosyalardan eişilmesini isteme sebebimiz veritabanına ekleme çıkarma yaparken bunu kullanmamız.

func init() { // init fonksiyonu main fonksiyonundan önce çalışır yani her şeyden önce bunun içindeki işlemler gerçekleşir.
	if err := godotenv.Load(); err != nil { // .env dosyasını yükler. Eğer hata varsa hata mesajını döndürür.
		fmt.Println(".env file not found") // .env dosyasını burada bir kere kontrol ediyoruz sonrasında kontrol etmeye gerek yok.
	}
	ConnectDB() // ConnectDB fonksiyonunu init fonksiyonu içinde çağırıyoruz ki main fonksiyonundan önce çalışsın.
}

// Veritabanı olarak postgresql kullandık.
// Veritabanını docker-compose.yml dosyası ile ayağa kaldırdık.
func ConnectDB() {
	var err error       // Hata kontrolü için bir değişken tanımladık. error modelinden oluşmuş err değişkeni.
	dsn := fmt.Sprintf( // Veritabanına bağlanmak için gerekli olan bilgileri dsn değişkenine atıyoruz.
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Veritabanına bağlanıyoruz.

	if err != nil {
		panic("failed to connect database")
	}

	if err := Db.AutoMigrate(&models.Book{}); err != nil { // models.Book{} ile Book modelini veritabanına ekliyoruz.
		panic("failed to migrate database")
	}

	fmt.Println("Database connected successfully!")
}
