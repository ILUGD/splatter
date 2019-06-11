/*
 *  Splatter
 *  Copyright (C) 2019 Kuntal Majumder
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or (at
 *  your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but
 *  WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 *  General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program; if not, write to the Free Software
 *  Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307
 *  USA.
 */

package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ILUGD/splatter/lib/logger"
	"github.com/ILUGD/splatter/lib/painter"

	"github.com/ILUGD/splatter/lib/readers"
)

//SplatterServer  Data Structure for holding the server
type SplatterServer struct {
	*logger.Logger
}

func (s *SplatterServer) havePoster(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	s.Must(err, "Read JSON body")

	// Unmarshal
	var msg readers.Document
	err = json.Unmarshal(b, &msg)
	s.Must(err, "Unmarshalled JSON")

	painter.GeneratePoster(msg)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I guess it worked\n"))
}

//Serve  Function to start the serving the API
func (s *SplatterServer) Serve(port string) {
	http.HandleFunc("/poster", s.havePoster)
	http.ListenAndServe(port, nil)
}

//NewServer  The splatter API server constructor
func NewServer(l *logger.Logger) *SplatterServer {
	s := SplatterServer{l}
	return &s
}
