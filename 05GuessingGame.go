package main

// Time is used to get a random number so it has a random number every time
import(
    "fmt"
    "math/rand"
    "time"
)

// Generates random number between the given range
func rando(min, max int) int {

    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {

	// Initialise variables
    myrand := rando(1, 50)
    prevGuess := 0
    var guess int
    var lastGuess int = 0
	
    fmt.Printf("Guess a number between 1 and 50\n")
    
    // Loops through 1-50 so the user gets as much guesses as needed
    for prevGuess < 50 {
        fmt.Println("Take a guess...")
		fmt.Scanf("%d", &guess)
		
        // Flushes the buffer
        fmt.Scanf("%d")
		prevGuess++
		
        // If statement checks to see if the user input a number between 1 and 50
        if (guess >= 1 && guess <= 50){

            // Checks if the user inputs the same number as they last entered
            if lastGuess == guess {
                fmt.Println("You've already guessed this, try again")
                prevGuess--
            }
            if guess < myrand {
                fmt.Println("Try higher")
            }
            if guess > myrand {
                fmt.Println("Go lower")
            }
            if guess == myrand {
                break
            }
            lastGuess = guess
        } else {
            fmt.Println("You must enter a number between 1 and 50");
        }
    } 
    if guess == myrand {
        fmt.Printf("Well done, you guessed the number in %d tries\n", prevGuess)
    } else {
        fmt.Printf("Next time, the number was %d\n", myrand)
    }
}