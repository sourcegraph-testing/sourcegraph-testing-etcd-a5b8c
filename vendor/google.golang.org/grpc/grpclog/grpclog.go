/*
 *
 * Copyright 2017 gRPC authors.
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
 *
 */

// Package grpclog defines logging for grpc.
//
// All logs in transport and grpclb packages only go to verbose level 2.
// All logs in other packages in grpc are logged in spite of the verbosity level.
//
// In the default logger,
// severity level can be set by environment variable GRPC_GO_LOG_SEVERITY_LEVEL,
// verbosity level can be set by GRPC_GO_LOG_VERBOSITY_LEVEL.
package grpclog // import "google.golang.org/grpc/grpclog"

import "os"

var logger = newLoggerV2()

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...any) {
	logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...any) {
	logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...any) {
	logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...any) {
	logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...any) {
	logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...any) {
	logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...any) {
	logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...any) {
	logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func Fatalln(args ...any) {
	logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...any) {
	logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...any) {
	logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func Println(args ...any) {
	logger.Infoln(args...)
}
