package requests

type Base struct {
	SessionID    string `json:"sessionId"`
	Version      string `json:"version"`
	RevivalVerID int64  `json:"revivalVerId,string"`
	Seq          int64  `json:"seq,string"`
}
