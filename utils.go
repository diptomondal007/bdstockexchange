package bdstockexchange

import (
	"log"
	"strconv"
	"strings"
)

// isValidCategoryName checks if the user input catergory is valid
func isValidCategoryName(categoryName string) bool {
	if categoryName == "A" || categoryName == "B" || categoryName == "G" || categoryName == "N" || categoryName == "Z" {
		return true
	}
	return false
}

// normalizeAmerican return the cleaned string if the input string is in american format ex : 10,000
func normalizeAmerican(old string) string {
	return strings.Replace(old, ",", "", -1)
}

// toFloat64 returns the float64 cleaning the input string
func toFloat64(text string) float64 {
	if text == "--" {
		text = strings.Replace(text, "--", "nan", -1)
	}
	if text == "N/A"{
		text = strings.Replace(text, "N/A", "nan", -1)
	}
	val, err := strconv.ParseFloat(normalizeAmerican(text), 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// toInt64 returns the int64 cleaning the input string
func toInt64(text string) int64 {
	if strings.Contains(text, " "){
		text = strings.Replace(text, " ","", -1)
	}
	val, err := strconv.ParseInt(normalizeAmerican(text), 10, 64)
	if err != nil {
		log.Println(err)
	}
	return val
}

// toInt parse the int from a input string
func toInt(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
