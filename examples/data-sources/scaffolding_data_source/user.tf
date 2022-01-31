data "mattermost_user" "example" {
  username = "amz"
}

output "email" {
  value = data.mattermost_user.example.email # amz@octo.com
}
