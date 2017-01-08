package webcrawler

const (
	CrawlerDownloadError    ErrorType = "Downloader Error"
	CrawlerAnalyzerError    ErrorType = "Analyzer Error"
	CrawlerItemProcessError ErrorType = "Item Process Error"
)

//条目
type Item map[string]interface{}

//数据接口
type Data interface {
	Valid() bool //数据是否有效
}
