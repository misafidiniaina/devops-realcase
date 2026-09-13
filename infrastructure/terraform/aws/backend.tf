terraform {
  backend "s3" {
    key          = "cloud-platform-lab/dev/terraform.tfstate"
    use_lockfile = true
  }
}
