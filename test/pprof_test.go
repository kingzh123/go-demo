package main


import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	fmt.Println("this is main func")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

//被测试的方法 #go test -v
func TestAdd(t *testing.T)  {
	fmt.Println("----------------------TEST FUNC Add-----------------------")
	if os.Getenv("xx") == "" {
		t.Skip("skipping test env not set xx")
	}
	fmt.Println("this is test func")
}

//#
func TestBdd(t *testing.T){
	fmt.Println("----------------------TEST FUNC Bdd-----------------------")
	fmt.Println("this is test func")
}

//#go test -timeout 1s 测试超过执行时间
func TestTimeOut(t *testing.T) {
	fmt.Println("----------------------TEST FUNC TimeOut-----------------------")
	time.Sleep(1*time.Second)
}

//并发执行 同时运行的测试的数量默认取决于 GOMAXPROCS。
//它可以通过 go test -parallel n（个数）
func TestParallel1(t *testing.T) {
	t.Parallel()
	fmt.Println("----------------------TEST FUNC Parallel1-----------------------")
}

func TestParallel2(t *testing.T) {
	t.Parallel()
	fmt.Println("----------------------TEST FUNC Parallel2-----------------------")
}

func TestParallel3(t *testing.T) {
	t.Parallel()
	fmt.Println("----------------------TEST FUNC Parallel3-----------------------")
}

func TestParallel4(t *testing.T) {
	t.Parallel()
	fmt.Println("----------------------TEST FUNC Parallel4-----------------------")
}