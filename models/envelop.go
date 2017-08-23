package models

import "time"

//Envelop is the main data structure of app logic.
//users pack data in this envelop to comunicate
type Envelop struct {
	datetime time.Time
	Tag      string `json:"tag"`
	Data     string `json:"data"`
	URLParam string `json:"urlparam"`
}
