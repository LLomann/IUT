package main

import(
	"errors"
	"fmt"
)
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

func main(){
	var b bool 
	var pos int
	var err error
	
	b, pos, err = RechElt(nil, 10)
	if err !=nil{
		fmt.Println(err)
	}

	b, pos, _ = RechElt([] int{}, 10)
	if pos == -1 && b == false{
		fmt.Println("Le nombre recherché n'existe pas")
	}

	b, pos, _ = RechElt([] int{1,2,3,10}, 10)
	if b == true{
		fmt.Println("le chiffre apartiens bien a l'ensemble et est a la position:", pos)
	}

	b, pos, _ = RechElt([] int{1,2,3}, 10)
	if b == false{
		fmt.Println("le chiffre n'apartiens pas a l'ensemble")
	}
}