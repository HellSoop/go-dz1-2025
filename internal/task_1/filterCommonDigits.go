package task1

func SplitDigits(number int) ([]int) { // function which getы digits from the number and return the slice of digits
  if (number == 0) { // checking for zero number
    return []int{0}
  }

  digits := []int{} // slice with digits

  for number > 0 {
    digit := number % 10 // take the last digits from number
    digits = append([]int{digit}, digits ...) // adding this digits in the begining of slice
    number = number / 10 // shorting the number
  }

  return digits
}

func FindCommonDigits(firstSet, secondSet []int) map[int]bool { // function witch find matches in both numbers
  commonDigits := make(map[int]bool) // map with common digits and bool type to filter them

  for _, digitsFirstSet := range firstSet {
    for _, digitsSecondSet := range secondSet {
      if digitsFirstSet == digitsSecondSet {
        commonDigits[digitsFirstSet] = true
        break
      }
    }
  }

  return commonDigits
}

func FilterDigits(digitsOfNumber []int, filter map[int]bool) []int { // function which filters the digits of one number by the digits of another
  filteredDigits := []int{}

  for _, digit := range digitsOfNumber {
    if !filter[digit] {
      filteredDigits = append(filteredDigits, digit) // adding the digits that isn't included in the map
    }
  }

  return filteredDigits
}

func BuildNumber (digitsNumber []int) int { // function witch build up a number from slice of digits
  buildNumber := 0
  
  for _, digit := range digitsNumber {
    buildNumber = buildNumber * 10 + digit
  }

  return buildNumber
}

func FilterCommonDigits(firstNumber, secondNumber int) (int, int, error) {
  if firstNumber < 0 || secondNumber < 0 {
    return 0, 0, ErrNegNums
  }

  // slice of digits in numbers
  firstNumberDigits := SplitDigits(firstNumber) 
  secondNumberDigits := SplitDigits(secondNumber)

  commonDigits := FindCommonDigits(firstNumberDigits, secondNumberDigits) // map witch containes the same digits in both numbers

  if len(commonDigits) == 0 { // chech for empty map
    return firstNumber, secondNumber, nil
  }

  // filtered digits
  filteredFirstDigits := FilterDigits(firstNumberDigits, commonDigits)
  filteredSecondDigits := FilterDigits(secondNumberDigits, commonDigits)

  if len(filteredFirstDigits) == 0 || len(filteredSecondDigits) == 0 { // check zero-numbers
    return 0, 0, ErrEmptyNum
  }

  // build up a number
  resultedFirstNumber := BuildNumber(filteredFirstDigits)
  resultedSecondNumber := BuildNumber(filteredSecondDigits)
  
  return resultedFirstNumber, resultedSecondNumber, nil
}
