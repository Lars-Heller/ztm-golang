package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to get item %v from queue", i)
		}
		if item != i {
			t.Errorf("got wrong item: %v, want %v", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("Queue should be empty, got %v", item)
	}
}
