package entites

// Status - тип для хранения статуса
type Status int

const (
	Live Status = iota
	Finished
	ReserveNotMet
)

// Строковое представление статусов
func (s Status) String() string {
	switch s {
	case Live:
		return "Live"
	case Finished:
		return "Finished"
	case ReserveNotMet:
		return "ReserveNotMed"
	default:
		return "Unknown"
	}
}
