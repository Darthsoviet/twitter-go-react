package models

//RespuestaLogin estruc del token que se devuelve
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
