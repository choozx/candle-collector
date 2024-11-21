package symbols

import (
	"candle-collector/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type Symbol struct {
	Code     int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:name" json:"name"`
	IsUpdate bool   `gorm:"column:is_update"`
}

var Symbols []Symbol

func InitCandle() {
	config.DB.Where("is_update=?", true).Find(&Symbols)
	fmt.Println("Symbols:", Symbols)
}

func SetSymbol(writer http.ResponseWriter, request *http.Request) {
	requestSymbol := new(Symbol)
	err := json.NewDecoder(request.Body).Decode(requestSymbol)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}

	if requestSymbol.Name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	symbol := new(Symbol)
	config.DB.Where("name = ?", requestSymbol.Name).First(symbol)
	if symbol.IsUpdate == true {
		writer.WriteHeader(http.StatusBadRequest)
		return // 중복 등록 체크
	}

	symbol.IsUpdate = true
	if symbol.Name == "" {
		symbol.Name = requestSymbol.Name
		config.DB.Create(symbol)
	} else {
		config.DB.Where("name=?", symbol.Name).Save(symbol)
	}

	Symbols = append(Symbols, *symbol)
}

func DeleteSymbol(writer http.ResponseWriter, request *http.Request) {
	requestSymbol := new(Symbol)
	err := json.NewDecoder(request.Body).Decode(requestSymbol)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}

	var result []Symbol
	for _, symbol := range Symbols {
		if symbol.Name != requestSymbol.Name {
			result = append(result, symbol)
		}
	}
	Symbols = result
	fmt.Println("Symbols:", Symbols)

	symbol := new(Symbol)
	config.DB.Where("name = ?", requestSymbol.Name).First(symbol)
	if symbol.Name == "" {
		return // 제거할 대상이 없다면
	}

	symbol.IsUpdate = false
	config.DB.Where("name=?", symbol.Name).Save(symbol)
}

func (Symbol) TableName() string {
	return "symbol"
}
