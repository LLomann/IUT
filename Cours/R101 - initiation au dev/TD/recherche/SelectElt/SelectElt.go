package SelectElt

import "errors"

var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")
var errPasNombre error = errors.New("le nombre proposer n'est pas dans le tablo.")

func SelectElt(tablo []int, n int)(b bool, tab []int, err error){

	b = false

	if tablo == nil{
		return b, []int{}, errFormaTablo
	}

	if tablo == [] int{} {
		return b, []int{}, errFormaTablo
	}

	var compte int = 0

	for i:=0; i<len(tablo); i++ { 
		if n > tablo[i] {
			tab[i] = tablo[i]
			compte += 1
			b = true
		}
	}

	if compte < 1{
		return false, []int{}, errPasNombre
	}

	return b, tab, errFormaTablo
}