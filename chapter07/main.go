package main
<<<<<<< HEAD

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"strings"
)

func main() {
	//Example 1 - Basic
	//data := [][]string{
	//	[]string{"A", "The Good", "500"},
	//	[]string{"B", "The Very very Bad Man", "288"},
	//	[]string{"C", "The Ugly", "120"},
	//	[]string{"D", "The Gopher", "800"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Name", "Sign", "Rating"})
	//for _, v := range data {
	//	table.Append(v)
	//}
	//table.Render() // Send output

	//Example 2 - Without Border / Footer / Bulk Append
	//data := [][]string{
	//	[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
	//	[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	//table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	//table.SetBorder(false)                                // Set Border to false
	//table.AppendBulk(data)                                // Add Bulk Data
	//table.Render()

	//Example 3 - CSV
	//table, _ := tablewriter.NewCSV(os.Stdout, "test_info.csv", true)
	//table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	//table.Render()

	//Example 4 - Custom Separator
	//table, _ := tablewriter.NewCSV(os.Stdout, "test.csv", true)
	//table.SetRowLine(true) // Enable row line
	//
	//// Change table lines
	//table.SetCenterSeparator("*")
	//table.SetColumnSeparator("╪")
	//table.SetRowSeparator("-")
	//
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	//table.Render()

	//Example 5 - Markdown Format
	//data := [][]string{
	//	[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
	//	[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	//table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	//table.SetCenterSeparator("|")
	//table.AppendBulk(data) // Add Bulk Data
	//table.Render()

	//Example 6 - Identical cells merging
	//data := [][]string{
	//	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
	//	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	//table.SetFooter([]string{"", "", "Total", "$146.93"})
	//table.SetAutoMergeCells(true)
	//table.SetRowLine(true)
	//table.AppendBulk(data)
	//table.Render()

	//Example 7 - Identical cells merging (specify the column index to merge)merge
	//data := [][]string{
	//	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "1234", "$10.98"},
	//	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	//table.SetFooter([]string{"", "", "Total", "$146.93"})
	//table.SetAutoMergeCellsByColumnIndex([]int{2, 3})
	//table.SetRowLine(true)
	//table.AppendBulk(data)
	//table.Render()

	//Example - 8 Table Cells with Color
	//data := [][]string{
	//	[]string{"Test1Merge", "HelloCol2 - 1", "HelloCol3 - 1", "HelloCol4 - 1"},
	//	[]string{"Test1Merge", "HelloCol2 - 2", "HelloCol3 - 2", "HelloCol4 - 2"},
	//	[]string{"Test1Merge", "HelloCol2 - 3", "HelloCol3 - 3", "HelloCol4 - 3"},
	//	[]string{"Test2Merge", "HelloCol2 - 4", "HelloCol3 - 4", "HelloCol4 - 4"},
	//	[]string{"Test2Merge", "HelloCol2 - 5", "HelloCol3 - 5", "HelloCol4 - 5"},
	//	[]string{"Test2Merge", "HelloCol2 - 6", "HelloCol3 - 6", "HelloCol4 - 6"},
	//	[]string{"Test2Merge", "HelloCol2 - 7", "HelloCol3 - 7", "HelloCol4 - 7"},
	//	[]string{"Test3Merge", "HelloCol2 - 8", "HelloCol3 - 8", "HelloCol4 - 8"},
	//	[]string{"Test3Merge", "HelloCol2 - 9", "HelloCol3 - 9", "HelloCol4 - 9"},
	//	[]string{"Test3Merge", "HelloCol2 - 10", "HelloCol3 -10", "HelloCol4 - 10"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Col1", "Col2", "Col3", "Col4"})
	//table.SetFooter([]string{"", "", "Footer3", "Footer4"})
	//table.SetBorder(false)
	//
	//table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
	//	tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	//	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	//	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})
	//
	//table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	//	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	//	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	//	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})
	//
	//table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
	//	tablewriter.Colors{tablewriter.Bold},
	//	tablewriter.Colors{tablewriter.FgHiRedColor})
	//
	//colorData1 := []string{"TestCOLOR1Merge", "HelloCol2 - COLOR1", "HelloCol3 - COLOR1", "HelloCol4 - COLOR1"}
	//colorData2 := []string{"TestCOLOR2Merge", "HelloCol2 - COLOR2", "HelloCol3 - COLOR2", "HelloCol4 - COLOR2"}
	//
	//for i, row := range data {
	//	if i == 4 {
	//		table.Rich(colorData1, []tablewriter.Colors{tablewriter.Colors{}, tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor}, tablewriter.Colors{}})
	//		table.Rich(colorData2, []tablewriter.Colors{tablewriter.Colors{tablewriter.Normal, tablewriter.FgMagentaColor}, tablewriter.Colors{}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor}, tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Italic, tablewriter.BgHiCyanColor}})
	//	}
	//	table.Append(row)
	//}
	//
	//table.SetAutoMergeCells(true)
	//table.Render()

	//Example 9 - Set table caption
	//data := [][]string{
	//	[]string{"A", "The Good", "500"},
	//	[]string{"B", "The Very very Bad Man", "288"},
	//	[]string{"C", "The Ugly", "120"},
	//	[]string{"D", "The Gopher", "800"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Name", "Sign", "Rating"})
	//table.SetCaption(true, "Movie ratings.")
	//
	//for _, v := range data {
	//	table.Append(v)
	//}
	//table.Render() // Send output

	//Example 10 - Set NoWhiteSpace and TablePadding option
	//data := [][]string{
	//	{"node1.example.com", "Ready", "compute", "1.11"},
	//	{"node2.example.com", "Ready", "compute", "1.11"},
	//	{"node3.example.com", "Ready", "compute", "1.11"},
	//	{"node4.example.com", "NotReady", "compute", "1.11"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Name", "Status", "Role", "Version"})
	//table.SetAutoWrapText(false)
	//table.SetAutoFormatHeaders(true)
	//table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	//table.SetCenterSeparator("")
	//table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	//table.SetHeaderLine(false)
	//table.SetBorder(false)
	//table.SetTablePadding("\t") // pad with tabs
	//table.SetNoWhiteSpace(true)
	//table.AppendBulk(data) // Add Bulk Data
	//table.Render()

	//Example 11 - Render table into a string
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Name", "Sign", "Rating"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	fmt.Println(tableString.String())
}
=======
>>>>>>> c3a1f4c2595f1faba5d7c90a1ab094c9bf9b5208
