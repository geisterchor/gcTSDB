package restapi

import (
	"geisterchor.com/gcTSDB/gctsdb"
)

type AppContext struct {
	GCTSDBServer *gctsdb.GCTSDBServer
}
