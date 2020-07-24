package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Example Sheet")
	f.SetColWidth("Example Sheet", "A", "A", 200)

	// Set value of a cell.
	f.SetCellValue("Example Sheet", "A2", "Hello world.")
	style, _ := f.NewStyle(`{"fill":{"type":"pattern","color":["#565664"],"pattern":1},"font":{"bold":true},"alignment":{"wrap_text":true}}`)
	_ = f.SetCellStyle("Sheet2", "A2", "A2", style)

	f.SetActiveSheet(index)
	// Remove default sheet
	f.DeleteSheet("Sheet1")

	// Save xlsx file
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
