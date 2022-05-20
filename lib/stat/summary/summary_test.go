package summary

import (
	"testing"
	"time"
)

func TestSummaryMinInterval(t *testing.T) {
	count := New(time.Second/2, 10)
	tk1 := time.NewTicker(5 * time.Millisecond)
	defer tk1.Stop()
	for i := 0; i < 100; i++ {
		<-tk1.C
		count.Add(2)
	}

	v, c := count.Value()
	t.Logf("count value: %d, %d\n", v, c)
	// 10% of error when bucket is 10
	if v < 190 || v > 210 {
		t.Errorf("expect value in [90-110] get %d", v)
	}
	// 10% of error when bucket is 10
	if c < 90 || c > 110 {
		t.Errorf("expect value in [90-110] get %d", v)
	}
}
