/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"net/rpc"
	"time"

	router "github.com/gorilla/mux"
)

// Routes paths for "minio control" commands.
const (
	controlPath = "/controller"
)

// Register controller RPC handlers.
func registerControllerRPCRouter(mux *router.Router, srvCmdConfig serverCmdConfig) {
	// Initialize Controller.
	ctrlHandlers := &controllerAPIHandlers{
		ObjectAPI:    newObjectLayerFn,
		StorageDisks: srvCmdConfig.storageDisks,
		timestamp:    time.Now().UTC(),
	}

	ctrlRPCServer := rpc.NewServer()
	ctrlRPCServer.RegisterName("Controller", ctrlHandlers)

	ctrlRouter := mux.NewRoute().PathPrefix(reservedBucket).Subrouter()
	ctrlRouter.Path(controlPath).Handler(ctrlRPCServer)
}

// Handler for object healing.
type controllerAPIHandlers struct {
	ObjectAPI    func() ObjectLayer
	StorageDisks []StorageAPI
	timestamp    time.Time
}
