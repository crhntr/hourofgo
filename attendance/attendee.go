package attendance

type Attendee struct {
	// ID    SomeSortOfUniqueID

	Name     Name
	Pass     Passer
	Contacts []*Attendee
}

type Name struct {
	NameSections []string
}

func (nm Name) Valid() bool {
	return true
}
