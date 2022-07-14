package models

import (
	"errors"
	"log"
)

type tea struct {
	ID    uint   `json:"id"` // TODO: Generate ID's
	Title string `json:"title"`
	Color color  `json:"color"`
}

// Tea color
type color struct {
	ID    byte   `json:"id"`
	Color string `json:"color"`
}

// Color for teaColor field
var teaColor = map[byte]string{
	0: "Красны",
	1: "Белый",
	2: "Зеленый",
	3: "Шу Пуэр",
	4: "Шен Пуэр",
	5: "Черный",
}

// Convert choosing color to color struct
func convertColor(col string) (*color, error) {
	for i, c := range teaColor {
		if c == col {
			return &color{
				ID:    i,
				Color: c,
			}, nil
		}
	}
	return nil, errors.New("такого чая не существует")
}

func NewTea(id uint, title string, color string) tea {
	col, err := convertColor(color)
	if err != nil {
		log.Println(err)
	}

	return tea{
		ID:    id,
		Title: title,
		Color: *col,
	}
}
