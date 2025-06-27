terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
  required_version = ">= 1.0"
}

provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

resource "google_project" "experiment01" {
  name            = "experiment01"
  project_id      = var.project_id
  billing_account = var.billing_account_id
  
  labels = {
    environment = "experiment"
    created-by  = "opentofu"
  }
}

resource "google_project_service" "compute_api" {
  project = google_project.experiment01.project_id
  service = "compute.googleapis.com"
  
  disable_dependent_services = true
}

resource "google_project_service" "container_api" {
  project = google_project.experiment01.project_id
  service = "container.googleapis.com"
  
  disable_dependent_services = true
  depends_on = [google_project_service.compute_api]
}

resource "google_compute_network" "main_vpc" {
  name                    = "main-vpc"
  project                 = google_project.experiment01.project_id
  auto_create_subnetworks = false
  mtu                     = 1460
  
  depends_on = [google_project_service.compute_api]
}

resource "google_compute_subnetwork" "main_subnet" {
  name          = "main-subnet"
  project       = google_project.experiment01.project_id
  ip_cidr_range = "10.0.1.0/24"
  region        = var.region
  network       = google_compute_network.main_vpc.id
  
  secondary_ip_range {
    range_name    = "k8s-pod-range"
    ip_cidr_range = "10.1.0.0/16"
  }
  
  secondary_ip_range {
    range_name    = "k8s-service-range"
    ip_cidr_range = "10.2.0.0/16"
  }
}

resource "google_compute_subnetwork" "private_subnet" {
  name          = "private-subnet"
  project       = google_project.experiment01.project_id
  ip_cidr_range = "10.0.2.0/24"
  region        = var.region
  network       = google_compute_network.main_vpc.id
}

resource "google_container_cluster" "primary" {
  name     = "experiment01-gke"
  project  = google_project.experiment01.project_id
  location = var.zone
  
  remove_default_node_pool = true
  initial_node_count       = 1
  
  network    = google_compute_network.main_vpc.name
  subnetwork = google_compute_subnetwork.main_subnet.name
  
  ip_allocation_policy {
    cluster_secondary_range_name  = "k8s-pod-range"
    services_secondary_range_name = "k8s-service-range"
  }
  
  network_policy {
    enabled = true
  }
  
  workload_identity_config {
    workload_pool = "${google_project.experiment01.project_id}.svc.id.goog"
  }
  
  depends_on = [
    google_project_service.container_api,
    google_compute_subnetwork.main_subnet
  ]
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "main-node-pool"
  project    = google_project.experiment01.project_id
  location   = var.zone
  cluster    = google_container_cluster.primary.name
  node_count = 2
  
  node_config {
    preemptible  = true
    machine_type = "e2-medium"
    
    service_account = google_service_account.gke_service_account.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
    
    labels = {
      environment = "experiment"
    }
    
    tags = ["gke-node", "experiment01-gke"]
    
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
  
  management {
    auto_repair  = true
    auto_upgrade = true
  }
  
  depends_on = [google_container_cluster.primary]
}

resource "google_service_account" "gke_service_account" {
  account_id   = "gke-service-account"
  project      = google_project.experiment01.project_id
  display_name = "GKE Service Account"
  
  depends_on = [google_project_service.compute_api]
}

resource "google_project_iam_member" "gke_service_account_roles" {
  for_each = toset([
    "roles/logging.logWriter",
    "roles/monitoring.metricWriter",
    "roles/monitoring.viewer",
    "roles/stackdriver.resourceMetadata.writer"
  ])
  
  project = google_project.experiment01.project_id
  role    = each.value
  member  = "serviceAccount:${google_service_account.gke_service_account.email}"
  
  depends_on = [google_service_account.gke_service_account]
}

resource "google_compute_firewall" "allow_internal" {
  name    = "allow-internal"
  project = google_project.experiment01.project_id
  network = google_compute_network.main_vpc.name
  
  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }
  
  allow {
    protocol = "udp"
    ports    = ["0-65535"]
  }
  
  allow {
    protocol = "icmp"
  }
  
  source_ranges = ["10.0.0.0/8"]
  
  depends_on = [google_compute_network.main_vpc]
}

resource "google_compute_firewall" "allow_ssh" {
  name    = "allow-ssh"
  project = google_project.experiment01.project_id
  network = google_compute_network.main_vpc.name
  
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
  
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["ssh-allowed"]
  
  depends_on = [google_compute_network.main_vpc]
}
