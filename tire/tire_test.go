package tire

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestTire(t *testing.T) {

	searcher := NewSearcher("../dics/main.dic","../dics/surname.dic")

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		keyword := request.URL.Query().Get("keyword")
		search := searcher.Search(keyword)
		marshal, err := json.Marshal(search)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write(marshal)
		}
	})

	http.HandleFunc("/analysis", func(writer http.ResponseWriter, request *http.Request) {
		keyword := request.URL.Query().Get("keyword")
		search := searcher.Analysis(keyword)
		marshal, err := json.Marshal(search)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write(marshal)
		}
	})

	http.ListenAndServe(":2560", nil)

}
