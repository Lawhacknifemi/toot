package toot

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPageInfo_NoPages(t *testing.T) {

	header := make(http.Header)
	PageInfo{}.SetHeader(header, "/api/v1/notifications")

	if got := header.Get("Link"); got != "" {
		t.Errorf("expected no Link header, got %q", got)
	}
}

func TestPageInfo_NextOnly(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=ZZZ>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestPageInfo_PreviousOnly(t *testing.T) {

	header := make(http.Header)
	PageInfo{MinID: "AAA"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?min_id=AAA>; rel="prev"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// Both links must survive. A "Set" per link would leave only the last one.
func TestPageInfo_BothPagesSurvive(t *testing.T) {

	header := make(http.Header)
	PageInfo{MinID: "AAA", MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=ZZZ>; rel="next", </api/v1/notifications?min_id=AAA>; rel="prev"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}

	if count := len(header.Values("Link")); count != 1 {
		t.Errorf("expected exactly 1 Link field, got %d", count)
	}
}

func TestPageInfo_OffsetIsNextPage(t *testing.T) {

	header := make(http.Header)
	PageInfo{Offset: "40"}.SetHeader(header, "/api/v1/directory")

	expected := `</api/v1/directory?offset=40>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// MaxID wins over Offset, because a cursor is more precise than a count.
func TestPageInfo_MaxIDBeatsOffset(t *testing.T) {

	page := PageInfo{MaxID: "ZZZ", Offset: "40"}

	if got := page.GetNextPage("/x"); got != "/x?max_id=ZZZ" {
		t.Errorf("expected max_id to win, got %q", got)
	}
}

// The base URL supplies the path only; each link writes its own query string.
func TestPageInfo_QueryStringIsReplaced(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications?max_id=OLD&limit=20")

	expected := `</api/v1/notifications?max_id=ZZZ>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// An ID carrying URL metacharacters must not break out of the link.
func TestPageInfo_IDsAreEscaped(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: `a b>"&c`}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=a+b%3E%22%26c>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestPageInfo_NilHeader(_ *testing.T) {
	PageInfo{MaxID: "ZZZ"}.SetHeader(nil, "/x") // must not panic
}

func TestPageInfo_SetHeadersNilGuards(_ *testing.T) {

	PageInfo{MaxID: "ZZZ"}.SetHeaders(nil)
	PageInfo{MaxID: "ZZZ"}.SetHeaders(&http.Response{Header: make(http.Header)})
	PageInfo{MaxID: "ZZZ"}.SetHeaders(&http.Response{
		Header:  make(http.Header),
		Request: &http.Request{},
	})
}

func TestPageInfo_SetHeadersFromClientResponse(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "https://example.com/api/v1/notifications?limit=20", nil)
	response := &http.Response{Header: make(http.Header), Request: request}

	PageInfo{MinID: "AAA", MaxID: "ZZZ"}.SetHeaders(response)

	expected := `<https://example.com/api/v1/notifications?max_id=ZZZ>; rel="next", ` +
		`<https://example.com/api/v1/notifications?min_id=AAA>; rel="prev"`

	if got := response.Header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
