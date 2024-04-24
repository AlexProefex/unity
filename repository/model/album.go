package model

type Album struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

/*
type Album struct {
	ID     uint   `gorm:"not null"`
	Title  string  `gorm:"type:varchar(20) not null"
	Artist string  `gorm:"type:varchar(20) not null"
	Price  float64 `gorm:"type:varchar(20) not null"
}*/
