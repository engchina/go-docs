package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go watch(ctx, "Monitor 1")
	go watch(ctx, "Monitor 2")

	fmt.Println("Now we'll wait for 8 seconds, time=", time.Now().Unix())
	time.Sleep(8 * time.Second)

	fmt.Println("After 8 seconds, prepare cancel function, found 2 routines have already complete, time=", time.Now().Unix())
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " received signal, monitor exits, time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "gorountine monitor, time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}

/*Now we'll wait for 8 seconds, time= 1672297004
Monitor 2 gorountine monitor, time= 1672297004
Monitor 1 gorountine monitor, time= 1672297004
Monitor 1 gorountine monitor, time= 1672297005
Monitor 2 gorountine monitor, time= 1672297005
Monitor 2 gorountine monitor, time= 1672297006
Monitor 1 gorountine monitor, time= 1672297006
Monitor 2  received signal, monitor exits, time= 1672297007
Monitor 1  received signal, monitor exits, time= 1672297007
After 8 seconds, prepare cancel function, found 2 routines have already complete, time= 1672297012
*/
