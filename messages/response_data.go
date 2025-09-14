package messages

type ResponseData struct {
	Ticket              string           `json:"ticket"`
	Cap                 map[string](any) `json:"cap"`
	Username            string           `json:"username"`
	CSRFPreventionToken string           `json:"CSRFPreventionToken"`
}
