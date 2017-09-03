package pizzabot

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		in      []string
		success bool
	}{
		{[]string{"exec_path"}, false},
		{[]string{"exec_path", "5x5 (1, 3) (4, 4)"}, true},
		{[]string{"exec_path", "5X5 (1, 3) (4, 4)"}, false},
		{[]string{"exec_path", "5x5 (1 3) (4 4)"}, false},
		{[]string{"exec_path", "5x5 (1,3) (4,4)"}, false},
		{[]string{"exec_path", "5x5 [1, 3] [4, 4]"}, false},
		{[]string{"exec_path", "5x5 (1, 3, 0) (4, 4, 0)"}, false},
		{[]string{"exec_path", "5x5x1 (1, 3) (4, 4)"}, false},
		{[]string{"exec_path", "Ax5 (1, 3) (4, 4)"}, false},
		{[]string{"exec_path", "5xB (1, 3) (4, 4)"}, false},
		{[]string{"exec_path", "5x5 (C, 3) (D, 4)"}, false},
		{[]string{"exec_path", "5x5 (1, E) (4, F)"}, false},
		{[]string{"exec_path", "5x5 (5, 3)"}, false},
		{[]string{"exec_path", "5x5 (1, 5)"}, false},
	}

	for _, c := range cases {
		config := Config{}
		err := config.Parse(c.in)
		if err != nil && c.success == true {
			t.Error("Case should have passed.")
		}

		if err == nil && c.success == false {
			t.Error("Case should have failed.")
		}
	}
}
