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
	Date             int       `json:"date"`
	StartTime        int       `json:"startTime"`
	EndTime          int       `json:"endTime"`
	Classes          []Class   `json:"kl"`
	Teachers         []Teacher `json:"te"`
	Subjects         []Subject `json:"su"`
	Rooms            []Room    `json:"ro"`
	LessonType       string    `json:"lstype,omitempty"`    //„ls“ (lesson) | „oh“ (office hour) | „sb“ (standby) | „bs“ (break supervision) | „ex“(examination)  omitted if lesson
	Code             string    `json:"code,omitempty"`      //„“ | „cancelled“ | „irregular“ omitted if empty
	Info             string    `json:"info"`                //Only in custom request
	SubstitutionText string    `json:"substText"`           //Only in custom request
	LessonText       string    `json:"lstext,omitempty"`    //omitempty only in non-custom request
	LessonNumber     int       `json:"lsnumber"`            //Only in custom request
	StatisticalFlags string    `json:"statflags,omitempty"` //omitempty only in non-custom request
	StudentGroup     string    `json:"sg"`
	BookingRemark    string    `json:"bkRemark"`
	BookingText      string    `json:"bkText"`
	ActivityType     string    `json:"activityType,omitempty"`
}
type IdObject struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	ExternalKey         string `json:"externalkey"`
	OriginalId          int    `json:"orgid"`
	OriginalName        string `json:"orgname"`
	OriginalExternalKey string `json:"orgexternalkey"`
}
type Substitutions struct {
	Type             string     `json:"type"` //"cancel" -> cancellation | subst -> teacher substitution | add -> additional period | shift -> shifted period | rmchg -> room change | rmlk -> locked period | bs -> break supervision | oh -> office hour | sb -> standby | other -> foreign substitutions | free -> free periods | ac -> activity | holi -> holiday | stxt -> substitution text
	LessonId         int        `json:"lsid"`
	Date             int        `json:"date"`      // Date conversion (current format: int YYYYMMDD)
	StartTime        int        `json:"startTime"` // Date / Time conversion (current format: int HHMM)
	EndTime          int        `json:"endTime"`   // Date / Time conversion (current format: int HHMM)
	Classes          []IdObject `json:"kl"`
	Teachers         []IdObject `json:"te"`
	Subjects         []IdObject `json:"su"`
	Rooms            []IdObject `json:"ro"`
	SubstitutionText string     `json:"txt"`
	Reschedule       struct {
		Date      int `json:"date"`      // Date conversion (current format: int YYYYMMDD)
		StartTime int `json:"startTime"` // Date / Time conversion (current format: int HHMM)
		EndTime   int `json:"endTime"`   // Date / Time conversion (current format: int HHMM)
	} `json:"reschedule"`
}
type ClassRegEvents struct {
	StudentId  int    `json:"studentid"`
	Surname    string `json:"surname"`
	Forname    string `json:"forname"`
	Date       int    `json:"date"` // Date conversion (current format: int YYYYMMDD)
	Subject    string `json:"subject"`
	Reason     string `json:"reason"`
	Text       string `json:"text"`
	CategoryId int    `json:"categoryId"`
}
type Exam struct {
	Id        int       `json:"id"`
	Classes   []Class   `json:"classes"`
	Teachers  []Teacher `json:"teachers"`
	Students  []Student `json:"students"`
	Subject   int       `json:"subject"`
	Date      int       `json:"date"`      // Date conversion (current format: int YYYYMMDD)
	StartTime int       `json:"startTime"` // Date / Time conversion (current format: int HHMM)
	EndTime   int       `json:"endTime"`   // Date / Time conversion (current format: int HHMM)
}

//TODO: to enum
