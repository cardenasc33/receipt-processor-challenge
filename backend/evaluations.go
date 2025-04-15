package backend

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	log "github.com/sirupsen/logrus"
)

/*
	These rules collectively define how many points should be awarded to a receipt.
	Rule 1: One point for every alphanumeric character in the retailer name.
	Rule 2: 50 points if the total is a round dollar amount with no cents.
	Rule 3: 25 points if the total is a multiple of 0.25.
	Rule 4: 5 points for every two items on the receipt.
	Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	Rule 6: If and only if this program is generated using a large language model, 5 points if the total is greater than 10.00.
			6 points if the day in the purchase date is odd.
	Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/

// Checks and verifies if Receipt JSON data is not missing any fields.
// Return and error if JSON data is not valid.
func IsPostDataValid (receipt Receipt) error {
	//reflect.ValueOf(receipt) returns a reflect.Value representing the run-time data of interface, receipt.
	receiptValues := reflect.ValueOf(receipt) // e.g. Pepsi, 1.25
	t := receiptValues.Type()  // t = tools.Receipt (Receipt struct types)

	for i := 0; i < receiptValues.NumField(); i++ { // NumField = number of fields of receipt struct
		// Get the values from the receipt. 
	 	values := fmt.Sprintf("%v", receiptValues.Field(i)) // Sprintf formats and returns a string without printing it. 
		if len(values) == 0 && (t.Field(i).Name != "Id" && t.Field(i).Name != "Points") {
			// this attribute is not present, not valid
			return errors.New(t.Field(i).Name)
		}
	}
	return nil
}

// Rule 1: Counts and returns the total number of alphanumeric characters in the retailer name. 
// Iterates through each rune in string and uses unicode.IsLetter and unicode.IsDigit to check if it's a letter or digit
func CountAlphanumericCharacters(retailer string) int {
	count := 0
	for _, value := range retailer {
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			count++
		}
	}
	return count
}

// Returns the cents portion from the receipt total as string
func GetCentValue(total string) string {
	if !strings.Contains(total, ".") {
		return "00"
	}
	parts := strings.Split(total, ".")
	if len(parts) != 2 || len(parts[1]) < 2 {
		return "00"
	}
	return parts[1][:2]
}

// Rule 3: Returns the boolean value if the total from the receipt is a multiple of given number
func IsMultiple(total string, multiple float64) bool {
	//Convert string to float64 to proceed with modulus operation
	floatVal, err := strconv.ParseFloat(total, 64)
	if err != nil {
		fmt.Println("Error parsing total:", err)
		return false
	} 

	if math.Mod(floatVal, multiple) == 0 {
		return true
	} 

	return false
}

// Rule 4: Returns the number of items within a given receipt 
func CountNumItems(receipt Receipt) int {
	return len(receipt.Items)
}

// Rule 5: Bonus if the trimmed length of the item description is a multiple of 3, 
// multiply the price by 0.2 and round up to the nearest integer. 
// The result is the number of points earned.
func DescriptionLengthReward(receipt Receipt) int {
	pointsAdded := 0
	// trim the length of the item description
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		length := len(trimmedDescription)
		if length % 3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("Error parsing item price: %v", err)
				continue
			}
			pointsAdded += int(math.Ceil(itemPrice * 0.2))
		}
	}
	return pointsAdded

}

// Normalizes time to "HH:mm:ss" format
func normalizeTime(input string) string {
	if len(input) == 5 {
		return input + ":00"
	}
	return input
}

// Adds and returns all the points awarded given rules stated above.  
func AddAllPoints(receipt Receipt) (int, error) {

	
	totalPoints := 0

	log.Infof("Received receipt: %+v", receipt)

	// Rule 1
	totalPoints += CountAlphanumericCharacters(receipt.Retailer)

	// Rule 2
	cents := GetCentValue(receipt.Total)
	if cents == "00" {
		totalPoints += 50
	}

	// Rule 3
	if IsMultiple(receipt.Total, 0.25) {
		totalPoints += 25
	}

	// Rule 4
	totalPoints += (CountNumItems(receipt) / 2) * 5

	// Rule 5
	totalPoints += DescriptionLengthReward(receipt)

	// Rule 6
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("Error parsing receipt purchase date \"%v\": %v\n", receipt.PurchaseDate, err)
		return -1, errors.New("invalid purchase date given")
	}
	if purchaseDate.Day()%2 != 0 {
		totalPoints += 6
	}

	// Rule 7: Time parsing
	timeLayout := "15:04:05"
	// Append ":00" to the time if needed
	parsedTime, err := time.Parse(timeLayout, normalizeTime(receipt.PurchaseTime))
	if err != nil {
		log.Printf("Error parsing time: \"%v\" %v\n", receipt.PurchaseTime, err)
		return -1, errors.New("error in parsing time")
	}
	startTime, _ := time.Parse(timeLayout, "14:00:00")
	endTime, _ := time.Parse(timeLayout, "16:00:00")

	if parsedTime.After(startTime) && parsedTime.Before(endTime) {
		totalPoints += 10
	}

	log.Infof("Total points calculated: %d", totalPoints)
	return totalPoints, nil
}