package main

import (
	"flag"
	"fmt"
	apihttp "mini-image-converter/api/http"
	"mini-image-converter/internal/conf"
	"mini-image-converter/internal/pkg/core/service/imageservice"
	"mini-image-converter/internal/pkg/platform/common"
	"mini-image-converter/internal/pkg/platform/gocv"
	"mini-image-converter/internal/pkg/platform/io"
	"mini-image-converter/internal/pkg/platform/os"
	"mini-image-converter/internal/pkg/transport"
)

func main() {

	//Initialize server
	cfgPath := flag.String("configpath", "config.yaml", "path to config file")
	flag.Parse()

	err := conf.Init(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	srv := transport.NewServer()

	//Init Services
	imageService := imageservice.NewImageService(imageservice.NewImageServiceReq{
		GoCV:   gocv.NewGoCV(),
		OS:     os.NewOS(),
		Common: common.NewCommon(io.NewIO()),
	})

	//Register APIs
	apihttp.API(srv.GetEcho(), imageService)

	srv.StartServer()
}
