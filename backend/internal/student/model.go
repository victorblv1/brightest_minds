package student

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

type Student struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Gender         Gender `json:"gender"`
	Nationality    string `json:"nationality"`
	Ethnicity      string `json:"ethnicity"`
	College        string `json:"college"`
	AcademicStatus string `json:"academic_status"` // e.g. bachelor, phd
	Area           string `json:"area"`            // e.g. physics
	StudyMajor     string `json:"study_major"`     // e.g. quantum mechanics
	Project        string `json:"project"`         // Game-changer plan
	PhotoURL       string `json:"photo_url,omitempty"`
}
