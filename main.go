package main

import (
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/pranavv1251/expense-tracker/utils"
    "github.com/pranavv1251/expense-tracker/storage"
    "github.com/pranavv1251/expense-tracker/expense"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: <add|list|delete|summary|export>")
        return
    }

    manager := expense.NewManager()
    expenses, _ := storage.LoadExpenses("data/data.json")
    manager.Load(expenses)

    switch os.Args[1] {
    case "add":
        if len(os.Args) < 5 {
            fmt.Println("Usage: add <description> <amount> <category>")
            return
        }
        description := os.Args[2]
        amount, err := strconv.ParseFloat(os.Args[3], 64)
        if err != nil {
            fmt.Println("Invalid amount")
            return
        }
        category := os.Args[4]
        manager.AddExpense(description, amount, category, time.Now())
        fmt.Println("Expense added!")
    case "list":
        manager.ListExpenses()
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: delete <id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid ID")
            return
        }
        manager.DeleteExpense(id)
        fmt.Println("Expense deleted!")
    case "summary":
        if len(os.Args) < 3 {
            fmt.Println("Usage: summary <month-number>")
            return
        }
        monthNum, err := strconv.Atoi(os.Args[2])
        if err != nil || monthNum < 1 || monthNum > 12 {
            fmt.Println("Invalid month")
            return
        }
        year := time.Now().Year()
        manager.SummarizeByMonth(time.Month(monthNum), year)
    case "export":
        if len(os.Args) < 3 {
            fmt.Println("Usage: export <filename>")
            return
        }
        err := utils.ExportToCSV(os.Args[2], manager.Expenses)
        if err != nil {
            fmt.Println("Export failed:", err)
        } else {
            fmt.Println("Exported successfully to", os.Args[2])
        }
    }

    storage.SaveExpenses("data/data.json", manager.Expenses)
}