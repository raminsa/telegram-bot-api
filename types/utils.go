package types

import (
	"io"
)

// Error is an error containing extra information returned by the Telegram API.
type Error struct {
	Code    int
	Message string
	ResponseParameters
}

// Chattable is any config type that can be sent.
type Chattable interface {
	Params() (Params, error)
	EndPoint() string
}

// Fileable is any config type that can be sent that includes a file.
type Fileable interface {
	Chattable
	Files() []RequestFile
}

// RequestFile represents a file associated with a field name.
type RequestFile struct {
	// The file field name.
	Name string
	// The file name.
	FileName string
	// The file data to include.
	Data RequestFileData
}

// RequestFileData represents the data to be used for a file.
type RequestFileData interface {
	// NeedsUpload shows if the file needs to be uploaded.
	NeedsUpload() bool

	// UploadData gets the file name and an `io.Reader` for the file to be uploaded. This
	// must only be called when the file needs to be uploaded.
	UploadData() (string, io.Reader, error)
	// SendData gets the file data to send when a file does not need to be uploaded. This
	// must only be called when the file does not need to be uploaded.
	SendData() string
}

// FileBytes contains information about a set of bytes to upload as a File.
type FileBytes struct {
	Name  string
	Bytes []byte
}

// FileReader contains information about a reader to upload as a File.
type FileReader struct {
	Name   string
	Reader io.Reader
}

// FilePath is a path to a local file.
type FilePath string

// FileURL is a URL to use as a file for a request.
type FileURL string

// FileID is an ID of a file already uploaded to Telegram.
type FileID string

// FileAttach is an internal file type used for processed media groups.
type FileAttach string

// Params represents a set of parameters that gets passed to a request.
type Params map[string]string
