package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" // can use the variable in any part of the application
const conferenceTickets int = 50     // number of tickets
var remainingTickets uint = 50       // remainig tickets
//var bookings []string{}         // creating a slice from an array by leaving the [] empty.
var bookings = make([]UserData, 0) // intiliazing slice. or list of maps

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//fmt.Println("Welcome to", conferenceName, "booking application") // longer way in putting placeholder
	//fmt.Printf("Welcome to %v booking application\n", conferenceName) // shorter way in putting placeholder
	//fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//fmt.Println("Get your tickets to attend conference!!!")

	//for loop is used for looping through the aopp
	//for {
	//Validate userinfo.
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// call functions
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) // add "go" to call function make it concurrency.

		firstNames := getFirstName()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end the program
			fmt.Println("Our conference is booked out. come back later")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("Name is too short!!")
		}
		if !isValidEmail {
			fmt.Println("email doesn't match")
		}
		if !isValidTicketNumber {
			fmt.Println("invalid ticket number")
		}

		wg.Wait()
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend conference!!!")
}

//func is going through the booking list.
func getFirstName() []string {
	firstNames := []string{}           // this loop gets the slice of the first name from the slices. $ local variable
	for _, booking := range bookings { //_ is a blank identifier
		//var names = strings.Fields(booking)

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask the user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // user input: basically whatever user enter as name will be stored here.
	//fmt.Println(&remainingTickets) Prints out memory address
	//& Pointer(var points to memory address of another var)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user and "make" create empty map
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTickets,
	}

	//bookings[0] = firstName + " " + lastName // stores the booking in an array.
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings) // list of booking structs
	fmt.Printf("Thank you %v %v for booked %v tickets.\nYou will receive a confirmation email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

//Concurrency
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("#############")
	wg.Done() // this decrease the counter
}
