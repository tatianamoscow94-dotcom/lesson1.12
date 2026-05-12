package main

import "testing"

func TestDivideSucces(t *testing.T) {
	result, err := divide(10.0, 2.0)
	if err != nil {
		t.Errorf("Неожиданная ошибка: %v", err)
		return
	}
	if result != 5.0 {
		t.Errorf("Ожидали 5.0, получили %v", result)
	}
}

func TestDivideByZero(t *testing.T) {
	result, err := divide(7.0, 0.0)
	if err == nil {
		t.Error("Ожидалась ошибка, но её не было")
	}
	if result != 0 {
		t.Errorf("Ожидали 0, получили %v", result)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = divide(100.5, 4.0)
	}
}
