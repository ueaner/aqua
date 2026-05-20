package config

import "os"

var (
	GitHubReleaseURLTemplate  string
	GitHubArchiveURLTemplate  string
	GitHubArchiveURLTemplate2 string
	GitHubCosignURLTemplate   string
)

func init() {
	// 在这里赋值，可以确保如果同包内其他 init() 修改了环境变量，能拿到最新值
	proxy := os.Getenv("GITHUB_PROXY")

	GitHubReleaseURLTemplate = proxy + "https://github.com/%s/%s/releases/download/%s/%s"
	GitHubArchiveURLTemplate = proxy + "https://github.com/%s/%s/archive/refs/tags/%s.tar.gz"
	GitHubArchiveURLTemplate2 = proxy + "https://github.com/%s/%s/archive/%s.tar.gz"
	GitHubCosignURLTemplate = proxy + "https://github.com/%s/%s/releases/download/{{.Version}}/"
}
