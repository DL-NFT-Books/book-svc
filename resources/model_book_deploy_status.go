package resources

type DeployStatus int64

const (
	DeployPending DeployStatus = iota + 1
	DeploySuccessful
	DeployFailed
)

var taskStates = map[DeployStatus]string{
	DeployPending:    "pending",
	DeploySuccessful: "successful",
	DeployFailed:     "failed",
}

func (s DeployStatus) String() string {
	return taskStates[s]
}
