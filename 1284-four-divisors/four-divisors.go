func sumFourDivisors(nums []int) int {
    maxNum := 0
    for _, num := range nums {
        maxNum = max(maxNum, num)
    }

    isPrime := getPrimeNumbers(maxNum)
    res := 0

    for _, num := range nums {
        numSqrt := int(math.Sqrt(float64(num)))

        for x := 2; x <= numSqrt; x++ {
            if !isPrime[x] {
                continue
            }

            if x * x == num {
                break
            }
            
            if num % x == 0 {
                quotient := num / x
                
                if isPrime[quotient] || quotient == x*x {
                    res += 1 + x + quotient + num
                    break
                }
            }
        }
    }

    return res
}

func getPrimeNumbers(upper int) []bool {
    isPrime := make([]bool, upper + 1)
    
    for i := 2; i <= upper; i++ {
        isPrime[i] = true
    }
    
    for num := 2; num <= int(math.Sqrt(float64(upper))); num++ {
        if !isPrime[num] {
            continue
        }
        
        for x := num * num; x <= upper; x += num {
            isPrime[x] = false
        }
    }
    
    return isPrime
}