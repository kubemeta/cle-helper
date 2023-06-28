package ilm

type ClusterLevelSettings struct {
	GlobalLifecycle *GlobalLifecycle `json:"lifecycle,omitempty"`
}

type GlobalLifecycle struct {
	HistoryIndexEnabled *bool  `json:"history_index_enabled,omitempty"`
	PollInterval        string `json:"poll_interval,omitempty"`
}

// IndexLevelSettings These index-level ILM settings are typically configured through index templates.
type IndexLevelSettings struct {
	Lifecycle *Lifecycle `json:"lifecycle,omitempty"`
}

type Lifecycle struct {
	// IndexingComplete Indicates whether the index has been rolled over. Automatically set to true
	IndexingComplete     *bool  `json:"indexing_complete,omitempty"`
	Name                 string `json:"name,omitempty,omitempty"`
	OriginationDate      int64  `json:"origination_date,omitempty"`
	ParseOriginationDate bool   `json:"parse_origination_date,omitempty"`
	Step                 *Step  `json:"step,omitempty"`
	RolloverAlias        string `json:"rollover_alias,omitempty"`
}

type Step struct {
	WaitTimeThreshold string `json:"wait_time_threshold,omitempty"`
}
