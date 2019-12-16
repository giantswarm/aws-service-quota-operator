package v1alpha1

import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindRelease = "Release"

	ChangelogKindAdded      ReleaseChangelogKind = "added"
	ChangelogKindChanged    ReleaseChangelogKind = "changed"
	ChangelogKindDeprecated ReleaseChangelogKind = "deprecated"
	ChangelogKindFixed      ReleaseChangelogKind = "fixed"
	ChangelogKindRemoved    ReleaseChangelogKind = "removed"
	ChangelogKindSecurity   ReleaseChangelogKind = "security"
)

const releaseCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: releases.release.giantswarm.io
spec:
  group: release.giantswarm.io
  scope: Cluster
  version: v1alpha1
  names:
    kind: Release
    plural: releases
    singular: release
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            changelog:
              type: array
              minItems: 1
              items:
                type: object
                properties:
                  component:
                    type: string
                    minLength: 3
                  description:
                    type: string
                    minLength: 3
                  kind:
                    enum:
                    - added
                    - changed
                    - deprecated
                    - fixed
                    - removed
                    - security
                required:
                - component
                - description
                - kind
            components:
              type: array
              minItems: 1
              items:
                type: object
                properties:
                  name:
                    type: string
                    minLength: 3
                  version:
                    type: string
                    minLength: 5
            parentVersion:
              type: string
              pattern: "^\\d+\\.\\d+\\.\\d+$"
            version:
              type: string
              minLength: 5
          required:
          - changelog
          - components
          - parentVersion
          - version
        status:
          type: object
          properties:
            cycle:
              type: object
              properties:
                disabledDate:
                  type: string
                  format: date
                enabledDate:
                  type: string
                  format: date
                phase:
                  enum:
                  - upcoming
                  - enabled
                  - disabled
                  - eol
              required:
              - phase
`

type ReleaseChangelogKind string

var releaseCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(releaseCRDYAML), &releaseCRD)
	if err != nil {
		panic(err)
	}
}

func NewReleaseCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return releaseCRD.DeepCopy()
}

func NewReleaseTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindRelease,
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Release CRs might look something like the following.
//
//	apiVersion: "release.giantswarm.io/v1alpha1"
//	kind: "Release"
//	metadata:
//	  name: "aws.v6.1.0"
//	  labels:
//	    giantswarm.io/managed-by: "app-operator"
//	    giantswarm.io/provider: "aws"
//	spec:
//	  changelog:
//	    - component: "cloudconfig"
//	      description: "Replace cloudinit with ignition."
//	      kind: "changed"
//	  components:
//	    - name: "aws-operator"
//	      version: "4.6.0"
//	    - name: "cert-operator"
//	      version: "0.1.0"
//	    - name: "chart-operator"
//	      version: "0.5.0"
//	    - name: "cluster-operator"
//	      version: "0.10.0"
//	  parentVersion: "6.0.1"
//	  version: "6.1.0"
//	status:
//	  cycle:
//	    phase: "eol"
//	    enabledDate: 2019-01-08
//	    disabledDate: 2019-01-12
//
type Release struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ReleaseSpec   `json:"spec" yaml:"spec"`
	Status            ReleaseStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ReleaseSpec struct {
	// Changelog is the changelog since ParentVersion.
	Changelog []ReleaseSpecChangelogEntry `json:"changelog" yaml:"changelog"`
	// Components describes components managing this release.
	Components []ReleaseSpecComponent `json:"components" yaml:"components"`
	// ParentVersion is a version from which the changes in changelog are
	// described. We need that because we may introduce bug fixes after
	// next major release and then taking previous semver version may not
	// render correct changelog. This should always be in the semver format
	// without the "v" prefix.
	ParentVersion string `json:"parentVersion" yaml:"parentVersion"`
	// Version is the version of the release. Releases with semver version
	// (without the "v" prefix) are taken from control-plane AppCatalog.
	// All other releases are taken from control-plane-test AppCatalog.
	Version string `json:"version" yaml:"version"`
}

type ReleaseSpecChangelogEntry struct {
	// Component name.
	Component string `json:"component" yaml:"component"`
	// Description of the component changes expressed in full sentence.
	Description string `json:"description" yaml:"description"`
	// Kind of the component changes. It can be one of: "added", "changed",
	// "deprecated", "fixed", "removed", "security".
	Kind ReleaseChangelogKind `json:"kind" yaml:"kind"`
}

type ReleaseSpecChangelogEntryKind string

type ReleaseSpecComponent struct {
	// Name of the release component.
	Name string `json:"name" yaml:"name"`
	// Version of the release component.
	Version string `json:"version" yaml:"version"`
}

type ReleaseStatus struct {
	// Cycle is the most recent observed copy of the specification of the
	// ReleaseCycle CR referencing this Release CR.
	Cycle ReleaseCycleSpec `json:"cycle,omitempty" yaml:"cycle,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Release `json:"items" yaml:"items"`
}
