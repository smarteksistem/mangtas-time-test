# mangtas-time-test

# How to run the code
go inside the root folder, then run the following commands
```
go mod tidy
go run main.go
```

# How to add more test case
To add more test case, add lines of code inside `main()` function. The values must be 4 digits, and integer type value
```
fmt.Println(possibleTimes([]int{A,B,C,D}))
```
with the A,B,C,D are integer values.

# How to run the test
go inside the root folder where `main.go` file exist. Then run the following commands
```
go mod tidy
go test main.go main_test.go -v
```

# Test coverage
* returns valid value where the valid possible 24hour times are > 0
* returns valid value where there's no valid possible 24hour times, return 0 value
* prove that the program won't return negative value, minimum returned value is 0
