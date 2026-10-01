package v1

const (
	// MediaTypeImageIndex specifies the media type for a Docker manifest list.
	//
	// OCI images use "application/vnd.oci.image.index.v1+json" for the
	// corresponding OCI image index format.
	MediaTypeImageIndex = "application/vnd.docker.distribution.manifest.list.v2+json"

	// MediaTypeImageManifest specifies the media type for a Docker image manifest.
	//
	// OCI images use "application/vnd.oci.image.manifest.v1+json" for the
	// corresponding OCI image manifest format.
	MediaTypeImageManifest = "application/vnd.docker.distribution.manifest.v2+json"

	// MediaTypeImageConfig specifies the media type for the image configuration.
	//
	// OCI images use "application/vnd.oci.image.config.v1+json" for the
	// corresponding OCI image configuration format.
	MediaTypeImageConfig = "application/vnd.docker.container.image.v1+json"
)

const (
	// MediaTypeImageLayer is the media type used for uncompressed layers
	// referenced by the manifest.
	//
	// OCI images use "application/vnd.oci.image.layer.v1.tar" for the
	// corresponding OCI image layer format.
	MediaTypeImageLayer = "application/vnd.docker.image.rootfs.diff.tar"

	// MediaTypeImageLayerGzip is the media type used for gzipped layers
	// referenced by the manifest.
	//
	// This media type is interchangeable with "application/vnd.oci.image.layer.v1.tar+gzip".
	MediaTypeImageLayerGzip = "application/vnd.docker.image.rootfs.diff.tar.gzip"
)

// Non-distributable layer media types.
//
// Deprecated: Non-distributable layers are deprecated, and not recommended
// for future use. Implementations SHOULD NOT produce new non-distributable
// layers.
const (
	// MediaTypeImageLayerNonDistributable is the media type for uncompressed
	// layers referenced by the manifest but with distribution restrictions.
	//
	// Deprecated: Non-distributable layers are deprecated, and not recommended
	// for future use. Implementations SHOULD NOT produce new non-distributable
	// layers.
	MediaTypeImageLayerNonDistributable = "application/vnd.docker.image.rootfs.foreign.diff.tar"

	// MediaTypeImageLayerNonDistributableGzip is the media type for gzipped
	// layers referenced by the manifest but with distribution restrictions.
	//
	// Deprecated: Non-distributable layers are deprecated, and not recommended
	// for future use. Implementations SHOULD NOT produce new non-distributable
	// layers.
	MediaTypeImageLayerNonDistributableGzip = "application/vnd.docker.image.rootfs.foreign.diff.tar.gzip"
)

// DockerOCIImageMediaType specifies the media type for the Docker image
// configuration.
//
// Deprecated: use [MediaTypeImageConfig].
const DockerOCIImageMediaType = MediaTypeImageConfig
