output "postgres_namespace" {
  description = "The namespace where PostgreSQL is deployed"
  value       = kubernetes_namespace.postgres.metadata[0].name
}

output "postgres_service_name" {
  description = "The name of the PostgreSQL service"
  value       = kubernetes_service.postgres.metadata[0].name
}

output "connection_string" {
  description = "PostgreSQL connection information"
  value       = "To connect to PostgreSQL, use: kubectl port-forward -n ${kubernetes_namespace.postgres.metadata[0].name} svc/${kubernetes_service.postgres.metadata[0].name} 5432:5432"
  sensitive   = false
}
