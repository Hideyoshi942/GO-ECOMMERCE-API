package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, _ := os.Create("users.csv")
	defer f.Close()

	w := bufio.NewWriterSize(f, 10*1024*1024) // buffer lớn
	now := time.Now().Unix()

	for i := 1; i <= 10_000_000; i++ {
		line := fmt.Sprintf(
			"user%d@gmail.com,090%07d,user_%d,$2a$10$fakehashedpassword,%d,%d,192.168.%d.%d,%d,10.0.%d.%d,%d,1\n",
			i,
			i,
			i,
			now,
			now,
			i%255,
			(i/255)%255,
			now,
			i%255,
			(i/255)%255,
			i%100,
		)
		w.WriteString(line)
	}
	w.Flush()
}
