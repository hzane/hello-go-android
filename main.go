// +build android
package main
// #cgo LDFLAGS: -llog
//
// #include <android/log.h>
//
// void hello() {
//   __android_log_print(
//     ANDROID_LOG_INFO, "MyProgram", "hello world");
// }
import "C"
func main() {
    C.hello()
}
