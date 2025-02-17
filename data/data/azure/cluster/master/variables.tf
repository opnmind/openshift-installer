variable "region" {
  type        = string
  description = "The region for the deployment."
}

variable "resource_group_name" {
  type        = string
  description = "The resource group name for the deployment."
}

variable "cluster_id" {
  type = string
}

variable "vm_size" {
  type = string
}

variable "disk_encryption_set_id" {
  type        = string
  default     = null
  description = "The ID of the Disk Encryption Set which should be used to encrypt OS disk."
}

variable "encryption_at_host_enabled" {
  type        = bool
  description = "Enables encryption at the VM host."
}

variable "vm_image" {
  type        = string
  description = "The resource id of the vm image used for masters."
}

variable "identity" {
  type        = string
  description = "The user assigned identity id for the vm."
}

variable "instance_count" {
  type = string
}

variable "elb_backend_pool_v4_id" {
  type = string
}

variable "elb_backend_pool_v6_id" {
  type = string
}

variable "ilb_backend_pool_v4_id" {
  type = string
}

variable "ilb_backend_pool_v6_id" {
  type = string
}

variable "ignition_master" {
  type    = string
  default = ""
}

variable "kubeconfig_content" {
  type    = string
  default = ""
}

variable "subnet_id" {
  type        = string
  description = "The subnet to attach the masters to."
}

variable "os_volume_type" {
  type        = string
  description = "The type of the volume for the root block device."
}

variable "os_volume_size" {
  type        = string
  description = "The size of the volume in gigabytes for the root block device."
}

variable "azure_extra_tags" {
  type = map(string)

  description = <<EOF
(optional) Extra Azure tags to be applied to created resources.

Example: `{ "key" = "value", "foo" = "bar" }`
EOF


  default = {}
}

variable "storage_account_name" {
  type = any
  description = "the name of the storage account for the cluster. It can be used for boot diagnostics."
}

variable "ignition" {
  type = string
}

variable "availability_zones" {
  type = list(string)
  description = "List of the availability zones in which to create the masters. The length of this list must match instance_count."
}

variable "private" {
  type = bool
  description = "This value determines if this is a private cluster or not."
}

variable "use_ipv4" {
  type = bool
  description = "This value determines if this is cluster should use IPv4 networking."
}

variable "use_ipv6" {
  type = bool
  description = "This value determines if this is cluster should use IPv6 networking."
}

variable "outbound_udr" {
  type = bool
  default = false

  description = <<EOF
This determined whether User defined routing will be used for egress to Internet.
When false, Standard LB will be used for egress to the Internet.

This is required because terraform cannot calculate counts during plan phase completely and therefore the `vnet/public-lb.tf`
conditional need to be recreated. See https://github.com/hashicorp/terraform/issues/12570
EOF
}

variable "ultra_ssd_enabled" {
  type        = bool
  description = "Determines if the control plane should have UltraSSD Enabled."
}

variable "vm_networking_type" {
  type        = bool
  description = <<EOF
networking_type specifies whether to enable accelerated networking. Accelerated networking
enables single root I/O virtualization (SR-IOV) to a VM, greatly improving its networking performance.
EOF
}
