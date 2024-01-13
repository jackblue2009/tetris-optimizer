printf "Bad Example 0:\n"
go run cmd/main.go "error-input/bad-example00.txt"
sleep 2

printf "Bad Example 1:\n"
go run cmd/main.go "error-input/bad-example01.txt"
sleep 2

printf "Bad Example 2:\n"
go run cmd/main.go "error-input/bad-example02.txt"
sleep 2

printf "Bad Example 3:\n"
go run cmd/main.go "error-input/bad-example03.txt"
sleep 2

printf "Bad Example 4:\n"
go run cmd/main.go "error-input/bad-example04.txt"
sleep 2

printf "Bad Format:\n"
go run cmd/main.go "error-input/bad-format.txt"
sleep 2

printf "Good Example 0:\n"
go run cmd/main.go "input/good-example00.txt"
sleep 2

printf "Good Example 1:\n"
go run cmd/main.go "input/good-example01.txt"
sleep 2

printf "Good Example 2:\n"
go run cmd/main.go "input/good-example02.txt"
sleep 2

printf "Good Example 3:\n"
go run cmd/main.go "input/good-example03.txt"
sleep 2

printf "Hard Example:\n"
go run cmd/main.go "input/hard-example.txt"