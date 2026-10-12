
package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Game struct {
	ID                 int64           `json:"id"`
	UserID             int64           `json:"user_id"`
	Name               string          `json:"name"`
	SettingDescription string          `json:"setting_description"`
	Ruleset            json.RawMessage `json:"ruleset"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type Entity struct {
	ID        int64           `json:"id"`
	GameID    int64           `json:"game_id"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Data      json.RawMessage `json:"data"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type Event struct {
	ID        int64           `json:"id"`
	GameID    int64           `json:"game_id"`
	Type      string          `json:"type"`
	ActorID   *int64          `json:"actor_id,omitempty"`
	TargetID  *int64          `json:"target_id,omitempty"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}
