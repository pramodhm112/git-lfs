package progress

import "github.com/git-lfs/git-lfs/tlog"

func Noop() Meter {
	return &nonMeter{
		updates: make(chan *tlog.Update),
	}
}

type nonMeter struct {
	updates chan *tlog.Update
}

func (m *nonMeter) Start()                                                               {}
func (m *nonMeter) Pause()                                                               {}
func (m *nonMeter) Add(size int64)                                                       {}
func (m *nonMeter) Skip(size int64)                                                      {}
func (m *nonMeter) StartTransfer(name string)                                            {}
func (m *nonMeter) TransferBytes(direction, name string, read, total int64, current int) {}
func (m *nonMeter) FinishTransfer(name string)                                           {}
func (m *nonMeter) Finish() {
	close(m.updates)
}

func (m *nonMeter) Updates() <-chan *tlog.Update { return m.updates }
func (m *nonMeter) Throttled() bool              { return false }
