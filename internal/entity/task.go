package entity

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID  `db:"id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	Status      TaskStatus `db:"status"`
	Type        TaskType   `db:"type"`
	Description string     `db:"description"`
}

type TaskType string

const (
	TaskTypeMark      TaskType = "mark"
	TaskTypeVerify    TaskType = "verify"
	TaskTypeClassify  TaskType = "classify"
	TaskTypeCompare   TaskType = "compare"
	TaskTypeTakePhoto TaskType = "take_photo"
	TaskTypeTakeVideo TaskType = "take_video"
	TaskTypeTakeAudio TaskType = "take_audio"
	TaskTypeLocation  TaskType = "location"
	TaskTypeDeliver   TaskType = "deliver"
	TaskTypeOther     TaskType = "other"
)

// Writer
func (ts TaskType) Value() (driver.Value, error) {
	switch ts {
	case TaskTypeMark,
		TaskTypeVerify,
		TaskTypeClassify,
		TaskTypeCompare,
		TaskTypeTakePhoto,
		TaskTypeTakeVideo,
		TaskTypeTakeAudio,
		TaskTypeLocation,
		TaskTypeDeliver,
		TaskTypeOther:
		return string(ts), nil
	}
	return nil, fmt.Errorf("invalid task type: %s", ts)
}

// Reader
func (ts *TaskType) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("task type cannot be null")
	}

	str, ok := value.(string)
	if !ok {
		// Some drivers can return []byte
		if bytes, ok := value.([]byte); ok {
			str = string(bytes)
		} else {
			return fmt.Errorf("invalid type for task type")
		}
	}

	taskType := TaskType(str)
	switch taskType {
	case TaskTypeMark,
		TaskTypeVerify,
		TaskTypeClassify,
		TaskTypeCompare,
		TaskTypeTakePhoto,
		TaskTypeTakeVideo,
		TaskTypeTakeAudio,
		TaskTypeLocation,
		TaskTypeDeliver,
		TaskTypeOther:
		*ts = taskType
		return nil
	}
	return fmt.Errorf("invalid task type from database: %s", str)
}

type TaskStatus string

const (
	TaskStatusNew      TaskStatus = "new"
	TaskStatusPrepaid  TaskStatus = "prepaid"
	TaskStatusAssigned TaskStatus = "assigned"
	TaskStatusProcess  TaskStatus = "process"
	TaskStatusDone     TaskStatus = "done"
	TaskStatusConfirm  TaskStatus = "confirm"
	TaskStatusFinish   TaskStatus = "finish"
	TaskStatusCancel   TaskStatus = "cancel"
)

// Writer
func (ts TaskStatus) Value() (driver.Value, error) {
	switch ts {
	case TaskStatusNew,
		TaskStatusPrepaid,
		TaskStatusAssigned,
		TaskStatusProcess,
		TaskStatusDone,
		TaskStatusConfirm,
		TaskStatusFinish,
		TaskStatusCancel:
		return string(ts), nil
	}
	return nil, fmt.Errorf("invalid task status: %s", ts)
}

// Reader
func (ts *TaskStatus) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("task status cannot be null")
	}

	str, ok := value.(string)
	if !ok {
		// Some drivers can return []byte
		if bytes, ok := value.([]byte); ok {
			str = string(bytes)
		} else {
			return fmt.Errorf("invalid type for task status")
		}
	}

	taskStatus := TaskStatus(str)
	switch taskStatus {
	case TaskStatusNew,
		TaskStatusPrepaid,
		TaskStatusAssigned,
		TaskStatusProcess,
		TaskStatusDone,
		TaskStatusConfirm,
		TaskStatusFinish,
		TaskStatusCancel:
		*ts = taskStatus
		return nil
	}
	return fmt.Errorf("invalid task status from database: %s", str)
}
