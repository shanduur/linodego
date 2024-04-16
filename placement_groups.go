package linodego

import "context"

// PlacementGroupAffinityType is an enum that determines the affinity policy
// for Linodes in a placement group.
type PlacementGroupAffinityType string

const (
	AffinityTypeAntiAffinityLocal PlacementGroupAffinityType = "anti_affinity:local"
)

// PlacementGroupMember represents a single Linode assigned to a
// placement group.
type PlacementGroupMember struct {
	LinodeID    int  `json:"linode_id"`
	IsCompliant bool `json:"is_compliant"`
}

// PlacementGroup represents a Linode placement group.
type PlacementGroup struct {
	ID           int                        `json:"id"`
	Label        string                     `json:"label"`
	Region       string                     `json:"region"`
	AffinityType PlacementGroupAffinityType `json:"affinity_type"`
	IsCompliant  bool                       `json:"is_compliant"`
	IsStrict     bool                       `json:"is_strict"`
	Members      []PlacementGroupMember     `json:"members"`
}

// PlacementGroupCreateOptions represents the options to use
// when creating a placement group.
type PlacementGroupCreateOptions struct {
	Label        string                     `json:"label"`
	Region       string                     `json:"region"`
	AffinityType PlacementGroupAffinityType `json:"affinity_type"`
	IsStrict     bool                       `json:"is_strict"`
}

// PlacementGroupUpdateOptions represents the options to use
// when updating a placement group.
type PlacementGroupUpdateOptions struct {
	Label string `json:"label,omitempty"`
}

// PlacementGroupAssignOptions represents options used when
// assigning Linodes to a placement group.
type PlacementGroupAssignOptions struct {
	Linodes       []int `json:"linodes"`
	CompliantOnly *bool `json:"compliant_only,omitempty"`
}

// PlacementGroupUnAssignOptions represents options used when
// unassigning Linodes from a placement group.
type PlacementGroupUnAssignOptions struct {
	Linodes []int `json:"linodes"`
}

// ListPlacementGroups lists placement groups under the current account
// matching the given list options.
func (c *Client) ListPlacementGroups(
	ctx context.Context,
	options *ListOptions,
) ([]PlacementGroup, error) {
	return getPaginatedResults[PlacementGroup](
		ctx,
		c,
		"placement/groups",
		options,
	)
}

// GetPlacementGroup gets a placement group with the specified ID.
func (c *Client) GetPlacementGroup(
	ctx context.Context,
	id int,
) (*PlacementGroup, error) {
	return doGETRequest[PlacementGroup](
		ctx,
		c,
		formatAPIPath("placement/groups/%d", id),
	)
}

// CreatePlacementGroup creates a placement group with the specified options.
func (c *Client) CreatePlacementGroup(
	ctx context.Context,
	options PlacementGroupCreateOptions,
) (*PlacementGroup, error) {
	return doPOSTRequest[PlacementGroup](
		ctx,
		c,
		"placement/groups",
		options,
	)
}

// UpdatePlacementGroup updates a placement group with the specified ID using the provided options.
func (c *Client) UpdatePlacementGroup(
	ctx context.Context,
	id int,
	options PlacementGroupUpdateOptions,
) (*PlacementGroup, error) {
	return doPUTRequest[PlacementGroup](
		ctx,
		c,
		formatAPIPath("placement/groups/%d", id),
		options,
	)
}

// AssignPlacementGroupLinodes assigns the specified Linodes to the given
// placement group.
func (c *Client) AssignPlacementGroupLinodes(
	ctx context.Context,
	id int,
	options PlacementGroupAssignOptions,
) (*PlacementGroup, error) {
	return doPOSTRequest[PlacementGroup](
		ctx,
		c,
		formatAPIPath("placement/groups/%d/assign", id),
		options,
	)
}

// UnAssignPlacementGroupLinodes un-assigns the specified Linodes from the given
// placement group.
func (c *Client) UnAssignPlacementGroupLinodes(
	ctx context.Context,
	id int,
	options PlacementGroupUnAssignOptions,
) (*PlacementGroup, error) {
	return doPOSTRequest[PlacementGroup](
		ctx,
		c,
		formatAPIPath("placement/groups/%d/unassign", id),
		options,
	)
}

// DeletePlacementGroup deletes a placement group with the specified ID.
func (c *Client) DeletePlacementGroup(
	ctx context.Context,
	id int,
) error {
	return doDELETERequest(
		ctx,
		c,
		formatAPIPath("placement/groups/%d", id),
	)
}
