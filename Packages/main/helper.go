package main

import "strings"

func validateUserInput(userName string, lastName string, email string, userTicket uint) (bool, bool ,bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidemail, isValidTicket
}
