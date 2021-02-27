package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 3},
		{1, 2, 4},
	}
	for _, c := range cases {
		result := add(c.first, c.second)

		if result != c.excepted {
			t.Fatalf("add function failed,first:%d,second:%d,result:%d", c.first, c.excepted, result)
		}
	}
}
func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	time.Sleep(10 * time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := upperCase(ut.in)
		t.Error(uc)
		if uc != ut.out {
			t.Errorf("uppercase(%s) = %s,must be %s", ut.in, uc, ut.out)
		}
	}
}
