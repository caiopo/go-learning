package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	// rand.Seed(42)

	for {
		answers := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes definitely",
			"You may rely on it",
			"As I see it yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}

		Print("Ask the Helix fossil: ")

		var input string

		Scanf("%s", &input)

		rand.Seed(int64(input[0]))

		Println("The Helix says:", answers[rand.Intn(len(answers))])

	}
}
