package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDefinitionsRepository tests the creation of a new repository
func TestDefinitionsRepository(t *testing.T) {
	t.Run("NewDefinitions()",
		func(t *testing.T) {
			// create a new repository
			repo := NewDefinitions()
			t.Logf("result of repo is %v", repo)
			// check if the repository is not nil
			if repo == nil {
				t.Errorf("expected a repository, got nil")
			}
		})
	t.Run("SetID()",
		func(t *testing.T) {
			// create a new repository
			repo := NewDefinitions()
			repo.SetID("definitions", "1234")
			assert.Equal(t, "Definitions_1234", *repo.GetID(), "want %v, got %v", "Definitions_1234", *repo.GetID())
		})
}
