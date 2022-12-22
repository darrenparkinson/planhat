package planhat

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// User represents a planhat user
type User struct {
	ID                         *string `json:"_id,omitempty"`
	SkippedGettingStartedSteps struct {
		Email     *bool `json:"email,omitempty"`
		Linkedin  *bool `json:"linkedin,omitempty"`
		Avatar    *bool `json:"avatar,omitempty"`
		All       *bool `json:"all,omitempty"`
		Team      *bool `json:"team,omitempty"`
		Customers *bool `json:"customers,omitempty"`
	} `json:"skippedGettingStartedSteps,omitempty"`
	Image struct {
		Path *string `json:"path,omitempty"`
	} `json:"image,omitempty"`
	FirstName               *string    `json:"firstName,omitempty"`
	LastName                *string    `json:"lastName,omitempty"`
	IsHidden                *bool      `json:"isHidden,omitempty"`
	Removed                 *bool      `json:"removed,omitempty"`
	Inactive                *bool      `json:"inactive,omitempty"`
	CompressedView          *bool      `json:"compressedView,omitempty"`
	CompanyFilter           *string    `json:"companyFilter,omitempty"`
	TaskFilter              *string    `json:"taskFilter,omitempty"`
	WorkflowFilter          *string    `json:"workflowFilter,omitempty"`
	PlayLogDisabled         *bool      `json:"playLogDisabled,omitempty"`
	RadarOneLine            *bool      `json:"radarOneLine,omitempty"`
	CollapsedFolders        []string   `json:"collapsedFolders,omitempty"`
	RevReportPeriodType     *string    `json:"revReportPeriodType,omitempty"`
	SplitLayoutDisabled     *bool      `json:"splitLayoutDisabled,omitempty"`
	DailyDigest             *bool      `json:"dailyDigest,omitempty"`
	FollowerUpdate          *bool      `json:"followerUpdate,omitempty"`
	InAppNotifications      *bool      `json:"inAppNotifications,omitempty"`
	LastVisitedCompanies    []string   `json:"lastVisitedCompanies,omitempty"`
	LastVisitedEndusers     []string   `json:"lastVisitedEndusers,omitempty"`
	Roles                   []string   `json:"roles,omitempty"`
	IsExposedAsSenderOption *bool      `json:"isExposedAsSenderOption,omitempty"`
	DefaultMeetingLength    *int       `json:"defaultMeetingLength,omitempty"`
	NickName                *string    `json:"nickName,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	CreateDate              *time.Time `json:"createDate,omitempty"`
	V                       *int       `json:"__v,omitempty"`
	RecentOpenTabs          *struct {
		Customers           *string `json:"customers,omitempty"`
		Bi                  *string `json:"bi,omitempty"`
		BiSystem            *string `json:"bi-system,omitempty"`
		BiAnalytics         *string `json:"bi-analytics,omitempty"`
		People              *string `json:"people,omitempty"`
		Tasks               *string `json:"tasks,omitempty"`
		Conversations       *string `json:"conversations,omitempty"`
		ConversationsOutbox *string `json:"conversations-outbox,omitempty"`
		Engage              *string `json:"engage,omitempty"`
		Nps                 *string `json:"nps,omitempty"`
		Revenue             *string `json:"revenue,omitempty"`
		Settings            *string `json:"settings,omitempty"`
		Team                *string `json:"team,omitempty"`
	} `json:"recentOpenTabs,omitempty"`
	RecentOpenPage      *string `json:"recentOpenPage,omitempty"`
	Segment             *string `json:"segment,omitempty"`
	BubbleChartSettings struct {
		XParam *string `json:"xParam,omitempty"`
		YParam *string `json:"yParam,omitempty"`
	} `json:"bubbleChartSettings,omitempty"`
	RecentTabSearches struct {
		BaseTasksAssigned *string `json:"base-tasks-assigned,omitempty"`
	} `json:"recentTabSearches,omitempty"`
	GoogleAPI struct {
		AccessEnabled *bool         `json:"accessEnabled,omitempty"`
		SyncEnabled   *bool         `json:"syncEnabled,omitempty"`
		SyncInitial   *bool         `json:"syncInitial,omitempty"`
		SyncedLabels  []interface{} `json:"syncedLabels,omitempty"`
	} `json:"googleApi,omitempty"`
	MsAPI struct {
		AccessEnabled *bool         `json:"accessEnabled,omitempty"`
		SyncEnabled   *bool         `json:"syncEnabled,omitempty"`
		SyncInitial   *bool         `json:"syncInitial,omitempty"`
		SyncedLabels  []interface{} `json:"syncedLabels,omitempty"`
	} `json:"msApi,omitempty"`
	GoogleCalendarAPI struct {
		AccessEnabled   *bool         `json:"accessEnabled,omitempty"`
		SyncEnabled     *bool         `json:"syncEnabled,omitempty"`
		SyncInitial     *bool         `json:"syncInitial,omitempty"`
		SyncedCalendars []interface{} `json:"syncedCalendars,omitempty"`
		CalendarToSave  struct {
		} `json:"calendarToSave,omitempty"`
	} `json:"googleCalendarApi,omitempty"`
}

// List returns a list of planhat users
func (s *UserService) List(ctx context.Context) ([]*User, error) {
	ur := []*User{}
	url := fmt.Sprintf("%s/users", s.client.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ur, err
	}
	if err := s.client.makeRequest(ctx, req, &ur); err != nil {
		return ur, err
	}
	return ur, nil
}

// Get returns a single user given it's planhat ID
func (s *UserService) Get(ctx context.Context, id string) (*User, error) {
	user := &User{}
	url := fmt.Sprintf("%s/users/%s", s.client.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	if err := s.client.makeRequest(ctx, req, &user); err != nil {
		return user, err
	}
	return user, nil
}
