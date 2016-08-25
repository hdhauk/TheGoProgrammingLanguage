package memo_test

import (
	"testing"

	"github.com/hdhauk/TheGoProgrammingLanguage/ch9/memo"
	"github.com/hdhauk/TheGoProgrammingLanguage/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
