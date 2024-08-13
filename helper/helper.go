package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, userEmail string, remainingTickets uint, userTickets uint) (bool, bool, bool) {
	// Validation of Details
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
