package RechMin

import "errors"

var errFormaTablo error = errors.New("hey ! le format du tablo est pas bon, gros nil-go.")

func RechMin(tablo[]int)(b bool, le_bon int, err error){

	b = false

	if tablo == nil{
		return false, -1, errFormaTablo
	}else if len(tablo) ==0{
		return false, -1, nil
	}else{
		le_bon = tablo[0]
		for i:=1; i<len(tablo); i++{
			if le_bon > tablo[i]{
				le_bon = tablo[i]
			}
		}
	}
	return true, le_bon, nil
}