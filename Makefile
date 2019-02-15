BASE_IMAGE_URL := scottcrawford03/bragsheet
GIT_COMMIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null)
IMAGE_URL := $(BASE_IMAGE_URL):$(GIT_COMMIT_SHA)

docker-build:
	docker build --pull -t ${IMAGE_URL} .

docker-push: docker-build
	docker tag "${IMAGE_URL}" "${BASE_IMAGE_URL}:latestapi"
	docker push "${IMAGE_URL}"

run-local:
	go run ./main.go

run-prod-migrations:
	ENVIRONMENT=production ./shmig.sh up
