package matrice

import (
	"errors"
)

var errPasMat error = errors.New("mat n'est pas vraiment une matrice")

/*
La fonction compte indiquera combien de fois un nombre n apparaît dans une matrice.

# Entrées
- n : le nombre à chercher
- mat : une matrice d'entiers (un tableau de tableaux qui font tous la même taille)

# Sorties
- num : le nombre de fois que n apparaît dans mat
- err : errPasMat si mat n'est pas une vraie matrice (toutes les lignes n'ont pas
la même longueur ou mat vaut nil ou une ligne vaut nil)

# Exemple
compte(0, [][]int{[]int{0, 1}, []int{0, 0}}) = 3
*/
func compte(n int, mat [][]int) (num int, err error) {

	if mat == nil {
		return 0, errPasMat
	}

	if len(mat) == 0 {
		return 0, nil
	}

	var res int
	var longueur int = len(mat[0])
	for _, tab := range mat { //	fo i:=0; i<len(mat); i++ { 	//	tab = mat[i]
		if len(tab) != longueur || tab == nil {
			return 0, errPasMat
		}
		for _, el := range tab { //	fo j:=0; len(tab); j++ 	//		el = tab[j]
			if el == n {
				res += 1
			}
		}
	}

	return res, nil
}
