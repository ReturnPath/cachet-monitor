package cachet

import "encoding/json"

// Component Cachet model
type Component struct {
	ID            json.Number `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Status        json.Number `json:"status_id"`
	HumanStatus   string      `json:"-"`
	IncidentCount int         `json:"-"`
	CreatedAt     *string     `json:"created_at"`
	UpdatedAt     *string     `json:"updated_at"`
}

// ComponentData json response model
type ComponentData struct {
	Component Component `json:"data"`
}
