package icinga

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// HostsService handles communication with the host object related methods of the Icinga2 API
type HostsService service

type hostServiceResults struct {
	Results []hostAttrs `json:"results"`
}

type hostAttrs struct {
	Attrs Host `json:"attrs" icinga:""`
}

// Host represents a Icinga2 Host object
type Host struct {
	Name                  string      `json:"__name,omitempty" icinga:"create"`
	Acknowledgement       json.Number `json:"acknowledgement,omitempty" icinga:""`
	AcknowledgementExpiry json.Number `json:"acknowledgement_expiry,omitempty" icinga:""`
	ActionURL             string      `json:"action_url,omitempty" icinga:""`
	Address               string      `json:"address,omitempty" icinga:"create"`
	Address6              string      `json:"address6,omitempty" icinga:""`
	CheckAttempt          json.Number `json:"check_attempt,omitempty" icinga:""`
	CheckCommand          string      `json:"check_command,omitempty" icinga:"create"`
	CheckInterval         json.Number `json:"check_interval,omitempty" icinga:""`
	CheckPeriod           string      `json:"check_period,omitempty" icinga:""`
	CheckTimeout          interface{} `json:"check_timeout,omitempty" icinga:""`
	CommandEndpoint       string      `json:"command_endpoint,omitempty" icinga:""`
	DisplayName           string      `json:"display_name,omitempty" icinga:""`
	DowntimeDepth         json.Number `json:"downtime_depth,omitempty" icinga:""`
	EventCommand          string      `json:"event_command,omitempty" icinga:""`
	FlappingLastChange    float64     `json:"flapping_last_change,omitempty" icinga:""`
	FlappingNegative      json.Number `json:"flapping_negative,omitempty" icinga:""`
	FlappingPositive      json.Number `json:"flapping_positive,omitempty" icinga:""`
	FlappingThreshold     json.Number `json:"flapping_threshold,omitempty" icinga:""`
	Groups                []string    `json:"groups,omitempty" icinga:""`
	HaMode                json.Number `json:"ha_mode,omitempty" icinga:""`
	IconImage             string      `json:"icon_image,omitempty" icinga:""`
	IconImageAlt          string      `json:"icon_image_alt,omitempty" icinga:""`
	LastCheck             float64     `json:"last_check,omitempty" icinga:""`
	LastHardState         json.Number `json:"last_hard_state,omitempty" icinga:""`
	LastHardStateChange   float64     `json:"last_hard_state_change,omitempty" icinga:""`
	LastState             json.Number `json:"last_state,omitempty" icinga:""`
	LastStateChange       float64     `json:"last_state_change,omitempty" icinga:""`
	LastStateDown         float64     `json:"last_state_down,omitempty" icinga:""`
	LastStateType         json.Number `json:"last_state_type,omitempty" icinga:""`
	LastStateUnreachable  json.Number `json:"last_state_unreachable,omitempty" icinga:""`
	LastStateUp           float64     `json:"last_state_up,omitempty" icinga:""`
	MaxCheckAttempts      json.Number `json:"max_check_attempts,omitempty" icinga:""`
	NextCheck             float64     `json:"next_check,omitempty" icinga:""`
	Notes                 string      `json:"notes,omitempty" icinga:""`
	NotesURL              string      `json:"notes_url,omitempty" icinga:""`
	OriginalAttributes    interface{} `json:"original_attributes,omitempty" icinga:""`
	Package               string      `json:"package,omitempty" icinga:""`
	RetryInterval         json.Number `json:"retry_interval,omitempty" icinga:""`
	State                 json.Number `json:"state,omitempty" icinga:""`
	StateType             json.Number `json:"state_type,omitempty" icinga:""`
	Templates             []string    `json:"templates,omitempty" icinga:""`
	Type                  string      `json:"type,omitempty" icinga:""`
	Version               float64     `json:"version,omitempty" icinga:""`
	Zone                  string      `json:"zone,omitempty" icinga:""`
	Flapping              bool        `json:"flapping,omitempty" icinga:""`
	Active                bool        `json:"active,omitempty" icinga:""`
	EnableActiveChecks    bool        `json:"enable_active_checks,omitempty" icinga:""`
	EnableEventHandler    bool        `json:"enable_event_handler,omitempty" icinga:""`
	EnableFlapping        bool        `json:"enable_flapping,omitempty" icinga:""`
	EnableNotifications   bool        `json:"enable_notifications,omitempty" icinga:""`
	EnablePassiveChecks   bool        `json:"enable_passive_checks,omitempty" icinga:""`
	EnablePerfdata        bool        `json:"enable_perfdata,omitempty" icinga:""`
	ForceNextCheck        bool        `json:"force_next_check,omitempty" icinga:""`
	ForceNextNotification bool        `json:"force_next_notification,omitempty" icinga:""`
	LastReachable         bool        `json:"last_reachable,omitempty" icinga:""`
	Paused                bool        `json:"paused,omitempty" icinga:""`
	Volatile              bool        `json:"volatile,omitempty" icinga:""`
	Vars                  interface{} `json:"vars,omitempty" icinga:""`
	LastCheckResult       *struct {
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
	} `json:"last_check_result,omitempty" icinga:""`
}

func (h *Host) objectForCreate() Host {
	sourceType, sourceValue := reflect.TypeOf(*h), reflect.ValueOf(*h)

	target := Host{}
	targetValue := reflect.ValueOf(&target).Elem()

	for i := 0; i < sourceType.NumField(); i++ {
		icingaKey := sourceType.Field(i).Tag.Get("icinga")
		if strings.Contains(icingaKey, "create") {
			targetValue.Field(i).Set(sourceValue.Field(i))
		}
	}

	return target
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

// Create Host.
func (s *HostsService) Create(host *Host) (*Host, *http.Response, error) {
	if host.Name == "" {
		return nil, nil, errors.New("Host name is empty")
	}

	// TODO validate / encode host name

	u := fmt.Sprintf("objects/hosts/%s", host.Name)

	requestBody := hostAttrs{
		Attrs: host.objectForCreate(),
	}

	req, err := s.client.NewRequest("PUT", u, requestBody)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	return nil, resp, nil
}
