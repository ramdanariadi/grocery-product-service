package utils

import (
	"encoding/csv"
	"go-tunas/helpers"
	"go-tunas/models"
	"os"
	"strconv"
)

func ReadCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	helpers.PanicIfError(err)

	reader := csv.NewReader(file)
	allcsv, err := reader.ReadAll()
	helpers.PanicIfError(err)

	return allcsv
}

func ProductsFromCSV(filePath string) []models.ProductModel {
	csvString := ReadCsv(filePath)
	var products []models.ProductModel
	for _, line := range csvString {
		productTmp := models.ProductModel{}
		for index, field := range line {
			switch index {
			case 0:
				productTmp.Id = field
			case 1:
				productTmp.Deleted, _ = strconv.ParseBool(field)
			case 6:
				productTmp.Price, _ = strconv.ParseInt(field, 10, 64)
			case 7:
				atoi, _ := strconv.Atoi(field)
				productTmp.Weight = uint(atoi)
			case 8:
				productTmp.CategoryId = field
			case 5:
				parseInt, _ := strconv.ParseInt(field, 10, 32)
				productTmp.PerUnit = int(parseInt)
			case 2:
				productTmp.Description = field
			case 3:
				productTmp.ImageUrl = field
			case 4:
				productTmp.Name = field
			}
		}
		products = append(products, productTmp)
	}
	return products
}

func ProductsFromCSVWithChannel(filePath string, channel chan models.ProductModel) {
	csvString := ReadCsv(filePath)
	for _, line := range csvString {
		productTmp := models.ProductModel{}
		for index, field := range line {
			switch index {
			case 0:
				productTmp.Id = field
			case 1:
				productTmp.Deleted, _ = strconv.ParseBool(field)
			case 6:
				productTmp.Price, _ = strconv.ParseInt(field, 10, 64)
			case 7:
				atoi, _ := strconv.Atoi(field)
				productTmp.Weight = uint(atoi)
			case 8:
				productTmp.CategoryId = field
			case 5:
				parseInt, _ := strconv.ParseInt(field, 10, 32)
				productTmp.PerUnit = int(parseInt)
			case 2:
				productTmp.Description = field
			case 3:
				productTmp.ImageUrl = field
			case 4:
				productTmp.Name = field
			}
		}
		channel <- productTmp
	}
}
