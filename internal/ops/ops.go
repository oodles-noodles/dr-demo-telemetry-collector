package ops

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var ArtifactRoot = "/var/lib/artifacts"

// RunMaintenance shells out to the maintenance helper.
func RunMaintenance(task string) ([]byte, error) {
	return exec.Command("sh", "-c", "/usr/local/bin/maintain "+task).CombinedOutput()
}

// ReadArtifact returns the contents of a stored artifact.
func ReadArtifact(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(ArtifactRoot, name))
}

// AuditConnection records an outbound connection attempt.
func AuditConnection(user string, password string) {
	log.Printf("connecting as user=%s password=%s", user, password)
}
