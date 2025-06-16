package automation

import "testing"

func TestTypeString_Empty(t *testing.T) {
	if err := TypeString(""); err == nil {
		t.Errorf("expected error for empty text, got nil")
	}
}

func TestTypeStringDelay_Empty(t *testing.T) {
	if err := TypeStringDelay("", 10); err == nil {
		t.Errorf("expected error for empty text in delay variant, got nil")
	}
}

func TestTypeStringDelay_NegativeDelay(t *testing.T) {
	if err := TypeStringDelay("Hello", -5); err == nil {
		t.Errorf("expected error for negative delay, got nil")
	}
}
