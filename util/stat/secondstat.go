package stat

import (
	"container/list"
	"time"
	"fmt"
	"sync"
)

const ElemMaxCount = 1000

type stElem struct {
	Tm int64
	Cot int64
}

type SecondStat struct {
	st list.List
	maxCount int
	maxElem *stElem	//
	lastDumpTm int

	lock sync.Mutex
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
	slf.lock.Lock()
	defer slf.lock.Unlock()

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
	// 更新最大值
	if slf.st.Len() > 0 {
		if slf.maxElem==nil || slf.maxElem.Cot < slf.st.Back().Value.(*stElem).Cot {
			slf.maxElem = slf.st.Back().Value.(*stElem)
		}

		// 中间时间补0
		lastem := slf.st.Back().Value.(*stElem)
		for i:=lastem.Tm+1;i<tm;i++{
			slf.st.PushBack(&stElem{Tm: i, Cot: 0})
		}
	}

	slf.st.PushBack(&stElem{
		Tm:tm,
		Cot: 1,
	})

	// 维护容器大小
	if slf.st.Len() > slf.maxCount {
		slf.st.Remove(slf.st.Front())
	}
}

func (slf *SecondStat) Dump() string {
	retstr := ""

	if slf.maxElem != nil{
		retstr += fmt.Sprintf(" Max: %d[%d]\n", slf.maxElem.Tm, slf.maxElem.Cot)
	}

	t := time.Now().Unix()

	c := 0
	cot := int(t)-slf.lastDumpTm
	slf.lastDumpTm=int(t)

	var total int64 = 0
	for i:=slf.st.Back(); i!=nil; i=i.Prev(){
		retstr += fmt.Sprintf("%d:%d\t", i.Value.(*stElem).Tm, i.Value.(*stElem).Cot)
		c++
		total += i.Value.(*stElem).Cot
		if c==cot {
			break
		}
	}
	if c>0 {
		retstr += fmt.Sprintf("average:%d/s", total/int64(c))
	}

	return retstr
}