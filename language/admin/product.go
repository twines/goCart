package languageAdmin

var (
	Product = map[string]string{}
)

func init() {
	Product["ProductName"] = "商品名称"
	Product["Keyword"] = "关键字"
	Product["Description"] = "商品描述"
	Product["CategoryId"] = "商品分类"
}
