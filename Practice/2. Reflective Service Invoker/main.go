package main

/*
This program implements a dynamic HTTP server for invoking methods on registered services.
It uses reflection to determine the service and method to invoke based on the HTTP request URL path.
The program supports the following workflow:
1. Extracts the `serviceName` and `methodName` from the URL.
2. Locates the corresponding service from a predefined `ServiceMap`.
3. Uses reflection to find and invoke the specified method on the service.
4. Parses the HTTP request body into the appropriate input type required by the method.
5. Calls the method and retrieves its response and error values.
6. Serializes the response (if valid) into JSON and sends it to the client.
7. Handles errors gracefully, including missing services, methods, or invalid input.
This design enables flexible and reusable service-method-based routing for APIs.
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"test/service"

	"github.com/gorilla/mux"
)

var ServiceMap = map[string]interface{}{
	"AuthService": &service.AuthService{},
}

func invokeServiceMethod(serviceName, methodName string, reqBody []byte, w http.ResponseWriter, r *http.Request) {
	service, exists := ServiceMap[serviceName]
	if !exists {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	method := reflect.ValueOf(service).MethodByName(methodName)
	if !method.IsValid() {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	reqType := method.Type().In(1)
	reqInstance := reflect.New(reqType).Interface()

	if err := json.Unmarshal(reqBody, reqInstance); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	results := method.Call([]reflect.Value{
		reflect.ValueOf(r.Context()),
		reflect.ValueOf(reqInstance).Elem(),
	})

	if len(results) == 2 {
		respValue := results[0]
		errValue := results[1]

		if respValue.IsValid() && !respValue.IsNil() {
			if respValue.Kind() == reflect.Ptr && respValue.Elem().Kind() == reflect.Struct {
				respJSON, err := json.Marshal(respValue.Interface())
				if err != nil {
					http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(respJSON)
			} else {
				http.Error(w, "Invalid response type", http.StatusInternalServerError)
			}
		}

		if errValue.IsValid() && !errValue.IsNil() {
			if errInterface, ok := errValue.Interface().(error); ok && errInterface != nil {
				http.Error(w, errInterface.Error(), http.StatusUnauthorized)
				return
			}
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/HPC/{serviceName}/{methodName}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		serviceName := vars["serviceName"]
		methodName := vars["methodName"]

		reqBody := []byte{}
		if r.Body != nil {
			defer r.Body.Close()
			reqBody, _ = io.ReadAll(r.Body)
		}

		invokeServiceMethod(serviceName, methodName, reqBody, w, r)
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
