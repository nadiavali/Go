 package main
 import ("fmt"
		"unicode")

// /*madam"
// "racecar"
// "level"*/

// func palindromeChecker(str string) bool {
// 	res := true
// 	mid := len(str) / 2
// 	if len(str) % 2 == 0{
// 		if str[:mid]==str[mid:]{
// 			return res
// 		}else {
// 			return false
// 		}
// 	}
// 	for i:= 0; i <= len(str); i++ {
// 		if len(str) % 2 != 0{
// 			if str[:mid] == str[mid+1:len(str)]{
// 				return res
// 			}
// 		}
// 	}
// 	return res
// }

// func mai() {
// 	fmt.Println(palindromeChecker("a man"))
// 	fmt.Print(len("a man")) // not working
// }


func palindromeChecker(str string) bool {
	// Normalize the string: remove spaces and convert to lowercase
	var cleaned []rune
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) { // Ignore non-alphanumeric characters
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	// Check if the cleaned string is a palindrome
	mid := len(cleaned) / 2
	for i := 0; i < mid; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindromeChecker("a man"))       // Output: false
	fmt.Println(palindromeChecker("A man a plan a canal Panama")) // Output: true
	fmt.Println(palindromeChecker("madam"))       // Output: true
	fmt.Println(palindromeChecker("racecar"))     // Output: true
	fmt.Println(palindromeChecker("hello"))       // Output: false
}
