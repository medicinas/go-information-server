package models

type MedicalSpecialty struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *MedicalSpecialty) ToString() string {
	return a.Name
}
