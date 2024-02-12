package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTicksets uint = 50
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTicksets)
	println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicketes uint

	var bookings []string //slice

	//for i := 0; i <= int(remainingTicksets); i++ {
	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter no of tickets to purchase")
		fmt.Scan(&userTicketes)

		remainingTicksets = remainingTicksets - userTicketes
		bookings = append(bookings, firstName+" "+lastName)

		//i = remainingTicksets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email on %v.", firstName, lastName, userTicketes, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTicksets, conferenceName)
		//fmt.Printf("the names of all bookings are %v\n", bookings)

		firstNames := []string{}
		for index, booking := range bookings {
			fmt.Printf("The booking index is %v and booking entry per index is %v\n", index, booking)
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The firstname of bookings are %v\n", firstNames)

		var noTicketsRemaining bool = remainingTicksets == 0
		if noTicketsRemaining {
			//end the program
			fmt.Println("Our conference is booked out. Come back ext year")
			break
		}
	}
}
