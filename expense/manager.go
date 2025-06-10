// package expense

// import (
//     "fmt"
//     "time"
// )

// type Manager struct {
//     Expenses []Expense
//     nextID   int
// }

// func NewManager() *Manager {
//     return &Manager{
//         Expenses: []Expense{},
//         nextID:   1,
//     }
// }

// func (m *Manager) Load(expenses []Expense) {
//     m.Expenses = expenses
//     maxID := 0
//     for _, e := range expenses {
//         if e.ID > maxID {
//             maxID = e.ID
//         }
//     }
//     m.nextID = maxID + 1
// }

// func (m *Manager) AddExpense(desc string, amt float64, cat string, date time.Time) {
//     exp := Expense{
//         ID:          m.nextID,
//         Description: desc,
//         Amount:      amt,
//         Category:    cat,
//         Date:        date,
//     }
//     m.Expenses = append(m.Expenses, exp)
//     m.nextID++
// }

// func (m *Manager) ListExpenses() {
//     for _, e := range m.Expenses {
//         fmt.Printf("ID: %d | %s | $%.2f | %s | %s\n", e.ID, e.Description, e.Amount, e.Category, e.Date.Format("2006-01-02"))
//     }
// }

// func (m *Manager) DeleteExpense(id int) {
//     for i, e := range m.Expenses {
//         if e.ID == id {
//             m.Expenses = append(m.Expenses[:i], m.Expenses[i+1:]...)
//             return
//         }
//     }
// }

package expense

import (
    "fmt"
    "time"
)

type Manager struct {
    Expenses []Expense
    nextID   int
}

func NewManager() *Manager {
    return &Manager{
        Expenses: []Expense{},
        nextID:   1,
    }
}

func (m *Manager) Load(expenses []Expense) {
    m.Expenses = expenses
    maxID := 0
    for _, e := range expenses {
        if e.ID > maxID {
            maxID = e.ID
        }
    }
    m.nextID = maxID + 1
}

func (m *Manager) AddExpense(desc string, amt float64, cat string, date time.Time) {
    exp := Expense{
        ID:          m.nextID,
        Description: desc,
        Amount:      amt,
        Category:    cat,
        Date:        date,
    }
    m.Expenses = append(m.Expenses, exp)
    m.nextID++
}

func (m *Manager) ListExpenses() {
    for _, e := range m.Expenses {
        fmt.Printf("ID: %d | %s | $%.2f | %s | %s\n", e.ID, e.Description, e.Amount, e.Category, e.Date.Format("2006-01-02"))
    }
}

func (m *Manager) DeleteExpense(id int) {
    for i, e := range m.Expenses {
        if e.ID == id {
            m.Expenses = append(m.Expenses[:i], m.Expenses[i+1:]...)
            return
        }
    }
}

func (m *Manager) SummarizeByMonth(month time.Month, year int) {
    var total float64
    fmt.Printf("Summary for %s %d:\n", month.String(), year)
    for _, e := range m.Expenses {
        if e.Date.Month() == month && e.Date.Year() == year {
            fmt.Printf("  %s | $%.2f | %s\n", e.Description, e.Amount, e.Category)
            total += e.Amount
        }
    }
    fmt.Printf("Total: $%.2f\n", total)
}