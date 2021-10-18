package nombrespremiers2

/*
La fonction selectionPremiers filtre le contenu d'un tableau d'entiers pour ne garder que ceux qui sont des nombres premiers, en éliminant les doublons.

# Entrée
- t : un tableau d'entiers

# Sortie
- p : un tableau contenant tous les nombres premiers de t, sans doublons.
Si t est vide, p doit être identique à t.

# Exemple
selectionPremiers([]int{1, 2, 2, 3, 4, 5}) = [2 3 5] (l'ordre n'a pas d'importance)
*/

func selectionPremiers(tab []int) (p []int) {

	if tab == nil {
		return nil
	}

	if len(tab) == 0 {
		return []int{}
	}

	var b bool
	p = []int{}

	for i := 0; i < len(tab); i++ {

		if tab[i] >= 2 {

			b = true

			for j := 2; j < tab[i]; j++ {
				if tab[i]%j == 0 {
					b = false
					break
				}
			}

			if b {
				for _, val := range p { // for i:=0 ; i<len(p); i++ {   ||   val = p[i]
					if val == tab[i] {
						b = false
						break
					}
				}
				if b {
					p = append(p, tab[i])
				}
			}
		}
	}

	return p
}
