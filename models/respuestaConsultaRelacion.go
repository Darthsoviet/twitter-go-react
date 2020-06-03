package models

//RespuestaConsultaRelacion contiene el true o false de comprobar
//la relacion de dos usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
