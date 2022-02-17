package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still availabel\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint
		// fmt.Println("Enter your first name:")
		// fmt.Scan(&firstName)
		// fmt.Println("Enter your last name:")
		// fmt.Scan(&lastName)
		// fmt.Println("Enter email address:")
		// fmt.Scan(&email)
		// fmt.Println("Enter number of tickets:")
		// fmt.Scan(&userTickets)
		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email,"@")
		// isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTickets

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumbers := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumbers{

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			// firstNames := []string{}
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }
			// fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			} 
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if ! isValidTicketNumbers {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}


		wg.Wait()
	// city := "London"
	// switch city {
	// case "New York","Mexico City":
	// 	//execute code for booking New York or Mexico City conference tickets
	// case "Singapore","Hong Kong":
	// 	//execute code for booking Singapore or Hong Kong conference tickets
	// case "London", "Berlin":
	// 	//execute code for booking London or berlin conference tickets
	// default:
	// 	fmt.Print("No valid city selected")
	// }

}

func greetUsers() {
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We have total of %v tickets and %v are still availabel\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)
			}
			return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email,"@")
// 	isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumbers
// }

func getUserInput() (string, string, string, uint){

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets	
}


func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################################################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("##################################################")

	wg.Done()
}