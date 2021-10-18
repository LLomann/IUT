package maxoccurrences

/*
Étant donné un tableau d'entiers t, la fonction maxoccurrences doit retourner
l'entier qui apparaît le plus souvent dans t et le nombre de fois que cet
entier apparaît. En cas d'égalité on choisira arbitrairement l'un des entiers
à égalité. Si le tableau est vide on retournera un entier quelconque et 0 pour son nombre d'apparitions.

# Entrées
- t : le tableau dans lequel chercher

# Sortie
- n : l'entier qui apparaît le plus de fois dans t
- occ : le nombre de fois que n apparaît dans t

# Exemple
maxoccurrences([]int{1, 2, 3, 4, 3}) = 3, 2
*/
func maxoccurrences(t []int) (n int, occ int) {

	if t == nil {
		return 69, 0
	}

	if occ == 1 {
		return 69, 1
	}

	var compteur int = 1

	for i := 0; i < len(t); i++ {

		for j := i + 1; j < len(t); j++ {
			if t[i] == t[j] {
				compteur++
			}
		}

		if compteur > occ {
			n = t[i]
			occ = compteur
		}
		compteur = 1
	}

	return n, occ
}
