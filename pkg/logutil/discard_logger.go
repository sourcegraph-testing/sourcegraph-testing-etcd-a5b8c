// Copyright 2018 The etcd Authors
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

package logutil

import (
	"log"

	"google.golang.org/grpc/grpclog"
)

// assert that "discardLogger" satisfy "Logger" interface
var _ Logger = &discardLogger{}

// NewDiscardLogger returns a new Logger that discards everything except "fatal".
func NewDiscardLogger() Logger { return &discardLogger{} }

type discardLogger struct{}

func (l *discardLogger) Info(args ...any)                    {}
func (l *discardLogger) Infoln(args ...any)                  {}
func (l *discardLogger) Infof(format string, args ...any)    {}
func (l *discardLogger) Warning(args ...any)                 {}
func (l *discardLogger) Warningln(args ...any)               {}
func (l *discardLogger) Warningf(format string, args ...any) {}
func (l *discardLogger) Error(args ...any)                   {}
func (l *discardLogger) Errorln(args ...any)                 {}
func (l *discardLogger) Errorf(format string, args ...any)   {}
func (l *discardLogger) Fatal(args ...any)                   { log.Fatal(args...) }
func (l *discardLogger) Fatalln(args ...any)                 { log.Fatalln(args...) }
func (l *discardLogger) Fatalf(format string, args ...any)   { log.Fatalf(format, args...) }
func (l *discardLogger) V(lvl int) bool {
	return false
}
func (l *discardLogger) Lvl(lvl int) grpclog.LoggerV2 { return l }
