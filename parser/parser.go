package main
import "fmt"
import "parser/pkg/structColumnDefs"

// see https://yuminlee2.medium.com/golang-import-local-packages-cbb2201c72e1
// for a brief explanation of how to define and use local imports

// cd to the root of the project and run:
// go mod init parser
// module parser
// go run .


func main() {
    fmt.Printf("\nstructColumnDefs\n")
	ret, error := structColumnDefs.ParseIt("VL1L2", "this is a line of data")
	fmt.Printf("returned %s with err %s\n", ret, error)
    ret, error = structColumnDefs.ParseIt("VAL1L2", "this is a line of data")
	fmt.Printf("returned %s with err %s\n", ret, error)
}