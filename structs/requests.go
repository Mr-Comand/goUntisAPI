package structs

type GetTimetableSimpleRequest struct {
	ElementId int    `json:"id"`
	Type      string `json:"type"`                // type of element, mandatory 	1 = klasse, 2 = teacher, 3 = subject, 4 = room, 5 = student
	StartDate int    `json:"startDate,omitempty"` // optional
	EndDate   int    `json:"endDate,omitempty"`   // optional
}
