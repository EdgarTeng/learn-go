package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://www.quii.co.uk"

	expect := fastURL
	actual := Racer(slowURL, fastURL)

	if actual != expect {
		t.Errorf("actual %s, but expect %s", actual, expect)
	}

}
