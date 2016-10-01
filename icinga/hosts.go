package icinga

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HostsService handles communication with the host object related methods of the Icinga2 API
type HostsService service

type hostServiceResults struct {
	Results []hostResults `json:"results"`
}

type hostResults struct {
	Attrs Host `json:"attrs"`
}

// Host represents a Icinga2 Host object
type Host struct {
	Name                  string      `json:"__name,omitempty"`
	Acknowledgement       json.Number `json:"acknowledgement,omitempty"`
	AcknowledgementExpiry json.Number `json:"acknowledgement_expiry,omitempty"`
	ActionURL             string      `json:"action_url,omitempty"`
	Address               string      `json:"address,omitempty"`
	Address6              string      `json:"address6,omitempty"`
	CheckAttempt          json.Number `json:"check_attempt,omitempty"`
	CheckCommand          string      `json:"check_command,omitempty"`
	CheckInterval         json.Number `json:"check_interval,omitempty"`
	CheckPeriod           string      `json:"check_period,omitempty"`
	CheckTimeout          interface{} `json:"check_timeout,omitempty"`
	CommandEndpoint       string      `json:"command_endpoint,omitempty"`
	DisplayName           string      `json:"display_name,omitempty"`
	DowntimeDepth         json.Number `json:"downtime_depth,omitempty"`
	EventCommand          string      `json:"event_command,omitempty"`
	FlappingLastChange    float64     `json:"flapping_last_change,omitempty"`
	FlappingNegative      json.Number `json:"flapping_negative,omitempty"`
	FlappingPositive      json.Number `json:"flapping_positive,omitempty"`
	FlappingThreshold     json.Number `json:"flapping_threshold,omitempty"`
	Groups                []string    `json:"groups,omitempty"`
	HaMode                json.Number `json:"ha_mode,omitempty"`
	IconImage             string      `json:"icon_image,omitempty"`
	IconImageAlt          string      `json:"icon_image_alt,omitempty"`
	LastCheck             float64     `json:"last_check,omitempty"`
	LastHardState         json.Number `json:"last_hard_state,omitempty"`
	LastHardStateChange   float64     `json:"last_hard_state_change,omitempty"`
	LastState             json.Number `json:"last_state,omitempty"`
	LastStateChange       float64     `json:"last_state_change,omitempty"`
	LastStateDown         float64     `json:"last_state_down,omitempty"`
	LastStateType         json.Number `json:"last_state_type,omitempty"`
	LastStateUnreachable  json.Number `json:"last_state_unreachable,omitempty"`
	LastStateUp           float64     `json:"last_state_up,omitempty"`
	MaxCheckAttempts      json.Number `json:"max_check_attempts,omitempty"`
	NextCheck             float64     `json:"next_check,omitempty"`
	Notes                 string      `json:"notes,omitempty"`
	NotesURL              string      `json:"notes_url,omitempty"`
	OriginalAttributes    interface{} `json:"original_attributes,omitempty"`
	Package               string      `json:"package,omitempty"`
	RetryInterval         json.Number `json:"retry_interval,omitempty"`
	State                 json.Number `json:"state,omitempty"`
	StateType             json.Number `json:"state_type,omitempty"`
	Templates             []string    `json:"templates,omitempty"`
	Type                  string      `json:"type,omitempty"`
	Version               float64     `json:"version,omitempty"`
	Zone                  string      `json:"zone,omitempty"`
	Flapping              bool        `json:"flapping,omitempty"`
	Active                bool        `json:"active,omitempty"`
	EnableActiveChecks    bool        `json:"enable_active_checks,omitempty"`
	EnableEventHandler    bool        `json:"enable_event_handler,omitempty"`
	EnableFlapping        bool        `json:"enable_flapping,omitempty"`
	EnableNotifications   bool        `json:"enable_notifications,omitempty"`
	EnablePassiveChecks   bool        `json:"enable_passive_checks,omitempty"`
	EnablePerfdata        bool        `json:"enable_perfdata,omitempty"`
	ForceNextCheck        bool        `json:"force_next_check,omitempty"`
	ForceNextNotification bool        `json:"force_next_notification,omitempty"`
	LastReachable         bool        `json:"last_reachable,omitempty"`
	Paused                bool        `json:"paused,omitempty"`
	Volatile              bool        `json:"volatile,omitempty"`
	Vars                  *struct {
		Notification struct {
			Mail struct {
				Groups []string `json:"groups,omitempty"`
			} `json:"mail,omitempty"`
		} `json:"notification,omitempty"`
		Os string `json:"os,omitempty"`
	} `json:"vars,omitempty"`
	LastCheckResult *struct {
		Active          bool        `json:"active,omitempty"`
		CheckSource     string      `json:"check_source,omitempty"`
		Command         []string    `json:"command,omitempty"`
		ExecutionEnd    float64     `json:"execution_end,omitempty"`
		ExecutionStart  float64     `json:"execution_start,omitempty"`
		ExitStatus      json.Number `json:"exit_status,omitempty"`
		Output          string      `json:"output,omitempty"`
		PerformanceData []string    `json:"performance_data,omitempty"`
		ScheduleEnd     float64     `json:"schedule_end,omitempty"`
		ScheduleStart   float64     `json:"schedule_start,omitempty"`
		State           json.Number `json:"state,omitempty"`
		Type            string      `json:"type,omitempty"`
		VarsAfter       *struct {
			Attempt   json.Number `json:"attempt,omitempty"`
			Reachable bool        `json:"reachable,omitempty"`
			State     json.Number `json:"state,omitempty"`
			StateType json.Number `json:"state_type,omitempty"`
		} `json:"vars_after,omitempty"`
		VarsBefore *struct {
			Attempt   json.Number `json:"attempt,omitempty"`
			Reachable bool        `json:"reachable,omitempty"`
			State     json.Number `json:"state,omitempty"`
			StateType json.Number `json:"state_type,omitempty"`
		} `json:"vars_before,omitempty"`
	} `json:"last_check_result,omitempty"`
}

// Get a single Host.
func (s *HostsService) Get(name string) (*Host, *http.Response, error) {
	// TODO check name (not empty)

	u := fmt.Sprintf("objects/hosts/%s", name)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	results := new(hostServiceResults)
	resp, err := s.client.Do(req, results)
	if err != nil {
		return nil, resp, err
	}

	return &results.Results[0].Attrs, resp, err
}
