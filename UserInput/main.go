package main

import "fmt"

func main() {
	conferenceName := "Akash Conference"
	const conferenceTickets = 55
	var remainingTickets uint = 50; //for positive integer 

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets,conferenceName)
	fmt.Printf("Welcome to %v booking application\n ", conferenceName)
	fmt.Println("We have total of" , conferenceTickets , "tickets and " , remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")


	var userName string
	var lastName string
	var email string
	var userTicket uint
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many Tickets you want to book: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
