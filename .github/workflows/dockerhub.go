name: Dockerhub Image CI

on: workflow_dispatch

jobs:

  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v2
      - name: Prepare
        id: prep
        run: |
          DOCKER_IMAGE=${{ secrets.DOCKER_USER }}/vk-callbackapi-to-telegram
          BUILD_VERSION=latest
          SHORTREF=${GITHUB_SHA::8}
          # If this is git tag, use the tag name as a docker tag
          if [[ $GITHUB_REF == refs/tags/* ]]; then
            BUILD_VERSION=${GITHUB_REF#refs/tags/v}
          fi
          TAGS="${DOCKER_IMAGE}:${VERSION},${DOCKER_IMAGE}:${SHORTREF}"
          # If the BUILD_VERSION looks like a version number, assume that
          # this is the most recent version of the image and also
          # tag it 'latest'.
          if [[ $BUILD_VERSION =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
            TAGS="$TAGS,${DOCKER_IMAGE}:latest"
          fi
          # Set output parameters.
          echo ::set-output name=tags::${TAGS}
          echo ::set-output name=docker_image::${DOCKER_IMAGE}
          echo "BUILD_TIMESTAMP=$(date +%Y-%m-%dT%H:%M:%S)" >> $GITHUB_ENV
          echo "BUILD_VERSION=${GITHUB_SHA::6}" >> $GITHUB_ENV
      - name: Set up QEMU
        uses: docker/setup-qemu-action@master
        with:
          platforms: all
      - name: Set up Docker Buildx
        id: buildx
        uses: docker/setup-buildx-action@master
      - name: Login to DockerHub
        if: github.event_name != 'pull_request'
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_USER }}
          password: ${{ secrets.DOCKER_PASSWORD }}
      - name: Cache Docker layers
        uses: actions/cache@v2
        with:
          path: /tmp/.buildx-cache
          key: ${{ runner.os }}-buildx-${{ github.sha }}
          restore-keys: ${{ runner.os }}-buildx-
      - name: Build
        uses: docker/build-push-action@v2
        with:
          builder: ${{ steps.buildx.outputs.name }}
          context: .
          file: ./Dockerfile
          platforms: linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64
          push: true
          tags: ${{ steps.prep.outputs.tags }}
          build-args: |
            BUILD_VERSION=${{ env.BUILD_VERSION }}
            BUILD_TIMESTAMP=${{ env.BUILD_TIMESTAMP }}
