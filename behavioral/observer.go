type Observer interface {
    Update(msg string)
}

type Subject struct {
    observers []Observer
}

func (s *Subject) Notify(msg string) {
    for _, obs := range s.observers {
        obs.Update(msg)
    }
}
