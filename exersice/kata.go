/* Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example

For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].*/
package main
import ( "fmt"
"strings")

var x =[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}


func main(){
  fmt.Println(CountPositivesSumNegatives(x))
  fmt.Println(AbbrevName("Nadi avali"))
  fmt.Println(QuarterOf(7))
  fmt.Println(AmIWilson(563))
}

func CountPositivesSumNegatives(numbers []int) []int {
  res := []int{}
  var count int
  var sum int
  for i := 0; i < len(numbers); i++{
    if numbers[i] > 0 {
      count++
    } else if numbers[i] < 0{
      sum += numbers[i]
    }
  }
  res = append(res, count, sum)
  return res
}


/*Write a function to convert a name into initials.
 This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F*/

func AbbrevName(name string) string{
  apparted := strings.Split(name, " ")
  stringAppart := fmt.Sprintf("%c.%c", apparted[0][0], apparted[1][0])

  return strings.ToUpper(stringAppart)
}


func QuarterOf(month int) int {
  return (month + 2) / 3
}


/*
Wilson primes satisfy the following condition. Let 
P
P represent a prime number.

Then,

(P−1)!+1/P∗P

​should give a whole number, where 
P! is the factorial of P.

Your task is to create a function that returns true if the given number is a Wilson prime and false otherwise.

*/

func AmIWilson(n int) bool {
  if n < 2 { //prime
    return false
  }
  factorial := 1 //factorial l
  for i := 2; i < n; i++ {
    factorial *= i
    factorial %= n * n // Using Modular Arithmetic to Simplify Factorial Calculation(toooo Large)
  }
 
  return  (factorial + 1) % n * n == 0
}

    
