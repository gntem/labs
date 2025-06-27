variable "project_id" {
  description = "The GCP project ID"
  type        = string
  default     = "experiment01-tofu"
}

variable "region" {
  description = "The GCP region for resources"
  type        = string
  default     = "us-central1"
}

variable "zone" {
  description = "The GCP zone for resources"
  type        = string
  default     = "us-central1-a"
}

variable "billing_account_id" {
  description = "The billing account ID to associate with the project"
  type        = string
  # This must be provided when running tofu apply
}
