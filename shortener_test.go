package main

import (
	"reflect"
	"testing"
)

func TestRandString_LengthZeroReturnsError(t *testing.T) {
	n := 0
	s, err := randString(n)
	if s != "" {
		t.Errorf("expected an empty string but returned %s", s)
	}
	if err == nil {
		t.Error("expected an error but returned nil")
	}
}

func TestRandString_LengthLessThanZeroReturnsError(t *testing.T) {
	n := -16
	s, err := randString(n)
	if s != "" {
		t.Errorf("expected an empty string but returned %s", s)
	}
	if err == nil {
		t.Error("expected an error but returned nil")
	}
}

func TestRandString_ReturnStringMatchesExpectedLength(t *testing.T) {
	ns := []int{1, 2, 3, 5, 8, 13, 16, 32, 64, 127}
	for _, n := range ns {
		expected := n
		s, err := randString(n)
		if err != nil {
			t.Errorf("expected to run without error but returned %v", err)
		}
		actual := len(s)
		if actual != expected {
			t.Errorf("expected string with length %d but returned string with length %d", expected, actual)
		}
	}
}

func TestCreateLink_EmptyURLReturnsError(t *testing.T) {
	url := ""
	desc := "A description for the URL"
	categoryID := 42
	l, err := CreateLink(url, desc, categoryID)
	if !reflect.DeepEqual(l, Link{}) {
		t.Errorf("expected an empty Link struct but returned %v", l)
	}
	if err == nil {
		t.Error("expected an error but returned nil")
	}
}

func TestCreateLink(t *testing.T) {
	url := "http://example.com/testcreatelink"
	desc := "A description for the URL"
	categoryID := 42
	expShortcodeLen := shortcodeLen
	l, err := CreateLink(url, desc, categoryID)
	if err != nil {
		t.Errorf("expected to run without error but returned %v", err)
	}
	if len(l.Shortcode) != expShortcodeLen {
		t.Errorf("expected Link to have a Shortcode of length %d but returned '%s' with length %d", expShortcodeLen, l.Shortcode, len(l.Shortcode))
	}
	if l.URL != url {
		t.Errorf("expected Link to have URL '%s' but returned '%s'", url, l.URL)
	}
	if l.Description != desc {
		t.Errorf("expected Link to have Description '%s' but returned '%s'", desc, l.Description)
	}
	if l.CategoryID != categoryID {
		t.Errorf("exepcted Link to have CategoryID %d but returned %d", categoryID, l.CategoryID)
	}
}
