package main

import (
	"fmt"
)

var confrenceName string = "Go Confrence"

const confrenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]User, 0) // slice of User struct to store bookings

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)
			// call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name must be at least 2 characters long.")
			}

			if !isValidEmail {
				fmt.Println("Email address must contain @ sign.")
			}

			if !isValidTicketNumber {
				fmt.Printf("Number of tickets must be between 1 and %v.\n", remainingTickets)
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{} // slice of strings to store first names

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName) // append the first name of each booking to the firstNames slice
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userDate = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userDate)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

}
