package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)
var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//? to run our server  type " go run serverCounter.go " into the terminal then head to  http://localhost:8081/increment .. every time we refresh the counter goes up.

//! let’s try incrementing a counter every time a specific URL is hit. Due to the fact that the web server is asynchronous, we’ll have to guard our counter using a mutex in order to prevent us from being hit with race-condition bugs.
//? A race condition occurs when two or more threads can access shared data and they try to change it at the same time. Because the thread scheduling algorithm can swap between threads at any time, you don't know the order in which the threads will attempt to access the shared data. Therefore, the result of the change in data is dependent on the thread scheduling algorithm, i.e. both threads are "racing" to access/change the data.
//* Problems often occur when one thread does a "check-then-act" (e.g. "check" if the value is X, then "act" to do something that depends on the value being X) and another thread does something to the value in between the "check" and the "act". E.g:

//*  if (x == 5) // The "Check"
//* {
//* 	y = x * 2; // The "Act"
//! If another thread changed x in between "if (x == 5)" and "y = x * 2" above,
//! y will not be equal to 10.
//* }

//? The point being, y could be 10, or it could be anything, depending on whether another thread changed x in between the check and act. You have no real way of knowing.
//? In order to prevent race conditions from occurring, you would typically put a lock around the shared data to ensure only one thread can access the data at a time. This would mean something like this:

// (Obtain lock for x)
// if (x == 5)
// {
//    y = x * 2; // Now, nothing can change x until the lock is released. Therefore y = 10
// }
// (release lock for x)

//* the other thread will have to wait for the lock to be released, this makes it important that the lock is released by the holding thread when it is finished with it. if it never releases it , the other thread will wait indefinitly