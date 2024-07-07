package mime

const (
	// CompressedTar represents MIME for compressed tar file.
	CompressedTar = "application/x-compressed-tar"
	// Gzip represents MIME for gzip file.
	Gzip = "application/gzip"
	// OctedStream represents MIME for binary file.
	OctedStream = "application/octet-stream"
	// Tar represents MIME for tar file.
	Tar = "application/x-tar"
	// Xz represents MIME for xz file.
	Xz = "application/x-xz"
	// Zip represents MIME for zip file.
	Zip = "application/zip"
)

// compressed is a list of MIME type of compressed file.
// See https://en.wikipedia.org/wiki/List_of_archive_formats more details.
var compressed = []MIME{
	"application/gzip",
	"application/java-archive",
	"application/vnd.android.package-archive",
	"application/vnd.genozip",
	"application/vnd.ms-cab-compressed",
	"application/x-7z-compressed",
	"application/x-ace-compressed",
	"application/x-alz-compressed",
	"application/x-apple-diskimage",
	"application/x-arj",
	"application/x-astrotite-afa",
	"application/x-b1",
	"application/x-brotli",
	"application/x-bzip2",
	"application/x-cfs-compressed",
	"application/x-compress",
	"application/x-dar",
	"application/x-dgc-compressed",
	"application/x-freearc",
	"application/x-gca-compressed",
	"application/x-gtar",
	"application/x-lzh",
	"application/x-lzip",
	"application/x-lzma",
	"application/x-lzop",
	"application/x-lzx",
	"application/x-ms-wim",
	"application/x-rar-compressed",
	"application/x-snappy-framed",
	"application/x-stuffit",
	"application/x-stuffitx",
	"application/x-xar",
	"application/x-xz",
	"application/x-zoo",
	"application/zip",
	"application/zstd",
}
