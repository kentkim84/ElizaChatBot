package main

import (
	//"bytes"
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
	"regexp"
	"strings"
)

func elizaResponse(message string) (string) {
	var response string
	responseList := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		"Why don’t you tell me more about your father?"}

	randNum := rand.Intn(len(responseList)-1)
	//fmt.Printf("%d\n",randNum)

	//fmt.Println(message)
	r, _ := regexp.Compile(`\b[Ff]ather`)
	
	if r.MatchString(message) {
			response = responseList[3]
	} else {
		r, _ = regexp.Compile("(?i)I am ([^.?!]*)[.?!]?")
		if r.MatchString(message){
			response = r.ReplaceAllString(message, "$1?")
			reflect := Reflect(response)			
			response = "How do you know you are " + reflect			
		} else {
			response = responseList[randNum]
		}
	}

	//fmt.Println("String match:",r.MatchString(message))
	return response
}

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `You`},
		{`You`, `I`},		
		{`my`, `your`},
		{`your`, `my`},
		{`me`, `you`},
		{`you`, `me`},
		{`am`, `are`},
		{`are`, `am`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}

/* func main() {

	testInputList01 := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!"}

		testInputList02 := []string{
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?"}
	
	
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Father --------------------------------------------------------\n")
	for testInput := range testInputList01 {
		fmt.Printf("%s\n\n", elizaResponse(testInputList01[testInput]))
	}
	fmt.Println("I am ----------------------------------------------------------\n")
	for testInput := range testInputList02 {
		fmt.Printf("%s\n\n", elizaResponse(testInputList02[testInput]))
	}
	fmt.Println("Reflect ----------------------------------------------------------\n")
	fmt.Printf("%s\n\n", Reflect("I am your friend."))
	fmt.Printf("%s\n\n", Reflect("You are my friend."))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Say something")
	for {		
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		fmt.Printf("%s\n", elizaResponse(text))

	}
} */