/* Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example

For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].*/
package main
import ( "fmt"
"strings"
"sort"
"strconv")

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
func GetSizes(w, h, d int) []int {  // Change return type to []int (slice)
  res := []int{}                   // Initialize res as an empty slice
  area := 2*(w*h + w*d + h*d)      // Correct surface area calculation
  volume := w * h * d
  res = append(res, area, volume)  // Append area and volume to the slice
  return res
}


func GetSizea(w, h, d int) [2]int {
  res := [2]int{}
  area := 2*(w*h + w*d + h*d) // Correct formula for surface area
  volume := w * h * d
  res[0] = area
  res[1] = volume
  return res
}

    
func ReverseSeqi(n int) []int {
  res := []int{}
  i := 0
  for i=1; i <= n; i++{
  res = append(res, i)
  }
  return res
}

func ReverseSeq(n int) (result []int) {
  for i := n; i > 0; i-- {
    result = append(result, i)
  }
  return
}


func ReverseSeqr(n int) []int {
  res := []int{}
  for i:=n; i >= 1; i-- {
  res = append(res, i)
  }
  return res
}


/*Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.*/


func GetCount(str string) (count int) {// return: init:count := 0
  vo := "aeiou"
  for _, char := range str{
    if strings.ContainsRune(vo, char){ //if char is a vo
      count++
     }
 }
  return count
}



/*Complete the function that takes a non-negative integer n as input,
and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).

Examples

n = 0  ==> [1]        # [2^0]
n = 1  ==> [1, 2]     # [2^0, 2^1]
n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]*/

func PowersOfTwo(n int) []uint64 {
  res :=[]uint64{}
  for i:=0; i<=n; i++{
    res = append(res, 1<<uint(i)) //2**i
  }
  return res
}

/*Complete the square sum function so that it squares each number passed into it and then sums the results together.

For example, for [1, 2, 2] it should return 9*/

func SquareSum(numbers []int) int {
    // your code here
  sum := 0
  for i:= range numbers{
    sum += numbers[i] * numbers[i]
  }
  return sum
}

/*Complete the solution so that it returns true if the first argument(string)
 passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false*/


func solution(str, ending string) bool {
  // Your code here!
  if len(str) < len(ending){
    return false
  }
  return str[len(str)-len(ending):] == ending
}

/*Deoxyribonucleic acid, DNA is the primary information storage molecule in biological systems.
 It is composed of four nucleic acid bases Guanine ('G'), Cytosine ('C'), Adenine ('A'), and Thymine ('T').

Ribonucleic acid, RNA, is the primary messenger molecule in cells.
RNA differs slightly from DNA its chemical structure and contains no Thymine.
 In RNA Thymine is replaced by another nucleic acid Uracil ('U').

Create a function which translates a given DNA string into RNA.

For example:

"GCAT"  =>  "GCAU"*/

func DNAtoRNA(dna string)string{
  return strings.ReplaceAll(dna, "T", "U")
}


/*In this little assignment you are given a string of space separated numbers,
 and have to return the highest and lowest number.

Examples

HighAndLow("1 2 3 4 5") // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes

All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.*/



func HighAndLow(in string) string {
  numStrings := strings.Fields(in) //numStrings := []string{"1", "2", "-3", "4", "5"}

  nums := []int{}
  
  for _, n := range numStrings {
    j, _ := strconv.Atoi(n)
    nums = append(nums, j)
  }
  sort.Ints(nums)
  return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}
/* Given an integer as input, can you round it to the next 
(meaning, "greater than or equal") multiple of 5?

Examples:

input:    output:
0    ->   0
2    ->   5
3    ->   5
12   ->   15
21   ->   25
30   ->   30
-2   ->   0
-5   ->   -5
etc.'''**/

func RoundToNext5(n int) int {
  if n % 5 == 0{
    return n
  }
  if n % 5 != 0 && n > 0{
    return n + (5 - n % 5)
  } else  if n % 5 != 0 && n < 0 {
    return (n - n % 5)
  }
  return 0
}
func RoundToNext5Celever(n int) int {
  for n % 5 != 0 {
    n++
  }
  return n
}

/*You are going to be given a non-empty string. Your job is to return the middle character(s) of the string.

If the string's length is odd, return the middle character.
If the string's length is even, return the middle 2 characters.
Examples:

"test" --> "es"
"testing" --> "t"
"middle" --> "dd"
"A" --> "A"
*/

func MiddleString(s string) string {
  length := len(s)
  middle := len(s) / 2
  if length % 2 == 0 {
    return s[middle-1  : middle+1]
  } else {
    return s[middle : middle+1]
  }
}

/*Your goal is to return multiplication table for number that is always an integer from 1 to 10.

For example, a multiplication table (string) for number == 5 looks like below:

1 * 5 = 5
2 * 5 = 10
3 * 5 = 15
4 * 5 = 20
5 * 5 = 25
6 * 5 = 30
7 * 5 = 35
8 * 5 = 40
9 * 5 = 45
10 * 5 = 50
P. S. You can use \n in string to jump to the next line.

Note: newlines should be added between rows, but there should be no trailing newline at the end.
 If you're unsure about the format, look at the sample tests.
 (("1 * 1 = 1\n2 * 1 = 2\n3 * 1 = 3\n4 * 1 = 4\n5 * 1 = 5\n6 * 1 = 6\n7 * 1 = 7\n8 * 1 = 8\n9 * 1 = 9\n10 * 1 = 10"))*/

func MultiTable(number int) string {
  res := ""
  for n:=1; n <= 10; n++{
    if n < 10 {
      res += fmt.Sprintf("%d * %d = %d\n", n, number, n*number)
    } else{
      res += fmt.Sprintf("%d * %d = %d", n, number, n*number)
    }
  }
  return res
}