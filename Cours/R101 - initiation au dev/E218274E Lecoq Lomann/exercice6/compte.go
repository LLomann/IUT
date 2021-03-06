package compte

/*
La fonction compte doit indiquer combien de fois une valeur v apparaît dans un
tableau tab.

# Entrées
- tab : un tableau d'entiers
- v : la valeur à chercher

# Sortie
- num : le nombre de fois que v apparaît dans tab
*/

func compte(tab []int, v int) (num int) {

	if tab == nil {
		return 0
	}

	for i:=0; i<len(tab); i++ {
		if tab[i] == v {
			num += 1
		}
	}

	return num
}
