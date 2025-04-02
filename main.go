package main

import (
	"fmt"
	"log"

	"github.com/ashikkabeer/expense-on-cli/models"
	"github.com/ashikkabeer/expense-on-cli/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// importing the models
	category := models.Category
	paymentMethod := models.PaymentMethod
	cashType := models.CashType

	// calling each questions
	categoryResult := utils.GetCategory(category)
	description, err := utils.GetText("description")
	if err != nil {
		log.Fatal(err)
	}
	paymentMethodResult := utils.GetCategory(paymentMethod)
	amount, err := utils.GetAmount()
	if err != nil {
		log.Fatal(err)
	}
	cashTypeResult := utils.GetCategory(cashType)
	remark, nil := utils.GetText("remark")
	if err != nil {
		log.Fatal(err)
	}
	// displaying it to the user
	fmt.Println(categoryResult)
	fmt.Println(description)
	fmt.Println(paymentMethodResult)
	fmt.Println(amount)

	fmt.Println(cashTypeResult)
	fmt.Println(remark)

	// get the bank balance
	// calculate the new balance
	// display it to user

}
