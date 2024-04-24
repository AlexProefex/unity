package service

import (
	"unity/repository/dao"
	"unity/types"
)

func GetDataAlbums() []types.Album {
	albums := dao.GetDataAlbums()

	albumsReturns := make([]types.Album, 0)

	for _, index := range albums {
		albumsReturns = append(albumsReturns, types.Album{
			ID:     index.ID,
			Title:  index.Title,
			Artist: index.Artist,
			Price:  index.Price,
		})
	}

	return albumsReturns

}
