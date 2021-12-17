package simian

import (
	//"regexp"
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
)

func Test_queue_init(t *testing.T) {
	q := make(EventHeap, 0)
	heap.Init(&q)
	//	t.Fatalf(`Hello(name) = %q, %v, want match for %#q, nil`, msg, err, want)

}

func Test_queue_push(t *testing.T) {
	events := map[string]float64{
		"event1": 10,
		"event2": 20,
		"event3": 15,
	}
	q := make(EventHeap, len(events))
	i := 0
	for value, timestamp := range events {
		q[i] = &Event{
			timestamp: timestamp,
			source:    value,
		}
		i++
	}
	heap.Init(&q)

	event := &Event{
		source:    "event4",
		timestamp: 5,
	}
	heap.Push(&q, event)
}

func Test_queue_pop(t *testing.T) {
	events := map[string]float64{
		"event1": 10,
		"event2": 20,
		"event3": 15,
	}
	q := make(EventHeap, len(events))
	i := 0
	for value, timestamp := range events {
		q[i] = &Event{
			source:    value,
			timestamp: timestamp,
		}
		i++
	}
	heap.Init(&q)

	event := &Event{
		source:    "event4",
		timestamp: 5,
	}
	heap.Push(&q, event)

	popped_event := heap.Pop(&q).(*Event)
	if event != popped_event {
		t.Fatalf(`heap.Pop(&q) = %f, wanted %v, error`, popped_event.timestamp, event.timestamp)
	}
}

func Test_queue_pop_simple(t *testing.T) {
	events := map[string]float64{
		"event1": 10,
		"event2": 20,
		"event3": 15,
	}
	q := make(EventHeap, len(events))
	i := 0
	for value, timestamp := range events {
		q[i] = &Event{
			source:    value,
			timestamp: timestamp,
		}
		i++
	}
	heap.Init(&q)
	event := &Event{
		source:    "event1",
		timestamp: 10,
	}

	popped_event := heap.Pop(&q).(*Event)
	if event.timestamp != popped_event.timestamp {
		t.Fatalf(`heap.Pop(&q) = %fp, wanted %f, error`, popped_event.timestamp, event.timestamp)
	}

	if event.source != popped_event.source {
		t.Fatalf(`heap.Pop(&q) = %q, wanted %v, error`, popped_event.source, event.source)
	}
}

func Test_queue_correctness_simple(t *testing.T) {
	events := map[string]float64{
		"event1":  10,
		"event2":  20,
		"event3":  15,
		"event4":  30,
		"event5":  25,
		"event6":  5,
		"event7":  1,
		"event8":  26,
		"event9":  12,
		"event10": 11,
		"event11": 23,
		"event12": 14,
		"event13": 18,
		"event14": 29,
		"event15": 3,
		"event16": 16,
		"event17": 7,
		"event18": 22,
	}
	sorted_events := [18]string{"event7", "event15", "event6", "event17", "event1", "event10", "event9", "event12", "event3", "event16", "event13", "event2", "event18", "event11", "event5", "event8", "event14", "event4"}
	q := make(EventHeap, len(events))
	i := 0
	for value, timestamp := range events {
		q[i] = &Event{
			source:    value,
			timestamp: timestamp,
		}
		i++
	}

	heap.Init(&q)

	for i := 0; i < len(events); i++ {
		event := sorted_events[i]
		popped_event := heap.Pop(&q).(*Event)
		if event != popped_event.source {
			t.Errorf(`hi`)

			t.Fatalf(`heap.Pop(&q) = %q, wanted %v, error`, popped_event.source, event)
		}

	}
}

func Test_queue_correctness(t *testing.T) {
	events := map[string]float64{
		"event1":  10,
		"event2":  20,
		"event3":  15,
		"event4":  30,
		"event5":  25,
		"event6":  5,
		"event7":  1,
		"event8":  26,
		"event9":  12,
		"event10": 11,
		"event11": 23,
		"event12": 14,
		"event13": 18,
		"event14": 29,
		"event15": 3,
		"event16": 16,
		"event17": 7,
		"event18": 22,
	}
	sorted_events := [18]string{"event7", "event15", "event6", "event17", "event1", "event10", "event9", "event12", "event3", "event16", "event13", "event2", "event18", "event11", "event5", "event8", "event14", "event4"}

	q := make(EventHeap, 0)
	heap.Init(&q)
	var event *Event

	i := 0
	for value, timestamp := range events {
		event = &Event{
			source:    value,
			timestamp: timestamp,
		}
		heap.Push(&q, event)
		i++
	}

	for i := 0; i < len(events); i++ {
		event := sorted_events[i]
		popped_event := heap.Pop(&q).(*Event)
		if event != popped_event.source {
			t.Fatalf(`heap.Pop(&q) = %q, wanted %v, error`, popped_event.source, event)
		}

	}
}

