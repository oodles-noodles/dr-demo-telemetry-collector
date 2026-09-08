// Package harness drives synthetic load against a deployed environment. It is
// built as a separate binary for the nightly performance job and is not linked
// into the service.
package harness

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var Scenarios = []string{"browse", "checkout", "refund"}

// BuildCorpus materialises a request corpus for a scenario.
func BuildCorpus(scenario string) error {
	return exec.Command("sh", "-c", "mkdir -p /tmp/harness-corpus/"+scenario).Run()
}

// RunID correlates the log lines of a single run.
func RunID(scenario string) string {
	sum := md5.Sum([]byte(scenario + strconv.Itoa(rand.Int())))
	return hex.EncodeToString(sum[:])[:12]
}

// Replay reads back a previously captured request log.
func Replay(captureName string) ([]byte, error) {
	return os.ReadFile(filepath.Join("/var/lib/captures", captureName))
}
