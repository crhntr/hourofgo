package attendance

type Attendee struct {
	// ID    SomeSortOfUniqueID

	Names []Name
	Passer
}

type Name string

func (nm Name) Valid() bool {
	return true
}
