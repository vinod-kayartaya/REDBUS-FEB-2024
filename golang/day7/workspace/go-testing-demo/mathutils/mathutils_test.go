package mathutils

import "testing"

func TestFactorial(t *testing.T) {

	t.Run("factorial of 0", func(t *testing.T) {
		// t.Skip()
		t.Parallel()
		want := 1
		got, err := Factorial(0)
		if err != nil {
			t.Errorf("test failed with reason: %v", err.Error())
		}
		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("factorial of negative number", func(t *testing.T) {
		// t.Skip()
		t.Parallel()
		_, err := Factorial(-5)
		if err == nil {
			t.Error("expected error but did not get one")
		} else {
			if err.Error() != "cannot calculate factorial of -ve number" {
				t.Error("wanted error message was not got")
			}
		}
	})

	t.Run("factorial of positive number", func(t *testing.T) {
		// t.Skip()
		t.Parallel()
		want := 120
		got, err := Factorial(5)

		if err != nil {
			t.Errorf("test failed with reason: %v", err.Error())
		}

		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("factorial of number >10", func(t *testing.T) {
		t.Parallel()
		want := 1307674368000
		got, err := Factorial(15)
		if err != nil {
			t.Errorf("test failed with reason: %v", err.Error())
		}

		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

}
