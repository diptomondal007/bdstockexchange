package bdstockexchange

import (
	"log"
	"strconv"
	"strings"
)

func isValidCategoryName(categoryName string) bool {
	if categoryName == "A" || categoryName == "B" || categoryName == "G" || categoryName == "N" || categoryName == "Z" {
		return true
	}
	return false
}

func normalizeAmerican(old string) string {
	return strings.Replace(old, ",", "", -1)
}

func toFloat64(text string) float64 {
	if text == "--" {
		text = strings.Replace(text, "--", "0", -1)
	}
	val, err := strconv.ParseFloat(normalizeAmerican(text), 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func toInt64(text string) int64 {
	//if strings.Contains(text, `"`){
	//	text = strings.Replace(text,`"`, "", -1)
	//}
	//if strings.Contains(text,":"){
	//	log.Println("contains :")
	//	text = strings.Replace(text, ":", "",1)
	//}
	val, err := strconv.ParseInt(normalizeAmerican(text), 10, 64)
	if err != nil {
		log.Println(err)
	}
	return val
}

func toInt(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
