package main

import(

	"fmt"
)
//var email string = "chawrainak_p@su.ac.th"

func main (){

	// var name string = "porawat"
	var age int = 21

	email := "chawrainak_p@su.ac.th"
    gpa := 3.75

	firstname,lastname := "porawat","chawrainak"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n,",firstname,lastname,age, email,
	gpa)
}