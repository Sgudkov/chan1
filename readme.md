# chan1.go Description

chan1.go is a Go program that demonstrates the use of goroutines and channels for concurrent programming. The program consists of three main functions: [main](cci:1://f:/GO/2/async/chan1/chan1.go:8:0-21:1), [runchannels](cci:1://f:/GO/2/async/chan1/chan1.go:23:0-31:1), and [getNumber](cci:1://f:/GO/2/async/chan1/chan1.go:33:0-38:1).

## main Function

The [main](cci:1://f:/GO/2/async/chan1/chan1.go:8:0-21:1) function is the entry point of the program. It creates a channel `chan1` with a buffer size of 1 and starts two goroutines: [runchannels](cci:1://f:/GO/2/async/chan1/chan1.go:23:0-31:1) and [getNumber](cci:1://f:/GO/2/async/chan1/chan1.go:33:0-38:1). The [runchannels](cci:1://f:/GO/2/async/chan1/chan1.go:23:0-31:1) goroutine calculates the hash of numbers from 0 to 99 and writes the results to the channel, while the [getNumber](cci:1://f:/GO/2/async/chan1/chan1.go:33:0-38:1) goroutine reads from the channel and prints the results. The [main](cci:1://f:/GO/2/async/chan1/chan1.go:8:0-21:1) function then sleeps for 2 seconds to allow the goroutines to complete.

## runchannels Function

The [runchannels](cci:1://f:/GO/2/async/chan1/chan1.go:23:0-31:1) function calculates the hash of numbers from 0 to 99 using the [hash](cci:1://f:/GO/2/async/chan1/chan1.go:40:0-44:1) function and writes the results to the channel. It uses a for loop to iterate over the numbers and writes each result to the channel using the `<-` operator.

## getNumber Function

The [getNumber](cci:1://f:/GO/2/async/chan1/chan1.go:33:0-38:1) function reads from the channel and prints the results. It uses a for loop with the `range` keyword to iterate over the values received from the channel and prints each value using `fmt.Println`.

## hash Function

The [hash](cci:1://f:/GO/2/async/chan1/chan1.go:40:0-44:1) function calculates the hash of a given string using the FNV-1a hash algorithm. It creates a new hash object using `fnv.New32a()` and writes the string to the hash object using `h.Write([]byte(s))`. The hash value is then returned using `h.Sum32()`.

Overall, chan1.go demonstrates the use of goroutines and channels for concurrent programming in Go, as well as the use of the FNV-1a hash algorithm for hashing strings.