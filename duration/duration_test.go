package duration

import "testing"

func TestConvertZero(t *testing.T) {
	resultInputZero := Convert(0)
	if resultInputZero != "0s" {
		t.Errorf("Convert(0) = %s; 0s", resultInputZero)
	}
}

func TestConvertOne(t *testing.T) {
	resultInputOne := Convert(1)
	if resultInputOne != "1s" {
		t.Errorf("Convert(1) = %s; 1s", resultInputOne)
	}
}

func TestConvertOneHundred(t *testing.T) {
	resultInputOneHundred := Convert(100)
	if resultInputOneHundred != "1m40s" {
		t.Errorf("Convert(100) = %s; 1m40s", resultInputOneHundred)
	}
}

func TestConvertOneThousand(t *testing.T) {
	resultInputOneThousand := Convert(1000)
	if resultInputOneThousand != "16m40s" {
		t.Errorf("Convert(1000) = %s; 16m40s", resultInputOneThousand)
	}
}

func TestConvertTenThousand(t *testing.T) {
	resultInputTenThousand := Convert(10000)
	if resultInputTenThousand != "2h46m40s" {
		t.Errorf("Convert(10000) = %s; 2h46m40s", resultInputTenThousand)
	}
}

func TestConvertOneHundredThousand(t *testing.T) {
	resultInputOneHundredThousand := Convert(100000)
	if resultInputOneHundredThousand != "1d3h46m40s" {
		t.Errorf("Convert(100000) = %s; 1d3h46m40s", resultInputOneHundredThousand)
	}
}

func TestConvertTenMillion(t *testing.T) {
	resultInputTenMillion := Convert(10000000)
	if resultInputTenMillion != "115d17h46m40s" {
		t.Errorf("Convert(10000000) = %s; 115d17h46m40s", resultInputTenMillion)
	}
}