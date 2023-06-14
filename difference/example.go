package difference

import "fmt"

func Difference1() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}
