package Tri

import "testing"

var tab2 []int
var b bool
var err error

func TestVide1(t *testing.T) {

	tab2, _, _= tri_ins(nil) 
		if tab2 == nil {
			t.Fail()
		}

}

func TestVide2(t *testing.T) {

	tab2, _, _= tri_ins([]int{}) 
		if len(tab2) != 0 {
			t.Fail()
		}

}