package dto

type RoomDTO struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RoomResponseDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	HostId  string `json:"host_id"`
	Created string `json:"created"`
}
