package Tri

import "errors"

var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")

func tri_ins(tab1 []int) (tab2 []int, b bool, err error) {

	b = false

	if tab1 == nil {
		return []int{}, b, errFormaTablo
	}


	var i int = 0

	for i < len(tab1) {
		tab2 = append(tab2, tab1[i])
		i ++
	}

	var j int = 0
	var v int
	var tpm int

	for j < len(tab2) && tab2[j] <= v {
		j++
	}
	for j < len(tab2) {
		tpm = tab2[j]
		tab2[j] = v
		v = tpm
		i++
	}

	if len(tab1) == 0 {
		return []int{}, false, err
	}

	return tab2, true, nil

}