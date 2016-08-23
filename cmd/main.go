package main

import (
	"flag"
	"github.com/CodisLabs/codis/pkg/utils/log"
	"github.com/fagongzi/collation-directory/pkg"
	"runtime"
)

var (
	cpus = flag.Int("cpus", 1, "use cpu nums")

	source = flag.String("dir-src", "", "source dir to collation")
	dest   = flag.String("dir-dest", "", "target dir to collation")

	logFile  = flag.String("log-file", "", "which file to record log, if not set stdout to use.")
	logLevel = flag.String("log-level", "info", "log level.")
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(*cpus)

	pkg.InitLog(*logFile)
	pkg.SetLogLevel(*logLevel)

	collationer := pkg.NewCollationer(*source, *dest)

	if err := collationer.Start(); err != nil {
		log.PanicErrorf(err, "Collation from <%s> to <%s> failure", *source, *dest)
	}

	log.Infof("Collation from <%s> to <%s> success", *source, *dest)
}
