package scheduler

type IScheduler interface {
	Name() string

	Set(event *Event)
	Has(eventName string)
	Get(eventName string) *Event

	Start()
	Stop()

	StartEvent(name string)
	StopEvent(name string)
}
