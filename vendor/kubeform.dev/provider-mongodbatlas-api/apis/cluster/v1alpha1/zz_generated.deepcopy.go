//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecAdvancedConfiguration) DeepCopyInto(out *ClusterSpecAdvancedConfiguration) {
	*out = *in
	if in.FailIndexKeyTooLong != nil {
		in, out := &in.FailIndexKeyTooLong, &out.FailIndexKeyTooLong
		*out = new(bool)
		**out = **in
	}
	if in.JavascriptEnabled != nil {
		in, out := &in.JavascriptEnabled, &out.JavascriptEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MinimumEnabledTlsProtocol != nil {
		in, out := &in.MinimumEnabledTlsProtocol, &out.MinimumEnabledTlsProtocol
		*out = new(string)
		**out = **in
	}
	if in.NoTableScan != nil {
		in, out := &in.NoTableScan, &out.NoTableScan
		*out = new(bool)
		**out = **in
	}
	if in.OplogSizeMb != nil {
		in, out := &in.OplogSizeMb, &out.OplogSizeMb
		*out = new(int64)
		**out = **in
	}
	if in.SampleRefreshIntervalBiConnector != nil {
		in, out := &in.SampleRefreshIntervalBiConnector, &out.SampleRefreshIntervalBiConnector
		*out = new(int64)
		**out = **in
	}
	if in.SampleSizeBiConnector != nil {
		in, out := &in.SampleSizeBiConnector, &out.SampleSizeBiConnector
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecAdvancedConfiguration.
func (in *ClusterSpecAdvancedConfiguration) DeepCopy() *ClusterSpecAdvancedConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecAdvancedConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecBiConnectorConfig) DeepCopyInto(out *ClusterSpecBiConnectorConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ReadPreference != nil {
		in, out := &in.ReadPreference, &out.ReadPreference
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecBiConnectorConfig.
func (in *ClusterSpecBiConnectorConfig) DeepCopy() *ClusterSpecBiConnectorConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecBiConnectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecConnectionStrings) DeepCopyInto(out *ClusterSpecConnectionStrings) {
	*out = *in
	if in.AwsPrivateLink != nil {
		in, out := &in.AwsPrivateLink, &out.AwsPrivateLink
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AwsPrivateLinkSrv != nil {
		in, out := &in.AwsPrivateLinkSrv, &out.AwsPrivateLinkSrv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = make([]ClusterSpecConnectionStringsPrivateEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateSrv != nil {
		in, out := &in.PrivateSrv, &out.PrivateSrv
		*out = new(string)
		**out = **in
	}
	if in.Standard != nil {
		in, out := &in.Standard, &out.Standard
		*out = new(string)
		**out = **in
	}
	if in.StandardSrv != nil {
		in, out := &in.StandardSrv, &out.StandardSrv
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecConnectionStrings.
func (in *ClusterSpecConnectionStrings) DeepCopy() *ClusterSpecConnectionStrings {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecConnectionStrings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecConnectionStringsPrivateEndpoint) DeepCopyInto(out *ClusterSpecConnectionStringsPrivateEndpoint) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]ClusterSpecConnectionStringsPrivateEndpointEndpoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SrvConnectionString != nil {
		in, out := &in.SrvConnectionString, &out.SrvConnectionString
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecConnectionStringsPrivateEndpoint.
func (in *ClusterSpecConnectionStringsPrivateEndpoint) DeepCopy() *ClusterSpecConnectionStringsPrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecConnectionStringsPrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecConnectionStringsPrivateEndpointEndpoints) DeepCopyInto(out *ClusterSpecConnectionStringsPrivateEndpointEndpoints) {
	*out = *in
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecConnectionStringsPrivateEndpointEndpoints.
func (in *ClusterSpecConnectionStringsPrivateEndpointEndpoints) DeepCopy() *ClusterSpecConnectionStringsPrivateEndpointEndpoints {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecConnectionStringsPrivateEndpointEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecLabels) DeepCopyInto(out *ClusterSpecLabels) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecLabels.
func (in *ClusterSpecLabels) DeepCopy() *ClusterSpecLabels {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecReplicationSpecs) DeepCopyInto(out *ClusterSpecReplicationSpecs) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(int64)
		**out = **in
	}
	if in.RegionsConfig != nil {
		in, out := &in.RegionsConfig, &out.RegionsConfig
		*out = make([]ClusterSpecReplicationSpecsRegionsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ZoneName != nil {
		in, out := &in.ZoneName, &out.ZoneName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecReplicationSpecs.
func (in *ClusterSpecReplicationSpecs) DeepCopy() *ClusterSpecReplicationSpecs {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecReplicationSpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecReplicationSpecsRegionsConfig) DeepCopyInto(out *ClusterSpecReplicationSpecsRegionsConfig) {
	*out = *in
	if in.AnalyticsNodes != nil {
		in, out := &in.AnalyticsNodes, &out.AnalyticsNodes
		*out = new(int64)
		**out = **in
	}
	if in.ElectableNodes != nil {
		in, out := &in.ElectableNodes, &out.ElectableNodes
		*out = new(int64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.ReadOnlyNodes != nil {
		in, out := &in.ReadOnlyNodes, &out.ReadOnlyNodes
		*out = new(int64)
		**out = **in
	}
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecReplicationSpecsRegionsConfig.
func (in *ClusterSpecReplicationSpecsRegionsConfig) DeepCopy() *ClusterSpecReplicationSpecsRegionsConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecReplicationSpecsRegionsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.AdvancedConfiguration != nil {
		in, out := &in.AdvancedConfiguration, &out.AdvancedConfiguration
		*out = new(ClusterSpecAdvancedConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoScalingComputeEnabled != nil {
		in, out := &in.AutoScalingComputeEnabled, &out.AutoScalingComputeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoScalingComputeScaleDownEnabled != nil {
		in, out := &in.AutoScalingComputeScaleDownEnabled, &out.AutoScalingComputeScaleDownEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoScalingDiskGbEnabled != nil {
		in, out := &in.AutoScalingDiskGbEnabled, &out.AutoScalingDiskGbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BackingProviderName != nil {
		in, out := &in.BackingProviderName, &out.BackingProviderName
		*out = new(string)
		**out = **in
	}
	if in.BackupEnabled != nil {
		in, out := &in.BackupEnabled, &out.BackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BiConnector != nil {
		in, out := &in.BiConnector, &out.BiConnector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.BiConnectorConfig != nil {
		in, out := &in.BiConnectorConfig, &out.BiConnectorConfig
		*out = new(ClusterSpecBiConnectorConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(string)
		**out = **in
	}
	if in.ConnectionStrings != nil {
		in, out := &in.ConnectionStrings, &out.ConnectionStrings
		*out = make([]ClusterSpecConnectionStrings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContainerID != nil {
		in, out := &in.ContainerID, &out.ContainerID
		*out = new(string)
		**out = **in
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(float64)
		**out = **in
	}
	if in.EncryptionAtRestProvider != nil {
		in, out := &in.EncryptionAtRestProvider, &out.EncryptionAtRestProvider
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]ClusterSpecLabels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongoDbMajorVersion != nil {
		in, out := &in.MongoDbMajorVersion, &out.MongoDbMajorVersion
		*out = new(string)
		**out = **in
	}
	if in.MongoDbVersion != nil {
		in, out := &in.MongoDbVersion, &out.MongoDbVersion
		*out = new(string)
		**out = **in
	}
	if in.MongoURI != nil {
		in, out := &in.MongoURI, &out.MongoURI
		*out = new(string)
		**out = **in
	}
	if in.MongoURIUpdated != nil {
		in, out := &in.MongoURIUpdated, &out.MongoURIUpdated
		*out = new(string)
		**out = **in
	}
	if in.MongoURIWithOptions != nil {
		in, out := &in.MongoURIWithOptions, &out.MongoURIWithOptions
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(int64)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.PitEnabled != nil {
		in, out := &in.PitEnabled, &out.PitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProviderAutoScalingComputeMaxInstanceSize != nil {
		in, out := &in.ProviderAutoScalingComputeMaxInstanceSize, &out.ProviderAutoScalingComputeMaxInstanceSize
		*out = new(string)
		**out = **in
	}
	if in.ProviderAutoScalingComputeMinInstanceSize != nil {
		in, out := &in.ProviderAutoScalingComputeMinInstanceSize, &out.ProviderAutoScalingComputeMinInstanceSize
		*out = new(string)
		**out = **in
	}
	if in.ProviderBackupEnabled != nil {
		in, out := &in.ProviderBackupEnabled, &out.ProviderBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProviderDiskIops != nil {
		in, out := &in.ProviderDiskIops, &out.ProviderDiskIops
		*out = new(int64)
		**out = **in
	}
	if in.ProviderDiskTypeName != nil {
		in, out := &in.ProviderDiskTypeName, &out.ProviderDiskTypeName
		*out = new(string)
		**out = **in
	}
	if in.ProviderEncryptEbsVolume != nil {
		in, out := &in.ProviderEncryptEbsVolume, &out.ProviderEncryptEbsVolume
		*out = new(bool)
		**out = **in
	}
	if in.ProviderEncryptEbsVolumeFlag != nil {
		in, out := &in.ProviderEncryptEbsVolumeFlag, &out.ProviderEncryptEbsVolumeFlag
		*out = new(bool)
		**out = **in
	}
	if in.ProviderInstanceSizeName != nil {
		in, out := &in.ProviderInstanceSizeName, &out.ProviderInstanceSizeName
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.ProviderRegionName != nil {
		in, out := &in.ProviderRegionName, &out.ProviderRegionName
		*out = new(string)
		**out = **in
	}
	if in.ProviderVolumeType != nil {
		in, out := &in.ProviderVolumeType, &out.ProviderVolumeType
		*out = new(string)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int64)
		**out = **in
	}
	if in.ReplicationSpecs != nil {
		in, out := &in.ReplicationSpecs, &out.ReplicationSpecs
		*out = make([]ClusterSpecReplicationSpecs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SnapshotBackupPolicy != nil {
		in, out := &in.SnapshotBackupPolicy, &out.SnapshotBackupPolicy
		*out = make([]ClusterSpecSnapshotBackupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SrvAddress != nil {
		in, out := &in.SrvAddress, &out.SrvAddress
		*out = new(string)
		**out = **in
	}
	if in.StateName != nil {
		in, out := &in.StateName, &out.StateName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecSnapshotBackupPolicy) DeepCopyInto(out *ClusterSpecSnapshotBackupPolicy) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.NextSnapshot != nil {
		in, out := &in.NextSnapshot, &out.NextSnapshot
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]ClusterSpecSnapshotBackupPolicyPolicies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReferenceHourOfDay != nil {
		in, out := &in.ReferenceHourOfDay, &out.ReferenceHourOfDay
		*out = new(int64)
		**out = **in
	}
	if in.ReferenceMinuteOfHour != nil {
		in, out := &in.ReferenceMinuteOfHour, &out.ReferenceMinuteOfHour
		*out = new(int64)
		**out = **in
	}
	if in.RestoreWindowDays != nil {
		in, out := &in.RestoreWindowDays, &out.RestoreWindowDays
		*out = new(int64)
		**out = **in
	}
	if in.UpdateSnapshots != nil {
		in, out := &in.UpdateSnapshots, &out.UpdateSnapshots
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecSnapshotBackupPolicy.
func (in *ClusterSpecSnapshotBackupPolicy) DeepCopy() *ClusterSpecSnapshotBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecSnapshotBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecSnapshotBackupPolicyPolicies) DeepCopyInto(out *ClusterSpecSnapshotBackupPolicyPolicies) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PolicyItem != nil {
		in, out := &in.PolicyItem, &out.PolicyItem
		*out = make([]ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecSnapshotBackupPolicyPolicies.
func (in *ClusterSpecSnapshotBackupPolicyPolicies) DeepCopy() *ClusterSpecSnapshotBackupPolicyPolicies {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecSnapshotBackupPolicyPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem) DeepCopyInto(out *ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(int64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem.
func (in *ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem) DeepCopy() *ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecSnapshotBackupPolicyPoliciesPolicyItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}
