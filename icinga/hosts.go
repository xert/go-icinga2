package icinga

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HostsService service

type hostServiceResults struct {
	Results []hostResults `json:"results"`
}

type hostResults struct {
	Attrs Host `json:"attrs"`
}

type Host struct {
	Name                  string      `json:"__name"`
	Acknowledgement       json.Number `json:"acknowledgement"`
	AcknowledgementExpiry json.Number `json:"acknowledgement_expiry"`
	ActionURL             string      `json:"action_url"`
	Active                bool        `json:"active"`
	Address               string      `json:"address"`
	Address6              string      `json:"address6"`
	CheckAttempt          json.Number `json:"check_attempt"`
	CheckCommand          string      `json:"check_command"`
	CheckInterval         json.Number `json:"check_interval"`
	CheckPeriod           string      `json:"check_period"`
	CheckTimeout          interface{} `json:"check_timeout"`
	CommandEndpoint       string      `json:"command_endpoint"`
	DisplayName           string      `json:"display_name"`
	DowntimeDepth         json.Number `json:"downtime_depth"`
	EnableActiveChecks    bool        `json:"enable_active_checks"`
	EnableEventHandler    bool        `json:"enable_event_handler"`
	EnableFlapping        bool        `json:"enable_flapping"`
	EnableNotifications   bool        `json:"enable_notifications"`
	EnablePassiveChecks   bool        `json:"enable_passive_checks"`
	EnablePerfdata        bool        `json:"enable_perfdata"`
	EventCommand          string      `json:"event_command"`
	Flapping              bool        `json:"flapping"`
	FlappingLastChange    float64     `json:"flapping_last_change"`
	FlappingNegative      json.Number `json:"flapping_negative"`
	FlappingPositive      json.Number `json:"flapping_positive"`
	FlappingThreshold     json.Number `json:"flapping_threshold"`
	ForceNextCheck        bool        `json:"force_next_check"`
	ForceNextNotification bool        `json:"force_next_notification"`
	Groups                []string    `json:"groups"`
	HaMode                json.Number `json:"ha_mode"`
	IconImage             string      `json:"icon_image"`
	IconImageAlt          string      `json:"icon_image_alt"`
	LastCheck             float64     `json:"last_check"`
	LastCheckResult       struct {
		Active          bool        `json:"active"`
		CheckSource     string      `json:"check_source"`
		Command         []string    `json:"command"`
		ExecutionEnd    float64     `json:"execution_end"`
		ExecutionStart  float64     `json:"execution_start"`
		ExitStatus      json.Number `json:"exit_status"`
		Output          string      `json:"output"`
		PerformanceData []string    `json:"performance_data"`
		ScheduleEnd     float64     `json:"schedule_end"`
		ScheduleStart   float64     `json:"schedule_start"`
		State           json.Number `json:"state"`
		Type            string      `json:"type"`
		VarsAfter       struct {
			Attempt   json.Number `json:"attempt"`
			Reachable bool        `json:"reachable"`
			State     json.Number `json:"state"`
			StateType json.Number `json:"state_type"`
		} `json:"vars_after"`
		VarsBefore struct {
			Attempt   json.Number `json:"attempt"`
			Reachable bool        `json:"reachable"`
			State     json.Number `json:"state"`
			StateType json.Number `json:"state_type"`
		} `json:"vars_before"`
	} `json:"last_check_result"`
	LastHardState        json.Number `json:"last_hard_state"`
	LastHardStateChange  float64     `json:"last_hard_state_change"`
	LastReachable        bool        `json:"last_reachable"`
	LastState            json.Number `json:"last_state"`
	LastStateChange      float64     `json:"last_state_change"`
	LastStateDown        float64     `json:"last_state_down"`
	LastStateType        json.Number `json:"last_state_type"`
	LastStateUnreachable json.Number `json:"last_state_unreachable"`
	LastStateUp          float64     `json:"last_state_up"`
	MaxCheckAttempts     json.Number `json:"max_check_attempts"`
	NextCheck            float64     `json:"next_check"`
	Notes                string      `json:"notes"`
	NotesURL             string      `json:"notes_url"`
	OriginalAttributes   interface{} `json:"original_attributes"`
	Package              string      `json:"package"`
	Paused               bool        `json:"paused"`
	RetryInterval        json.Number `json:"retry_interval"`
	State                json.Number `json:"state"`
	StateType            json.Number `json:"state_type"`
	Templates            []string    `json:"templates"`
	Type                 string      `json:"type"`
	Vars                 struct {
		Notification struct {
			Mail struct {
				Groups []string `json:"groups"`
			} `json:"mail"`
		} `json:"notification"`
		Os string `json:"os"`
	} `json:"vars"`
	Version  float64 `json:"version"`
	Volatile bool    `json:"volatile"`
	Zone     string  `json:"zone"`
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
