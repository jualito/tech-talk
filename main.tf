provider "mattermost" {
  url = "https://dev.mattermost.octo.tools"
}

data "mattermost_user" "this" {
  username = "mattermost"
}

output "result" {
  value = data.mattermost_user.this
}
