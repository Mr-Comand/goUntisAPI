package structs

type GetTimetableSimpleRequest struct {
	ElementId int `json:"id"`
	Type      int `json:"type"`                // type of element, mandatory 	1 = klasse, 2 = teacher, 3 = subject, 4 = room, 5 = student
	StartDate int `json:"startDate,omitempty"` // optional
	EndDate   int `json:"endDate,omitempty"`   // optional
}
type GetTimetableRequest struct {
	Element struct {
		id      int    `json:"id"`
		Type    int    `json:"Type"`
		keyType string `json:"keyType"`
	} `json:"element"`
	Type              int      `json:"type"`                        // type of element, mandatory 	1 = klasse, 2 = teacher, 3 = subject, 4 = room, 5 = student
	StartDate         int      `json:"startDate,omitempty"`         // number, format: YYYYMMDD, optional (default: current date)
	EndDate           int      `json:"endDate,omitempty"`           // number, format: YYYYMMDD, optional (default: current date)
	OnlyBaseTimetable bool     `json:"onlyBaseTimetable,omitempty"` //  boolean, returns only the base timetable (without bookings etc.)(default:false)
	ShowBooking       bool     `json:"showBooking,omitempty"`       //  returns the period's booking info if available (default: false)
	ShowInfo          bool     `json:"showInfo,omitempty"`          //  returns the period information if available (default: false)
	ShowSubstText     bool     `json:"showSubstText,omitempty"`     //  returns the Untis substitution text if available (default: false)
	ShowLsText        bool     `json:"showLsText,omitempty"`        //  returns the text of the period's lesson (default: false)
	ShowLsNumber      bool     `json:"showLsNumber,omitempty"`      //  returns the number of the period's lesson (default: false)
	ShowStudentgroup  bool     `json:"showStudentgroup,omitempty"`  //  returns the name(s) of the studentgroup(s) (default: false)
	KlasseFields      []string `json:"klasseFields,omitempty"`      //  optional, values: „id“, „name“, „longname“, „externalkey“ //TODO: to enum
	RoomFields        []string `json:"roomFields,omitempty"`        //  optional, values: „id“, „name“, „longname“, „externalkey“//TODO: to enum
	SubjectFields     []string `json:"subjectFields,omitempty"`     //  optional, values: „id“, „name“, „longname“, „externalkey“//TODO: to enum
	TeacherFields     []string `json:"teacherFields,omitempty"`     //  optional, values: „id“, „name“, „longname“, „externalkey“//TODO: to enum

}
type GetPersonIdRequest struct {
	Type      int    `json:"type"` // type of element, mandatory 2 = teacher, 5 = student
	Surname   string `json:"sn"`   //surname, mandatory
	Forename  string `json:"fn"`   //forename, mandatory
	Birthdata string `json:"dob"`  //birthdata, mandatory, use 0 if unknown

}
type GetSubstitutionsRequest struct {
	StartDate    int    `json:"startDate"`    // mandatory
	EndDate      string `json:"endDate"`      //mandatory
	DepartmentId string `json:"departmentId"` //mandatory, use 0 for all departments or if not applicable
}
type StartAndEndDate struct {
	StartDate int    `json:"startDate"` // mandatory
	EndDate   string `json:"endDate"`   //mandatory
}
type GetExamsRequest struct {
	ExamTypeId int    `json:"examTypeId"` // mandatory
	StartDate  int    `json:"startDate"`  // mandatory
	EndDate    string `json:"endDate"`    //mandatory
}
