package requests

type Base struct {
	SessionID    string `json:"sessionId"`
	Version      string `json:"version"`
	RevivalVerID int64  `json:"revivalVerId,omitempty"`
	Seq          int64  `json:"seq,string"`
}
