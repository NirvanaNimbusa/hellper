package model

import (
	"database/sql"
	"time"
)

const (
	StatusOpen     = "open"
	StatusCancel   = "canceled"
	StatusResolved = "resolved"
	StatusClosed   = "closed"
)

type Incident struct {
	Id                      int64         `db:"id,omitempty"`
	Title                   string        `db:"title,omitempty"`
	StartTimestamp          *time.Time    `db:"start_ts,omitempty"`
	EndTimestamp            *time.Time    `db:"end_ts,omitempty"`
	IdentificationTimestamp *time.Time    `db:"identification_ts,omitempty"`
	Responsibility          string        `db:"responsibility,omitempty"`
	Team                    string        `db:"team,omitempty"`
	Functionality           string        `db:"functionality,omitempty"`
	RootCause               string        `db:"root_cause,omitempty"`
	CustomerImpact          sql.NullInt64 `db:"customer_impact,omitempty"`
	StatusPageUrl           string        `db:"status_page_url,omitempty"`
	PostMortemUrl           string        `db:"post_mortem_url,omitempty"`
	Status                  string        `db:"status,omitempty"`
	Product                 string        `db:"product,omitempty"`
	SeverityLevel           int64         `db:"severity_level,omitempty"`
	ChannelName             string        `db:"channel_name,omitempty"`
	UpdatedAt               *time.Time    `db:"updated_at,omitempty"`
	SnoozedUntil            sql.NullTime  `db:"snoozed_until,omitempty"`
	DescriptionStarted      string        `db:"description_started,omitempty"`
	DescriptionCancelled    string        `db:"description_cancelled,omitempty"`
	DescriptionResolved     string        `db:"description_resolved,omitempty"`
	ChannelId               string        `db:"channel_id,omitempty"`
	IncidentAuthor          string        `db:"incident_author_id,omitempty"`
	CommanderId             string        `db:"commander_id,omitempty"`
	CommanderEmail          string        `db:"commander_email,omitempty"`
}
