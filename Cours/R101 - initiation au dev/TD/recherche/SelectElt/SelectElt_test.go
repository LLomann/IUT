package SelectElt

import "testing"

var b bool 
var tab []int
var err error

func TestVide1(t *testing.T){

	b, _, err = SelectElt(nil, 10)
	if err == nil || b == false {
		t.Fail()
	}
}

func TestVide2(t *testing.T){

	b, _, err = SelectElt([] int{}, 10)
	if err == nil || b == false {
		t.Fail()
	}
}

func NoNombre( t *testing.T){

	b, _, _  = SelectElt([] int{3,4,8,9}, 2)
	if b == false {
		t.Fail()
	}
}

func TestOk(t *testing.T){

	b, tab, err = SelectElt([] int{3,4,5,6,7,8}, 6)
	if b == true || tab[0] == 3 && tab[1] == 4 && tab[2] == 5 || err != nil{
		t.Fail()
	}
}