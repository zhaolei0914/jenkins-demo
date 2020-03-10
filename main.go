package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World 粗发, Kubernetes！I'm from Jenkins CI！")
	time.Sleep(time.Second*10000)
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}