func Test_queue_rand_determinism(t *testing.T) {

	r := rand.New(rand.NewSource(1))
	expected_value := r.Float64()

	var generated_value float64
	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(1))
		generated_value = r.Float64()
		if expected_value != generated_value {
			t.Fatalf(`r.float64() expected: "%f", got "%f"`, expected_value, generated_value)
		}
	}
}

func Test_queue_determinism(t *testing.T) {

	var p int
	for p = 0; p < 10; p++ {
		r := rand.New(rand.NewSource(1))
		events := map[string]float64{
			"event1":  10 + r.Float64(),
			"event2":  20 + r.Float64(),
			"event3":  15 + r.Float64(),
			"event4":  30 + r.Float64(),
			"event5":  25 + r.Float64(),
			"event6":  5 + r.Float64(),
			"event7":  1 + r.Float64(),
			"event8":  26 + r.Float64(),
			"event9":  12 + r.Float64(),
			"event10": 11 + r.Float64(),
			"event11": 23 + r.Float64(),
			"event12": 14 + r.Float64(),
			"event13": 18 + r.Float64(),
			"event14": 29 + r.Float64(),
			"event15": 3 + r.Float64(),
			"event16": 16 + r.Float64(),
			"event17": 5 + r.Float64(),
			"event18": 22 + r.Float64(),
		}
		sorted_events := [18]string{"event7", "event15", "event17", "event6", "event1", "event10", "event9", "event12", "event3", "event16", "event13", "event2", "event18", "event11", "event5", "event8", "event14", "event4"}

		q := make(EventHeap, 0)
		heap.Init(&q)
		var event *Event

		i := 0
		for value, timestamp := range events {
			event = &Event{
				source:    value,
				timestamp: timestamp,
			}
			//fmt.Printf("timestamp: %f", event.timestamp)
			heap.Push(&q, event)
			i++
		}

		for i := 0; i < len(events); i++ {
			event := sorted_events[i]
			popped_event := heap.Pop(&q).(*Event)
			if event != popped_event.source {
				t.Fatalf(`heap.Pop(&q) = %q, wanted %v, error`, popped_event.source, event)
			}

		}
	}
}

func Benchmark_queue_set_column____________(b *testing.B) {

}
func Benchmark_queue_simple_push_pop(b *testing.B) {

	r := rand.New(rand.NewSource(1))
	events := map[string]float64{
		"event1":  10 + r.Float64(),
		"event2":  20 + r.Float64(),
		"event3":  15 + r.Float64(),
		"event4":  30 + r.Float64(),
		"event5":  25 + r.Float64(),
		"event6":  5 + r.Float64(),
		"event7":  1 + r.Float64(),
		"event8":  26 + r.Float64(),
		"event9":  12 + r.Float64(),
		"event10": 11 + r.Float64(),
		"event11": 23 + r.Float64(),
		"event12": 14 + r.Float64(),
		"event13": 18 + r.Float64(),
		"event14": 29 + r.Float64(),
		"event15": 3 + r.Float64(),
		"event16": 16 + r.Float64(),
		"event17": 5 + r.Float64(),
		"event18": 22 + r.Float64(),
	}
	//sorted_events := [18]string{"event7", "event15", "event17", "event6", "event1", "event10", "event9", "event12", "event3", "event16", "event13", "event2", "event18", "event11", "event5", "event8", "event14", "event4"}

	q := make(EventHeap, 0)
	heap.Init(&q)
	var event *Event
	var popped_event *Event

	i := 0
	for value, timestamp := range events {
		event = &Event{
			source:    value,
			timestamp: timestamp,
		}
		heap.Push(&q, event)
		i++
	}

	for j := 0; j < b.N; j++ {

		popped_event = heap.Pop(&q).(*Event)
		heap.Push(&q, popped_event)
	}

}

func Benchmark_queue_push_scale(b *testing.B) {

	scales := [6]int{100, 1000, 10000, 100000, 1000000, 10000000}
	for _, tc := range scales {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := rand.New(rand.NewSource(1))
				q := make(EventHeap, 0)
				heap.Init(&q)
				var event *Event
				for i := 0; i < tc; i++ {
					event = &Event{
						source:    "Event Placeholder",
						timestamp: float64(r.Intn(10000000)) + r.Float64(),
					}
					heap.Push(&q, event)
				}
			}
		})
	}
}
func Benchmark_queue_pop_scale(b *testing.B) {

	scales := [5]int{100, 1000, 10000, 100000, 1000000}
	for _, tc := range scales {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				r := rand.New(rand.NewSource(1))
				q := make(EventHeap, 0)
				heap.Init(&q)
				var event *Event
				for i := 0; i < tc; i++ {
					event = &Event{
						source:    "Event Placeholder",
						timestamp: float64(r.Intn(10000000)) + r.Float64(),
					}
					heap.Push(&q, event)
				}
				b.StartTimer()
				for p := 0; p < tc; p++ {
					heap.Pop(&q)
				}

			}

		})
	}
}
