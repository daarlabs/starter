docker_compose('docker-compose.yaml')

docker_build('starter-example', '.',
	build_args = { "NAME": "example", "PORT": "8000" },
  live_update = [
    sync('.', '/app/example'),
  ],
)