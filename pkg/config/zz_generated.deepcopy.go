// +build !ignore_autogenerated

/*
 * Copyright 2019 SUSE LINUX GmbH, Nuernberg, Germany..
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfiguration) DeepCopyInto(out *AuthConfiguration) {
	*out = *in
	out.OIDC = in.OIDC
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfiguration.
func (in *AuthConfiguration) DeepCopy() *AuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindConfiguration) DeepCopyInto(out *BindConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindConfiguration.
func (in *BindConfiguration) DeepCopy() *BindConfiguration {
	if in == nil {
		return nil
	}
	out := new(BindConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapConfiguration) DeepCopyInto(out *BootstrapConfiguration) {
	*out = *in
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapConfiguration.
func (in *BootstrapConfiguration) DeepCopy() *BootstrapConfiguration {
	if in == nil {
		return nil
	}
	out := new(BootstrapConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertsConfiguration) DeepCopyInto(out *CertsConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertsConfiguration.
func (in *CertsConfiguration) DeepCopy() *CertsConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFormationConfiguration) DeepCopyInto(out *ClusterFormationConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFormationConfiguration.
func (in *ClusterFormationConfiguration) DeepCopy() *ClusterFormationConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterFormationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CniConfiguration) DeepCopyInto(out *CniConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CniConfiguration.
func (in *CniConfiguration) DeepCopy() *CniConfiguration {
	if in == nil {
		return nil
	}
	out := new(CniConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfiguration) DeepCopyInto(out *DNSConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfiguration.
func (in *DNSConfiguration) DeepCopy() *DNSConfiguration {
	if in == nil {
		return nil
	}
	out := new(DNSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfiguration) DeepCopyInto(out *EtcdConfiguration) {
	*out = *in
	if in.LocalEtcd != nil {
		in, out := &in.LocalEtcd, &out.LocalEtcd
		*out = new(LocalEtcdConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfiguration.
func (in *EtcdConfiguration) DeepCopy() *EtcdConfiguration {
	if in == nil {
		return nil
	}
	out := new(EtcdConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesConfiguration) DeepCopyInto(out *FeaturesConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesConfiguration.
func (in *FeaturesConfiguration) DeepCopy() *FeaturesConfiguration {
	if in == nil {
		return nil
	}
	out := new(FeaturesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubicInitConfiguration) DeepCopyInto(out *KubicInitConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Network = in.Network
	out.Paths = in.Paths
	out.ClusterFormation = in.ClusterFormation
	out.Certificates = in.Certificates
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Runtime = in.Runtime
	out.Features = in.Features
	out.Services = in.Services
	out.Auth = in.Auth
	in.Bootstrap.DeepCopyInto(&out.Bootstrap)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubicInitConfiguration.
func (in *KubicInitConfiguration) DeepCopy() *KubicInitConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubicInitConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubicInitConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalEtcdConfiguration) DeepCopyInto(out *LocalEtcdConfiguration) {
	*out = *in
	if in.ServerCertSANs != nil {
		in, out := &in.ServerCertSANs, &out.ServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PeerCertSANs != nil {
		in, out := &in.PeerCertSANs, &out.PeerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalEtcdConfiguration.
func (in *LocalEtcdConfiguration) DeepCopy() *LocalEtcdConfiguration {
	if in == nil {
		return nil
	}
	out := new(LocalEtcdConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mirror) DeepCopyInto(out *Mirror) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mirror.
func (in *Mirror) DeepCopy() *Mirror {
	if in == nil {
		return nil
	}
	out := new(Mirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfiguration) DeepCopyInto(out *NetworkConfiguration) {
	*out = *in
	out.Bind = in.Bind
	out.Cni = in.Cni
	out.DNS = in.DNS
	out.Proxy = in.Proxy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfiguration.
func (in *NetworkConfiguration) DeepCopy() *NetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(NetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCConfiguration) DeepCopyInto(out *OIDCConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCConfiguration.
func (in *OIDCConfiguration) DeepCopy() *OIDCConfiguration {
	if in == nil {
		return nil
	}
	out := new(OIDCConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PathsConfigration) DeepCopyInto(out *PathsConfigration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathsConfigration.
func (in *PathsConfigration) DeepCopy() *PathsConfigration {
	if in == nil {
		return nil
	}
	out := new(PathsConfigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfiguration) DeepCopyInto(out *ProxyConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfiguration.
func (in *ProxyConfiguration) DeepCopy() *ProxyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProxyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.Mirrors != nil {
		in, out := &in.Mirrors, &out.Mirrors
		*out = make([]Mirror, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeConfiguration) DeepCopyInto(out *RuntimeConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeConfiguration.
func (in *RuntimeConfiguration) DeepCopy() *RuntimeConfiguration {
	if in == nil {
		return nil
	}
	out := new(RuntimeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesConfiguration) DeepCopyInto(out *ServicesConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesConfiguration.
func (in *ServicesConfiguration) DeepCopy() *ServicesConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServicesConfiguration)
	in.DeepCopyInto(out)
	return out
}
