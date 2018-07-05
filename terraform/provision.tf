provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "ca-central-1"
}

output "master-url" {
  value = "https://${module.openshift-cluster.master-public_ip}.xip.io:8443"
}

output "master-public_dns" {
  value = "${module.openshift-cluster.master-public_dns}"
}

output "master-public_ip" {
  value = "${module.openshift-cluster.master-public_ip}"
}

output "bastion-public_dns" {
  value = "${module.openshift-cluster.bastion-public_dns}"
}

output "bastion-public_ip" {
  value = "${module.openshift-cluster.bastion-public_ip}"
}

output "master-private_dns" {
  value = "${module.openshift-cluster.master-private_dns}"
}

output "master-private_ip" {
  value = "${module.openshift-cluster.master-private_ip}"
}

output "node1-private_dns" {
  value = "${module.openshift-cluster.node1-private_dns}"
}

output "node1-private_ip" {
  value = "${module.openshift-cluster.node1-private_ip}"
}

output "node2-private_dns" {
  value = "${module.openshift-cluster.node2-private_dns}"
}

output "node2-private_ip" {
  value = "${module.openshift-cluster.node2-private_ip}"
}
