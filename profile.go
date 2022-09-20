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
	jsonProfile, err := json.Marshal(map[string]interface{}{
		"profile": profile,
	})

	if err != nil {
		return err
	}

	path := "users.profile.set"
	if user != "" {
		path += "?user=" + user
	}

	response := &userResponseFull{}
	if err = api.jsonMethod(ctx, path, jsonProfile, response); err != nil {
		return err
	}

	return response.Err()
}
