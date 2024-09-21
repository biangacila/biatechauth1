package valueobjects

type Status string

const (
	StatusPending    Status = "Pending"
	StatusInProgress Status = "In Progress"
	StatusCompleted  Status = "Completed"
	StatusFailed     Status = "Failed"
)
