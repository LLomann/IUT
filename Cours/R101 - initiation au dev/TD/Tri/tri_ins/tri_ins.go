package Tri

//var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")

//func tri_ins(tab1 []int) (tab2 []int, b bool, err error) {

//	b = false

//	if tab1 == nil {
//		return []int{}, b, errFormaTablo
//	}

//	var i int = 0

//	for i < len(tab1) {
//		tab2 = append(tab2, tab1[i])
//		i ++
//	}

//	var j int = 0
//	var v int
//	var tpm int

//	for j < len(tab2) && tab2[j] <= v {
//		j++
//	}
//	for j < len(tab2) {
//		tpm = tab2[j]
//		tab2[j] = v
//		v = tpm
//		i++
//	}

//	if len(tab1) == 0 {
//		return []int{}, false, err
//	}

//	return tab2, true, nil

//}

func tri_ins(tab []int) (table []int) {

	table = []int{}
	var v int
	var tmp int

	for i := 0; i < len(table); i++ {

		v = tab[i]
		var j int = 0

		for j < len(table) && table[i] < v {

			j++
		}

		for j < len(table) {

			tmp = table[j]
			table[j] = v
			v = tmp
			j++
		}

		table = append(table, v)
	}
	return table
}
