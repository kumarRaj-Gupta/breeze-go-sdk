package breeze

type CustomerDetails struct {
	AppKey       string `json:"AppKey"`
	SessionToken string `json:"SessionToken"`
}
type SuccessResponse struct {
	Session_token string `json:"session_token"`
}
type CustomerDetailsResponse struct {
	Success SuccessResponse `json:"Success"`
	Status  int             `json:"Status"`
	Error   any             `json:"Error"`
}
