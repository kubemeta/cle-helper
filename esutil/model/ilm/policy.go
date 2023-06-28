package ilm

import (
	"fmt"

	"github.com/kubemeta/cle-helper/esutil/model/unit"
)

type IlmObject struct {
	Policy Policy `json:"policy"`
}

type Policy struct {
	Meta   *Meta  `json:"_meta,omitempty"`
	Phases Phases `json:"phases"`
}

type Meta struct {
	Description string   `json:"description,omitempty"`
	Project     *Project `json:"project,omitempty"`
}

type Project struct {
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
}

type Phases struct {
	Hot    *PhaseSetting `json:"hot,omitempty"`
	Warm   *PhaseSetting `json:"warm,omitempty"`
	Cold   *PhaseSetting `json:"cold,omitempty"`
	Frozen *PhaseSetting `json:"frozen,omitempty"`
	Delete *PhaseSetting `json:"delete,omitempty"`
}

type PhaseSetting struct {
	// a minimum age for each phase for an index to move to the next phase
	MinAge string `json:"min_age"`
	// all actions in the current phase must be complete and the index must be older than the minimum age of the next phase
	Actions *Actions `json:"actions"`
}

func NewPhaseSetting(actionsFunc ...ActionsFunc) *PhaseSetting {
	p := &PhaseSetting{
		MinAge:  "0s",
		Actions: new(Actions),
	}
	for _, f := range actionsFunc {
		f(p.Actions)
	}
	return p
}

func (s *PhaseSetting) SetMinAge(minAge int, unit unit.TimeUnit) *PhaseSetting {
	s.MinAge = fmt.Sprintf("%d%s", minAge, unit)
	return s
}

type ActionsFunc func(actions *Actions)

type Actions struct {
	// Phases allowed: hot, warm, cold
	SetPriority *SetPriority `json:"set_priority,omitempty"`
	// Phases allowed: hot, warm, cold, frozen
	Unfollow *struct{} `json:"unfollow,omitempty"`
	// Phases allowed: hot
	Rollover *Rollover `json:"rollover,omitempty"`
	// Phases allowed: hot, warm, cold
	ReadOnly *struct{} `json:"read_only,omitempty"`
	// Phases allowed: hot, warm
	Shrink *Shrink `json:"shrink,omitempty"`
	// Phases allowed: warm, cold
	Allocate *Allocate `json:"allocate,omitempty"`
	// Phases allowed: hot, warm
	ForceMerge *ForceMerge `json:"forcemerge,omitempty"`
	// Phases allowed: warm, cold
	// ILM will auto-inject migrate action
	Migrate *Migrate `json:"migrate,omitempty"`
	// Phases allowed: hot, cold, frozen
	SearchableSnapshot *SearchableSnapshot `json:"searchable_snapshot,omitempty"`
	// Phases allowed: cold
	Freeze *struct{} `json:"freeze,omitempty"`
	// Phases allowed: delete
	WaitForSnapshot *WaitForSnapshot `json:"wait_for_snapshot,omitempty"`
	// Phases allowed: delete
	Delete *Delete `json:"delete,omitempty"`
}

type SetPriority struct {
	Priority int `json:"priority,omitempty"`
}

type Rollover struct {
	MaxAge              string `json:"max_age,omitempty"`
	MaxDocs             int    `json:"max_docs,omitempty"`
	MaxSize             string `json:"max_size,omitempty"`
	MaxPrimaryShardSize string `json:"max_primary_shard_size,omitempty"`
}

type Shrink struct {
	NumberOfShards      int    `json:"number_of_shards,omitempty"`
	MaxPrimaryShardSize string `json:"max_primary_shard_size,omitempty"`
}

type Allocate struct {
	NumberOfReplicas   int `json:"number_of_replicas,omitempty"`
	TotalShardsPerNode int `json:"total_shards_per_node,omitempty"`
	// Assigns an index to nodes that have at least one of the specified custom attributes.
	Include CustomAttribute `json:"include,omitempty"`
	// Assigns an index to nodes that have none of the specified custom attributes.
	Exclude CustomAttribute `json:"exclude,omitempty"`
	// Assigns an index to nodes that have all the specified custom attributes.
	Require CustomAttribute `json:"require,omitempty"`
}

