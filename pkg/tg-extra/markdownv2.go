package tg_extra

import (
	"regexp"
	"strings"
)

func findAllIndex(str, pattern string) []int {
	indexList := []int{0}
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatchIndex(str, -1)
	for _, match := range matches {
		if match[2] != -1 {
			start := match[2]
			end := match[3]
			indexList = append(indexList, start, end)
		}
	}
	indexList = append(indexList, len(str))
	return indexList
}

func replaceAll(text, pattern string, function func(string) string) string {
	posList := findAllIndex(text, pattern)
	strList := make([]string, 0)
	originStr := make([]string, 0)
	for i := 1; i < len(posList)-1; i += 2 {
		start := posList[i]
		end := posList[i+1]
		strList = append(strList, function(text[start:end]))
	}
	for i := 0; i < len(posList); i += 2 {
		j := posList[i]
		k := posList[i+1]
		originStr = append(originStr, text[j:k])
	}
	if len(strList) < len(originStr) {
		strList = append(strList, "")
	} else {
		originStr = append(originStr, "")
	}
	newSlice := make([]string, 0, len(originStr)+len(strList))
	for i := range originStr {
		newSlice = append(newSlice, originStr[i], strList[i])
	}
	return strings.Join(newSlice, "")
}

func escapeShape(text string) string {
	parts := strings.Fields(text)
	return "▎*" + parts[1] + "*"
}

func escapeMinus(text string) string {
	return "\\" + text
}

func escapeBackquote(text string) string {
	return "\\`\\`"
}

func escapePlus(text string) string {
	return "\\" + text
}

func EscapeMarkdownV2(text string, flag int) string {
	text = strings.ReplaceAll(text, "\\[", "@->@")
	text = strings.ReplaceAll(text, "\\]", "@<-@")
	text = strings.ReplaceAll(text, "\\(", "@-->@")
	text = strings.ReplaceAll(text, "\\)", "@<--@")
	if flag == 1 {
		text = strings.ReplaceAll(text, "\\\\", "@@@")
	}
	text = strings.ReplaceAll(text, "\\", "\\\\")
	if flag == 1 {
		text = strings.ReplaceAll(text, "@@@", "\\\\")
	}
	text = strings.ReplaceAll(text, "_", "\\_")
	re := regexp.MustCompile(`\*{2}(.*?)\*{2}`)
	text = re.ReplaceAllString(text, "@@@$1@@@")
	text = strings.ReplaceAll(text, "\n\n*", "\n\n• ")
	text = strings.ReplaceAll(text, "*", "\\*")
	text = re.ReplaceAllString(text, "*$1*")
	re = regexp.MustCompile(`!?\[(.*?)\]\((.*?)\)`)
	text = re.ReplaceAllString(text, "@@@$1@@@^^^$2^^^")
	text = strings.ReplaceAll(text, "[", "\\[")
	text = strings.ReplaceAll(text, "]", "\\]")
	text = strings.ReplaceAll(text, "(", "\\(")
	text = strings.ReplaceAll(text, ")", "\\)")
	text = strings.ReplaceAll(text, "@->@", "[")
	text = strings.ReplaceAll(text, "@<-@", "]")
	text = strings.ReplaceAll(text, "@-->", "(")
	text = strings.ReplaceAll(text, "@<--@", ")")
	text = strings.ReplaceAll(text, "@@@$1@@@^^^$2^^^", "[$1]($2)")
	text = strings.ReplaceAll(text, "~", "\\~")
	text = strings.ReplaceAll(text, ">", "\\>")
	text = replaceAll(text, "(^#+\\s.+?$)|```[\\D\\d\\s]+?```", escapeShape)
	text = strings.ReplaceAll(text, "#", "\\#")
	text = replaceAll(text, "(\\+)|\n[\\s]*-\\s|```[\\D\\d\\s]+?```|`[\\D\\d\\s]*?`", escapePlus)
	text = strings.ReplaceAll(text, "\n\n-", "\n\n• ")
	text = strings.ReplaceAll(text, "\n\n\\d{1,2}\\.\\s", "\n\n")
	text = replaceAll(text, "(-)|\n[\\s]*-\\s|```[\\D\\d\\s]+?```|`[\\D\\d\\s]*?`", escapeMinus)
	re = regexp.MustCompile("```([\\D\\d\\s]+?)```")
	text = re.ReplaceAllString(text, "@@@$1@@@")
	text = replaceAll(text, "(``)", escapeBackquote)
	re = regexp.MustCompile("@@@([\\D\\d\\s]+?)@@@")
	text = re.ReplaceAllString(text, "```$1```")
	text = strings.ReplaceAll(text, "=", "\\=")
	text = strings.ReplaceAll(text, "|", "\\|")
	text = strings.ReplaceAll(text, "{", "\\{")
	text = strings.ReplaceAll(text, "}", "\\}")
	text = strings.ReplaceAll(text, ".", "\\.")
	text = strings.ReplaceAll(text, "!", "\\!")
	return text
}
