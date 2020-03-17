package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("It's the guessing game. Let GO!!!")

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(20)

	maxTrail := 5
	var guess int
	var counter int

	for {
		if counter < maxTrail {
			fmt.Println("Please guess a number: ")
			fmt.Scan(&guess)

			if guess < secretNumber {
				fmt.Println("Too small")

			} else if guess > secretNumber {
				fmt.Println("Too big")
			
			} else {
				fmt.Println("Yipee!!! you guess right.")
				fmt.Printf("You guess %d times", counter)
				break
			}
		} else {
			fmt.Println("You have exceeded your trail. Please try again")
			break
		}

		counter++
	}
	
}
