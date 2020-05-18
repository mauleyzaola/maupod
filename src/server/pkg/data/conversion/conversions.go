package conversion

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func MediaToORM(v *pb.Media) *orm.Medium {
	return &orm.Medium{
		ID:                    v.Id,
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
		FileModifiedDate:      helpers.TsToTime2(v.FileModifiedDate),
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
		LastScan:              helpers.TsToTime2(v.LastScan),
		ModifiedDate:          helpers.TsToTime2(v.ModifiedDate),
		Sha:                   v.Sha,
	}
}

func MediaFromORM(v *orm.Medium) *pb.Media {
	return &pb.Media{
		Id:                    v.ID,
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
		FileModifiedDate:      helpers.TimeToTs(&v.FileModifiedDate),
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
		LastScan:              helpers.TimeToTs(&v.LastScan),
		ModifiedDate:          helpers.TimeToTs(&v.ModifiedDate),
		Sha:                   v.Sha,
	}
}

func MediasFromORM(a ...*orm.Medium) []*pb.Media {
	var result []*pb.Media
	for _, v := range a {
		result = append(result, MediaFromORM(v))
	}
	return result
}
