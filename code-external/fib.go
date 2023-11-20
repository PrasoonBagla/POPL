package main
import (
    "fmt"
    "net/http"
    "time"

    "github.com/tealeg/xlsx" // Excel file package
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Your handler logic here
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    go http.ListenAndServe(":8080", nil)

    // Create a new Excel file
    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell

    file = xlsx.NewFile()
    sheet, err := file.AddSheet("Request Times")
    if err != nil {
        fmt.Printf("Failed to create sheet: %v", err)
        return
    }

    for i := 0; i < 1000; i++ {
        start := time.Now()
        resp, err := http.Get("http://localhost:8080")
        if err != nil {
            fmt.Printf("Request failed: %v", err)
            continue
        }
        resp.Body.Close()
        elapsed := time.Since(start)

        // Write the time taken for each request to the Excel sheet in milliseconds with 3 decimal places
        row = sheet.AddRow()
        cell = row.AddCell()
        cell.SetValue(fmt.Sprintf("%.3f", float64(elapsed)/float64(time.Millisecond)))
    }

    // Save the Excel file
    err = file.Save("RequestTimes.xlsx")
    if err != nil {
        fmt.Printf("Failed to save Excel file: %v", err)
    }
}
