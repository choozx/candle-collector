package symbols

import (
	"candle-collector/internal/config"
	"fmt"
)

type Symbol struct {
	Code int    `gorm:"column:code"`
	Name string `gorm:"column:name"`
}

var Symbols []Symbol

func init() {
	config.DB.Find(&Symbols)
	fmt.Println("Symbols:", Symbols)
}

func (Symbol) TableName() string {
	return "symbol"
}
