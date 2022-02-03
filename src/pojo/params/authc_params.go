package params

type AuthcParams struct {
	Principal   string `json:"principal" binding:"required"`
	Certificate string `json:"certificate" binding:"required"`
	GrantType   string `json:"grant_type" binding:"required"`
}
