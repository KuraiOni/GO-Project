package main

import "fmt"

func main() {

	var conferenceName = "GO conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("You can book yuor ticket here")

	var firstName string
	var lastName string
	var email string
	var tickets int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&tickets)

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, tickets)
}
