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
	var userTicket int


	userName = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)
}
