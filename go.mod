module github.com/moby/docker-image-spec

go 1.18

require github.com/opencontainers/image-spec v1.0.2

require github.com/opencontainers/go-digest v1.0.0 // indirect

retract v1.3.0 // Package github.com/moby/docker-image-spec/specs-go has the wrong package name.
