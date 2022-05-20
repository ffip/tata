package trace

import (
	"context"
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"bitbucket.org/pwq/tata/lib/conf/env"
	"bitbucket.org/pwq/tata/lib/net/metadata"
)

var _hostHash byte

func init() {
	rand.Seed(time.Now().UnixNano())
	_hostHash = byte(oneAtTimeHash(env.System.Hostname))
}

func extendTag() (tags []Tag) {
	tags = append(tags,
		TagString("hostname", env.System.Hostname),
		TagString("ip", env.System.IP),
		TagString("zone", env.System.Zone),
		TagString("region", env.System.Region),
	)
	return
}

func serviceNameFromEnv() string {
	return env.System.AppID
}

func isUATEnv() bool {
	return env.System.DeployEnv == env.DeployEnvUat
}

func genID() uint64 {
	var b [8]byte
	// i think this code will not survive to 2106-02-07
	binary.BigEndian.PutUint32(b[4:], uint32(time.Now().Unix())>>8)
	b[4] = _hostHash
	binary.BigEndian.PutUint32(b[:4], uint32(rand.Int31()))
	return binary.BigEndian.Uint64(b[:])
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type ctxKey string

var _ctxkey ctxKey = "go-common/net/trace.trace"

// FromContext returns the trace bound to the context, if any.
func FromContext(ctx context.Context) (t Trace, ok bool) {
	if v := metadata.Value(ctx, metadata.Trace); v != nil {
		t, ok = v.(Trace)
		return
	}
	t, ok = ctx.Value(_ctxkey).(Trace)
	return
}

// NewContext new a trace context.
// NOTE: This method is not thread safe.
func NewContext(ctx context.Context, t Trace) context.Context {
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Trace] = t
		return ctx
	}
	return context.WithValue(ctx, _ctxkey, t)
}
