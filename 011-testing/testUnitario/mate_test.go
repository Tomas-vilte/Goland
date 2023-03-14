package testUnitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}

func TestSuma1(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}
	for _, value := range tabla {
		total := Suma(value.a, value.b)
		if total != value.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, value.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 3, 5},
		{2, 3, 3},
		{3, 5, 5},
	}
	for _, value := range tabla {
		max := GetMax(value.a, value.b)
		if max != value.c {
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", max, value.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 100},
	}
	for _, value := range tabla {
		fibo := Fibonacci(value.n)
		if fibo != value.r {
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", fibo, value.r)
		}
	}

}
