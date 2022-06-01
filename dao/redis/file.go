package redis

import (
	"JByun/models"
	"go.uber.org/zap"
	"strconv"
)

func InsertChunkInfo(f *models.ChunkInitParam) error {
	err := rdb.HSet("MP_"+strconv.FormatInt(f.UpLoadID, 10), "file_size", f.FileSize).Err()
	if err != nil {
		zap.L().Error("rdb.HSet failed", zap.Error(err))
		return err
	}
	err = rdb.HSet("MP_"+strconv.FormatInt(f.UpLoadID, 10), "file_sha1", f.FileSha1).Err()
	if err != nil {
		zap.L().Error("rdb.HSet failed", zap.Error(err))
		return err
	}
	err = rdb.HSet("MP_"+strconv.FormatInt(f.UpLoadID, 10), "chunk_size", f.ChunkSize).Err()
	if err != nil {
		zap.L().Error("rdb.HSet failed", zap.Error(err))
		return err
	}
	err = rdb.HSet("MP_"+strconv.FormatInt(f.UpLoadID, 10), "chunk_count", f.ChunkCount).Err()
	if err != nil {
		zap.L().Error("rdb.Set failed", zap.Error(err))
		return err
	}
	return nil
}
