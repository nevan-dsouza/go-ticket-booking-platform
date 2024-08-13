package main

import (
	"fmt"
	"go-ticket-booking-platform/helper"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userEmail   string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, remainingTickets, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, userEmail, userTickets)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, userEmail)

			// Total tickets update
			fmt.Printf("There are only %v tickets remaining\n", remainingTickets)

			firstNames := getFirstNames()

			// Total users update
			fmt.Printf("The people in the conference are: %v\n\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("The %v conference is booked out! Come back next year :')\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Your number of tickets are invalid. There are only %v tickets remaining", remainingTickets)
			}
		}
		wg.Wait()
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets. %v tickets are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var userEmail string

	// ask for user info
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Println("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(firstName string, lastName string, userEmail string, userTickets uint) ([]UserData, string, string, string, uint, uint) {

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userEmail:   userEmail,
		userTickets: userTickets,
	}

	// User booking confirmation
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	// Total tickets update
	remainingTickets -= userTickets

	return bookings, firstName, lastName, userEmail, remainingTickets, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############################################")
	fmt.Printf("Sending ticket %v to email address: %v\n", ticket, userEmail)
	fmt.Println("###############################################")

	wg.Done()
}
