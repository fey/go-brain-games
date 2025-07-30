package main

import (
		"bufio"
		"fmt"
		"os"
)
// import (
// 	"fmt"

// 	"github.com/manifoldco/promptui"
// )

// func main() {
// 	prmpt := (promptui.Prompt{
// 		Label: "Mah I hamve your name?",
// 	})
// 	name, _ := prmpt.Run()

// 	fmt.Printf("Hello, %s!\n", name)
// }

func main() {
	fmt.Println("Welcome to the Brain Games!")
	name := Question("May i have your name? ")

	fmt.Printf("Hello, %s\n", name)
}

func Question(query string) string {
	fmt.Print(query)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	return answer
}
