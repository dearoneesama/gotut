package main

import "fmt"

func MergeChannels(a, b chan int) chan int {
	out := make(chan int)
	go func () {
		for a != nil || b != nil {
			select {
			case v, ok := <- a:
				if !ok {
					a = nil
					fmt.Printf("%p is closed\n", &a)
					continue
				}
				out <- v
			case v, ok := <- b:
				if !ok {
					b = nil
					fmt.Printf("%p is closed\n", &b)
					continue
				}
				out <- v
			default:
			}
		}
		close(out)
	}()
	return out
}

func MergeManyChannels(cs []chan int) chan int {
	switch len(cs) {
	case 0:
		return nil
	case 1:
		return cs[0]
	case 2:
		return MergeChannels(cs[0], cs[1])
	default:
		c := MergeChannels(cs[0], cs[1])
		for i := 2; i < len(cs); i++ {
			c = MergeChannels(c, cs[i])
		}
		return c
	}
}

func MergeManyChannelsv(cs... chan int) chan int {
	return MergeManyChannels(cs)
}

func main() {
	chans := make([]chan int, 5)

	loc := func (ch chan int, mu int) {
		for i := 0; i < 5; i++ {
			ch <- i * mu
		}
		close(ch)
	}

	for i := 0; i < 5; i++ {
		chans[i] = make(chan int)
		go loc(chans[i], i + 1)
	}

	for v := range MergeManyChannels(chans) {
		fmt.Println(v)
	}
}
