package main

import (
	"fmt"
	"time"
)

type job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 결과 %d\n", j.index, j.index*j.index)
	wg.Done()
}
