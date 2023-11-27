package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type Slab struct {
	Litres int64 `json:"litres"`
	Cost   int64 `json:"cost"`
	Charge int64
}
type Note struct {
	Note string `json:"note"`
}
type Bill struct {
	ProjectName           string
	Invoice               int64
	Admin                 Admin
	UserInfo              UserInfo
	BillGenerated         string
	DueDate               string
	OptGst                bool
	SlabRates             []Slab
	MeterReadingData      []MeterReadingData
	TotalDiffernce        int64
	ChargeSummary         ChargeSummary
	TotalChargeWithOutGst int64
	TotalChargeWithGst    int64

	TotalConsumption int64
	GstCharges       int64
	Total            int64
	Note             []AddNotes
}
type AddNotes struct {
	Id    int64  `json:"Id"`
	Notes []Note `json:"notes"`
}

type ChargeSummary struct {
	SlabRates []Slab
}

type MeterReadingData struct {
	SerialNo        string
	PreviousReading int64
	CurrentReading  int64
	Difference      int64
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Pincode int64  `json:"pincode"`
}
type UserInfo struct {
	Name     string
	MobileNo string
	Address  Address
	Email    string
}

type Admin struct {
	AdminName string
	MobileNo  string
	Address   string
	Email     string
	GstPin    string
}

func ReadingDifference(previous, current int64) int64 {
	return previous - current
}
func calculateTotalDifference(meterReadings []MeterReadingData) int64 {
	var totalDifference int64
	for _, reading := range meterReadings {
		totalDifference += reading.Difference
	}
	return totalDifference
}
func slabCharge(litre, cost int64) int64 {
	return litre * cost
}
func calucateSlabTotal(slabCharges []Slab) int64 {
	var total int64
	for _, charge := range slabCharges {
		total += charge.Charge

	}
	return total

}

func TotalCostWithGst(totalCost int64) (int64, int64) {
	gstPercentage := 18
	gstAmount := (totalCost * int64(gstPercentage)) / 100
	totalAmount := totalCost + gstAmount
	return totalAmount, gstAmount
}

func main() {
	bill := Bill{
		ProjectName: "BuildingA",
		Admin: Admin{
			AdminName: "Izaz",
			MobileNo:  "12345566",
			Address:   "123444",
			Email:     "shaikIzaz",
			GstPin:    "1234",
		},
		UserInfo: UserInfo{
			Name:  "shaik",
			Email: "shaik@1123.com",
			Address: Address{
				Street:  "btm",
				City:    "banglore",
				Country: "india",
				Pincode: 12345,
			},
			MobileNo: "9199112",
		},
		Invoice:       1,
		BillGenerated: time.Now().Format("2006-01-02"),
		DueDate:       time.Now().Format("2006-01-02"),
		OptGst:        false,
		SlabRates: []Slab{
			{
				Litres: 12,
				Cost:   233,
			},
			{
				Litres: 12,
				Cost:   233,
			},
		},
		MeterReadingData: []MeterReadingData{
			{
				SerialNo:        "12345",
				PreviousReading: 1234,
				CurrentReading:  456,
				Difference:      ReadingDifference(1234, 456),
			},
			{
				SerialNo:        "12345",
				PreviousReading: 1234,
				CurrentReading:  456,
				Difference:      ReadingDifference(1234, 456),
			},
		},

		TotalConsumption: 12355,
		Total:            12345,
		Note: []AddNotes{
			{
				Id: 1,
				Notes: []Note{
					{
						Note: "array of json",
					},
				},
			},
		},
		ChargeSummary: ChargeSummary{
			SlabRates: []Slab{
				{
					Litres: 12,
					Cost:   233,
					Charge: slabCharge(12, 233),
				},
				{
					Litres: 12,
					Cost:   233,
					Charge: slabCharge(12, 233),
				},
			},
		},
	}
	bill.TotalDiffernce = calculateTotalDifference(bill.MeterReadingData)
	bill.TotalChargeWithOutGst = calucateSlabTotal(bill.ChargeSummary.SlabRates)
	if bill.OptGst {
		bill.TotalChargeWithGst, bill.GstCharges = TotalCostWithGst(bill.TotalChargeWithOutGst)
	}
	file, err := os.Create("output_bill.html")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmplFile := "bill.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(file, bill)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
