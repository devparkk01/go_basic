package main 
import (
	"fmt"
)

//  will look into structs in this 

type Bill struct {
	name string 
	items map[string]float64
	tip float64
}

func newBill(name string ) (Bill) {

	b := Bill {
		name : name ,
		items : map[string]float64{} ,
		tip : 0 ,
	}


	return b 
}
// receiver function for adding new item
func (b *Bill)addItem(item string , price float64){
	b.items[item] = price 
}

// receiver function for adding  tip 
func (b *Bill)addTip(tip float64) {
	b.tip = tip 
}

// receiver function for calculating bill 
func (b *Bill)calculateBill() (string) {
	receipt := fmt.Sprintf("%30v\n\n" , "_____Bill breakdown_____") 
	total := 0.00

	for item , price := range b.items {
		total += price 
		receipt += fmt.Sprintf("%-20v : %-5v\n" , item , price) 
	}
	total += b.tip 
	receipt += fmt.Sprintf("%-20v : %-5v\n" ,"Tip" , b.tip)
	receipt += fmt.Sprintf("%20v : %-5v" , "Total " , total) 

	return receipt 
}




func main() {

	b := newBill("Dev")
	// fmt.Println(b) 

	b.addItem("Soup" , 34.44)
	b.addItem("Ice cream" , 35)
	b.addItem("Juice" , 20) 
	b.addTip(10) 

	// fmt.Println(b) 
	receipt := b.calculateBill() 
	fmt.Println(receipt)
	





}