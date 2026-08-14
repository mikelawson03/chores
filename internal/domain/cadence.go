package domain

type Cadence string

const (
	CadenceDaily   Cadence = "daily"
	CadenceWeekly  Cadence = "weekly"
	CadenceMonthly Cadence = "monthly"
)

func (c Cadence) IsValid() bool {
	switch c {
	case CadenceDaily, CadenceWeekly, CadenceMonthly:
		return true
	default:
		return false
	}
}
