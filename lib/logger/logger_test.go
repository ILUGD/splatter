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

package logger

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogWriter(t *testing.T) {
	filename := "testf"
	log := NewLogger(filename, false)
	log.Must(nil, "abc")

	input := formatLog("abc\n")
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Errorf("log file with name testf not created")
	}

	testFormat(t, input, string(content))
	os.Remove(filename)
}

func testFormat(t *testing.T, input, output string) {

	if output[0] != input[0] {
		t.Errorf("output is in the wrong format, expected %b , got %b", input[0], output[0])
	}

	if output[26] != input[26] {
		t.Errorf("output is in the wrong format, expected %b , got %b", input[26], output[26])
	}

	if output[28:] != input[28:] {
		t.Errorf("output is in the wrong format, expected %s , got %s", input[28:], output[28:])
	}
}
