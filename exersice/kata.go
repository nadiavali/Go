/* Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example

For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].*/
package main
import "fmt"

var x =[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}


func main(){
  fmt.Println(CountPositivesSumNegatives(x))
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
