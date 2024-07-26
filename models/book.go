package models

type Book struct {
	ID     int    `json:"id"`    // json olarak verdiğimiz isimler ile struct içindeki isimler aynı olmalı
	Title  string `json:"title"` // json verme sebebimiz json formatında veri döndürürken bu isimlerin kullanılmasını sağlamak
	Author string `json:"author"`
}
