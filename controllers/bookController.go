package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/example/config"
	"github.com/nazzarr03/example/models"
)

// Buraya yazacağımız fonksiyonlar routes klasöründeki çağırılacak.
// Bu fonksiyonlar "endpoint" veya "handler" olarak adlandırılır.

// GetBooks fonksiyonu tüm kitapları getirir.
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book // models.Book{} modelinden oluşmuş bir slice tanımladık ve ismi books.

	config.Db.Find(&books) // Veritabanındaki tüm kitapları books değişkenine atıyoruz.

	return c.Status(fiber.StatusOK).JSON(fiber.Map{ // books değişkenini json formatında döndürüyoruz.
		"status": "success", // status anahtarına success değerini atıyoruz.
		"data":   books})    // data anahtarına books değişkenini atıyoruz.
}

// GetBook fonksiyonu id'si verilen kitabı getirir. id'yi url'den alır.
func GetBookByID(c *fiber.Ctx) error {
	var book models.Book // models.Book{} modelinden oluşmuş bir değişken tanımladık ve ismi book.
	id := c.Params("id") // id'yi url'den alıyoruz.

	config.Db.Find(&book, id) // Veritabanındaki id'si url'den alınan kitabı bularak book değişkenine atıyoruz.
	if book.ID == 0 {         // Eğer kitap bulunamazsa yani id'si 0 ise
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ // 404 not found hatası döndürür.
			"status":  "error",           // status anahtarına error değerini atıyoruz.
			"message": "book not found"}) // message anahtarına book not found değerini atıyoruz.
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{ // book değişkenini json formatında döndürüyoruz.
		"status": "success", // status anahtarına success değerini atıyoruz.
		"data":   book})     // data anahtarına book değişkenini atıyoruz.
}

// CreateBook fonksiyonu yeni bir kitap ekler.
func CreateBook(c *fiber.Ctx) error {
	var book models.Book                        // models.Book{} modelinden oluşmuş bir değişken tanımladık ve ismi book.
	if err := c.BodyParser(&book); err != nil { // Request body'sini book değişkenine atıyoruz. Gelen isteği parse ediyoruz.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // 400 bad request hatası döndürür.
			"status":  "error",              // status anahtarına error değerini atıyoruz.
			"message": "cannot parse body"}) // message anahtarına cannot parse body değerini atıyoruz.
	}

	// Gelen istek doğru mu yani benim modelimdeki alanları vermiş mi diye kontrol yapıyoruz.
	if book.Title == "" || book.Author == "" { // Eğer title veya author boş ise
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // 400 bad request hatası döndürür.
			"status":  "error",                     // status anahtarına error değerini atıyoruz.
			"message": "title or author is empty"}) // message anahtarına title or author is empty değerini atıyoruz.
	}

	config.Db.Create(&book) // Veritabanına book değişkenini ekliyoruz.

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{ // 201 created durum kodu döndürür.
		"status": "success", // status anahtarına success değerini atıyoruz.
		"data":   book})     // data anahtarına book değişkenini atıyoruz.
}
