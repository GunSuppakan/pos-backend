package domain

type FilePath struct {
	Model
	TypeFolder string `json:"type_folder"`
	FileUUID   string `json:"file_uuid"`
	FileName   string `json:"file_name"`
	Hash       string `json:"hash"`
}
