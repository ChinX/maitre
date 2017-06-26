package tests

import "testing"

type Mystruct struct {
	A string `equal:""`
	B *string `equal:""`
	c int
	D map[string]string
}

func BenchmarkEqualStruct(t *testing.B) {
	str := "bbb"
	a := &Mystruct{
		A: "aaa",
		B: &str,
		c: 1,
		D: map[string]string{"aa":"aaa"},
	}

	b := &Mystruct{
		A: "aaa",
		B: &str,
		c: 2,
		D: map[string]string{"aa":"aaa"},
	}
	for i:=0; i< t.N ; i++ {
		EqualStruct(a, b)
	}
}

func TestEqualStruct(t *testing.T) {
	str := "bbb"
	a := &Mystruct{
		A: "aaa",
		B: &str,
		c: 1,
		D: map[string]string{"aa":"aaa"},
	}

	b := &Mystruct{
		A: "aaa",
		B: &str,
		c: 2,
		D: map[string]string{"aa":"aaa"},
	}

	t.Log(EqualStruct(a, b))
}
