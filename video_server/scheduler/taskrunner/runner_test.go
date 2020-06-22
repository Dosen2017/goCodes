package taskrunner

import (
	"testing"
	"log"
	"time"
	// "errors"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0;i<30; i++ {
			dc <- i;
			log.Printf("Dispatcher sent:%v", i)
		}
		return nil
	}

	e := func(dc dataChan) error {

		//不建议使用 for range
		forloop:
			for {
				select {
				case d :=<- dc:
					log.Printf("Executor received:%v", d)
				default:
					break forloop
				}
			}

		return nil
		// return errors.New("Executor")   //如果用这句，那么执行一轮就结束了
	}

	runner := NewRunner(30, false, d, e) 
	go runner.StartAll()  //这里加goroute的主要目的是执行的代码块中有死循环
	time.Sleep(10 * time.Second)
}