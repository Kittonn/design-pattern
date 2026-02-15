package prototype

import "testing"

func TestUserClone(t *testing.T) {
	original := &User{Name: "Alice", Age: 30}
	clone := original.Clone().(*User)

	clone.Name = "Bob"

	if original.Name != "Alice" {
		t.Errorf("Expected original name to be 'Alice', got '%s'", original.Name)
	}

	if clone.Name != "Bob" {
		t.Errorf("Expected clone name to be 'Bob', got '%s'", clone.Name)
	}

	if original.Age != clone.Age {
		t.Errorf("Expected ages to be the same, got %d and %d", original.Age, clone.Age)
	}
}
