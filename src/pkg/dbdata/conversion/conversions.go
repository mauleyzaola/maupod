package conversion

import (
	"math"
	"strings"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/volatiletech/null/v8"
)

// covert playlist back and forth
func PlaylistToORM(v *protos.PlayList) *orm.Playlist {
	return &orm.Playlist{
		ID:   v.Id,
		Name: v.Name,
	}
}

func PlaylistFromORM(v *orm.Playlist) *protos.PlayList {
	return &protos.PlayList{
		Id:   v.ID,
		Name: v.Name,
	}
}

func PlaylistsFromORM(a ...*orm.Playlist) []*protos.PlayList {
	var result []*protos.PlayList
	for _, v := range a {
		result = append(result, PlaylistFromORM(v))
	}
	return result
}

// covert playlistItemcls back and forth
func PlaylistItemToORM(v *protos.PlaylistItem) *orm.PlaylistItem {
	return &orm.PlaylistItem{
		ID:         v.Id,
		PlaylistID: v.Playlist.Id,
		Position:   int(v.Position),
		MediaID:    v.Media.Id,
	}
}

func PlaylistItemFromORM(v *orm.PlaylistItem) *protos.PlaylistItem {
	return &protos.PlaylistItem{
		Id: v.ID,
		Playlist: &protos.PlayList{
			Id: v.PlaylistID,
		},
		Position: int32(v.Position),
		Media: &protos.Media{
			Id: v.MediaID,
		},
	}
}

func PlaylistItemsFromORM(a ...*orm.PlaylistItem) []*protos.PlaylistItem {
	var result []*protos.PlaylistItem
	for _, v := range a {
		result = append(result, PlaylistItemFromORM(v))
	}
	return result
}

// covert media back and forth
func MediaToORM(v *protos.Media) *orm.Medium {
	return &orm.Medium{
		Album:                 v.Album,
		AlbumPerformer:        v.AlbumPerformer,
		AudioCount:            v.AudioCount,
		BitDepth:              v.BitDepth,
		BitDepthString:        v.BitDepthString,
		BitRate:               v.BitRate,
		BitRateMode:           v.BitRateMode,
		ChannelLayout:         v.ChannelLayout,
		ChannelPositions:      v.ChannelPositions,
		Channels:              v.Channels,
		CommercialName:        v.CommercialName,
		Comment:               v.Comment,
		CompleteName:          v.CompleteName,
		CompressionMode:       v.CompressionMode,
		CountOfAudioStreams:   v.CountOfAudioStreams,
		Duration:              v.Duration,
		EncodedLibraryDate:    v.EncodedLibraryDate,
		EncodedLibraryName:    v.EncodedLibraryName,
		EncodedLibraryVersion: v.EncodedLibraryVersion,
		FileExtension:         strings.ToLower(v.FileExtension),
		FileSize:              v.FileSize,
		Format:                v.Format,
		Genre:                 v.Genre,
		ID:                    v.Id,
		LastScan:              helpers.TsToTime2(v.LastScan),
		Location:              v.Location,
		OverallBitRate:        v.OverallBitRate,
		OverallBitRateMode:    v.OverallBitRateMode,
		Performer:             v.Performer,
		RecordedDate:          v.RecordedDate,
		SamplingCount:         v.SamplingCount,
		SamplingRate:          v.SamplingRate,
		Sha:                   v.Sha,
		StreamSize:            v.StreamSize,
		Title:                 v.Title,
		Track:                 v.Track,
		TrackNameTotal:        v.TrackNameTotal,
		TrackPosition:         v.TrackPosition,
		FileName:              v.FileName,
		FolderName:            v.FolderName,
		FormatInfo:            v.FormatInfo,
		FormatURL:             v.FormatUrl,
		InternetMediaType:     v.InternetMediaType,
		KindOfStream:          v.KindOfStream,
		Part:                  v.Part,
		PartTotal:             v.PartTotal,
		StreamIdentifier:      v.StreamIdentifier,
		WritingLibrary:        v.WritingLibrary,
		ModifiedDate:          helpers.TsToTime2(v.ModifiedDate),
		Composer:              v.Composer,
		LastImageScan:         null.TimeFromPtr(helpers.TsToTime(v.LastImageScan)),
		ImageLocation:         v.ImageLocation,
		AlbumIdentifier:       v.AlbumIdentifier,
		Directory:             v.Directory,
	}
}

