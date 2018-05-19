package ctypes

// Auth0UserInfoResponse is the basic user info response.
type Auth0UserInfoResponse struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
