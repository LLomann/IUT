package facteurspremiers

/*
La fonction facteursPremiers doit retourner un tableau contenant la liste de tous
les facteurs premiers d'un entier n, doublons compris

# Entrée
- n : un nombre entier positif

# Sortie
- facteurs : un tableau contenant tous les facteurs premiers de n, si n vaut 0 il
faut retourner un tableau à zéro éléments.

# Exemple
premiers(24) = [2 2 2 3] (l'ordre n'a pas d'importance)
*/
func facteursPremiers(n uint) (facteurs []uint) {

	var tablo []uint

	if n <= 1 {
		return []uint{}
	}else{
		var i uint = uint(2)
		for n != 1 {
			if n%i == 0{
				tablo = append(tablo, i)
			n = n/i
			}else{
				i++
			}
			
		}

	}
	return tablo
}
