package untisApi

import (
	"encoding/json"
	"fmt"

	"github.com/Mr-Comand/goUntisAPI/structs"
)

func (c *Client) Test() error {
	rpcResp, err := c.CallRPC("getSchoolyears", struct{}{})
	data := string(rpcResp.Result)
	fmt.Println(data, err)
	return err
}

// Get list of teachers form Api
func (c *Client) GetTeachers() ([]structs.Teacher, error) {
	rpcResp, err := c.CallRPC("getTeachers", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var teachers []structs.Teacher
	err = json.Unmarshal(rpcResp.Result, &teachers)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

// Get a list of all students
func (c *Client) GetStudents() ([]structs.Student, error) {
	rpcResp, err := c.CallRPC("getStudents", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var students []structs.Student
	err = json.Unmarshal(rpcResp.Result, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

// Get all classes of the current schoolYear from the Api
func (c *Client) GetClasses(params ...interface{}) ([]structs.Class, error) {

	rpcResp, err := c.CallRPC("getKlassen", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var classes []structs.Class
	err = json.Unmarshal(rpcResp.Result, &classes)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

// Get all classes of the given schoolYear from the Api
func (c *Client) GetClassesOfSchoolYear(schoolYearID int) ([]structs.Class, error) {

	rpcResp, err := c.CallRPC("getKlassen", struct {
		SchoolYearID int `json:"schoolyearId"`
	}{SchoolYearID: schoolYearID})
	if err != nil {
		return nil, err
	}

	var classes []structs.Class
	err = json.Unmarshal(rpcResp.Result, &classes)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *Client) GetSubjects() ([]structs.Subject, error) {

	rpcResp, err := c.CallRPC("getSubjects", struct{}{})
	if err != nil {
		return nil, err
	}

	var subjects []structs.Subject
	err = json.Unmarshal(rpcResp.Result, &subjects)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
func (c *Client) GetRooms() ([]structs.Room, error) {

	rpcResp, err := c.CallRPC("getRooms", struct{}{})
	if err != nil {
		return nil, err
	}

	var rooms []structs.Room
	err = json.Unmarshal(rpcResp.Result, &rooms)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (c *Client) GetDepartments() ([]structs.Department, error) {

	rpcResp, err := c.CallRPC("getDepartments", struct{}{})
	if err != nil {
		return nil, err
	}

	var departments []structs.Department
	err = json.Unmarshal(rpcResp.Result, &departments)
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (c *Client) GetHolidays() ([]structs.Holiday, error) {

	rpcResp, err := c.CallRPC("getHolidays", struct{}{})
	if err != nil {
		return nil, err
	}

	var holidays []structs.Holiday
	err = json.Unmarshal(rpcResp.Result, &holidays)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}

// Get timegrid
func (c *Client) GetTimegridUnits() ([]structs.TimegridUnit, error) {

	rpcResp, err := c.CallRPC("getTimegridUnits", struct{}{})
	if err != nil {
		return nil, err
	}

	var timegridUnits []structs.TimegridUnit
	err = json.Unmarshal(rpcResp.Result, &timegridUnits)
	if err != nil {
		return nil, err
	}

	return timegridUnits, nil
}

// Information about lesson types and period codes and their colors
func (c *Client) GetStatusData() (structs.StatusData, error) {

	rpcResp, err := c.CallRPC("getStatusData", struct{}{})
	if err != nil {
		return structs.StatusData{}, err
	}

	var StatusData structs.StatusData
	err = json.Unmarshal(rpcResp.Result, &StatusData)
	if err != nil {
		return structs.StatusData{}, err
	}

	return StatusData, nil
}

// Data for the current schoolyear
func (c *Client) GetCurrentSchoolyear() (structs.SchoolYear, error) {

	rpcResp, err := c.CallRPC("getCurrentSchoolyear", struct{}{})
	if err != nil {
		return structs.SchoolYear{}, err
	}

	var currentSchoolYear structs.SchoolYear
	err = json.Unmarshal(rpcResp.Result, &currentSchoolYear)
	if err != nil {
		return structs.SchoolYear{}, err
	}

	return currentSchoolYear, nil
}

// List of all available schoolyears
func (c *Client) GetSchoolyears() ([]structs.SchoolYear, error) {

	rpcResp, err := c.CallRPC("getSchoolyears", struct{}{})
	if err != nil {
		return nil, err
	}

	var SchoolYears []structs.SchoolYear
	err = json.Unmarshal(rpcResp.Result, &SchoolYears)
	if err != nil {
		return nil, err
	}

	return SchoolYears, nil
}

// Get timetable for Classes, teacher, student, room, subject
func (c *Client) GetTimetableSimple(params structs.GetTimetableSimpleRequest) ([]structs.Period, error) {

	rpcResp, err := c.CallRPC("getTimetable", params)
	if err != nil {
		return nil, err
	}

	var periods []structs.Period
	err = json.Unmarshal(rpcResp.Result, &periods)
	if err != nil {
		return nil, err
	}

	return periods, nil
}

// Get a customizable timetable for classes, teacher, student, room, subject..
func (c *Client) GetTimetable(params structs.GetTimetableRequest) ([]structs.Period, error) {

	rpcResp, err := c.CallRPC("getTimetable", struct {
		Options structs.GetTimetableRequest `json:"options"`
	}{Options: params})
	if err != nil {
		return nil, err
	}

	var periods []structs.Period
	err = json.Unmarshal(rpcResp.Result, &periods)
	if err != nil {
		return nil, err
	}

	return periods, nil
}

// Import time of the last lesson/timetable or substitution import from Untis, returns Unix timestamp
func (c *Client) GetLatestImportTime() (int, error) {

	rpcResp, err := c.CallRPC("getLatestImportTime", struct{}{})
	if err != nil {
		return 0, err
	}

	var Timestamp int
	err = json.Unmarshal(rpcResp.Result, &Timestamp)
	if err != nil {
		return 0, err
	}

	return Timestamp, nil
}

// Get Id of the person (teacher or student) from the name, returns: Id of person or 0 if no fitting person found
func (c *Client) GetPersonId(params structs.GetPersonIdRequest) (int, error) {

	rpcResp, err := c.CallRPC("getPersonId", params)
	if err != nil {
		return 0, err
	}

	var PersonID int
	err = json.Unmarshal(rpcResp.Result, &PersonID)
	if err != nil {
		return 0, err
	}

	return PersonID, nil
}

// Request substitutions for the given date range
func (c *Client) GetSubstitutions(params structs.GetSubstitutionsRequest) ([]structs.Substitutions, error) {

	rpcResp, err := c.CallRPC("getSubstitutions", params)
	if err != nil {
		return nil, err
	}

	var substitution []structs.Substitutions
	err = json.Unmarshal(rpcResp.Result, &substitution)
	if err != nil {
		return nil, err
	}

	return substitution, nil
}

// Request classregevents for the given date range
func (c *Client) getClassregEvents(params structs.StartAndEndDate) ([]structs.ClassRegEvents, error) {

	rpcResp, err := c.CallRPC("getClassregEvents", params)
	if err != nil {
		return nil, err
	}

	var ClassregEvents []structs.ClassRegEvents
	err = json.Unmarshal(rpcResp.Result, &ClassregEvents)
	if err != nil {
		return nil, err
	}

	return ClassregEvents, nil
}

// Request classregevents for the given date range
func (c *Client) GetExams(params structs.GetExamsRequest) ([]structs.Exam, error) {

	rpcResp, err := c.CallRPC("getExams", params)
	if err != nil {
		return nil, err
	}

	var substitution []structs.Exam
	err = json.Unmarshal(rpcResp.Result, &substitution)
	if err != nil {
		return nil, err
	}

	return substitution, nil
}

// Request classregevents for the given date range
func (c *Client) GetExamTypes() (json.RawMessage, error) {

	rpcResp, err := c.CallRPC("getExams", struct{}{})
	if err != nil {
		return nil, err
	}

	var substitution json.RawMessage
	err = json.Unmarshal(rpcResp.Result, &substitution)
	if err != nil {
		return nil, err
	}

	return substitution, nil
}

// Retrieves the timetable for all students together with absene information for a given daterange.
func (c *Client) GetTimetableWithAbsences(params structs.StartAndEndDate) ([]structs.PeriodWithAbsenceObject, error) {

	rpcResp, err := c.CallRPC("getTimetableWithAbsences", struct {
		Objects structs.StartAndEndDate `json:"options"`
	}{})
	if err != nil {
		return nil, err
	}

	var PeriodWithAbsence []structs.PeriodWithAbsenceObject
	err = json.Unmarshal(rpcResp.Result, &PeriodWithAbsence)
	if err != nil {
		return nil, err
	}

	return PeriodWithAbsence, nil
}
