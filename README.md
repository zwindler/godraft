# go draft

godraft is a simple webapp to help players during Magic the Gathering(tm) drafts

## test code

```
make dev
```

## compile code in a binary

```
VERSION=x.y.z make build
```

## run binary stored in bin/

```
make run
```

## docker build + push new image

```
TAG=x.y.z make dockerpush
```

## run from a docker image

```
TAG=x.y.z make dockerrun
```

## Use goreleaser

See [goreleaser.com/install](https://goreleaser.com/install/) and [goreleaser.com/quick-start](https://goreleaser.com/quick-start/)

```
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo apt install goreleaser

goreleaser init
  • Generating .goreleaser.yaml file
  • config created; please edit accordingly to your needs file=.goreleaser.yaml
```

Create a new Github token (https://github.com/settings/tokens/new) with ** write:packages** permissions

Each time you want to release, run 

```
git tag -a 2.1.0 -m "2.1.0 release"
git push origin 2.1.0
export GITHUB_TOKEN="ghx_xxxxxxxxxxxxxxxxxx"
goreleaser release
```
