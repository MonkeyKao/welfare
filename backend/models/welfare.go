package models

type Welfare struct {
	Id       			int   		`json:"id"`
	Category 			int    		`json:"category"`
	City     			int    		`json:"city"`
	Date     			string 		`json:"date"`
	Priority 			int 		`json:"priority"`
	Title	 			string		`json:"title"`
	Detail	 			string		`json:"detail"`
	DetailCondition 	string		`json:"detailCodition"`
	DetailDocument		string		`json:"detailDocument"`
	DetailLink			string		`json:"detailLink"`
}

