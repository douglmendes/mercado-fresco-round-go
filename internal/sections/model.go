package sections

type Section struct {
	Id                 int     `json:"id,omitempty"`
	SectionNumber      int     `json:"section_number,omitempty"`
	CurrentTemperature float64 `json:"current_temperature,omitempty"`
	MinimumTemperature float64 `json:"minimum_temperature,omitempty"`
	CurrentCapacity    int     `json:"current_capacity,omitempty"`
	MinimumCapacity    int     `json:"minimum_capacity,omitempty"`
	MaximumCapacity    int     `json:"maximum_capacity,omitempty"`
	WarehouseId        int     `json:"warehouse_id,omitempty"`
	ProductTypeId      int     `json:"product_type_id,omitempty"`
}
