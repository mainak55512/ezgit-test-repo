package main

// import "fmt"
import (
	"os"
	"os/exec"
)

// func main() {
// 	var (
// 		text1, text2 string
// 	)
// 	fmt.Print("Enter some text: ")
// 	fmt.Scanln(&text1)
// 	fmt.Print("Enter some other text: ")
// 	fmt.Scanln(&text2)
//
// 	fmt.Println("User Inputs: ", text1, text2)
// }

func main() {
	cmd := exec.Command("git", "--version")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
