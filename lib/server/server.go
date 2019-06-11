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

	"github.com/ILUGD/splatter/lib/painter"

	"github.com/ILUGD/splatter/lib/readers"
)

//SplatterServer  Data Structure for holding the server
type SplatterServer struct {
	//Don't worry about the empty struct this would hold the logger
}

func havePoster(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg readers.Document
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	painter.GeneratePoster(msg)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I guess it worked\n"))
}

//Serve  Function to start the serving the API
func (s *SplatterServer) Serve(port string) {
	http.HandleFunc("/poster", havePoster)
	http.ListenAndServe(port, nil)
}

//NewServer  The splatter API server constructor
func NewServer() *SplatterServer {
	s := SplatterServer{}
	return &s
}
