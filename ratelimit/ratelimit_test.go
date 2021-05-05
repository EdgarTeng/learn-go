package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestTake(t *testing.T) {
	t.Run("take", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer func() {
			cancel()
			t.Log("Done")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				Take()
				t.Log(time.Now())
			}

		}

	})
}

func TestAllow(t *testing.T) {
	t.Run("take", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer func() {
			cancel()
			t.Log("Done")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				t.Log(Allow(), time.Now())
			}

		}

	})
}
