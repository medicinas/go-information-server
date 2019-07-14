package models

type MedicalSpecialty struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ServedBy string `json:"servedBy"`
	Quote    Quote  `json:"quote"`
}

type Quote struct {
	Text     string `json:"quote"`
	ServedBy string `json:"ipAddress"`
	Language string `json:"language"`
}

func (a *MedicalSpecialty) ToString() string {
	return a.Name
}
