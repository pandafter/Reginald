package tests

import (
	"reginald_backend/internal/tasks"
	"testing"
)

func TestCreateTask(t *testing.T) {
	req := tasks.CreateTaskRequest{
		Title:     "Test Task",
		CreatedBy: "user1",
	}

	task, err := tasks.CreateTask(req)
	if err != nil {
		t.Fatal("expected no error")
	}

	if task.Status != "pending" {
		t.Errorf("expected status pending, got %s", task.Status)
	}
}

func TestInvalidStatusChange(t *testing.T) {
	err := tasks.ChangeTaskStatus("fake-id", "invalid", "test@user.com")
	if err == nil {
		t.Fatal("expected error for invalid status")
	}
}

func TestUpdateStatus(t *testing.T) {
	err := tasks.ChangeTaskStatus("fake-id", "done", "test@user.com")
	if err == nil {
		t.Fatal("expected error for invalid status")
	}
}
