package main
import ("fmt"
	"os/exec"
)
func main() {
	command := "echo -en \"\\a\" > /dev/tty1"
	err := exec.Command("sh", "-c", command).Run()
	if err != nil {
		fmt.Println("no beeboop :( | ", err)
	} else {
		fmt.Println("beeboop :)")
	}
}
