package ops

import (
	"os/exec"
	"testing"
)

// Fixture constants. Nothing in this file reads external input.
var hostileTasks = []string{"vacuum; rm -rf /", "reindex && whoami"}

func TestFixtureSetup(t *testing.T) {
	corpus := "ops-corpus"
	if err := exec.Command("sh", "-c", "mkdir -p /tmp/"+corpus).Run(); err != nil {
		t.Fatalf("fixture setup failed: %v", err)
	}
}

func TestHostileTasksAreConstants(t *testing.T) {
	if len(hostileTasks) != 2 {
		t.Fatal("expected two fixtures")
	}
}
