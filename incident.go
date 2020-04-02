package statuspage

type Incident struct {
	Name                 *string  `json:"name"`
	Status               *string  `json:"status"`
	Body                 *string  `json:"y_axis_min"`
	ComponentIDs         []string `json:"component_ids"`
	DeliverNotifications *bool    `json:"deliver_notifications"`
}

/*
Full request contract for an incident, from https://developer.statuspage.io/#operation/postPagesPageIdIncidents
{
  "incident": {
    "name": "string",
    "status": "postmortem",
    "impact_override": "maintenance",
    "scheduled_for": "2020-04-02T21:41:26Z",
    "scheduled_until": "2020-04-02T21:41:26Z",
    "scheduled_remind_prior": true,
    "scheduled_auto_in_progress": true,
    "scheduled_auto_completed": true,
    "metadata": {},
    "deliver_notifications": true,
    "auto_transition_deliver_notifications_at_end": true,
    "auto_transition_deliver_notifications_at_start": true,
    "auto_transition_to_maintenance_state": true,
    "auto_transition_to_operational_state": true,
    "auto_tweet_at_beginning": true,
    "auto_tweet_on_completion": true,
    "auto_tweet_on_creation": true,
    "auto_tweet_one_hour_before": true,
    "backfill_date": "string",
    "backfilled": true,
    "body": "string",
    "components": {
      "component_id": "operational"
    },
    "component_ids": [
      "wmvkjlvhw3cf"
    ],
    "scheduled_auto_transition": true
  }
}
*/

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
