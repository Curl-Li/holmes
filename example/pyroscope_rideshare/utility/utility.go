package utility

import (
	"os"
	"time"
)

const durationConstant = time.Duration(200 * time.Millisecond)

func mutexLock(n int64) {
	var i int64 = 0

	// start time is number of seconds since epoch
	startTime := time.Now()

	// This changes the amplitude of cpu bars
	for time.Since(startTime) < time.Duration(n*30)*durationConstant {
		i++
	}
}

func checkDriverAvailability(n int64) {
	var i int64 = 0

	// start time is number of seconds since epoch
	startTime := time.Now()

	for time.Since(startTime) < time.Duration(n)*durationConstant {
		i++
	}

	// Every other minute this will artificially create make requests.py in eu-north region slow
	// this is just for demonstration purposes to show how performance impacts show up in the
	// flamegraph
	force_mutex_lock := time.Now().Minute()%2 == 0
	if os.Getenv("REGION") == "eu-north" && force_mutex_lock {
		mutexLock(n)
	}

}

func FindNearestVehicle(searchRadius int64, vehicle string) {
	//pyroscope.TagWrapper(context.Background(), pyroscope.Labels("vehicle", vehicle), func(ctx context.Context) {
	var i int64 = 0

	startTime := time.Now()
	for time.Since(startTime) < time.Duration(searchRadius)*durationConstant {
		i++
	}

	if vehicle == "car" {
		checkDriverAvailability(searchRadius)
		go func() {
			go AllocMem()
		}()
	}
	if vehicle == "bike" {
		for i := 1; i < 10; i++ {
			go func() {
				time.Sleep(15 * time.Second)
			}()
		}
	}
	//})
}

func AllocMem() {
	var a = make([]byte, 1073741824)
	_ = a
	time.Sleep(10 * time.Second)
}
