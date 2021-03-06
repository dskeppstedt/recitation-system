package models

import "strings"

type Course struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	NumTracks int    `json:"tracks"`
}

func (this *Course) Validate() bool {
	return this.Name != "" && this.NumTracks > 0
}

type Recitation struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Track  string `json:"track"`
	Points int
}

type RecitationSub struct {
	CourseId   int       `json:"course_id"`
	Name       string    `json:"name"`
	NrProblems string    `json:"nr_problems"`
	Problems   []Problem `json:"problems"`
}
type Problem struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Com  string `json:"com"`
}

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (this *Student) Format() {
	this.Name = strings.ToLower(this.Name)
}

type Enrollment struct {
	Student int   `json:"student"`
	Courses []int `json:"courses"`
}

type DisplayProblem struct {
	Problem     string
	Compulsory  int
	Subproblems []Subproblem
}

type Subproblem struct {
	Letter string
}

type Solved struct {
	Problems       map[string][]string `json:"problems"`
	RecitationName string              `json:"recitation_name"`
	Course         int                 `json:"course_id"`
	Track          int                 `json:"track"`
}

type CloseRec struct {
	CourseId       int    `json:"course"`
	RecitationName string `json:"recitation"`
}
