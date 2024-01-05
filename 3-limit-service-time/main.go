//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"fmt"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := make(chan struct{})
	now := time.NewTicker(time.Second)
	defer now.Stop()

	go func() {
		for range now.C {
			u.TimeUsed++
			if hasExceededLimit(u) {
				done <- struct{}{}
				return
			}
		}
	}()

	go func() {
		process()
		done <- struct{}{}
	}()

	<-done

	return hasExceededLimit(u)
}

func hasExceededLimit(u *User) bool {
	return !u.IsPremium && u.TimeUsed > 10
}

func main() {
	now := time.Now()
	RunMockServer()
	fmt.Println("Total time:", time.Since(now))
}
