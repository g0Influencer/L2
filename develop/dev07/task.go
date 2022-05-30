package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)


// 1 способ
func or (channels ...<- chan interface{}) <- chan interface{} {
	out:=make(chan interface{})
	defer close(out)

	wg:=&sync.WaitGroup{}
	wg.Add(len(channels))

	for _,c:= range channels{
		go func(c <-chan interface{}) {
			for val:=range c {
				out <- val
			}
			wg.Done()
		}(c)
	}

		wg.Wait()

	return out
}

// 2 способ с помощью отражения

func mergeReflect(channels ...<-chan interface{}) <-chan interface{}{
	out:=make(chan interface{})
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out

}


func main(){
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))

}
