package dateutils

import "testing"

func TestGetMaxDays(t *testing.T) {

	subtests := []struct {
		name  string
		month int
		year  int
		want  int
	}{
		{"January 2024", 1, 2024, 31},
		{"February 2024", 2, 2024, 29},
		{"February 2025", 2, 2025, 28},
		{"September 2024", 9, 2024, 30},
	}
	for _, st := range subtests {
		t.Run(st.name, func(t *testing.T) {
			got, err := GetMaxDays(st.month, st.year)
			if err != nil {
				t.Error("was not expecting any error, but got one!")
			}
			if got != st.want {
				t.Errorf("wanted %v, but got %v", st.want, got)
			}
		})
	}

	negativeTests := []struct {
		name   string
		month  int
		year   int
		errmsg string
	}{
		{"month < 1", 0, 2024, "month should be >= 1"},
		{"month > 12", 20, 2024, "month should be <=12"},
		{"year < 1900", 10, 1890, "year should be >=1900"},
	}

	for _, nt := range negativeTests {
		t.Run(nt.name, func(t *testing.T) {
			_, err := GetMaxDays(nt.month, nt.year)
			if err == nil {
				t.Error("was expecting an error; did not get one")
			} else {
				if err.Error() != nt.errmsg {
					t.Errorf("wanted error is '%v' got error is '%v'", nt.errmsg, err.Error())
				}
			}
		})
	}
}