func MediaFromORM(v *orm.Medium) *protos.Media {
	return &protos.Media{
		Album:                 v.Album,
		AlbumPerformer:        v.AlbumPerformer,
		AudioCount:            v.AudioCount,
		BitDepth:              v.BitDepth,
		BitDepthString:        v.BitDepthString,
		BitRate:               v.BitRate,
		BitRateMode:           v.BitRateMode,
		ChannelLayout:         v.ChannelLayout,
		ChannelPositions:      v.ChannelPositions,
		Channels:              v.Channels,
		CommercialName:        v.CommercialName,
		Comment:               v.Comment,
		CompleteName:          v.CompleteName,
		CompressionMode:       v.CompressionMode,
		CountOfAudioStreams:   v.CountOfAudioStreams,
		Duration:              v.Duration,
		EncodedLibraryDate:    v.EncodedLibraryDate,
		EncodedLibraryName:    v.EncodedLibraryName,
		EncodedLibraryVersion: v.EncodedLibraryVersion,
		FileExtension:         v.FileExtension,
		FileSize:              v.FileSize,
		Format:                v.Format,
		Genre:                 v.Genre,
		Id:                    v.ID,
		LastScan:              helpers.TimeToTs2(v.LastScan),
		Location:              v.Location,
		OverallBitRate:        v.OverallBitRate,
		OverallBitRateMode:    v.OverallBitRateMode,
		Performer:             v.Performer,
		RecordedDate:          v.RecordedDate,
		SamplingCount:         v.SamplingCount,
		SamplingRate:          v.SamplingRate,
		Sha:                   v.Sha,
		StreamSize:            v.StreamSize,
		Title:                 v.Title,
		Track:                 v.Track,
		TrackNameTotal:        v.TrackNameTotal,
		TrackPosition:         v.TrackPosition,
		FileName:              v.FileName,
		FolderName:            v.FolderName,
		FormatInfo:            v.FormatInfo,
		FormatUrl:             v.FormatURL,
		InternetMediaType:     v.InternetMediaType,
		KindOfStream:          v.KindOfStream,
		Part:                  v.Part,
		PartTotal:             v.PartTotal,
		StreamIdentifier:      v.StreamIdentifier,
		WritingLibrary:        v.WritingLibrary,
		ModifiedDate:          helpers.TimeToTs2(v.ModifiedDate),
		Composer:              v.Composer,
		LastImageScan:         helpers.TimeToTs(v.LastImageScan.Ptr()),
		ImageLocation:         v.ImageLocation,
		AlbumIdentifier:       v.AlbumIdentifier,
		Directory:             v.Directory,
		Seconds:               int64(math.Floor(v.Duration / 1000)),
	}
}

func MediasFromORM(a ...*orm.Medium) []*protos.Media {
	var result []*protos.Media
	for _, v := range a {
		result = append(result, MediaFromORM(v))
	}
	return result
}

func ViewAlbumToMedia(v *orm.ViewAlbum) *protos.Media {
	m := &protos.Media{
		Id:              v.ID.String,
		Format:          v.Format.String,
		FileSize:        v.FileSize.Int64,
		Album:           v.Album.String,
		Performer:       v.Performer.String,
		Genre:           v.Genre.String,
		RecordedDate:    v.RecordedDate.Int64,
		SamplingRate:    v.SamplingRate.Int64,
		BitRate:         v.BitRate.Int64,
		TrackNameTotal:  v.TrackNameTotal.Int64,
		AlbumIdentifier: v.AlbumIdentifier.String,
		Duration:        float64(v.Duration.Int64),
		ImageLocation:   v.ImageLocation.String,
	}
	return m
}

func ViewAlbumsToMedia(a ...*orm.ViewAlbum) []*protos.Media {
	var result []*protos.Media
	for _, v := range a {
		result = append(result, ViewAlbumToMedia(v))
	}
	return result
}
