package comrequest

type CreateDepartment struct {
	Name     string `json:"name"`
	NameEn   string `json:"name_en,omitempty"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order,omitempty"`
	ID       int    `json:"id,omitempty"`
}
