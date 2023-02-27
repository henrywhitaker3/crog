package circuits

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestItRetiresAGivenNumberOfTimes(t *testing.T) {
	tries := 0
	retry := Retry(func(ctx context.Context) (any, error) {
		tries++
		return nil, errors.New("fail")
	}, 3, 1*time.Microsecond)
	retry(context.Background())
	if tries != 3 {
		t.Fail()
	}
}

func TestItErrorsWhenContextTimesOut(t *testing.T) {
	retry := Retry(func(ctx context.Context) (any, error) {
		time.Sleep(time.Millisecond * 5)
		return nil, errors.New("error")
	}, 3, 1*time.Microsecond)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1)
	defer cancel()

	if _, err := retry(ctx); !errors.Is(context.DeadlineExceeded, err) {
		t.Fail()
	}
}
