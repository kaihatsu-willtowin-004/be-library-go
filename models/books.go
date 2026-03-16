package models 
type Books struct{
	ID int `json:"id"`
	TITLE string `json:"title"`
	TACGIA string `json:"tacgia"`
	CATEGORYID int `json:"category_id"`
	TOTAL_COPIES string `json:"total_copies"`
	AVAILABLE_COPIES string `json:"available_copies"`
}