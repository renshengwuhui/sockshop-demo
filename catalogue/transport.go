package catalogue

// transport.go contains the binding from endpoints to a concrete transport.
// In our case we just use a REST-y HTTP transport.

import (
	"net/http"
	"strconv"
	"strings"
	rf "github.com/ServiceComb/go-chassis/server/restful"
)

// MakeHTTPHandler mounts the endpoints into a REST-y HTTP handler.
func (s *catalogueService) URLPatterns() []rf.Route {
	return []rf.Route{
		{http.MethodGet, "/catalogue", "ListEndpoint"},
		{http.MethodGet, "/catalogue/size", "CountEndpoint"},
		{http.MethodGet, "/catalogue/{id}", "GetEndpoint"},
		{http.MethodGet, "/tags", "TagsEndpoint"},
		{http.MethodGet, "/health", "HealthEndpoint"},
	}
}
func (s *catalogueService)ListEndpoint(b *rf.Context){
	listReq,err := decodeListRequest(b)
	if err!=nil{
		encodeError(err,b)
		return
	}
	req := listReq.(listRequest)
	socks, err := s.List(req.Tags, req.Order, req.PageNum, req.PageSize)
	if err!=nil{
		encodeError(err,b)
		return
	}
	listResp := listResponse{Socks: socks, Err: err}
	err = encodeListResponse(b,listResp)
	if err!=nil{
		encodeError(err,b)
		return
	}
}
func(s *catalogueService)CountEndpoint(b *rf.Context){
	countreq,err :=	decodeCountRequest(b)
	if err!=nil{
		encodeError(err,b)
		return
	}
	req := countreq.(countRequest)
	n, err := s.Count(req.Tags)
	if err!=nil{
		encodeError(err,b)
		return
	}
	countResp := countResponse{N: n, Err: err}
	err = encodeResponse(b,countResp)
	if err!=nil{
		encodeError(err,b)
		return
	}


}
func(s *catalogueService)GetEndpoint(b *rf.Context){
	getReq := decodeGetRequest(b)

	req := getReq.(getRequest)
	sock, err := s.Get(req.ID)
	if err!=nil{
		encodeError(err,b)
		return
	}
	getResp :=  getResponse{Sock: sock, Err: err}
	err = encodeGetResponse(b,getResp)
	if err!=nil{
		encodeError(err,b)
		return
	}

}
func(s *catalogueService)TagsEndpoint(b *rf.Context){
	_,err := decodeTagsRequest(b)
	if err!=nil{
		encodeError(err,b)
		return
	}
	tags, err := s.Tags()
	if err!=nil{
		encodeError(err,b)
		return
	}
	tagResp := tagsResponse{Tags: tags, Err: err}
	err = encodeResponse(b,tagResp)
	if err!=nil{
		encodeError(err,b)
		return
	}

}
func(s *catalogueService)HealthEndpoint(b *rf.Context){
	_,err := decodeHealthRequest(b)
	if err!=nil{
		encodeError(err,b)
		return
	}
	health := s.Health()
	healthResp := healthResponse{Health: health}
	err = encodeHealthResponse(b,healthResp)


}
func encodeError( err error, w *rf.Context) {
	code := http.StatusInternalServerError
	switch err {
	case ErrNotFound:
		code = http.StatusNotFound
	}
	body :=map[string]interface{}{
		"error":       err.Error(),
		"status_code": code,
		"status_text": http.StatusText(code),
	}
	w.WriteHeaderAndJSON(code,body,"application/json")
}

func decodeListRequest( r *rf.Context) (interface{}, error) {
	pageNum := 1
	if page := r.ReadRequest().FormValue("page"); page != "" {
		pageNum, _ = strconv.Atoi(page)
	}
	pageSize := 10
	if size := r.ReadRequest().FormValue("size"); size != "" {
		pageSize, _ = strconv.Atoi(size)
	}
	order := "id"
	if sort := r.ReadRequest().FormValue("sort"); sort != "" {
		order = strings.ToLower(sort)
	}
	tags := []string{}
	if tagsval := r.ReadRequest().FormValue("tags"); tagsval != "" {
		tags = strings.Split(tagsval, ",")
	}
	return listRequest{
		Tags:     tags,
		Order:    order,
		PageNum:  pageNum,
		PageSize: pageSize,
	}, nil
}

// encodeListResponse is distinct from the generic encodeResponse because our
// clients expect that we will encode the slice (array) of socks directly,
// without the wrapping response object.
func encodeListResponse(w *rf.Context, response interface{}) error {
	resp := response.(listResponse)
	return encodeResponse(w, resp.Socks)
}

func decodeCountRequest( r *rf.Context) (interface{}, error) {
	tags := []string{}
	if tagsval := r.ReadRequest().FormValue("tags"); tagsval != "" {
		tags = strings.Split(tagsval, ",")
	}
	return countRequest{
		Tags: tags,
	}, nil
}

func decodeGetRequest( r *rf.Context) (interface{}) {
	return getRequest{
		ID: r.ReadPathParameter("id"),
	}
}

// encodeGetResponse is distinct from the generic encodeResponse because we need
// to special-case when the getResponse object contains a non-nil error.
func encodeGetResponse( w *rf.Context, response interface{}) error {
	resp := response.(getResponse)
	return encodeResponse(w, resp.Sock)
}

func decodeTagsRequest( r *rf.Context) (interface{}, error) {
	return struct{}{}, nil
}

func decodeHealthRequest(r *rf.Context) (interface{}, error) {
	return struct{}{}, nil
}

func encodeHealthResponse( w *rf.Context, response interface{}) error {
	return encodeResponse(w, response.(healthResponse))
}

func encodeResponse( w *rf.Context, response interface{}) error {

	return w.WriteJSON(response,"application/json")
}
