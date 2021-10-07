package RechElt

import "errors"

var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")

func RechElt(tablo []int, n int)(b bool, pos int, err error){

	b = false

	if tablo == nil{
		return false, -1, errFormaTablo
	}else if len(tablo) == 0{
			return false, -1, nil		
	}else{
		for i:=0; i<len(tablo); i++{
			if tablo[i] == n{
				return true, i, nil
			}
		}
	}
	return false, -1, nil 
}