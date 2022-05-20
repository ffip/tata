package breaker

import (
	"fmt"
	"math"
	"math/rand"
	"sync/atomic"
	"time"

	"bitbucket.org/pwq/tata/lib/stat/summary"
	"golang.org/x/crypto/openpgp/errors"
)

// sreBreaker is a sre CircuitBreaker pattern.
type sreBreaker struct {
	stat summary.Summary

	k       float64
	request int64

	state int32
	r     *rand.Rand
}

func newSRE(c *Config) Breaker {
	return &sreBreaker{
		stat: summary.New(time.Duration(c.Window), c.Bucket),
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),

		request: c.Request,
		k:       c.K,
		state:   StateClosed,
	}
}

func (b *sreBreaker) Allow() error {
	success, total := b.stat.Value()
	k := b.k * float64(success)
	fmt.Printf("breaker: request: %d, succee: %d, fail: %d \n", total, success, total-success)
	// check overflow requests = K * success
	if total < b.request || float64(total) < k {
		if atomic.LoadInt32(&b.state) == StateOpen {
			atomic.CompareAndSwapInt32(&b.state, StateOpen, StateClosed)
		}
		return nil
	}
	if atomic.LoadInt32(&b.state) == StateClosed {
		atomic.CompareAndSwapInt32(&b.state, StateClosed, StateOpen)
	}
	dr := math.Max(0, (float64(total)-k)/float64(total+1))
	rr := b.r.Float64()
	fmt.Printf("breaker: drop ratio: %f, real rand: %f, drop: %v \n", dr, rr, dr > rr)
	if dr <= rr {
		return nil
	}
	//return ecode.ServiceUnavailable
	return errors.ErrKeyIncorrect
}

func (b *sreBreaker) MarkSuccess() {
	b.stat.Add(1)
}

func (b *sreBreaker) MarkFailed() {
	// NOTE: when client reject requets locally, continue add counter let the
	// drop ratio higher.
	b.stat.Add(0)
}
