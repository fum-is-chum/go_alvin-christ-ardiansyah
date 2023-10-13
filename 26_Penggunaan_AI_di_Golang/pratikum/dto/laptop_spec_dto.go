package dto

type LaptopSpecRequest struct {
	Budget           int    `json:"budget"` // in Rupiah
	Brand            string `json:"brand"`
	Purpose          string `json:"purpose"`
	Ram              int    `json:"ram"`
	CpuCores         int    `json:"cpu_cores"`
	DisplayDimension int    `json:"display_dimension"` // in inch
}