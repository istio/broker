// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server provides HTTP open service broker API server bindings.
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/golang/glog"
	"istio.io/broker/pkg/controller"
)

type Server struct {
	controller *controller.Controller
}

func CreateServer() (*Server, error) {
	controller, err := controller.CreateController()
	if err != nil {
		return nil, err
	}

	return &Server{
		controller: controller,
	}, nil
}

func (s *Server) Start(port uint16) {
	router := mux.NewRouter()

	router.HandleFunc("/v2/catalog", s.controller.Catalog).Methods("GET")

	http.Handle("/", router)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		glog.Errorf("Unable to start server: %v", err)
	}
}
