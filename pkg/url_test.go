package pkg

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsReadmePathURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{
			url:  "README.md",
			want: false,
		},
		{
			url:  "https://github.com/kevincobain2000/re",
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewURLHandler().IsRemotePath(test.url)
			assert.Equal(t, test.want, got)

		})
	}
}
func TestGithubBranchedURLToRawReadmeURL(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			url:  "https://github.com/kevincobain2000/re",
			want: "https://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
		{
			url:  "https://github.com/kevincobain2000/re/",
			want: "https://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
		{
			url:  "http://github.com/kevincobain2000/re/",
			want: "http://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewURLHandler().githubBranchedURLToRawReadmeURL(test.url, "master")
			assert.Equal(t, test.want, got)

		})
	}
}
func TestGithubBlobURLToRawReadmeURL(t *testing.T) {
	tests := []struct {
		url             string
		want            string
		isFullReadmeURL bool
	}{
		{
			url:             "https://github.com/kevincobain2000/re",
			want:            "",
			isFullReadmeURL: false,
		},
		{
			url:             "https://github.com/kevincobain2000/gobrew/blob/feature/cache/README.md",
			want:            "https://raw.githubusercontent.com/kevincobain2000/gobrew/feature/cache/README.md",
			isFullReadmeURL: true,
		},
		{
			url:             "https://github.com/kevincobain2000/gobrew/blob/master/README.md",
			want:            "https://raw.githubusercontent.com/kevincobain2000/gobrew/master/README.md",
			isFullReadmeURL: true,
		},
		{
			url:             "https://ghe.company.com/kevincobain2000/re",
			want:            "",
			isFullReadmeURL: false,
		},
		{
			url:             "https://ghe.company.com/kevincobain2000/gobrew/blob/feature/cache/README.md",
			want:            "https://raw.ghe.company.com/kevincobain2000/gobrew/feature/cache/README.md",
			isFullReadmeURL: true,
		},
		{
			url:             "https://ghe.company.com/kevincobain2000/gobrew/blob/master/README.md",
			want:            "https://raw.ghe.company.com/kevincobain2000/gobrew/master/README.md",
			isFullReadmeURL: true,
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			isFullReadmeURL := NewURLHandler().isFullReadmeURL(test.url)
			assert.Equal(t, test.isFullReadmeURL, isFullReadmeURL)
			if isFullReadmeURL {
				got := NewURLHandler().githubBlobURLToRawReadmeURL(test.url)
				assert.Equal(t, test.want, got)
			} else {
				assert.Equal(t, test.want, "") // basically skip this assertion
			}

		})
	}
}

func TestExtractDomainFromURL(t *testing.T) {
	tests := []struct {
		url string
		got string
	}{
		{
			url: "https://github.com/kevincobain2000/re",
			got: "github.com",
		},
		{
			url: "https://github.com/kevincobain2000/re/",
			got: "github.com",
		},
		{
			url: "http://github.com/kevincobain2000/re/",
			got: "github.com",
		},
		{
			url: "http://ghe.company.com/kevincobain2000/re/",
			got: "ghe.company.com",
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewURLHandler().extractDomainFromURL(test.url)
			assert.Equal(t, test.got, got)

		})
	}
}
func TestGithubURLToRawURL(t *testing.T) {
	tests := []struct {
		url string
		got string
	}{
		{
			url: "https://github.com/kevincobain2000/re",
			got: "https://raw.githubusercontent.com/kevincobain2000/re",
		},
		{
			url: "https://github.com/kevincobain2000/re/",
			got: "https://raw.githubusercontent.com/kevincobain2000/re/",
		},
		{
			url: "http://github.com/kevincobain2000/re/",
			got: "http://raw.githubusercontent.com/kevincobain2000/re/",
		},
		{
			url: "https://ghe.company.com/kevincobain2000/re/",
			got: "https://raw.ghe.company.com/kevincobain2000/re/",
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewURLHandler().githubURLToRawURL(test.url)
			assert.Equal(t, test.got, got)

		})
	}
}
