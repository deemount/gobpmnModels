package core

import (
	"testing"
)

func TestDefinitionsRepository(t *testing.T) {
	// create a new repository
	repo := NewDefinitions()
	t.Logf("result of repo is %v", repo)
	// check if the repository is not nil
	if repo == nil {
		t.Errorf("expected a repository, got nil")
	}
}

func TestDefinitionsRepositoryID(t *testing.T) {
	// create a new repository
	repo := NewDefinitions()
	repo.SetID("definitions", "1234")
	repo.SetCollaboration()
	repo.SetProcess(0)
	t.Logf("result of repo is %v", repo)
	// check if the repository has the correct id
	if *repo.GetID() != "Definitions_1234" {
		t.Errorf("expected a correct id, got something else")
	}
}
