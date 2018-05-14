package dream

import (
	"fmt"
	"os"
)

func Nonsense(n int) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < n; i++ {
		go func() {
			for {
				fmt.Fprintf(f, ".")
			}
		}()
	}
}
