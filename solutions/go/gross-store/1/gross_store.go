package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	newMap := make(map[string]int)

	newMap["quarter_of_a_dozen"] =	3
    newMap["half_of_a_dozen"]	=6
    newMap["dozen"]	=12
    newMap["small_gross"]	=120
    newMap["gross"]	=144
    newMap["great_gross"]	=1728
    return newMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    if exists == false {
        return false
    } else {
        bill[item] += value
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, bExists := bill[item]
    _, uExists := units[unit]

    if bExists == false || uExists == false {
        return false
    }

    if bill[item] < units[unit] {
        return false
    }

    bill[item] -= units[unit]
    if bill[item] == 0{
        delete(bill, item)
    } 

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	b, exists := bill[item]

    if exists == false {
        return 0, false
    } else {
        return b, true
    }
}
