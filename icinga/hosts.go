package icinga

import (
	"fmt"
	"net/http"
)

type HostsService service

type Host struct {
	Name                  string      `json:"__name"`
	Acknowledgement       int         `json:"acknowledgement"`
	AcknowledgementExpiry int         `json:"acknowledgement_expiry"`
	ActionURL             string      `json:"action_url"`
	Active                bool        `json:"active"`
	Address               string      `json:"address"`
	Address6              string      `json:"address6"`
	CheckAttempt          int         `json:"check_attempt"`
	CheckCommand          string      `json:"check_command"`
	CheckInterval         int         `json:"check_interval"`
	CheckPeriod           string      `json:"check_period"`
	CheckTimeout          interface{} `json:"check_timeout"`
	CommandEndpoint       string      `json:"command_endpoint"`
	DisplayName           string      `json:"display_name"`
	DowntimeDepth         int         `json:"downtime_depth"`
	EnableActiveChecks    bool        `json:"enable_active_checks"`
	EnableEventHandler    bool        `json:"enable_event_handler"`
	EnableFlapping        bool        `json:"enable_flapping"`
	EnableNotifications   bool        `json:"enable_notifications"`
	EnablePassiveChecks   bool        `json:"enable_passive_checks"`
	EnablePerfdata        bool        `json:"enable_perfdata"`
	EventCommand          string      `json:"event_command"`
	Flapping              bool        `json:"flapping"`
	FlappingLastChange    float64     `json:"flapping_last_change"`
	FlappingNegative      int         `json:"flapping_negative"`
	FlappingPositive      int         `json:"flapping_positive"`
	FlappingThreshold     int         `json:"flapping_threshold"`
	ForceNextCheck        bool        `json:"force_next_check"`
	ForceNextNotification bool        `json:"force_next_notification"`
	Groups                []string    `json:"groups"`
	HaMode                int         `json:"ha_mode"`
	IconImage             string      `json:"icon_image"`
	IconImageAlt          string      `json:"icon_image_alt"`
	LastCheck             float64     `json:"last_check"`
	LastCheckResult       struct {
		Active          bool     `json:"active"`
		CheckSource     string   `json:"check_source"`
		Command         []string `json:"command"`
		ExecutionEnd    float64  `json:"execution_end"`
		ExecutionStart  float64  `json:"execution_start"`
		ExitStatus      int      `json:"exit_status"`
		Output          string   `json:"output"`
		PerformanceData []string `json:"performance_data"`
		ScheduleEnd     float64  `json:"schedule_end"`
		ScheduleStart   float64  `json:"schedule_start"`
		State           int      `json:"state"`
		Type            string   `json:"type"`
		VarsAfter       struct {
			Attempt   int  `json:"attempt"`
			Reachable bool `json:"reachable"`
			State     int  `json:"state"`
			StateType int  `json:"state_type"`
		} `json:"vars_after"`
		VarsBefore struct {
			Attempt   int  `json:"attempt"`
			Reachable bool `json:"reachable"`
			State     int  `json:"state"`
			StateType int  `json:"state_type"`
		} `json:"vars_before"`
	} `json:"last_check_result"`
	LastHardState        int         `json:"last_hard_state"`
	LastHardStateChange  float64     `json:"last_hard_state_change"`
	LastReachable        bool        `json:"last_reachable"`
	LastState            int         `json:"last_state"`
	LastStateChange      float64     `json:"last_state_change"`
	LastStateDown        float64     `json:"last_state_down"`
	LastStateType        int         `json:"last_state_type"`
	LastStateUnreachable int         `json:"last_state_unreachable"`
	LastStateUp          float64     `json:"last_state_up"`
	MaxCheckAttempts     int         `json:"max_check_attempts"`
	NextCheck            float64     `json:"next_check"`
	Notes                string      `json:"notes"`
	NotesURL             string      `json:"notes_url"`
	OriginalAttributes   interface{} `json:"original_attributes"`
	Package              string      `json:"package"`
	Paused               bool        `json:"paused"`
	RetryInterval        int         `json:"retry_interval"`
	State                int         `json:"state"`
	StateType            int         `json:"state_type"`
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
	u := fmt.Sprintf("objects/hosts/%s", name)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	host := new(Host)
	resp, err := s.client.Do(req, host)
	if err != nil {
		return nil, resp, err
	}

	return host, resp, err
}
