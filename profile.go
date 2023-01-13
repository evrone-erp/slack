package slack

import (
	"context"
	"encoding/json"
)

// SetUserProfile will set a profile for the provided user
func (api *Client) SetUserProfile(user string, profile *UserProfile) error {
	return api.SetUserProfileContext(context.Background(), user, profile)
}

// SetUserProfileContext will set a profile for the provided user with a custom context
func (api *Client) SetUserProfileContext(ctx context.Context, user string, profile *UserProfile) error {
	data := map[string]interface{}{
		"profile": profile,
	}

	if user != "" {
		data["user"] = user
	}

	jsonProfile, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response := &userResponseFull{}
	if err = api.jsonMethod(ctx, "users.profile.set", jsonProfile, response); err != nil {
		return err
	}

	return response.Err()
}
