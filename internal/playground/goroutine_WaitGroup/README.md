Using WaitGroup to delay execution of main function until all goroutines are
complete: [Article](https://www.golangprograms.com/advance-programs/how-to-use-waitgroup-to-delay-execution-of-the-main-function-until-after-all-goroutines-are-complete.html)

WaitGroup is a counting semaphore that can be used to maintain a record of 
running goroutines. When the value of a WaitGroup is greater than zero, the 
Wait method will block.
