package main

import "errors"

func generateBinaryNumbersMiddleware(next func(int) []string) func(int) ([]string, error) {
	return func(n int) ([]string, error) {
		if n <= 0 {
			return nil, errors.New("number must be bigger then zero")
		}

		return next(n), nil
	}
}

func (q *Queue) generateBinaryNumbers(n int) []string {
	result := make([]string, n)
	q.enqueue("1")
	for i := 0; i < n; i++ {
		result[i] = q.dequeue()
		n1, n2 := result[i]+"0", result[i]+"1"
		q.enqueue(n1)
		q.enqueue(n2)
	}
	return result
}
