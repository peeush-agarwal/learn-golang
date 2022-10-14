package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidNames := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidNames, isValidEmail, isValidUserTickets
}
