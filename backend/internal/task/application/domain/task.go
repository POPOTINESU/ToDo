package domain

import "github.com/google/uuid"

type Task struct {
	id uuid.UUID
	title Title
	description Description
	priority Priority
	stage Stage
}

func NewTask(
	id *uuid.UUID,
	title string,
	description string,
	priority string,
	stage string,
) (Task, error){
	var taskID uuid.UUID
	if id == nil {
		newID, err := uuid.NewV7()
		if err != nil {
			return Task{}, err
		}
		taskID = newID
	} else {
		taskID = *id
	}

	taskTitle, err := NewTitle(title)
	if err != nil {
		return Task{}, err
	}
	taskDescription, err := NewDescription(description)
	if err != nil {
		return Task{}, err
	}
	taskPriority, err := NewPriority(priority)
	if err != nil {
		return Task{}, err
	}
	taskStage, err := NewStage(stage)
	if err != nil {
		return Task{}, err
	}

	return Task{
		id: taskID,
		title: taskTitle,
		description: taskDescription,
		priority: taskPriority,
		stage: taskStage,
	}, nil
}

func (t *Task) ID() uuid.UUID {
	return t.id
}

func (t *Task) Title() Title {
	return t.title
}

func (t *Task) Description() Description {
	return t.description
}

func (t *Task) Priority() Priority {
	return t.priority
}

func (t *Task) Stage() Stage {
	return t.stage
}

func (t *Task) Equal(target *Task) bool {
	return t.id == target.id
}