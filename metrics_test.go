package metrics

import (
	"testing"
	"time"

	"github.com/Shyp/goshyp/test"
	"github.com/rcrowley/go-metrics"
)

func TestNamespace(t *testing.T) {
	originalNamespace := Namespace
	originalRealm := Realm
	defer func() {
		Namespace = originalNamespace
		Realm = originalRealm
	}()
	Namespace = "foo"
	Realm = "test"
	if getWithNamespace("bar") != "test.foo.bar" {
		t.Errorf("expected getWithNamespace(bar) to be test.foo.bar, was %s", getWithNamespace("bar"))
	}

	Namespace = ""
	test.AssertEquals(t, getWithNamespace("bar"), "test.bar")
}

func TestIncrementIncrements(t *testing.T) {
	Increment("bar")
	Increment("bar")
	Increment("bar")
	mn := metrics.GetOrRegisterMeter("bar", nil)
	if mn.Count() != 3 {
		t.Errorf("expected Count() to be 3, was %d", mn.Count())
	}
}

func ExampleIncrement() {
	Start("web")
	Increment("dequeue.success")
}

func ExampleMeasure() {
	Start("web")
	Measure("workers.active", 6)
}

func ExampleTime() {
	Start("web")
	start := time.Now()
	time.Sleep(3)
	Time("auth.latency", time.Since(start))
}
