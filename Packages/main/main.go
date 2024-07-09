package main

import (
	"fmt"
	"strings"
)
var conferenceName = "Akash Conference"
const conferenceTickets = 55
var remainingTickets uint = 50 //for positive integer
var bookings = []string{}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")

	for {
		
		userName, lastName, email, userTicket := getUserInput()


		isValidName, isValidemail, isValidTicket := validateUserInput(userName, lastName, email, userTicket)
		

		
		if isValidName && isValidemail && isValidTicket {

			bookTicket( userTicket , userName, lastName, email)

			
			
			firstNames := getFirstNames()
			fmt.Printf("These are all your bookings: %v\n", firstNames)
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
						fmt.Printf("number of tickets you entered is Invalid\n")
					}
					
				}

				}

			}

	func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstName := name[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}



func getUserInput() (string,string, string, uint) {
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
	return userName, lastName, email , userTicket
}


func bookTicket( userTicket uint,  userName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, userName+" "+lastName)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
}
