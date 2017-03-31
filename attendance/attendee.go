package attendance

type Attendee struct {
	// ID    SomeSortOfUniqueID

	Name     []string
	Pass     Passer
	Contacts []*Attendee
}

type Name []string

func (nm Name) Valid() bool {
	return true
}
