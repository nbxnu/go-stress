package manager

import "testing"

func TestChannel(t *testing.T){
	ch := make(chan int, 1000)
	go func(){
		for i := 0; i < 500; i ++{
			ch <- i*2
		}
		close(ch)
	}()
	for{
		select {
		case i := <- ch:
			t.Log(i)
		default:

		}
	}
}