package usecase

import (
	"testing"

	"github.com/dzakaeryan20/lemonilo/config"
	cachedal "github.com/dzakaeryan20/lemonilo/dal/cache"
	dbdal "github.com/dzakaeryan20/lemonilo/dal/db"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg   *config.Config
		db    dbdal.DBDAL
		cache cachedal.CacheDAL
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg, tt.args.db, tt.args.cache)
		})
	}
}
