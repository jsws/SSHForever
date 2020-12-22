package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 1600; i++ {
// 		wg.Add(1)
// 		time.Sleep(time.Millisecond * 10)
// 		go func() {
// 			conn, err := net.Dial("tcp", "127.0.0.1:8080")

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			for i := 0; i < 10; i++ {
// 				rec, err := bufio.NewReader(conn).ReadString('\n')
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				fmt.Println(rec)
// 			}

// 			fmt.Println("hiu")
// 		}()

// 	}
// 	wg.Wait()

// }
