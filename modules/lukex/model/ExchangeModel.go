package model

type ExchangeModel struct {
	IDExchange int `gorm:"PRIMARY_KEY";"Column:id_exchange"`
	Timestamp string `gorm:"Column:timestamp"`
	Provider string `gorm:"Column:provider"`
	Exchange string `gorm:"Column:exchange"`
	Note	string `gorm:"Column:note"`
}
