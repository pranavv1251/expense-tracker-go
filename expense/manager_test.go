package expense

import (
    "testing"
    "time"
)

func TestAddExpense(t *testing.T) {
    mgr := NewManager()
    mgr.AddExpense("Coffee", 3.5, "Food", time.Now())

    if len(mgr.Expenses) != 1 {
        t.Errorf("Expected 1 expense, got %d", len(mgr.Expenses))
    }

    if mgr.Expenses[0].Description != "Coffee" {
        t.Errorf("Expected description 'Coffee', got '%s'", mgr.Expenses[0].Description)
    }
}

func TestDeleteExpense(t *testing.T) {
    mgr := NewManager()
    mgr.AddExpense("Coffee", 3.5, "Food", time.Now())
    id := mgr.Expenses[0].ID

    mgr.DeleteExpense(id)

    if len(mgr.Expenses) != 0 {
        t.Errorf("Expected 0 expenses after delete, got %d", len(mgr.Expenses))
    }
}
