package types

type Album struct {
	ID     uint    `form:"id" json:"id"`
	Title  string  `form:"title" json:"title"`
	Artist string  `form:"artist" json:"artist"`
	Price  float64 `form:"price" json:"price"`
}
