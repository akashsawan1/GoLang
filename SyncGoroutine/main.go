package main

import (
	"fmt"
	"packages/helper"
	"time"
	"sync"
)
var conferenceName = "Akash Conference"
const conferenceTickets = 55
var remainingTickets uint = 50 //for positive integer
var bookings = make([]UserData,0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")


		
		userName, lastName, email, userTicket := getUserInput()


		isValidName, isValidemail, isValidTicket := helper.ValidateUserInput(userName, lastName, email, userTicket, remainingTickets)
		

		
		if isValidName && isValidemail && isValidTicket {

			bookTicket( userTicket , userName, lastName, email)
			wg.Add(1)
			go sendTicket( userTicket , userName, lastName, email)

			
			
			firstNames := getFirstNames()
			fmt.Printf("These are all your bookings: %v\n", firstNames)
			} else if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
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
				wg.Wait()
				}

			

	func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, " are remaining")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		// var name = strings.Fields(booking)
		// firstName := name[0]
		firstNames = append(firstNames, booking.firstName)
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

	var userData =  UserData {
		firstName : userName,
		lastName : lastName,
		email : email, 
		numberOfTickets : userTicket,
	}



	bookings = append(bookings, userData)

	fmt.Printf("Here,Are all the list of booking %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
}


func sendTicket(userTicket uint, userName string, lastName string, email string) { 
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTicket, userName , lastName)
	fmt.Println("########")
	fmt.Printf("Sending Ticket : %v, to email %v\n", ticket,  email)
	fmt.Println("########")
	wg.Done()
}