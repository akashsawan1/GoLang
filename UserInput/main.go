package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Akash Conference"
	const conferenceTickets = 55
	var remainingTickets uint = 50; //for positive integer 

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets,conferenceName)
	fmt.Printf("Welcome to %v booking application\n ", conferenceName)
	fmt.Println("We have total of" , conferenceTickets , "tickets and " , remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var lastName string
		var email string
		var userTicket uint
		bookings := []string{}
		
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many Tickets you want to book: ")
		fmt.Scan(&userTicket)

		isValidName := len(userName) >=2 && len(lastName) >=2
		isValidemail := strings.Contains(email, "@")
		isValidTicket := userTicket > 0 && userTicket <= remainingTickets 

		// isValidCity := city == "Singapore" || city == "London"

		// isInValidCity := city != "Singapore" || city != "London"

		if isValidName && isValidemail && isValidTicket {
			
		remainingTickets = remainingTickets - userTicket

		bookings = append(bookings, userName + " " + lastName)
		
		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type : %T\n", bookings)
		fmt.Printf("array length : %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstName := name[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)
		} else if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short")
			}
			if !isValidemail {
				fmt.Println("Email is invalid")
			}
			if !isValidTicket {
				fmt.Printf("number of tickets you entered is Invalid")
			}
			
			
		}
		// or
		// noTicketsRemaining := remainingTickets == 0
		// if noTicketsRemaining { 
		// fmt.Println("Our conference is booked out,come back next year.")
		// }
	}

	
}
