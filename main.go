package main

import (
	"fmt"
	"os"
)

var defaultPIN = 0000

var accountBalance = 00.00

var menuNo int

func main() {
	welcome()
	start()
	menu()
}

func newLine(numberOfLInes int) {
	i := 0
	for i < numberOfLInes {
		fmt.Print("\n")
		i++
	}
}
func welcome() {
	fmt.Println("******{Welcome to This ATM Machine}*******")
	newLine(2)
}

func start() {
	fmt.Println("Please enter your Pin: ")
	_, err := fmt.Scan(&defaultPIN)
	if err != nil {
		return
	}
	if defaultPIN == 0000 {
		menu()
	} else {
		fmt.Println("Error: Incorrect Pin")
		start()
	}
}

func exitProgram() {
	fmt.Println("Thank You For Using This ATM Machine.\n Please take your card.")
	os.Exit(0)
}

func changePin() {
	newLine(1)
	fmt.Println("Please Enter your new Pin : ")
	var newItem int
	_, err := fmt.Scan(&newItem)
	if err != nil {
		fmt.Println("Error: Invalid Input! ")
	}
	defaultPIN = newItem
	fmt.Println("Pin Changed Successfully.")
	
	anotherOperation()
}

func viewAccountBalance() {
	newLine(1)
	fmt.Println("Your Account Balance is: ", accountBalance)
	anotherOperation()
}

func withdrawFunds() {
	newLine(1)
	fmt.Println("Enter the amount you want to Withdraw :")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid Input !")
	}
	if amount > accountBalance {
		fmt.Println("Error: Insufficient Funds !")
	} else {
		accountBalance -= amount
		fmt.Println("Successful Withdrawal")
	}
	anotherOperation()
}

func depositFunds() {
	newLine(1)
	fmt.Println("Enter Deposit Amount: ")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid Input")
	}
	accountBalance += amount
	fmt.Println("Deposit successful. ")
	anotherOperation()
}

func anotherOperation() {
	fmt.Println("DO YOU WANT TO PERFORM ANOTHER OPERATION ? \n\n For Yes,Please click 'y' \n\n For No,Please Click 'n'. ")
	var solution string
	_, err := fmt.Scan(&solution)
	if err != nil {
		fmt.Println("Error: Invalid Input!")
	}
	if solution == "y" {
		menu()
	} else if solution == "n" {
		exitProgram()
	} else {
		fmt.Println("Error: Invalid Input! ")
	}

}
func menu() {
	newLine(2)
	fmt.Println("Select Your Option: ")
	fmt.Println("1.Change Pin      2.Account Balance ")
	fmt.Println("3.Withdraw Funds  4.Deposit Funds ")
	fmt.Println("0.Cancel/Exit ATM")
	newLine(2)

	_, err := fmt.Scan(&menuNo)
	if err != nil {
		fmt.Println("Error: Please Select the correct Menu Item")

		menu()
	}

	switch menuNo {
	case 0:
		exitProgram()
	case 1:
		changePin()
	case 2:
		viewAccountBalance()
	case 3:
		withdrawFunds()
	case 4:
		depositFunds()
	default:
		fmt.Println("Error: Invalid Input")
		menu()
	}
}
