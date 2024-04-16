package user

import (
	"fmt"

	"github.com/abdurahmonrixsiboyev1998/tahsir/validateuser"
)

func Validate() {
	var name string
	var age int
	var email string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	err := validateuser.ValidateUser(name, age, email)
	if err != nil {
		for _, e := range err {
			fmt.Println(e)
		}
	} else {
		fmt.Println("Successfully!!")
	}

}
