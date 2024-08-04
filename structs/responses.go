package structs

import "encoding/json"

type Teacher struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ForeName  string `json:"foreName"`
	LongName  string `json:"longName"`
	ForeColor string `json:"foreColor,omitempty"`
	BackColor string `json:"backColor,omitempty"`
}

type Student struct {
	ID       int    `json:"id"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	ForeName string `json:"foreName"`
	LongName string `json:"longName"`
	Gender   string `json:"gender"`
}

type Class struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	ForeColor string `json:"foreColor,omitempty"`
	BackColor string `json:"backColor,omitempty"`
	Did       string `json:"did"`
	Teacher1  int    `json:"teacher1"`
	Teacher2  int    `json:"teacher2"`
}

type Subject struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	ForeColor string `json:"foreColor,omitempty"`
	BackColor string `json:"backColor,omitempty"`
}

type Room struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	ForeColor string `json:"foreColor,omitempty"`
	BackColor string `json:"backColor,omitempty"`
}

type Department struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	ForeColor string `json:"foreColor,omitempty"`
	BackColor string `json:"backColor,omitempty"`
}

type Holiday struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	StartDate int    `json:"startDate"`
	EndDate   int    `json:"endDate"`
}

type TimegridUnit struct {
	Day       int `json:"day"` //1 = sunday, 2 = monday, ..., 7 = saturday
	TimeUnits struct {
		StartTime int `json:"startTime"`
		EndTime   int `json:"endTime"`
	} `json:"timeUnits"`
}
type StatusData struct {
	Lstypes json.RawMessage `json:"lstypes"`
	Codes   json.RawMessage `json:"codes"`
}
type SchoolYear struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartDate int    `json:"startDate"`
	EndDate   int    `json:"endDate"`
}
type Period struct {
	Id               int       `json:"id"`
	Date             string    `json:"date"`
	StartTime        int       `json:"startTime"`
	EndTime          int       `json:"endTime"`
	ClassIDs         []Class   `json:"kl"`
	TeacherIDs       []Teacher `json:"te"`
	SubjectIDs       []Subject `json:"su"`
	RoomIDs          []Room    `json:"ro"`
	LessonType       string    `json:"lstype,omitempty"` //„ls“ (lesson) | „oh“ (office hour) | „sb“ (standby) | „bs“ (break supervision) | „ex“(examination)  omitted if lesson
	Code             string    `json:"code,omitempty"`   //„“ | „cancelled“ | „irregular“ omitted if empty
	LessonText       string    `json:"lstext,omitempty"`
	StatisticalFlags string    `json:"statflags,omitempty"`
	ActivityType     string    `json:"activityType,omitempty"`
}
