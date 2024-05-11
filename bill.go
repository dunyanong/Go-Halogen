package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, // empty map/objecct
		tip:   0,
	}

	return b
}

// Receiver Function: To format the bill
func (my_bill *bill) format() string {
	formatted_string := "Bill breakdown: \n"

	var total float64 = 0

	for key, value := range my_bill.items {
		formatted_string += fmt.Sprintf("%v...$%v \n", key+":", value)
		total += value
	}

	// add tip
	formatted_string += fmt.Sprintf("%-25v ...$%v\n", "tip:", my_bill.tip)

	// output total
	formatted_string += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+my_bill.tip)

	return formatted_string
}

// update tip
func (my_bill *bill) updateTip(tip float64) {
	my_bill.tip = tip // dereference
}

// add an item to the bill
func (my_bill *bill) addItem(name string, price float64) {
	my_bill.items[name] = price
}

// to save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
