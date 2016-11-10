package main

import (
	"fmt"

	"github.com/jgrelet/oceano2oceansites/fileExtractor"
	_ "github.com/y0ssar1an/q"
)

// usefull macro
var p = fmt.Println
var pf = fmt.Printf

// example main
func main() {

	// CTD
	opts := fileExtractor.NewFileExtractOptions()
	//files := "ctd/dfr2600?.cnv"
	//opts.SetVarsList("PRES,3,DEPTH,4,ETDD,2,TEMP,5,PSAL,17,DENS,20,SVEL,22,DOX2,15,FLU2,13,TUR3,14,NAVG,23")
	files := "D:/campagnes/CASSIOPEE/data-processing/CTD/data/cnv/dcsp055??.cnv"
	opts.SetVarsList("PRES,3,DEPTH,4,ETDD,2,TEMP,5,PSAL,21,DENS,23,SVEL,25,DOX2,19,FLU2,12,TUR3,13,NAVG,21")
	//opts.SetSkipLine(354)
	opts.SetHdrDelimiter("*END*")
	// print options
	p(opts)

	prf := fileExtractor.NewProfilesExtractor(files, *fileExtractor.NewFileExtractor(opts))
	prf.Read()
	p(prf)

	//TSG

	opts = fileExtractor.NewFileExtractOptions()
	opts.SetVarsList("DATE,3,TIME,4,LATITUDE,6,LONGITUDE,9,SSTP,19,SSJT,20,COND,21,SSPS,22")
	opts.SetSkipLine(6)
	opts.SetSeparator(",")
	p(opts)
	traj := fileExtractor.NewTrajectoriesExtractor("tsg/*-TS_COLCOR.COLCOR", *fileExtractor.NewFileExtractor(opts))
	traj.Read()
	p(traj)

}
