package conversion

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func MediaToORM(v *pb.Media) *orm.Medium {
	return &orm.Medium{
		Album:                 v.Album,
		BitDepth:              v.BitDepth,
		BitRate:               v.BitRate,
		BitRateMode:           v.BitRateMode,
		ChannelLayout:         v.ChannelLayout,
		ChannelPositions:      v.ChannelPositions,
		Channels:              v.Channels,
		Comment:               v.Comment,
		CompressionMode:       v.CompressionMode,
		Duration:              v.Duration,
		EncodedLibrary:        v.EncodedLibrary,
		EncodedLibraryName:    v.EncodedLibraryName,
		EncodedLibraryVersion: v.EncodedLibraryVersion,
		FileExtension:         v.FileExtension,
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
		Track:                 v.Track,
		TrackPosition:         v.TrackPosition,
	}
}

func MediaFromORM(v *orm.Medium) *pb.Media {
	return &pb.Media{
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
		EncodedLibrary:        v.EncodedLibrary,
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
	}
}

func MediasFromORM(a ...*orm.Medium) []*pb.Media {
	var result []*pb.Media
	for _, v := range a {
		result = append(result, MediaFromORM(v))
	}
	return result
}
