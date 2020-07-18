package bd_sotck_exchange


func isValidGroupName(categoryName string) bool{
	if categoryName == "A" || categoryName == "B" || categoryName == "G" || categoryName == "N" || categoryName == "Z"{
		return true
	}
	return false
}
