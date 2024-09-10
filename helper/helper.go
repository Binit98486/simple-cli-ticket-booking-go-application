package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingConferenceTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingConferenceTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
