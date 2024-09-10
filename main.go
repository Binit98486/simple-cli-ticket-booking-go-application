package main

import (
	"fmt"
	"go-code/helper"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingConferenceTickets int = 50

// list of map
var bookings = make([]UserData, 0)

// var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// get user input
	firstName, lastName, email, userTickets := getUserInput()

	// validate user input
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingConferenceTickets)

	if isValidEmail && isValidName && isValidTicketNumber {
		booking(userTickets, firstName, lastName, email)

		wg.Add(1)
		// concurrency
		go sendTickets(userTickets, firstName, lastName, email)
		// print firstname
		firstNames := getFirstName()
		fmt.Printf("first name %v\n", firstNames)

		if remainingConferenceTickets == 0 {
			fmt.Printf("%v all ticket is sold out.Come back next year\n", conferenceName)
			// break

		}

	} else {
		if !isValidName {
			fmt.Println("First Name and Last Name is too short")

		}
		if !isValidEmail {
			fmt.Println("Email must contain @ sign")

		}
		if !isValidTicketNumber {
			fmt.Println("Ticket number is not valid")
		}
		// fmt.Println("Your input data is invalid, try again")
	}

	// city := "Ktm"
	// switch city {
	// case "kathmandu":

	// case "lalitpur":
	// case "pokhara":
	// default:
	// 	fmt.Println("Enter city is not valid")
	// }

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v remaining tickets\n", conferenceTickets, remainingConferenceTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println(("Enter your last name"))
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter your ticket number")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func booking(userTickets int, firstName string, lastName string, email string) {
	remainingConferenceTickets = remainingConferenceTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v %v\n", firstName, lastName, userTickets, email, bookings)

}
func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################################")
	fmt.Printf("Sending ticket \n %v \n to email address %v\n", ticket, email)
	fmt.Println("#################################")
	wg.Done()

}
