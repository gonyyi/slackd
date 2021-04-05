package slackmsg

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	SLACK_ENDPOINT_MSG         = "https://slack.com/api/chat.postMessage"
	SLACK_ENDPOINT_FILE_UPLOAD = "https://slack.com/api/files.upload"
	TEST_ENDPOINT              = "http://localhost:8080"
)

// Slack only supports a single file upload at a time
func postFile(slackToken, channel, threadTS, dir, filename, fileTitle, comment string, doGzip bool) error {
	p := newPostFile()

	if doGzip {
		f, err := ioutil.ReadFile(path.Join(dir, filename))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		zw.Name = filename
		// println("zw.Name = ", filename)

		zw.ModTime = time.Now()
		if _, err := zw.Write(f); err != nil {
			return err
		}
		if err := zw.Close(); err != nil {
			return err
		}

		filename = filename + ".gz"
		if _, err := p.AddFileReader(filename, &buf); err != nil {
			return err
		}
	} else {
		f, err := os.Open(path.Join(dir, filename))
		if f != nil {
			defer f.Close()
		}
		if err != nil {
			return err
		}
		if _, err := p.AddFileReader(filename, f); err != nil {
			return err
		}
	}

	p.AddField("channels", channel)
	p.AddField("filename", filename)

	if threadTS != "" {
		p.AddField("thread_ts", threadTS)
	}
	if fileTitle != "" {
		p.AddField("title", fileTitle)
	}
	if comment != "" {
		p.AddField("initial_comment", comment)
	}

	p.AddHTTPHeader("authorization", "Bearer "+slackToken)

	var buf bytes.Buffer
	if _, err := p.Send("POST", SLACK_ENDPOINT_FILE_UPLOAD, &buf); err != nil { // SLACK_ENDPOINT_FILE_UPLOAD
		return err
	}

	p = nil
	return nil
}

// 12/22/2020, gonyi, wrote it for slack file upload.
//    unlike i thought, it seems like i cannot upload multiple fils to slack at once..?
//    maybe a bug on my end..

type postMultipartForm struct {
	buf *bytes.Buffer
	w   *multipart.Writer
	h   map[string][]string
}

func newPostFile() *postMultipartForm {
	m := postMultipartForm{
		buf: &bytes.Buffer{},
		h:   make(map[string][]string),
	}
	m.w = multipart.NewWriter(m.buf)
	return &m
}

func (m *postMultipartForm) AddField(key, value string) {
	m.w.WriteField(key, value)
}

func (m *postMultipartForm) AddFile(filename string) (n int64, err error) {
	if f, err := m.w.CreateFormFile("file", filename); err != nil {
		n = 0
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return 0, err
		}
		n, err = io.Copy(f, file)
		file.Close()
	}
	return
}

func (m *postMultipartForm) AddFileReader(filename string, ior io.Reader) (n int64, err error) {
	if f, err := m.w.CreateFormFile("file", filename); err != nil {
		n = 0
	} else {
		n, err = io.Copy(f, ior)
	}
	return
}

func (m *postMultipartForm) Reader() io.Reader {
	return m.buf
}

func (m *postMultipartForm) ContentType() string {
	return m.w.FormDataContentType()
}

func (m *postMultipartForm) SetHTTPHeader(key, value string) {
	m.h[key] = []string{value}
}

func (m *postMultipartForm) AddHTTPHeader(key, value string) {
	if v, ok := m.h[key]; ok {
		m.h[key] = append(v, value)
	} else {
		m.SetHTTPHeader(key, value)
	}
}

func (m *postMultipartForm) GetRequest(method, url string) (*http.Request, error) {
	if err := m.w.Close(); err != nil {
		return nil, err
	}
	return http.NewRequest(method, url, m.buf)
}

func (m *postMultipartForm) Send(method, url string, resp *bytes.Buffer) (received int64, err error) {
	req, err := m.GetRequest(method, url)
	if err != nil {
		return 0, err
	}
	// content type will be automatically prepared by multipart writer
	// and add authroization keyee if any

	req.Header.Set("Content-Type", m.ContentType())
	for k, v := range m.h {
		if len(v) > 1 {
			for _, v2 := range v {
				req.Header.Add(k, v2)
			}
		} else {
			req.Header.Set(k, v[0])
		}
	}

	// send the request
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	// out, err := ioutil.ReadAll(r.Body)
	if resp != nil {
		received, err = io.Copy(resp, r.Body)
	}

	if r != nil && r.Body != nil {
		r.Body.Close()
	}

	return received, err
}
