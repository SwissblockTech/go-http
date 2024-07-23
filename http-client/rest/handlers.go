package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/swisslockTech/go-http/http-client/commons"
	"github.com/swisslockTech/go-http/http-client/logging"
)

func (s *Server) getProducts(writer http.ResponseWriter, request *http.Request) {
	startTimer := time.Now()

	logging.Log.Info("Get products")

	endpointUrl := &url.URL{Path: rootProductsEndpoint}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodGet, path.String(), nil)
	if reqErr != nil {
		errMsg := "Get products failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	s.setRequestHeaders(restRequest)

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Get products failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}

	// TODO do not unmarshal response
	var products []*commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&products)
	if unmarshErr != nil {
		errMsg := "Get products failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	defer closeBody(response.Body)

	sendJsonResponse(writer, http.StatusOK, products)

	IncreaseRestRequests("getProducts")
	ObserveRestRequestsTime("getProducts", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) getProduct(writer http.ResponseWriter, request *http.Request) {
	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Get product failed: Invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
		return
	}
	logging.SugaredLog.Infof("Get product by ID: %d", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodGet, path.String(), nil)
	if reqErr != nil {
		errMsg := "Get product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	s.setRequestHeaders(restRequest)

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Get product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Get product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	defer closeBody(response.Body)

	sendJsonResponse(writer, http.StatusOK, product)

	IncreaseRestRequests("getProduct")
	ObserveRestRequestsTime("getProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) createProduct(writer http.ResponseWriter, request *http.Request) {
	startTimer := time.Now()

	logging.SugaredLog.Infof("Create product")

	endpointUrl := &url.URL{Path: rootProductsEndpoint}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodPost, path.String(), request.Body)
	defer closeBody(request.Body)

	if reqErr != nil {
		errMsg := "Create product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	s.setRequestHeaders(restRequest)

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Create product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Create product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
		return
	}
	defer closeBody(response.Body)

	sendJsonResponse(writer, http.StatusCreated, product)

	IncreaseRestRequests("createProduct")
	ObserveRestRequestsTime("createProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {
	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Update product failed: invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
		return
	}

	logging.SugaredLog.Infof("Update product with ID %d", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodPut, path.String(), request.Body)
	defer closeBody(request.Body)

	if reqErr != nil {
		errMsg := "Update product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	s.setRequestHeaders(restRequest)

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Update product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Update product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
		return
	}
	defer closeBody(response.Body)

	sendJsonResponse(writer, http.StatusCreated, product)

	IncreaseRestRequests("updateProduct")
	ObserveRestRequestsTime("updateProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Delete product failed: invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
		return
	}

	logging.SugaredLog.Infof("Delete product by ID: %d", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodDelete, path.String(), nil)
	if reqErr != nil {
		errMsg := "Delete product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}
	s.setRequestHeaders(restRequest)

	_, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Update product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
		return
	}

	sendJsonResponse(writer, http.StatusOK, map[string]string{"result": "success"})

	IncreaseRestRequests("deleteProduct")
	ObserveRestRequestsTime("deleteProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}
