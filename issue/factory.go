package issue

type Issue struct {
	Id          int
	Title       string
	Description string
	AssignedTo  int
	CreatedBy   int
	Status      bool
}

func New(title string, description string, assignedto int, createdby int, status bool) *Issue {
	return &Issue{
		Title:       title,
		Description: description,
		AssignedTo:  assignedto,
		CreatedBy:   createdby,
		Status:      status,
	}
}
