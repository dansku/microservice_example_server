package main

import (
  "testing"
)

// TestCalculateValues for SUM
func TestCalculateValuesForSUMAndMultipleWritings(t *testing.T) {
	values := []float32{10.2,19,31.3}
	var expected float32 = 60.5
	actual:= CalculateValues(values, "sum")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "Sum")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "SUM")

	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}
}

// TestCalculateValues for Multiply
func TestCalculateValuesForMULTIPLYAndMultipleWritings(t *testing.T) {
	values := []float32{10.2,19,13.3}
	var expected float32 = 2577.54
	actual := CalculateValues(values, "multiply")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "Multiply")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "mUltipLY")

	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}
}

// TestCalculateValues for Divide
func TestCalculateValuesForDIVIDEAndMultipleWritings(t *testing.T) {
	values := []float32{10.2,19,13.3}
	var expected float32 = 0.0003879668
	actual := CalculateValues(values, "divide")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "Divide")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "divIdE")

	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}
}

// TestCalculateValues for Subtract
func TestCalculateValuesForSUBTRACTAndMultipleWritings(t *testing.T) {
	values := []float32{10.2,19,13.3}
	var expected float32 = -42.5
	actual := CalculateValues(values, "subtract")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "Subtract")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}

	actual = CalculateValues(values, "SubtracT")

	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}
}


// TestCalculateValues for Wrong Operation should return zero
func TestCalculateValuesForWrongOperationAndMultipleWritings(t *testing.T) {
	values := []float32{10.2,19,13.3}
	var expected float32 = 0
	actual := CalculateValues(values, "wrongOperation")
  
	if actual != expected {
		t.Errorf("Expected: %v but got %v", expected, actual)
	}


}


