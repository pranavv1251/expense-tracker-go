package utils

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/pranavv1251/expense-tracker/expense"
)

func ExportToCSV(filename string, expenses []expense.Expense) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write header
    writer.Write([]string{"ID", "Description", "Amount", "Category", "Date"})

    for _, e := range expenses {
        record := []string{
            strconv.Itoa(e.ID),
            e.Description,
            fmt.Sprintf("%.2f", e.Amount),
            e.Category,
            e.Date.Format(time.RFC3339),
        }
        writer.Write(record)
    }

    return nil
}