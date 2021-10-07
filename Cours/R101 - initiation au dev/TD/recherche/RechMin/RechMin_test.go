package RechMin

import "testing"

var b bool 
var le_bon int
var err error

func TestNil(t *testing.T){

	b,le_bon,err = RechMin(nil)
	if err == nil{
		t.Fail()
	}
}

func TestVide(t *testing.T){

	b,_,_ = RechMin([]int{})
	if !b {
		t.Fail()
	}
}

func Test1(t *testing.T){

	b,_,_ = RechMin([]int{10,56,3,84})
	if b {
		t.Fail()
	}
}

	