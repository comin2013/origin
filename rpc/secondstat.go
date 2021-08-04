package rpc

import (
	"container/list"
	"time"
	"fmt"
)

const ElemMaxCount = 1000

type stElem struct {
	Tm int64
	Cot int64
}

type SecondStat struct {
	st list.List
	maxCount int
}

func NewTimeST() *SecondStat{
	st := SecondStat{}
	st.init()
	st.maxCount = ElemMaxCount
	return &st
}

func (slf *SecondStat) SetMaxCot(cot int) {
	slf.maxCount = cot
}

func (slf *SecondStat) init(){
	slf.st.Init()
}

func (slf *SecondStat) Add(){
	t := time.Now().Unix()

	if slf.st.Len()==0 {
		slf.newE(t)
		return
	}

	e := slf.st.Back().Value.(*stElem)

	if e.Tm == t {
		e.Cot++
	} else {
		slf.newE(t)
		return
	}
}

func (slf *SecondStat) newE(tm int64) {
	slf.st.PushBack(&stElem{
		Tm:tm,
		Cot: 1,
	})

	if slf.st.Len() > slf.maxCount {
		slf.st.Remove(slf.st.Front())
	}
}

func (slf *SecondStat) Dump(cot int) string {
	retstr := ""
	c := 0
	for i:=slf.st.Back(); i!=nil; i=i.Prev(){
		retstr += fmt.Sprintf("%d:%d\t", i.Value.(*stElem).Tm, i.Value.(*stElem).Cot)
		c++
		if c==cot {
			break
		}
	}

	return retstr
}