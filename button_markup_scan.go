package gotdbot

import "strings"

func readButtonMarkdownLink(text string, start int) (string, string, int, bool) {
	depth, middle := 1, -1
	for pos := start + 1; pos < len(text); pos++ {
		switch text[pos] {
		case '\\':
			pos++
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				middle = pos
			}
		}
		if middle >= 0 {
			break
		}
	}
	if middle < 0 || middle+1 >= len(text) || text[middle+1] != '(' {
		return "", "", 0, false
	}
	depth = 1
	for pos := middle + 2; pos < len(text); pos++ {
		switch text[pos] {
		case '\\':
			pos++
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return text[start+1 : middle], text[middle+2 : pos], pos + 1, true
			}
		}
	}
	return "", "", 0, false
}

func unescapeButtonMarkdown(text string) string {
	var out strings.Builder
	for pos := 0; pos < len(text); pos++ {
		if text[pos] == '\\' && pos+1 < len(text) && text[pos+1] >= '!' && text[pos+1] <= '~' {
			pos++
		}
		out.WriteByte(text[pos])
	}
	return out.String()
}

func buttonMarkdownCodeEnd(text string, start int) int {
	end := start
	for end < len(text) && text[end] == '`' {
		end++
	}
	delimiter := text[start:end]
	for pos := end; pos < len(text); pos++ {
		if text[pos] == '\\' {
			pos++
			continue
		}
		if strings.HasPrefix(text[pos:], delimiter) {
			return pos + len(delimiter)
		}
	}
	return len(text)
}

type buttonHTMLTag struct {
	name    string
	closing bool
	end     int
}

func readButtonHTMLTag(text string, start int) (buttonHTMLTag, bool) {
	var tag buttonHTMLTag
	pos := start + 1
	if pos < len(text) && text[pos] == '/' {
		tag.closing = true
		pos++
	}
	nameStart := pos
	for pos < len(text) && ((text[pos] >= 'a' && text[pos] <= 'z') || (text[pos] >= 'A' && text[pos] <= 'Z') || text[pos] == '-') {
		pos++
	}
	if pos == nameStart {
		return tag, false
	}
	tag.name = strings.ToLower(text[nameStart:pos])
	for pos < len(text) {
		for pos < len(text) && strings.ContainsRune(" \t\r\n/", rune(text[pos])) {
			pos++
		}
		if pos == len(text) {
			return tag, false
		}
		if text[pos] == '>' {
			tag.end = pos + 1
			return tag, true
		}
		attrStart := pos
		for pos < len(text) && !strings.ContainsRune(" \t\r\n=><", rune(text[pos])) {
			pos++
		}
		if attrStart == pos {
			return tag, false
		}
		for pos < len(text) && strings.ContainsRune(" \t\r\n", rune(text[pos])) {
			pos++
		}
		if pos < len(text) && text[pos] == '=' {
			pos++
			for pos < len(text) && strings.ContainsRune(" \t\r\n", rune(text[pos])) {
				pos++
			}
			if pos == len(text) {
				return tag, false
			}
			if text[pos] == '\'' || text[pos] == '"' {
				quote := text[pos]
				pos++
				for pos < len(text) && text[pos] != quote {
					pos++
				}
				if pos == len(text) {
					return tag, false
				}
				pos++
			} else {
				for pos < len(text) && !strings.ContainsRune(" \t\r\n><", rune(text[pos])) {
					pos++
				}
			}
		}
	}
	return tag, false
}

func buttonHTMLElementEnd(text string, tag buttonHTMLTag) int {
	depth := 1
	for pos := tag.end; pos < len(text); {
		index := strings.IndexByte(text[pos:], '<')
		if index < 0 {
			break
		}
		pos += index
		if next, ok := readButtonHTMLTag(text, pos); ok {
			if next.name == tag.name {
				if next.closing {
					depth--
				} else {
					depth++
				}
				if depth == 0 {
					return next.end
				}
			}
			pos = next.end
		} else {
			pos++
		}
	}
	return len(text)
}
