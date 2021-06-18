package gremlin

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mitchellh/mapstructure"
)

type ListKubernetesAttacksResponse struct {
	KubernetesAttacks []KubernetesAttack
}

type ImpactDefinition struct {
	CliArgs []string `json:"cliArgs,omitempty"`
}

type K8sObjectInternal struct {
	ClusterId          string            `json:"clusterId,omitempty"`
	Uid                string            `json:"uid,omitempty"`
	Namespace          string            `json:"namespace,omitempty"`
	Name               string            `json:"name,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	Labels             map[string]string `json:"labels,omitempty"`
	Annotations        map[string]string `json:"annotations,omitempty"`
	OwnerReferences    string            `json:"ownerReferences,omitempty"`
	CreatedAt          string            `json:"createdAt,omitempty"`
	ResolvedContainers []string          `json:"resolvedContainers,omitempty"`
	Data               string            `json:"data,omitempty"`
	ClientId           string            `json:"clientId,omitempty"`
	AttackUid          string            `json:"attackUid,omitempty"`
	TeamId             string            `json:"teamId,omitempty"`
	Resolved           bool              `json:"resolved,omitempty"`
}

type K8sObject struct {
	ClusterId          string            `json:"clusterId,omitempty"`
	OwnerReferences    string            `json:"ownerReferences,omitempty"`
	CreatedAt          string            `json:"createdAt,omitempty"`
	ResolvedContainers []string          `json:"resolvedContainers,omitempty"`
	Uid                string            `json:"uid,omitempty"`
	Namespace          string            `json:"namespace,omitempty"`
	Name               string            `json:"name,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	Labels             map[string]string `json:"labels,omitempty"`
	Annotations        map[string]string `json:"annotations,omitempty"`
	PodPhase           string            `json:"podPhase,omitempty"`
}

type Strategy struct {
	K8sObjects         []K8sObject         `json:"k8sObjects,omitempty"`
	K8sObjectsInternal []K8sObjectInternal `json:"k8sObjectsInternal,omitempty"`
	Count              int32               `json:"count,omitempty"`
	Percentage         int32               `json:"percentage,omitempty"`
}

type TargetDefinition struct {
	Strategy Strategy `json:"strategy,omitempty"`
}

type AttackDefinition struct {
	ImpactDefinition ImpactDefinition `json:"impactDefinition"`
	TargetDefinition TargetDefinition `json:"targetDefinition"`
}

type KubernetesAttack struct {
	Uid        string           `json:"uid,omitempty"`
	AttackId   string           `json:"attackId,omitempty"`
	Stage      string           `json:"stage,omitempty"`
	CreateUser string           `json:"createUser,omitempty"`
	CreatedAt  string           `json:"createdAt,omitempty"`
	UpdatedAt  string           `json:"updatedAt,omitempty"`
	Error      string           `json:"error,omitempty"`
	Attack     AttackDefinition `json:"attack"`
}

func (c *Client) ListKubernetesAttacks(team_id string) (*ListKubernetesAttacksResponse, error) {
	resp, err := c.get("/kubernetes/attacks?source=Adhoc&teamId=" + team_id)
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(d), &results)
	var attacks []KubernetesAttack
	for _, result := range results {
		var attack KubernetesAttack
		err := mapstructure.Decode(result, &attack)
		if err != nil {
			return nil, err
		}
		attacks = append(attacks, attack)
	}
	result := ListKubernetesAttacksResponse{KubernetesAttacks: attacks}

	return &result, nil
}
