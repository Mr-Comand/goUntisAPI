package untisApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Mr-Comand/goUntisAPI/structs"
)

func (c *Client) Test() error {
	_, err := c.CallRPC("getSchoolyears", struct{}{})
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
	var periodsTemp []interface{}

	err = json.Unmarshal(rpcResp.Result, &periodsTemp)
	if err != nil {
		return nil, err
	}

	periods := make([]structs.Period, len(periodsTemp))

	// Loop over periodsTemp and convert each element into a Period struct
	// Loop over periodsTemp and convert each element into a Period struct
	for i, item := range periodsTemp {
		// Marshal the item to JSON
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON into the Period struct
		var period structs.Period
		err = json.Unmarshal(itemJSON, &period)
		if err != nil {
			return nil, err
		}

		// Check if original fields exist and populate them
		var tempMap map[string]interface{}
		err = json.Unmarshal(itemJSON, &tempMap)
		if err != nil {
			return nil, err
		}
		// Populate OriginalClasses, OriginalTeachers, etc., if they exist in the response
		if orgClasses, ok := tempMap["orgkl"]; ok {
			orgClassesJSON, _ := json.Marshal(orgClasses)
			var classes []structs.Class
			_ = json.Unmarshal(orgClassesJSON, &classes)
			period.OriginalClasses = classes
		}

		if orgTeachers, ok := tempMap["orgte"]; ok {
			orgTeachersJSON, _ := json.Marshal(orgTeachers)
			var teachers []structs.Teacher
			_ = json.Unmarshal(orgTeachersJSON, &teachers)
			period.OriginalTeachers = teachers
		}

		if orgSubjects, ok := tempMap["orgsu"]; ok {
			orgSubjectsJSON, _ := json.Marshal(orgSubjects)
			var subjects []structs.Subject
			_ = json.Unmarshal(orgSubjectsJSON, &subjects)
			period.OriginalSubjects = subjects
		}

		if orgRooms, ok := tempMap["orgro"]; ok {
			orgRoomsJSON, _ := json.Marshal(orgRooms)
			var rooms []structs.Room
			_ = json.Unmarshal(orgRoomsJSON, &rooms)
			period.OriginalRooms = rooms
		}

		// Add the populated Period struct to the periods slice
		periods[i] = period
	}

	return periods, nil
}

// Ensure the structs implement the OriginalData interface
func extractOriginalData[T structs.OriginalData](rawPeriod []interface{}, responseElements []T, jsonKey string) ([]T, error) {
	var originalDataList []T

	for i := range responseElements {
		var tempMap map[string]interface{}

		// Marshal the element to JSON to extract the needed data
		elementJSON, err := json.Marshal(rawPeriod[i])
		if err != nil {
			return nil, errors.New("failed to marshal element to JSON")
		}

		// Unmarshal JSON into a temporary map
		if err = json.Unmarshal(elementJSON, &tempMap); err != nil {
			return nil, errors.New("failed to unmarshal JSON to tempMap")
		}

		// Extract the specified jsonKey (e.g., "kl" for classes)
		if data, exists := tempMap[jsonKey]; exists {
			// Marshal the extracted data to JSON
			dataJSON, err := json.Marshal(data)
			if err != nil {
				return nil, errors.New("failed to marshal data to JSON")
			}

			// Parse the JSON array into a slice of interfaces
			var dataArray []interface{}
			if err = json.Unmarshal(dataJSON, &dataArray); err != nil {
				return nil, errors.New("failed to unmarshal dataJSON to dataArray")
			}
			fmt.Println("data:", dataArray)

			// Iterate through each item in dataArray
			for _, item := range dataArray {
				var itemMap map[string]interface{}
				itemJSON, err := json.Marshal(item)
				if err != nil {
					return nil, errors.New("failed to marshal item to JSON")
				}

				if err = json.Unmarshal(itemJSON, &itemMap); err != nil {
					return nil, errors.New("failed to unmarshal itemJSON to itemMap")
				}

				// Create a temporary map for org fields
				orgFields := make(map[string]interface{})

				// Populate orgFields for the original item
				for key, value := range itemMap {
					fmt.Println(key, itemMap, item)
					if strings.HasPrefix(key, "org") {
						orgField := strings.TrimPrefix(key, "org") // Remove "org" prefix
						orgFields[orgField] = value
					}
				}

				// Marshal orgFields back to JSON to unmarshal into originalItem
				if len(orgFields) > 0 {
					// Initialize a new instance of T
					var originalItem T
					orgFieldsJSON, err := json.Marshal(orgFields)
					if err != nil {
						return nil, errors.New("failed to marshal orgFields to JSON")
					}
					if err := json.Unmarshal(orgFieldsJSON, &originalItem); err != nil {
						return nil, errors.New("failed to unmarshal orgFieldsJSON to originalItem")
					}
					originalDataList = append(originalDataList, originalItem)
				}
			}
		} else {
			fmt.Println(tempMap)
			return nil, errors.New("specified jsonKey not found in tempMap" + jsonKey)
		}
	}

	return originalDataList, nil
}

// Get a customizable timetable for classes, teacher, student, room, subject..
func (c *Client) GetTimetable(params structs.GetTimetableRequest) ([]structs.Period, error) {

	rpcResp, err := c.CallRPC("getTimetable", struct {
		Options structs.GetTimetableRequest `json:"options"`
	}{Options: params})
	if err != nil {
		return nil, err
	}

	var periodsTemp []interface{}

	err = json.Unmarshal(rpcResp.Result, &periodsTemp)
	if err != nil {
		return nil, err
	}

	periods := make([]structs.Period, len(periodsTemp))
	fmt.Println(periodsTemp)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// Loop over periodsTemp and convert each element into a Period struct
	// Loop over periodsTemp and convert each element into a Period struct
	for i, item := range periodsTemp {
		// Marshal the item to JSON
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON into the Period struct
		var period structs.Period
		err = json.Unmarshal(itemJSON, &period)
		if err != nil {
			return nil, err
		}

		// Check if original fields exist and populate them
		var tempMap map[string]interface{}
		err = json.Unmarshal(itemJSON, &tempMap)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(rpcResp.Result, &periodsTemp)
		if err != nil {
			return nil, err
		}
		// Populate OriginalClasses, OriginalTeachers, etc., if they exist in the response
		// period.OriginalClasses, err = extractOriginalData(periodsTemp, period.Classes, "kl")
		// if err != nil {
		// 	return nil, err
		// }
		period.OriginalRooms, err = extractOriginalData(periodsTemp, period.Rooms, "ro")
		if err != nil {
			return nil, err
		}
		// period.OriginalTeachers, err = extractOriginalData(periodsTemp, period.Teachers, "te")
		// if err != nil {
		// 	return nil, err
		// }
		// period.OriginalSubjects, err = extractOriginalData(periodsTemp, period.Subjects, "su")
		// if err != nil {
		// 	return nil, err
		// }

		// Add the populated Period struct to the periods slice
		periods[i] = period
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
