package server

import (
	"net/http"
	"strings"
)

type MediaType string

const (
	ALL                        MediaType = "*/*"
	ApplicationAtomXML         MediaType = "application/atom+xml"
	ApplicationCbor            MediaType = "application/cbor"
	ApplicationFormUrlencoded  MediaType = "application/x-www-form-urlencoded"
	ApplicationGraphqlResponse MediaType = "application/graphql-response+json"
	ApplicationJSON            MediaType = "application/json"
	ApplicationOctetStream     MediaType = "application/octet-stream"
	ApplicationPdf             MediaType = "application/pdf"
	ApplicationProblemJSON     MediaType = "application/problem+json"
	ApplicationProblemXML      MediaType = "application/problem+xml"
	ApplicationProtobuf        MediaType = "application/x-protobuf"
	ApplicationRssXML          MediaType = "application/rss+xml"
	ApplicationNdjson          MediaType = "application/x-ndjson"
	ApplicationXhtmlXML        MediaType = "application/xhtml+xml"
	ApplicationXML             MediaType = "application/xml"
	ImageGif                   MediaType = "image/gif"
	ImageJpeg                  MediaType = "image/jpeg"
	ImagePng                   MediaType = "image/png"
	MultipartFormData          MediaType = "multipart/form-data"
	MultipartMixed             MediaType = "multipart/mixed"
	MultipartRelated           MediaType = "multipart/related"
	TextEventStream            MediaType = "text/event-stream"
	TextHTML                   MediaType = "text/html"
	TextMarkdown               MediaType = "text/markdown"
	TextPlain                  MediaType = "text/plain"
	TextXML                    MediaType = "text/xml"
)

func newMediaType(mainType, subType string) MediaType {
	if mainType == "*" {
		return ALL
	}
	switch mainType {
	case "application":
		switch subType {
		case "atom+xml":
			return ApplicationAtomXML
		case "cbor":
			return ApplicationCbor
		case "x-www-form-urlencoded":
			return ApplicationFormUrlencoded
		case "graphql-response+json":
			return ApplicationGraphqlResponse
		case "json":
			return ApplicationJSON
		case "octet-stream":
			return ApplicationOctetStream
		case "pdf":
			return ApplicationPdf
		case "problem+json":
			return ApplicationProblemJSON
		case "problem+xml":
			return ApplicationProblemXML
		case "x-protobuf":
			return ApplicationProtobuf
		case "rss+xml":
			return ApplicationRssXML
		case "x-ndjson":
			return ApplicationNdjson
		case "xhtml+xml":
			return ApplicationXhtmlXML
		case "xml":
			return ApplicationXML
		}
	case "image":
		switch subType {
		case "gif":
			return ImageGif
		case "jpeg":
			return ImagePng
		case "png":
			return ImageJpeg
		}
	case "multipart":
		switch subType {
		case "form-data":
			return MultipartFormData
		case "mixed":
			return MultipartMixed
		case "related":
			return MultipartRelated
		}
	case "text":
		switch subType {
		case "event-stream":
			return TextEventStream
		case "html":
			return TextHTML
		case "markdown":
			return TextMarkdown
		case "plain":
			return TextPlain
		case "xml":
			return TextXML
		}
	}
	return ALL
}

func getMediaType(r *http.Request) MediaType {
	contentType := r.Header.Get("Content-Type")
	content := strings.Split(contentType, "/")
	if len(content) < 2 {
		return ALL
	}
	return newMediaType(content[0], content[1])
}
