package pkg

import (
	"net/url"
	"strings"
)

const (
	GITHUB_DOMAIN     = "github.com"
	GITHUB_RAW_DOMAIN = "raw.githubusercontent.com"
	README_MD         = "README.md"
	EXTENSION_MD      = ".md"
)

type URLHandler struct {
	defaultBranches []string
}

func NewURLHandler() *URLHandler {
	return &URLHandler{
		defaultBranches: []string{"main", "master", "develop"}, // in order of preference
	}
}

func (h *URLHandler) isFullReadmeURL(u string) bool {
	return strings.HasSuffix(u, EXTENSION_MD)
}

func (h *URLHandler) IsRemotePath(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func (h *URLHandler) removeBlobFromURL(u string) string {
	// remote /blob/ from u
	return strings.Replace(u, "/blob/", "/", 1)
}

func (h *URLHandler) githubBlobURLToRawReadmeURL(u string) string {
	u = h.removeBlobFromURL(u)
	return h.githubURLToRawURL(u)
}
func (h *URLHandler) githubBranchedURLToRawReadmeURL(u string, branch string) string {
	// if url doesn't end with / then append /
	if !strings.HasSuffix(u, "/") {
		u += "/"
	}

	u = h.githubURLToRawURL(u)
	u += branch + "/" + README_MD
	return u
}

func (h *URLHandler) githubURLToRawURL(u string) string {
	domain := h.extractDomainFromURL(u)
	// replace github.com with raw.githubusercontent.com
	if domain == GITHUB_DOMAIN {
		u = strings.Replace(u, domain, GITHUB_RAW_DOMAIN, 1)
		return u
	}

	// if not github.com then probably enterprise github
	// append raw. to domain
	// eg. https://ghe.company.com/kevincobain2000/re -> https://raw.ghe.company.com/kevincobain2000/re
	u = strings.Replace(u, domain, "raw."+domain, 1)
	return u
}

func (h *URLHandler) extractDomainFromURL(u string) string {
	// Parse the URL
	parsedURL, err := url.Parse(u)
	if err != nil {
		return ""
	}

	// Extract the domain
	domain := parsedURL.Hostname()
	return domain
}
