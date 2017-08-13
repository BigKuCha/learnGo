package webcrawler

type PageDownloader interface {
	Id() uint32                               //获取ID
	Download(req *Request) (*Response, error) //下载网页，返回请求
}

type PageDownloaderPool interface {
	Take() (PageDownloader, error)  //从池中获取一个网页下载器
	Return(dl PageDownloader) error //把一个网页下载器还给池子
	Total() uint32                  //获取池的总容量
	Used() uint32                   //获取正在被使用的网页下载器的数量
}
