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
	"os"
	"time"
)

//Logger  Data Structure responsible for containing the details regarding the logger
type Logger struct {
	FileName string
	verbose  bool
}

func (l *Logger) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(l.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	if l.verbose == true {
		n, err = os.Stdout.Write(p)
	}
	n, err = file.Write(p)
	file.Close()
	return n, err
}

//Must  Function for reporting and storing logs/errors
func (l *Logger) Must(e error, logstring string) {
	if e != nil {
		l.Write([]byte(e.Error()))
		panic(e)
	}

	l.Write([]byte(formatLog(logstring) + "\n"))
}

func formatLog(logstring string) string {
	tm := time.Now()
	logtime := tm.Format("2/Jan/2006:15:04:05 -0700")
	return "[" + logtime + "] " + logstring
}

//NewLogger  The splatter logger constructor
func NewLogger(fname string, verbosity bool) *Logger {
	temp := Logger{
		FileName: fname,
		verbose:  verbosity,
	}
	return &temp
}
