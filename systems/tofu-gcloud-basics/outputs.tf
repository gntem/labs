output "project_id" {
  description = "The project ID"
  value       = google_project.experiment01.project_id
}

output "project_number" {
  description = "The project number"
  value       = google_project.experiment01.number
}

output "vpc_name" {
  description = "The name of the VPC"
  value       = google_compute_network.main_vpc.name
}

output "vpc_id" {
  description = "The ID of the VPC"
  value       = google_compute_network.main_vpc.id
}

output "main_subnet_name" {
  description = "The name of the main subnet"
  value       = google_compute_subnetwork.main_subnet.name
}

output "main_subnet_cidr" {
  description = "The CIDR range of the main subnet"
  value       = google_compute_subnetwork.main_subnet.ip_cidr_range
}

output "private_subnet_name" {
  description = "The name of the private subnet"
  value       = google_compute_subnetwork.private_subnet.name
}

output "private_subnet_cidr" {
  description = "The CIDR range of the private subnet"
  value       = google_compute_subnetwork.private_subnet.ip_cidr_range
}

output "gke_cluster_name" {
  description = "The name of the GKE cluster"
  value       = google_container_cluster.primary.name
}

output "gke_cluster_endpoint" {
  description = "The endpoint of the GKE cluster"
  value       = google_container_cluster.primary.endpoint
  sensitive   = true
}

output "gke_cluster_ca_certificate" {
  description = "The CA certificate of the GKE cluster"
  value       = google_container_cluster.primary.master_auth.0.cluster_ca_certificate
  sensitive   = true
}

output "gke_get_credentials_command" {
  description = "Command to configure kubectl"
  value       = "gcloud container clusters get-credentials ${google_container_cluster.primary.name} --zone ${google_container_cluster.primary.location} --project ${google_project.experiment01.project_id}"
}
