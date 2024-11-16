package gross

// pre-defined units map for immutability.
var predefinedUnits = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units returns a copy of the unit measurements to ensure immutability.
func Units() map[string]int {
	units := make(map[string]int, len(predefinedUnits))
	for k, v := range predefinedUnits {
		units[k] = v
	}
	return units
}

// NewBill creates and returns a new empty bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds a unit's quantity of an item to the bill.
// Returns false if the unit is invalid.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if quantity, validUnit := units[unit]; validUnit {
		bill[item] += quantity
		return true
	}
	return false
}

// RemoveItem removes a unit's quantity of an item from the bill.
// Returns false if the item doesn't exist, the unit is invalid, or the resulting quantity is negative.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQuantity, itemExists := bill[item]
	unitQuantity, validUnit := units[unit]

	if !itemExists || !validUnit {
		return false
	}

	newQuantity := currentQuantity - unitQuantity
	switch {
	case newQuantity < 0:
		return false
	case newQuantity == 0:
		delete(bill, item)
	default:
		bill[item] = newQuantity
	}
	return true
}

// GetItem retrieves the quantity of an item in the bill and indicates if it exists.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
