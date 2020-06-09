package entity

type ReviewerResp struct {
	Id        uint   `json:"id"`
	StaffId   uint64 `json:"staff_id"`
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}
