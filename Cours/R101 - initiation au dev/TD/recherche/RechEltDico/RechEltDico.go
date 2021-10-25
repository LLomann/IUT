package RechEltDico

import "errors"

var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")

func RechElt(tab []int, n int)(b bool, pos int, err error){

	var debut int = 0
	var fin int = len(tab)
	var milieu int 
	for debut < fin {
		milieu = (debut + fin)/ 2 
		if tab[milieu] == x {
			return true, milieu, nil
		}
		if tab[milieu] > x {
			fin = milieu

		}else {
			debut = milieu +1
		}

	}
	return false, -1, nil
	
}