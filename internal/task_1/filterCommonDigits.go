package task1


func FilterCommonDigits(firstNumber, secondNumber int) (int, int, error) {
  if firstNumber < 0 || secondNumber < 0 {
    return 0, 0, ErrNegNums
  }

  // slice of digits in numbers
  firstNumberDigits := splitDigits(firstNumber) 
  secondNumberDigits := splitDigits(secondNumber)

  commonDigits := findCommonDigits(firstNumberDigits, secondNumberDigits) // map witch containes the same digits in both numbers

  // filtered digits
  filteredFirstDigits := filterDigits(firstNumberDigits, commonDigits)
  filteredSecondDigits := filterDigits(secondNumberDigits, commonDigits)

  if len(filteredFirstDigits) == 0 || len(filteredSecondDigits) == 0 { // check zero-numbers
    return 0, 0, ErrEmptyNum
  }

  // build up a number
  resultedFirstNumber := buildNumber(filteredFirstDigits)
  resultedSecondNumber := buildNumber(filteredSecondDigits)
  
  return resultedFirstNumber, resultedSecondNumber, nil
}


func splitDigits(number int) ([]int) { // function which gets digits from the number and return the slice of digits
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


func findCommonDigits(firstSet, secondSet []int) map[int]bool { // function witch find matches in both numbers
  commonDigits := make(map[int]bool) // map with common digits and bool type to filter them

  for _, firstSetDigit := range firstSet {
    for _, secondSetDigit := range secondSet {
      if firstSetDigit == secondSetDigit {
        commonDigits[firstSetDigit] = true
        break
      }
    }
  }

  return commonDigits
}


func filterDigits(digitsOfNumber []int, filter map[int]bool) []int { // function which filters the digits of one number by the digits of another
  filteredDigits := []int{}

  for _, digit := range digitsOfNumber {
    if !filter[digit] {
      filteredDigits = append(filteredDigits, digit) // adding the digits that isn't included in the map
    }
  }

  return filteredDigits
}


func buildNumber(digitsNumber []int) int { // function witch build up a number from slice of digits
  buildNumber := 0
  
  for _, digit := range digitsNumber {
    buildNumber = buildNumber * 10 + digit
  }

  return buildNumber
}
