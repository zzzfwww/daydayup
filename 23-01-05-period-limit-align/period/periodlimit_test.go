package period

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"period-limit/redistest"
	"testing"
)

func TestPeriodLimit_Take(t *testing.T) {
	testPeriodLimit(t)
}

func TestPeriodLimit_TakeWithAlign(t *testing.T) {
	testPeriodLimit(t, Align())
}

func TestPeriodLimit_RedisUnavailable(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	const (
		seconds = 1
		quota   = 5
	)
	l := NewPeriodLimit(seconds, quota, redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	}), "periodlimit")
	s.Close()
	val, err := l.Take("first")
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)
}

func testPeriodLimit(t *testing.T, opts ...PeriodOption) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	const (
		seconds = 1
		total   = 100
		quota   = 5
	)
	l := NewPeriodLimit(seconds, quota, store, "periodlimit", opts...)
	var allowed, hitQuota, overQuota int
	for i := 0; i < total; i++ {
		val, err := l.Take("first")
		if err != nil {
			t.Error(err)
		}
		switch val {
		case Allowed:
			allowed++
		case HitQuota:
			hitQuota++
		case OverQuota:
			overQuota++
		default:
			t.Error("unknown status")
		}
	}

	assert.Equal(t, quota-1, allowed)
	assert.Equal(t, 1, hitQuota)
	assert.Equal(t, total-quota, overQuota)
}

func TestQuotaFull(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	l := NewPeriodLimit(1, 1, redis.NewClient(&redis.Options{Addr: s.Addr()}), "periodlimit")
	val, err := l.Take("first")
	assert.Nil(t, err)
	assert.Equal(t, HitQuota, val)
}
