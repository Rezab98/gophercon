package automation

import "testing"

// TestClickNegativeCoordinates ensures that the validation layer inside Click correctly rejects
// negative coordinates without invoking the underlying robotgo call (which would otherwise move
// the real cursor during CI runs).
func TestClickNegativeCoordinates(t *testing.T) {
	if err := Click(-1, 10); err == nil {
		t.Errorf("expected error for negative X coordinate, got nil")
	}

	if err := Click(10, -1); err == nil {
		t.Errorf("expected error for negative Y coordinate, got nil")
	}
}

// TestParseCoordinates validates the happy-path for ParseCoordinates and the error conditions
// for malformed input.
func TestParseCoordinates(t *testing.T) {
	x, y, err := ParseCoordinates("50", "100")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if x != 50 || y != 100 {
		t.Fatalf("expected (50,100), got (%d,%d)", x, y)
	}

	if _, _, err := ParseCoordinates("not-an-int", "10"); err == nil {
		t.Errorf("expected error for non-numeric X coordinate, got nil")
	}

	if _, _, err := ParseCoordinates("10", "not-an-int"); err == nil {
		t.Errorf("expected error for non-numeric Y coordinate, got nil")
	}
}

// TestMoveNegativeCoordinates ensures that Move and MoveSmooth reject negative coordinates.
func TestMoveNegativeCoordinates(t *testing.T) {
	if err := Move(-1, 0); err == nil {
		t.Errorf("expected error for negative X coordinate in Move, got nil")
	}
	if err := Move(0, -1); err == nil {
		t.Errorf("expected error for negative Y coordinate in Move, got nil")
	}

	if err := MoveSmooth(-1, 0, 1000); err == nil {
		t.Errorf("expected error for negative X coordinate in MoveSmooth, got nil")
	}
	if err := MoveSmooth(0, -1, 1000); err == nil {
		t.Errorf("expected error for negative Y coordinate in MoveSmooth, got nil")
	}
}
