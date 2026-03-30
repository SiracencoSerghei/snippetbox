package examples

import (
	"context"
	"testing"
	"time"
)

func TestRunWorkers(t *testing.T) {
	ctx := context.Background()
	output := RunWorkers(ctx, 3, 10)

	if len(output) != 10 {
		t.Fatalf("expected 10 results, got %d", len(output))
	}

	for _, res := range output {
		t.Log(res)
	}
}

func TestRunWorkersWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	output := RunWorkers(ctx, 3, 10)

	if len(output) >= 10 {
		t.Fatalf("expected fewer than 10 results due to timeout, got %d", len(output))
	}

	t.Logf("got %d results before timeout", len(output))
}