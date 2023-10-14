package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func loopFibonacci(num int) int {
	ans := 0
	for i := 0; i <= num; i++ {
		if i <= 2 {
			ans++
		} else {
			ans += i
		}
	}

	return ans
}

func recursiveFibbonacci(num int) int {
	if num <= 2 {
		return 1
	}
	return recursiveFibbonacci(num-1) + recursiveFibbonacci(num-2)
}

func main() {
	fmt.Printf("previous GOMAXPROCS = %v\n", runtime.GOMAXPROCS(1))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Println("Golang server started...")
	err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		recursive := true
		if qp := query.Get("recursive"); qp == "" {
			recursive = false
		}

		number := 10
		if qp := query.Get("number"); qp != "" {
			numInt, err := strconv.Atoi(qp)
			if err == nil {
				number = numInt
			}
		}

		ans := 0
		if recursive {
			ans = recursiveFibbonacci(number)
		} else {
			ans = loopFibonacci(number)
		}

		byts, _ := json.Marshal(map[string]any{
			"number":    number,
			"recursive": recursive,
			"fibonacci": ans,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(byts)
	}))

	if err != nil {
		panic(err)
	}
}
