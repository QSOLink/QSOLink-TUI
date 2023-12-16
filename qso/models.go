package qso

import "time"

type Contact struct {
	ID        int        `json:"ID"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
	DateOn    string     `json:"dateon"`
	TimeOn    time.Time  `json:"timeon"`
	Callsign  string     `json:"callsign"`
	Band      int        `json:"band"`
	Mode      string     `json:"mode"`
	City      string     `json:"city"`
	State     string     `json:"state"`
	County    string     `json:"county"`
	Country   string     `json:"country"`
	Name      string     `json:"name"`
	QslR      bool       `json:"qslr"`
	QslS      bool       `json:"qsls"`
	Rstr      int        `json:"rstr"`
	Rsts      int        `json:"rsts"`
	Power     int        `json:"power"`
	Remarks   string     `json:"remarks"`
}

type ResponseStruct struct {
	Contacts []Contact `json:"contacts"`
}
