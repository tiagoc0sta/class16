/*package main

import (
	"fmt"
)

func main() {

 //Anonymous function to calculate the total cost of the trip

 calculateTotalCost := func(pricePerDay float64, days int) float64 {
  return pricePerDay * float64(days) //converting days to float64
 }

 //Anonymous function to check if the budget is sufficient for the trip
 isBudgetSufficient := func(totalCost, budget float64) bool {
  return budget >= totalCost
 }

 //User input for travelling scenario
 fmt.Println("You are planning a trip to the Mediterranean Sea. Let's plan your journey")

 var days int

 fmt.Println("How many days do you plan to stay?")

 fmt.Scanln(&days)

 var pricePerDay float64

 fmt.Println("What is the estimated daily cost (in CAD)?")

 fmt.Scanln(&pricePerDay)

 var budget float64

 fmt.Println("What is your budget for the trip (in CAD)?")

 fmt.Scanln(&budget)

//Creating a slice of places to visit
placesToVisit := []string{"Barcelona", "Athena", "Rome", "Ibiza", "cote d'azure"}

//Creating a map of activities and their costs
activitiesCost := map[string]float64{

 "Sightseeing":    50.0,
 "Fine Dining":    100.0,
 "Boat Tour":      150.0,
 "Shopping Spree": 200.0,
}

//Calculating the total cost of the trip
totalCost := calculateTotalCost(pricePerDay, days)

//Checking is the budget is sufficient
sufficientBudget := isBudgetSufficient(totalCost, budget)

if sufficientBudget {

  fmt.Printf("You are planning a %d-day trip to the Mediterranean Sea with a budget of $%.2f.\n", days, budget)
  fmt.Printf("The estimated total cost of the trip is $%.2f.\n", totalCost)
  fmt.Println("Places to visit: ")
  for _, place := range placesToVisit {
   fmt.Println("-", place)
  }
  fmt.Println("Activities and their cost")
  for activity, cost := range activitiesCost {
   fmt.Printf("- %s: $%.2f.\n", activity, cost)
  }
  fmt.Println("Enjoy your journey!")
 } else {
  fmt.Printf("Unfortunately, your budget of  $%.2f is not sufficient for the trip.\n", budget)
  fmt.Printf("The estimated total cost of the trip is  $%.2f.\n", totalCost)
  fmt.Println("Please adjust your budget for the duration of your trip")
 }

}*/

// Variadicle functions

/*package main

import "fmt"

func sum(numbers ...int) int {
 total := 0
 for _, num := range numbers {
  total += num
 }
 return total
}

func main() {
 result := sum(1, 2, 3, 4, 5)
 fmt.Println("Result:", result) // Output: Result: 15
 result = sum(10, 20, 30)
 fmt.Println("Result:", result) // Output: Result: 60
 result = sum()
 fmt.Println("Result:", result) // Output: Result: 0
}*/

/*package main

import "fmt"

//Variadic function to concatenate strings

func concatenate(strings ...string) string {

 var result string
 for _, str := range strings {

  result += str

 }
 return result

}

func main() {

 greeting := concatenate("Hello", ", ", "everyone", "!")
 fmt.Println("Greeting:", greeting)

 names := concatenate("James", " ", "Bond")
 fmt.Println("Full name: ", names)

 empty := concatenate()
 fmt.Println("Empty string: ", empty)
}*/

/*package main

import "fmt"

//Variadic function to combine slices of integers

func combine(slices ...[]int) []int {

 var result []int
 for _, slice := range slices {
  result = append(result, slice...)
 }

 return result

}

func main() {

 combined := combine([]int{1, 2, 3}, []int{4, 5}, []int{6, 7, 8})
 fmt.Println("Combined Slice: ", combined)

 combined = combine([]int{9, 10}, []int{11})
 fmt.Println("Combined Slice: ", combined)

 combined = combine()
 fmt.Println("Empty Combined Slice: ", combined)

}*/

/*package main

import (
 "fmt"
 "math"
)

//Variadic function to find the maximum of float64

func max(values ...float64) float64 {

 maxValue := math.Inf(-1)
 for _, value := range values {
  if value > maxValue {
   maxValue = value
  }
 }
 return maxValue
}

func main() {

 maxVal := max(1.2, 3.4, 5.6, 7.8)
 fmt.Println("Maximum Value: ", maxVal)

 maxVal = max(-10, -20, -30)
 fmt.Println("Maximum Value", maxVal)

}
*/

/*package main

import (
 "fmt"
 "strings"
)

//Variadic function to join elements with a separator

func join(separator string, elements ...string) string {

 return strings.Join(elements, separator)

}

func main() {

 sentence := join(" ", "This", "is", "a", "sentence.")
 fmt.Println("Sentence:", sentence)

 csv := join(", ", "apple", "banana", "cherry")
 fmt.Println("CSV: ", csv)

 single := join("-", "single")
 fmt.Println("Single Element: ", single)
}*/

/*package main

import "fmt"

//Variadic function to check if a value exists in a slice of integers

func exists(value int, slice ...int) bool {

 for _, v := range slice {
  if v == value {
   return true
  }
 }
 return false
}

func main() {
 found := exists(5, 1, 2, 3, 4, 5)
 fmt.Println("Found 5:", found)
 
}*/

package main

import "fmt"

//Variadic function to calculate the average of float64 numbers

func average(numbers ...float64) float64 {

 if len(numbers) == 0 {
  return 0
 }

 total := 0.0
 for _, num := range numbers {
  total += num
 }

 return total / float64(len(numbers))

}

func main() {

 avg := average(10.5, 20.75, 30.25)
 fmt.Println("Average: ", avg)

 avg = average(5.5, 6.5)
 fmt.Println("Average: ", avg)

 avg = average()
 fmt.Println("Average with no inputs: ", avg)

}
