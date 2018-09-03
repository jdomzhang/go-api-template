package enc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

type Values map[string]string

var expXML = regexp.MustCompile(`<(?P<key>\w+)>(<\!\[CDATA\[)?(?P<value>[^\<\>]*?)(\]\]>)?<\/\w+>`)
var expXMLGroupNames = expXML.SubexpNames()

// ParseXML will ignore the first <xml> node
func (obj Values) ParseXML(xml string) (Values, error) {
	matches := expXML.FindAllStringSubmatch(xml, -1)
	if matches == nil {
		return obj, fmt.Errorf("Not a valid xml")
	}

	nv := map[string]string{}

	for _, m := range matches {
		for i, v := range m {
			if expXMLGroupNames[i] != "" {
				nv[expXMLGroupNames[i]] = v
			}
		}
		obj.Add(nv["key"], nv["value"])
	}

	return obj, nil
}

// ParseJSON will add all key value pairs
func (obj Values) ParseJSON(jsonStr string) (Values, error) {
	kv := map[string]interface{}{}
	if err := json.Unmarshal([]byte(jsonStr), &kv); err != nil {
		return obj, err
	}

	for k, v := range kv {
		obj.Add(k, fmt.Sprintf("%v", v))
	}

	return obj, nil
}

// ParseObject will add all fields and values
func (obj Values) ParseObject(object interface{}) (Values, error) {
	rawData, err := json.Marshal(object)
	if err != nil {
		return obj, err
	}

	return obj.ParseJSON(string(rawData))
}

// ParseQuery will parse url query
func (obj Values) ParseQuery(query string) (Values, error) {
	var err error
	for query != "" {
		key := query
		if i := strings.Index(key, "?"); i >= 0 {
			key = key[i+1:]
		}
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		// obj[key] = append(obj[key], value)
		obj[key] = value
	}
	return obj, err
}

// Add key value pair
func (obj Values) Add(k string, v string) {
	obj[k] = v
}

// Get value by key
func (obj Values) Get(k string) string {
	return obj[k]
}

// Del one key value pair
func (obj Values) Del(k string) {
	delete(obj, k)
}

// String return string
func (obj Values) String() string {
	return string(obj.Bytes())
}

// Bytes returns []byte
func (obj Values) Bytes() []byte {
	var buf bytes.Buffer
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := obj[k]
		if v == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k + "=")
		buf.WriteString(v)
	}
	return buf.Bytes()
}

// XML returns xml format
func (obj Values) XML() string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf.WriteString("<xml>")
	for _, k := range keys {
		v := obj[k]
		if v == "" {
			continue
		}
		buf.WriteString("<")
		buf.WriteString(k)
		buf.WriteString(">")

		buf.WriteString(v)

		buf.WriteString("</")
		buf.WriteString(k)
		buf.WriteString(">")
	}
	buf.WriteString("</xml>")
	return buf.String()
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (obj Values) Encode() string {
	if obj == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := obj[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
	}
	return buf.String()
}
