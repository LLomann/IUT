package doublons

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement une fois chaque nombre compris entre 1 et n. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

Coder la fonction doublons de manière à ne parcourir le tableau tab qu'une seule
fois rapportera plus de points.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si tab contient bien exactement une
fois chaque entier entre 1 et len(tab) et false sinon
*/

func doublons(tab []int) (ok bool) {

	ok = true

	 for i:=0; i<len(tab); i++ {

		if tab[0] != 1 {
			return false
		}
		for j:=i+1; j<len(tab); j++ { 
			if tab[i] == tab[j] {
				ok = false
				break
			}
		}
	break
	}

	return ok
}
