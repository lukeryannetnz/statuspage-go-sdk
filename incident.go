package statuspage

type Incident struct {
	Name         *string  `json:"name"`
	GroupID      *string  `json:"group_id"`
	UpdateStatus *string  `json:"update_status"`
	Title        *string  `json:"suffix"`
	Body         *string  `json:"y_axis_min"`
	ComponentIDs []string `json:"component_ids"`
	ShouldTweet  *bool    `json:"should_tweet"`
}

type IncidentFull struct {
	Incident
	ID        *string `json:"id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func CreateIncident(client *Client, pageID string, incident *Incident) (*IncidentFull, error) {
	var i IncidentFull
	err := createResource(
		client,
		pageID,
		"incident",
		struct {
			Incident *Incident `json:"incident"`
		}{incident},
		&i,
	)

	return &i, err
}

func GetIncident(client *Client, pageID, incidentID string) (*IncidentFull, error) {
	var i IncidentFull
	err := readResource(client, pageID, incidentID, "incident", &i)

	return &i, err
}

func UpdateIncident(client *Client, pageID, incidentID string, incident *Incident) (*IncidentFull, error) {
	var i IncidentFull

	err := updateResource(
		client,
		pageID,
		"incident",
		incidentID,
		struct {
			Incident *Incident `json:"incident"`
		}{incident},
		&i,
	)

	return &i, err
}
