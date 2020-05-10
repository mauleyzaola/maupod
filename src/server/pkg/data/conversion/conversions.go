package conversion

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
)

func MediaToORM(v *domain.Media) *orm.Medium {
	return &orm.Medium{
		ID:                    v.ID,
		Location:              v.Location,
		FileExtension:         v.FileExtension,
		Format:                v.Format,
		FileSize:              v.FileSize,
		Duration:              v.Duration,
		OverallBitRateMode:    v.OverallBitRateMode,
		OverallBitRate:        v.OverallBitRate,
		StreamSize:            v.StreamSize,
		Album:                 v.Album,
		Track:                 v.Track,
		TrackPosition:         v.TrackPosition,
		Performer:             v.Performer,
		Genre:                 v.Genre,
		RecordedDate:          v.RecordedDate,
		FileModifiedDate:      v.FileModifiedDate,
		Comment:               v.Comment,
		Channels:              v.Channels,
		ChannelPositions:      v.ChannelPositions,
		ChannelLayout:         v.ChannelLayout,
		SamplingRate:          v.SamplingRate,
		SamplingCount:         v.SamplingCount,
		BitDepth:              v.BitDepth,
		CompressionMode:       v.CompressionMode,
		EncodedLibrary:        v.EncodedLibrary,
		EncodedLibraryName:    v.EncodedLibraryName,
		EncodedLibraryVersion: v.EncodedLibraryVersion,
		BitRateMode:           v.BitRateMode,
		BitRate:               v.BitRate,
	}
}

func MediaFromORM(v *orm.Medium) *domain.Media {
	return &domain.Media{
		ID:                    v.ID,
		Location:              v.Location,
		FileExtension:         v.FileExtension,
		Format:                v.Format,
		FileSize:              v.FileSize,
		Duration:              v.Duration,
		OverallBitRateMode:    v.OverallBitRateMode,
		OverallBitRate:        v.OverallBitRate,
		StreamSize:            v.StreamSize,
		Album:                 v.Album,
		Track:                 v.Track,
		TrackPosition:         v.TrackPosition,
		Performer:             v.Performer,
		Genre:                 v.Genre,
		RecordedDate:          v.RecordedDate,
		FileModifiedDate:      v.FileModifiedDate,
		Comment:               v.Comment,
		Channels:              v.Channels,
		ChannelPositions:      v.ChannelPositions,
		ChannelLayout:         v.ChannelLayout,
		SamplingRate:          v.SamplingRate,
		SamplingCount:         v.SamplingCount,
		BitDepth:              v.BitDepth,
		CompressionMode:       v.CompressionMode,
		EncodedLibrary:        v.EncodedLibrary,
		EncodedLibraryName:    v.EncodedLibraryName,
		EncodedLibraryVersion: v.EncodedLibraryVersion,
		BitRateMode:           v.BitRateMode,
		BitRate:               v.BitRate,
	}
}

func MediasFromORM(a ...*orm.Medium) []*domain.Media {
	var result []*domain.Media
	for _, v := range a {
		result = append(result, MediaFromORM(v))
	}
	return result
}
