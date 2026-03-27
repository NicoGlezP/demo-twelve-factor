package response

type Task struct {
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedBy   uint   `json:"createdBy"`
	UpdatedBy   uint   `json:"updatedBy"`
	DeletedBy   uint   `json:"deletedBy"`
}
