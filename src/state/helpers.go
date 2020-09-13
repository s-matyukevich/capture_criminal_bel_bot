package state

import (
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
)

func removeAllWatches(chatId int64, db *dbpkg.DB) error {
	for key, _ := range common.Dist {
		err := db.DeleteWatch(chatId, key)
		if err != nil {
			return err
		}
	}
	return nil
}
