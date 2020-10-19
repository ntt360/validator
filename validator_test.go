package validator

import "testing"

func TestNew(t *testing.T) {

	data := map[string][]string{
		"hello": {"2"},
	}

	rules := map[string][]string{
		"golang": {"min:1", "regex:^\\w$"},
	}

	_, err := New(data, rules)
	if err != nil {
		println(err)
	}
}
