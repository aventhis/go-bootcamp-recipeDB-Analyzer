package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
)

// PrettyPainting - Функция для конвертации и вывода данных в противоположном формате
func PrettyPainting(recipe *Recipes, fType Filetype) error {
	switch fType {
	case JSON:
		data, err := xml.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return errors.New("ошибка при кодировании данных в JSON")
		}
		fmt.Println(string(data))
	case XML:
		data, err := json.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return errors.New("ошибка при кодировании данных в XML")
		}
		fmt.Println(string(data))
	}
	return nil
}
