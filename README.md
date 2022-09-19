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
