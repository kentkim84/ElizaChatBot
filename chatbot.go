package main

import (
	"net/http"
	"log"
	"html/template"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

const RAND_MAX_NUMBER int = 20

type PageVariables struct {
	HeadTitle string
	BodyTitle string
	Response string
	Message string
}

type GameVariables struct {
	TryNumber int
	RandomNumber int
}

var gv GameVariables

// create slices(=dynamic array)
var numContainer []int = make([]int, 0)

func initGameVar() {
	gv = GameVariables{TryNumber: 0, RandomNumber: 0}
	//fmt.Println("gv: ", gv.TryNumber, gv.RandomNumber)
}

// generate a random number
func genRandInt(minNum int, maxNum int) int {
    return rand.Intn(maxNum) + minNum
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game index page"
	bodyTitle := "Guessing Game"

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		BodyTitle : bodyTitle,
	}

	// parse the template
	t, err := template.ParseFiles("index.html")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)
	// fmt.Fprintf(w, "Execute indexHandler")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game page"
	message := "Guess a number between 1 and 20"
	save := true

	// set random number
	rand.Seed(time.Now().UnixNano())
	randNum := genRandInt(1, RAND_MAX_NUMBER)

	// set up cookie
	cookie := http.Cookie{Name: "target", Value: strconv.Itoa(randNum)}
	_, err := r.Cookie("target")
	if err == nil { // if there is no error
		http.SetCookie(w, &cookie)
	}

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		Message : message,
	}

	// parse the template
	t, err := template.ParseFiles("guess.tmpl")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)


	// read the value from the input box in the form
	guessNum := r.FormValue("guess")
	iGuessNum, err := strconv.Atoi(guessNum)
	
	

	// increase the try number by 1
	if gv.TryNumber == 0 && gv.RandomNumber == 0 {
		gv.RandomNumber = randNum
	}
	
	fmt.Fprintf(w, "<h3>guess number: %d</h3>", iGuessNum)

	if (iGuessNum != 0 && iGuessNum != gv.RandomNumber) {

		fmt.Fprintf(w, "Number of tries: %d<br>", gv.TryNumber)

		if (iGuessNum < gv.RandomNumber) {
			fmt.Fprintf(w, "Your number is lower than the random number<br>")
		} else if (iGuessNum > gv.RandomNumber) {
			fmt.Fprintf(w, "Your number is higher than the random number<br>")
		}

		for i := 0; i < len(numContainer); i++  {
			if (iGuessNum != 0 && iGuessNum == numContainer[i]) {
				fmt.Fprintf(w, "<h2>Enter a different guessing number</h2>")
				save = false
			}
		}

		if save == true {
			// store the guessing number in the number container when it's not a repeated number
			numContainer = append(numContainer, iGuessNum)
		}

		fmt.Fprintf(w, "A number in the container: %d<br>", numContainer)
	}

	if (iGuessNum == gv.RandomNumber) {
		fmt.Fprintf(w, "You have tried: %d<br>", gv.TryNumber)
		fmt.Fprintf(w, "Numbers matched<br>")
	}
	gv.TryNumber += 1
}

func main() {
	// call initGamevar to initialise game variables
	initGameVar()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/guess/", guessHandler)
    http.ListenAndServe(":8080", nil)
}