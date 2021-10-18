package suite

/*
On considère la suite (un) définie par u(n) = 3 * u(n-1) + 5 et u(0) = 0
La fonction suite doit calculer le n-ième terme de (un) de manière récursive :
une fonction qui n'est pas récursive rapportera moins de points.

# Entrée
- n : un entier indiquant le terme de la suite à calculer

# Sortie
- un : le n-ième terme de la suite
*/

func suite(n uint) (un uint) {

	var x uint = 0
	var y uint = 0

	for i  :=0 ; y < n; i++ {
		x = x *3 +5
		y = y+1
	}

	un = x

	return un
}
