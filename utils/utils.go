package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/eiannone/keyboard"
)
var (
	reader = bufio.NewReader(os.Stdin)
)

func SortAndPrint(mapContent map[int]string) int {
	mapLength := len(mapContent)
	sortedContent := make([]int, 0, mapLength)
	fmt.Println(mapLength)
	for k := range mapContent {
		sortedContent = append(sortedContent, k)
	}
	sort.Ints(sortedContent) //

	for _, k := range sortedContent {
		fmt.Printf("[%d]: %s\n", k, mapContent[k])
	}
	fmt.Printf("\nPress a number (1-%d) to choose a category:", mapLength)
	return mapLength

}

func GetCategory(mapContent map[int]string) string {
	mapLength := SortAndPrint(mapContent)
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal("Error reading category selection:", err)
		}

		mapIndex := int(char - '0') // Convert ASCII to integer
		if mapIndex >= 1 && mapIndex <= mapLength {
			fmt.Println("You selected:", mapContent[mapIndex])
			return mapContent[mapIndex]
		} else {
			fmt.Println("Invalid selection. Please press a number between 1 and 8.")
		}
	}
}


func GetText(title string) (string, error) {
	fmt.Printf("Enter the %s: ", title)
	line, err := reader.ReadString('\n')

	if line == "" || err != nil {
		return "", err
	}
	return line, nil
}

func GetAmount() (int, error) {
	var amount int
	fmt.Printf("Enter the Amount: ")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("error reading the amount")
		return -1, err
	}
	fmt.Println("the amount is: ", amount)
	return amount, nil
}