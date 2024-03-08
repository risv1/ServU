package models

type Request struct {
	ID              string `json:"id"`
	Problem_Details string `json:"problem_details"`
	Since_When      string `json:"since_when"`
	Customer_Name   string `json:"c_name"`
	Customer_Email  string `json:"c_email"`
	Customer_Phone  string `json:"c_phone"`
	Severity        string `json:"severity"`
	Category        string `json:"category"`
}
