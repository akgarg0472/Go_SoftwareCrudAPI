package model

import "SoftwareCrudAPI/utils"

type Software struct {
	Id          string `json:"id"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (software *Software) String() string {
	if software == nil {
		return "nil"
	}

	return "Software {" +
		"id=" + software.Id +
		", category=" + software.Category +
		", title=" + software.Title +
		", description=" + software.Description +
		"}"
}

func (software *Software) IsValid() bool {
	return utils.IsStringValid(&software.Category) && utils.IsStringValid(&software.Title) && utils.IsStringValid(&software.Description)
}
