package utils

import "strings"

func VerifyHtml(content string) string {
	content = strings.Replace(content,"{{","&123456;",-1)
	content = strings.Replace(content,"}}","&654321;",-1)
	content = strings.Replace(content, "<script>", "&lt;script&gt;", -1)
	content = strings.Replace(content, "</script>", "&lt;/script&gt;", -1)
	content = strings.Replace(content, "<css>", "&lt;script&gt;", -1)
	content = strings.Replace(content, "</css>", "&lt;/script&gt;", -1)
	content = strings.Replace(content,"\"","", -1)
	/*content = strings.Replace(content, "<a", "<lianjie", -1)
	content = strings.Replace(content, "</a>", "</lianjie>", -1)*/
	return content
}
