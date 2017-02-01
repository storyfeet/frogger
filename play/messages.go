package play

type RowMessage struct {
	dir int
}

func (RowMessage) Type() string { return "RowMessage" }

type StopMessage struct {
}

func (StopMessage) Type() string { return "StopMessage" }

type ResetMessage struct {
	Score bool
}

func (ResetMessage) Type() string { return "ResetMessage" }
