package RechElt

import "testing"

var b bool 
var pos int
var err error

func TestVide(t *testing.T){

	b, pos, err = RechElt(nil, 10)
	if err == nil || pos != -1 || b == true {
		t.Fail()
	}
}

func Test1(t *testing.T){

	b, pos,_ = RechElt([] int{}, 10)
	if pos != -1 && b != false{
		t.Fail()
	}
}	

func Test2(t *testing.T){

	b, pos ,_ = RechElt([] int{1,2,3,10}, 10)
	if b == false {
		t.Fail()
	}
}	

func Test3(t *testing.T){

	b, pos ,_ = RechElt([] int{1,2,3,4}, 10)
	if b == true{
		t.Fail()
	}
}	
