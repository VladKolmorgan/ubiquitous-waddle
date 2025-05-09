import (
	"fmt"
	"os"
	//"github.com/go-openapi/strfmt"
	add import for github.com/go-openapi/strfmt
)

func main() {

	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}

	//a := strfmt.IsEmail("vlad@test.com")
	//fmt.Println(a)
}