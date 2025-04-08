package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
 "unicode"
 "sync"
)

// Input - Mimics Python's input() with prompt and error handling, running asynchronously.
func Input(prompt string, result chan<- string) {
 fmt.Print(prompt + ": ")
 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {
  input := scanner.Text()
  if input == "" {
   // Handle empty input or prompt the user again.
   fmt.Print("Input cannot be empty. Please try again: ")
  } else {
   result <- input // Send result to channel
   return
  }
 }
 // Handle scanner errors
 if err := scanner.Err(); err != nil {
  fmt.Println("Error reading input:", err)
 }
 result <- "" // Send empty string on error
}

// Lower - Converts the string to lowercase asynchronously.
func Lower(input string, result chan<- string) {
 var transformed []rune
 for _, char := range input {
  if char >= 'A' && char <= 'Z' {
   transformed = append(transformed, char+'a'-'A')
  } else {
   transformed = append(transformed, char)
  }
 }
 result <- string(transformed)
}

// Upper - Converts the string to uppercase asynchronously.
func Upper(input string, result chan<- string) {
 var transformed []rune
 for _, char := range input {
  if char >= 'a' && char <= 'z' {
   transformed = append(transformed, char-'a'+'A')
  } else {
   transformed = append(transformed, char)
  }
 }
 result <- string(transformed)
}

// Strip - Removes leading and trailing spaces from the string asynchronously.
func Strip(input string, result chan<- string) {
 start, end := 0, len(input)-1

 // Trim leading spaces
 for start <= end && unicode.IsSpace(rune(input[start])) {
  start++
 }

 // Trim trailing spaces
 for end >= start && unicode.IsSpace(rune(input[end])) {
  end--
 }

 result <- input[start : end+1]
}

// Replace - Replaces occurrences of a substring with a new string asynchronously.
func Replace(input, old, new string, result chan<- string) {
 var resultString []rune
 oldLen := len(old)
 i := 0

 for i < len(input) {
  // Find the substring to replace
  if i+oldLen <= len(input) && input[i:i+oldLen] == old {
   resultString = append(resultString, []rune(new)...)
   i += oldLen
  } else {
   resultString = append(resultString, rune(input[i]))
   i++
  }
 }
 result <- string(resultString)
}

// Split - Splits the string into a slice based on a separator asynchronously.
func Split(input, separator string, result chan<- []string) {
 var resultSlice []string
 start := 0

 for i := 0; i < len(input); i++ {
  if input[i:i+len(separator)] == separator {
   resultSlice = append(resultSlice, input[start:i])
   start = i + len(separator)
   i += len(separator) - 1
  }
 }

 // Append the remaining string after the last separator
 if start < len(input) {
  resultSlice = append(resultSlice, input[start:])
 }

 result <- resultSlice
}

// Join - Joins elements of a slice into a string, separated by a specified delimiter asynchronously.
func Join(elements []string, separator string, result chan<- string) {
 var resultString []rune
 for i, element := range elements {
  if i > 0 {
   resultString = append(resultString, []rune(separator)...)
  }
  resultString = append(resultString, []rune(element)...)
 }
 result <- string(resultString)
}

// Count - Counts the number of occurrences of a substring in the string asynchronously.
func Count(input, substr string, result chan<- int) {
 count := 0
 substrLen := len(substr)
 for i := 0; i <= len(input)-substrLen; i++ {
  if input[i:i+substrLen] == substr {
   count++
  }
 }
 result <- count
}

// Find - Finds the first occurrence of a substring and returns its index or -1 asynchronously.
func Find(input, substr string, result chan<- int) {
 substrLen := len(substr)
 for i := 0; i <= len(input)-substrLen; i++ {
  if input[i:i+substrLen] == substr {
   result <- i
   return
  }
 }
 result <- -1
}

// StartsWith - Checks if the string starts with a given substring asynchronously.
func StartsWith(input, substr string, result chan<- bool) {
 if len(input) < len(substr) {
  result <- false
  return
 }
 result <- input[:len(substr)] == substr
}

// EndsWith - Checks if the string ends with a given substring asynchronously.
func EndsWith(input, substr string, result chan<- bool) {
 if len(input) < len(substr) {
  result <- false
  return
 }
 result <- input[len(input)-len(substr):] == substr
}

// IsAlpha - Checks if the string consists only of alphabetic characters asynchronously.
func IsAlpha(input string, result chan<- bool) {
 for _, r := range input {
  if !unicode.IsLetter(r) {
   result <- false
   return
  }
 }
 result <- true
}

// IsDigit - Checks if the string consists only of digits asynchronously.
func IsDigit(input string, result chan<- bool) {
 for _, r := range input {
  if !unicode.IsDigit(r) {
   result <- false
   return
  }
 }
 result <- true
}

// IsSpace - Checks if the string consists only of whitespace characters asynchronously.
func IsSpace(input string, result chan<- bool) {
 for _, r := range input {
  if !unicode.IsSpace(r) {
   result <- false
   return
  }
 }
 result <- true
}

// CapitalizeFirstLetter - Capitalizes the first letter of the string and lowercases the rest asynchronously.
func CapitalizeFirstLetter(input string, result chan<- string) {
 if len(input) == 0 {
  result <- input
  return
 }
 result <- string(unicode.ToUpper(rune(input[0]))) + Lower(input[1:], result)
}

func main() {
 // Create a channel for async results
 resultChan := make(chan string, 1)
 resultBoolChan := make(chan bool, 1)
 resultSliceChan := make(chan []string, 1)
 resultIntChan := make(chan int, 1)
 var wg sync.WaitGroup

 // Example usage of async functions
 wg.Add(1)
 go func() {
  defer wg.Done()
  userInput := Input("Enter some text", resultChan)
  fmt.Println("You entered:", userInput)

  wg.Add(1)
  go func() {
   defer wg.Done()
   lowered := Lower(userInput, resultChan)
   fmt.Println("Lowercase:", lowered)

   uppered := Upper(userInput, resultChan)
   fmt.Println("Uppercase:", uppered)

   stripped := Strip(userInput, resultChan)
   fmt.Println("Stripped:", stripped)

   replaced := Replace(userInput, "go", "Golang", resultChan)
   fmt.Println("Replaced 'go' -> 'Golang':", replaced)

   splitted := Split(userInput, " ", resultSliceChan)
   fmt.Println("Split:", splitted)

   joined := Join(splitted, "-", resultChan)
   fmt.Println("Joined:", joined)
  }()
 }()

 wg.Wait()
}
