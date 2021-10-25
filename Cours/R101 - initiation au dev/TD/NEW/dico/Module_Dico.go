package dico

import "errors"

var errFormatTableau error = errors.New("tu n'a pas rentrer un tableau")

func RechEltDico(tab []int, el int) (b bool, pos int, err error) {

	b = false
	var milieu int
	var debut int = 0
	var fin int = len(tab)

	if tab == nil {
		return b, -1, errFormatTableau
	}

	for i := 0; i < len(tab); i++ {
		milieu = (debut + fin) / 2
		if tab[milieu] == el {
			pos = milieu
			b = true
			return b, pos, nil
		}
		if tab[milieu] > el {
			fin = milieu
		}
		if tab[milieu] < el {
			debut = milieu + 1
		}

	}
	return b, -1, err

}
