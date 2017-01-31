package play

type RowMessage struct {
	dir int
}

func (RowMessage) Type() string { return "RowMessage" }

type ResetMessage struct {
}

func (ResetMessage) Type() string { return "ResetMessage" }
