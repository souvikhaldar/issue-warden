package issue

type Issue struct {
	Id          int
	Title       string
	Description string
	AssignedTo  string
	CreatedBy   string
	Status      bool
}

func New(title string, description string, assignedto, createdby string, status bool) *Issue {
	return &Issue{
		Title:       title,
		Description: description,
		AssignedTo:  assignedto,
		CreatedBy:   createdby,
		Status:      status,
	}
}
