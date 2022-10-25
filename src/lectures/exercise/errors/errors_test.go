package timeparse

import "testing"

func TestSuccess(t *testing.T) {
	testCases := []struct {
		desc     string
		s        string
		expected Time
	}{
		{"parse some random time", "10:20:42", Time{10, 20, 42}},
		{"parse midnight", "00:00:00", Time{0, 0, 0}},
		{"big hours", "1976239:00:00", Time{1976239, 0, 0}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			time, err := ParseTime(tC.s)
			if err != nil {
				t.Errorf("parsing throw error %v, wanted %v", err, tC.expected)
			}
			if time != tC.expected {
				t.Errorf("parsing did return wrong value %v, wanted %v", time, tC.expected)
			}
		})
	}
}

func TestError(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
	}{
		{"missing colon", "12:3223"},
		{"no colon", "123232"},
		{"too many colons", "12:23:34:45"},
		{"hour not a number", "1d:23:12"},
		{"minute not a number", "11:ds:12"},
		{"second not a number", "11:2:asdf"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			time, err := ParseTime(tC.s)
			if err == nil {
				t.Error("expected an error")
			}
			defaultTime := Time{}
			if time != defaultTime {
				t.Errorf("din't get default time value: %v, wanted %v", time, defaultTime)
			}
		})
	}
}
