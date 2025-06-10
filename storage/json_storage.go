package storage

import (
    "encoding/json"
    "os"
    "github.com/pranavv1251/expense-tracker/expense"
)

func SaveExpenses(filename string, expenses []expense.Expense) error {
    os.MkdirAll("data", os.ModePerm)
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    return json.NewEncoder(file).Encode(expenses)
}

func LoadExpenses(filename string) ([]expense.Expense, error) {
    var expenses []expense.Expense
    file, err := os.Open(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return expenses, nil
        }
        return nil, err
    }
    defer file.Close()
    err = json.NewDecoder(file).Decode(&expenses)
    return expenses, err
}