type CustomAttribute map[string]string

type Migrate struct {
}

type ForceMerge struct {
	MaxNumSegments int `json:"max_num_segments,omitempty"`
	// optional
	IndexCodec string `json:"index_codec,omitempty"`
}

type SearchableSnapshot struct {
	SnapshotRepository string `json:"snapshot_repository,omitempty"`
	ForceMergeIndex    *bool  `json:"force_merge_index,omitempty"`
}

type WaitForSnapshot struct {
	// Name of the SLM policy that the delete action should wait for.
	Policy string `json:"policy,omitempty"`
}

type Delete struct {
	DeleteSearchableSnapshot *bool `json:"delete_searchable_snapshot,omitempty"`
}

func WithPriority(priority int) ActionsFunc {
	return func(actions *Actions) {
		actions.SetPriority = new(SetPriority)
		actions.SetPriority.Priority = priority
	}
}

func WithUnfollow() ActionsFunc {
	return func(actions *Actions) {
		actions.Unfollow = new(struct{})
	}
}

func WithRollover(rollover *Rollover) ActionsFunc {
	return func(actions *Actions) {
		actions.Rollover = rollover
	}
}

func DefaultRollover() *Rollover {
	return &Rollover{
		MaxAge:              fmt.Sprintf("%d%s", 1, unit.Days),
		MaxDocs:             100000,
		MaxPrimaryShardSize: fmt.Sprintf("%d%s", 100, unit.Gigabytes),
	}
}

func WithReadOnly() ActionsFunc {
	return func(actions *Actions) {
		actions.ReadOnly = new(struct{})
	}
}

func WithShrink(shards int, maxPrimaryShardSize int, maxPrimaryShardSizeUnit string) ActionsFunc {
	return func(actions *Actions) {
		actions.Shrink = new(Shrink)
		if shards != 0 {
			actions.Shrink.NumberOfShards = shards
		}

		if maxPrimaryShardSize != 0 {
			actions.Shrink.MaxPrimaryShardSize = fmt.Sprintf("%d%s", maxPrimaryShardSize, maxPrimaryShardSizeUnit)
		}
	}
}

// WithAllocate Assign index to nodes based on multiple attributes
func WithAllocate(replicas *int, totalShards int, keyvals ...interface{}) ActionsFunc {
	return func(actions *Actions) {
		actions.Allocate = new(Allocate)
		if replicas != nil {
			actions.Allocate.NumberOfReplicas = *replicas
		}
		if totalShards != 0 {
			actions.Allocate.TotalShardsPerNode = totalShards
		}

		if len(keyvals) != 0 && len(keyvals)%2 == 0 {
			for i := 0; i < len(keyvals); i += 2 {
				switch keyvals[i] {
				case Include:
					val, ok := keyvals[i+1].(CustomAttribute)
					if ok {
						actions.Allocate.Include = val
					}
				case Exclude:
					val, ok := keyvals[i+1].(CustomAttribute)
					if ok {
						actions.Allocate.Exclude = val
					}
				case Require:
					val, ok := keyvals[i+1].(CustomAttribute)
					if ok {
						actions.Allocate.Require = val
					}
				}
			}
		}
	}
}

const (
	Include = "include"
	Exclude = "exclude"
	Require = "require"
)

func WithForceMerge(maxNumSegments int) ActionsFunc {
	return func(actions *Actions) {
		actions.ForceMerge = new(ForceMerge)
		actions.ForceMerge.MaxNumSegments = maxNumSegments
	}
}

func WithSearchableSnapshot() ActionsFunc {
	return func(actions *Actions) {

	}
}

func WithFreeze() ActionsFunc {
	return func(actions *Actions) {
		actions.Freeze = new(struct{})
	}
}

func WithWaitForSnapshot() ActionsFunc {
	return func(actions *Actions) {
		actions.WaitForSnapshot = new(WaitForSnapshot)
	}
}

func WithDelete(deleteSearchableSnapshot *bool) ActionsFunc {
	return func(actions *Actions) {
		actions.Delete = new(Delete)
		if deleteSearchableSnapshot != nil {
			actions.Delete.DeleteSearchableSnapshot = deleteSearchableSnapshot
		}
	}
}
