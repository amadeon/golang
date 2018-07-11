package meter

import (
	"sync"
	"time"
)

type Meter struct {
	Interval     int64 //calculate the average interval every time
	Cnt          int64 //number of measured data
	StartTime    int64 //interval fact start
	ComputedRate int64 //speed
	mu           sync.Mutex
}

const NANO = 1e9

func New(interval int64) Meter {
	meter := Meter{}
	meter.StartTime = time.Now().UnixNano()
	meter.Interval = interval * NANO
	return meter
}

func (m *Meter) SetInterval(interval int64) {
	m.Interval = interval
}

func (m *Meter) Reset(interval int64) {
	m.mu.Lock()
	m.Interval = interval * NANO
	m.Cnt = 0
	m.ComputedRate = 0
	m.StartTime = time.Now().UnixNano()
	m.mu.Unlock()
}

func (m *Meter) Count(increment int64) int64 {
	m.mu.Lock()
	curTime := time.Now().UnixNano()
	curInt := curTime - m.StartTime

	m.Cnt += increment
	if curInt >= m.Interval {
		oldRate := m.ComputedRate
		m.ComputedRate = (NANO * m.Cnt + NANO / 2) / curInt
		if oldRate != 0{
			m.ComputedRate = (oldRate + m.ComputedRate + 1) / 2
		}
		m.StartTime = curTime
		m.Cnt = 0
	}

	m.mu.Unlock()
	return m.ComputedRate
}

func (m *Meter) Get() int64 {
	return m.ComputedRate
}
