// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	Metadata
	BurndownSparseMatrixRow
	BurndownSparseMatrix
	FilesOwnership
	BurndownAnalysisResults
	CompressedSparseRowMatrix
	Couples
	TouchedFiles
	CouplesAnalysisResults
	UASTChange
	UASTChangesSaverResults
	ShotnessRecord
	ShotnessAnalysisResults
	FileHistory
	FileHistoryResultMessage
	LineStats
	DevDay
	DayDevs
	DevsAnalysisResults
	Sentiment
	CommentSentimentResults
	CommitFile
	Commit
	CommitsAnalysisResults
	Typo
	TyposDataset
	AnalysisResults
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	// this format is versioned
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// git hash of the revision from which Hercules is built
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// repository's name
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	// UNIX timestamp of the first analysed commit
	BeginUnixTime int64 `protobuf:"varint,4,opt,name=begin_unix_time,json=beginUnixTime,proto3" json:"begin_unix_time,omitempty"`
	// UNIX timestamp of the last analysed commit
	EndUnixTime int64 `protobuf:"varint,5,opt,name=end_unix_time,json=endUnixTime,proto3" json:"end_unix_time,omitempty"`
	// number of processed commits
	Commits int32 `protobuf:"varint,6,opt,name=commits,proto3" json:"commits,omitempty"`
	// duration of the analysis in milliseconds
	RunTime int64 `protobuf:"varint,7,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	// time taken by each pipeline item in seconds
	RunTimePerItem map[string]float64 `protobuf:"bytes,8,rep,name=run_time_per_item,json=runTimePerItem" json:"run_time_per_item,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{0} }

func (m *Metadata) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Metadata) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Metadata) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Metadata) GetBeginUnixTime() int64 {
	if m != nil {
		return m.BeginUnixTime
	}
	return 0
}

func (m *Metadata) GetEndUnixTime() int64 {
	if m != nil {
		return m.EndUnixTime
	}
	return 0
}

func (m *Metadata) GetCommits() int32 {
	if m != nil {
		return m.Commits
	}
	return 0
}

func (m *Metadata) GetRunTime() int64 {
	if m != nil {
		return m.RunTime
	}
	return 0
}

func (m *Metadata) GetRunTimePerItem() map[string]float64 {
	if m != nil {
		return m.RunTimePerItem
	}
	return nil
}

type BurndownSparseMatrixRow struct {
	// the first `len(column)` elements are stored,
	// the rest `number_of_columns - len(column)` values are zeros
	Columns []uint32 `protobuf:"varint,1,rep,packed,name=columns" json:"columns,omitempty"`
}

func (m *BurndownSparseMatrixRow) Reset()                    { *m = BurndownSparseMatrixRow{} }
func (m *BurndownSparseMatrixRow) String() string            { return proto.CompactTextString(m) }
func (*BurndownSparseMatrixRow) ProtoMessage()               {}
func (*BurndownSparseMatrixRow) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{1} }

func (m *BurndownSparseMatrixRow) GetColumns() []uint32 {
	if m != nil {
		return m.Columns
	}
	return nil
}

type BurndownSparseMatrix struct {
	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumberOfRows    int32  `protobuf:"varint,2,opt,name=number_of_rows,json=numberOfRows,proto3" json:"number_of_rows,omitempty"`
	NumberOfColumns int32  `protobuf:"varint,3,opt,name=number_of_columns,json=numberOfColumns,proto3" json:"number_of_columns,omitempty"`
	// `len(row)` matches `number_of_rows`
	Rows []*BurndownSparseMatrixRow `protobuf:"bytes,4,rep,name=rows" json:"rows,omitempty"`
}

func (m *BurndownSparseMatrix) Reset()                    { *m = BurndownSparseMatrix{} }
func (m *BurndownSparseMatrix) String() string            { return proto.CompactTextString(m) }
func (*BurndownSparseMatrix) ProtoMessage()               {}
func (*BurndownSparseMatrix) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{2} }

func (m *BurndownSparseMatrix) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BurndownSparseMatrix) GetNumberOfRows() int32 {
	if m != nil {
		return m.NumberOfRows
	}
	return 0
}

func (m *BurndownSparseMatrix) GetNumberOfColumns() int32 {
	if m != nil {
		return m.NumberOfColumns
	}
	return 0
}

func (m *BurndownSparseMatrix) GetRows() []*BurndownSparseMatrixRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type FilesOwnership struct {
	// The sum always equals to the total number of lines in the file.
	Value map[int32]int32 `protobuf:"bytes,1,rep,name=value" json:"value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *FilesOwnership) Reset()                    { *m = FilesOwnership{} }
func (m *FilesOwnership) String() string            { return proto.CompactTextString(m) }
func (*FilesOwnership) ProtoMessage()               {}
func (*FilesOwnership) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{3} }

func (m *FilesOwnership) GetValue() map[int32]int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type BurndownAnalysisResults struct {
	// how many days are in each band [burndown_project, burndown_file, burndown_developer]
	Granularity int32 `protobuf:"varint,1,opt,name=granularity,proto3" json:"granularity,omitempty"`
	// how frequently we measure the state of each band [burndown_project, burndown_file, burndown_developer]
	Sampling int32 `protobuf:"varint,2,opt,name=sampling,proto3" json:"sampling,omitempty"`
	// always exists
	Project *BurndownSparseMatrix `protobuf:"bytes,3,opt,name=project" json:"project,omitempty"`
	// this is included if `--burndown-files` was specified
	Files []*BurndownSparseMatrix `protobuf:"bytes,4,rep,name=files" json:"files,omitempty"`
	// these two are included if `--burndown-people` was specified
	People []*BurndownSparseMatrix `protobuf:"bytes,5,rep,name=people" json:"people,omitempty"`
	// rows and cols order correspond to `burndown_developer`
	PeopleInteraction *CompressedSparseRowMatrix `protobuf:"bytes,6,opt,name=people_interaction,json=peopleInteraction" json:"people_interaction,omitempty"`
	// How many lines belong to relevant developers for each file. The order is the same as in `files`.
	FilesOwnership []*FilesOwnership `protobuf:"bytes,7,rep,name=files_ownership,json=filesOwnership" json:"files_ownership,omitempty"`
}

func (m *BurndownAnalysisResults) Reset()                    { *m = BurndownAnalysisResults{} }
func (m *BurndownAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*BurndownAnalysisResults) ProtoMessage()               {}
func (*BurndownAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{4} }

func (m *BurndownAnalysisResults) GetGranularity() int32 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *BurndownAnalysisResults) GetSampling() int32 {
	if m != nil {
		return m.Sampling
	}
	return 0
}

func (m *BurndownAnalysisResults) GetProject() *BurndownSparseMatrix {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *BurndownAnalysisResults) GetFiles() []*BurndownSparseMatrix {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *BurndownAnalysisResults) GetPeople() []*BurndownSparseMatrix {
	if m != nil {
		return m.People
	}
	return nil
}

func (m *BurndownAnalysisResults) GetPeopleInteraction() *CompressedSparseRowMatrix {
	if m != nil {
		return m.PeopleInteraction
	}
	return nil
}

func (m *BurndownAnalysisResults) GetFilesOwnership() []*FilesOwnership {
	if m != nil {
		return m.FilesOwnership
	}
	return nil
}

type CompressedSparseRowMatrix struct {
	NumberOfRows    int32 `protobuf:"varint,1,opt,name=number_of_rows,json=numberOfRows,proto3" json:"number_of_rows,omitempty"`
	NumberOfColumns int32 `protobuf:"varint,2,opt,name=number_of_columns,json=numberOfColumns,proto3" json:"number_of_columns,omitempty"`
	// https://en.wikipedia.org/wiki/Sparse_matrix#Compressed_sparse_row_.28CSR.2C_CRS_or_Yale_format.29
	Data    []int64 `protobuf:"varint,3,rep,packed,name=data" json:"data,omitempty"`
	Indices []int32 `protobuf:"varint,4,rep,packed,name=indices" json:"indices,omitempty"`
	Indptr  []int64 `protobuf:"varint,5,rep,packed,name=indptr" json:"indptr,omitempty"`
}

func (m *CompressedSparseRowMatrix) Reset()                    { *m = CompressedSparseRowMatrix{} }
func (m *CompressedSparseRowMatrix) String() string            { return proto.CompactTextString(m) }
func (*CompressedSparseRowMatrix) ProtoMessage()               {}
func (*CompressedSparseRowMatrix) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{5} }

func (m *CompressedSparseRowMatrix) GetNumberOfRows() int32 {
	if m != nil {
		return m.NumberOfRows
	}
	return 0
}

func (m *CompressedSparseRowMatrix) GetNumberOfColumns() int32 {
	if m != nil {
		return m.NumberOfColumns
	}
	return 0
}

func (m *CompressedSparseRowMatrix) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CompressedSparseRowMatrix) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *CompressedSparseRowMatrix) GetIndptr() []int64 {
	if m != nil {
		return m.Indptr
	}
	return nil
}

type Couples struct {
	// name of each `matrix`'s row and column
	Index []string `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	// is always square
	Matrix *CompressedSparseRowMatrix `protobuf:"bytes,2,opt,name=matrix" json:"matrix,omitempty"`
}

func (m *Couples) Reset()                    { *m = Couples{} }
func (m *Couples) String() string            { return proto.CompactTextString(m) }
func (*Couples) ProtoMessage()               {}
func (*Couples) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{6} }

func (m *Couples) GetIndex() []string {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Couples) GetMatrix() *CompressedSparseRowMatrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

type TouchedFiles struct {
	Files []int32 `protobuf:"varint,1,rep,packed,name=files" json:"files,omitempty"`
}

func (m *TouchedFiles) Reset()                    { *m = TouchedFiles{} }
func (m *TouchedFiles) String() string            { return proto.CompactTextString(m) }
func (*TouchedFiles) ProtoMessage()               {}
func (*TouchedFiles) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{7} }

func (m *TouchedFiles) GetFiles() []int32 {
	if m != nil {
		return m.Files
	}
	return nil
}

type CouplesAnalysisResults struct {
	FileCouples   *Couples `protobuf:"bytes,6,opt,name=file_couples,json=fileCouples" json:"file_couples,omitempty"`
	PeopleCouples *Couples `protobuf:"bytes,7,opt,name=people_couples,json=peopleCouples" json:"people_couples,omitempty"`
	// order corresponds to `people_couples::index`
	PeopleFiles []*TouchedFiles `protobuf:"bytes,8,rep,name=people_files,json=peopleFiles" json:"people_files,omitempty"`
	// order corresponds to `file_couples::index`
	FilesLines []int32 `protobuf:"varint,9,rep,packed,name=files_lines,json=filesLines" json:"files_lines,omitempty"`
}

func (m *CouplesAnalysisResults) Reset()                    { *m = CouplesAnalysisResults{} }
func (m *CouplesAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*CouplesAnalysisResults) ProtoMessage()               {}
func (*CouplesAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{8} }

func (m *CouplesAnalysisResults) GetFileCouples() *Couples {
	if m != nil {
		return m.FileCouples
	}
	return nil
}

func (m *CouplesAnalysisResults) GetPeopleCouples() *Couples {
	if m != nil {
		return m.PeopleCouples
	}
	return nil
}

func (m *CouplesAnalysisResults) GetPeopleFiles() []*TouchedFiles {
	if m != nil {
		return m.PeopleFiles
	}
	return nil
}

func (m *CouplesAnalysisResults) GetFilesLines() []int32 {
	if m != nil {
		return m.FilesLines
	}
	return nil
}

type UASTChange struct {
	FileName   string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	SrcBefore  string `protobuf:"bytes,2,opt,name=src_before,json=srcBefore,proto3" json:"src_before,omitempty"`
	SrcAfter   string `protobuf:"bytes,3,opt,name=src_after,json=srcAfter,proto3" json:"src_after,omitempty"`
	UastBefore string `protobuf:"bytes,4,opt,name=uast_before,json=uastBefore,proto3" json:"uast_before,omitempty"`
	UastAfter  string `protobuf:"bytes,5,opt,name=uast_after,json=uastAfter,proto3" json:"uast_after,omitempty"`
}

func (m *UASTChange) Reset()                    { *m = UASTChange{} }
func (m *UASTChange) String() string            { return proto.CompactTextString(m) }
func (*UASTChange) ProtoMessage()               {}
func (*UASTChange) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{9} }

func (m *UASTChange) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UASTChange) GetSrcBefore() string {
	if m != nil {
		return m.SrcBefore
	}
	return ""
}

func (m *UASTChange) GetSrcAfter() string {
	if m != nil {
		return m.SrcAfter
	}
	return ""
}

func (m *UASTChange) GetUastBefore() string {
	if m != nil {
		return m.UastBefore
	}
	return ""
}

func (m *UASTChange) GetUastAfter() string {
	if m != nil {
		return m.UastAfter
	}
	return ""
}

type UASTChangesSaverResults struct {
	Changes []*UASTChange `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
}

func (m *UASTChangesSaverResults) Reset()                    { *m = UASTChangesSaverResults{} }
func (m *UASTChangesSaverResults) String() string            { return proto.CompactTextString(m) }
func (*UASTChangesSaverResults) ProtoMessage()               {}
func (*UASTChangesSaverResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{10} }

func (m *UASTChangesSaverResults) GetChanges() []*UASTChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ShotnessRecord struct {
	Type     string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name     string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	File     string          `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Counters map[int32]int32 `protobuf:"bytes,4,rep,name=counters" json:"counters,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ShotnessRecord) Reset()                    { *m = ShotnessRecord{} }
func (m *ShotnessRecord) String() string            { return proto.CompactTextString(m) }
func (*ShotnessRecord) ProtoMessage()               {}
func (*ShotnessRecord) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{11} }

func (m *ShotnessRecord) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ShotnessRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShotnessRecord) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *ShotnessRecord) GetCounters() map[int32]int32 {
	if m != nil {
		return m.Counters
	}
	return nil
}

type ShotnessAnalysisResults struct {
	Records []*ShotnessRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *ShotnessAnalysisResults) Reset()                    { *m = ShotnessAnalysisResults{} }
func (m *ShotnessAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*ShotnessAnalysisResults) ProtoMessage()               {}
func (*ShotnessAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{12} }

func (m *ShotnessAnalysisResults) GetRecords() []*ShotnessRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type FileHistory struct {
	Commits            []string             `protobuf:"bytes,1,rep,name=commits" json:"commits,omitempty"`
	ChangesByDeveloper map[int32]*LineStats `protobuf:"bytes,2,rep,name=changes_by_developer,json=changesByDeveloper" json:"changes_by_developer,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FileHistory) Reset()                    { *m = FileHistory{} }
func (m *FileHistory) String() string            { return proto.CompactTextString(m) }
func (*FileHistory) ProtoMessage()               {}
func (*FileHistory) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{13} }

func (m *FileHistory) GetCommits() []string {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *FileHistory) GetChangesByDeveloper() map[int32]*LineStats {
	if m != nil {
		return m.ChangesByDeveloper
	}
	return nil
}

type FileHistoryResultMessage struct {
	Files map[string]*FileHistory `protobuf:"bytes,1,rep,name=files" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FileHistoryResultMessage) Reset()                    { *m = FileHistoryResultMessage{} }
func (m *FileHistoryResultMessage) String() string            { return proto.CompactTextString(m) }
func (*FileHistoryResultMessage) ProtoMessage()               {}
func (*FileHistoryResultMessage) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{14} }

func (m *FileHistoryResultMessage) GetFiles() map[string]*FileHistory {
	if m != nil {
		return m.Files
	}
	return nil
}

type LineStats struct {
	Added   int32 `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
	Removed int32 `protobuf:"varint,2,opt,name=removed,proto3" json:"removed,omitempty"`
	Changed int32 `protobuf:"varint,3,opt,name=changed,proto3" json:"changed,omitempty"`
}

func (m *LineStats) Reset()                    { *m = LineStats{} }
func (m *LineStats) String() string            { return proto.CompactTextString(m) }
func (*LineStats) ProtoMessage()               {}
func (*LineStats) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{15} }

func (m *LineStats) GetAdded() int32 {
	if m != nil {
		return m.Added
	}
	return 0
}

func (m *LineStats) GetRemoved() int32 {
	if m != nil {
		return m.Removed
	}
	return 0
}

func (m *LineStats) GetChanged() int32 {
	if m != nil {
		return m.Changed
	}
	return 0
}

type DevDay struct {
	Commits   int32                 `protobuf:"varint,1,opt,name=commits,proto3" json:"commits,omitempty"`
	Stats     *LineStats            `protobuf:"bytes,2,opt,name=stats" json:"stats,omitempty"`
	Languages map[string]*LineStats `protobuf:"bytes,3,rep,name=languages" json:"languages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DevDay) Reset()                    { *m = DevDay{} }
func (m *DevDay) String() string            { return proto.CompactTextString(m) }
func (*DevDay) ProtoMessage()               {}
func (*DevDay) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{16} }

func (m *DevDay) GetCommits() int32 {
	if m != nil {
		return m.Commits
	}
	return 0
}

func (m *DevDay) GetStats() *LineStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *DevDay) GetLanguages() map[string]*LineStats {
	if m != nil {
		return m.Languages
	}
	return nil
}

type DayDevs struct {
	Devs map[int32]*DevDay `protobuf:"bytes,1,rep,name=devs" json:"devs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DayDevs) Reset()                    { *m = DayDevs{} }
func (m *DayDevs) String() string            { return proto.CompactTextString(m) }
func (*DayDevs) ProtoMessage()               {}
func (*DayDevs) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{17} }

func (m *DayDevs) GetDevs() map[int32]*DevDay {
	if m != nil {
		return m.Devs
	}
	return nil
}

type DevsAnalysisResults struct {
	Days     map[int32]*DayDevs `protobuf:"bytes,1,rep,name=days" json:"days,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	DevIndex []string           `protobuf:"bytes,2,rep,name=dev_index,json=devIndex" json:"dev_index,omitempty"`
}

func (m *DevsAnalysisResults) Reset()                    { *m = DevsAnalysisResults{} }
func (m *DevsAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*DevsAnalysisResults) ProtoMessage()               {}
func (*DevsAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{18} }

func (m *DevsAnalysisResults) GetDays() map[int32]*DayDevs {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *DevsAnalysisResults) GetDevIndex() []string {
	if m != nil {
		return m.DevIndex
	}
	return nil
}

type Sentiment struct {
	Value    float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Comments []string `protobuf:"bytes,2,rep,name=comments" json:"comments,omitempty"`
	Commits  []string `protobuf:"bytes,3,rep,name=commits" json:"commits,omitempty"`
}

func (m *Sentiment) Reset()                    { *m = Sentiment{} }
func (m *Sentiment) String() string            { return proto.CompactTextString(m) }
func (*Sentiment) ProtoMessage()               {}
func (*Sentiment) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{19} }

func (m *Sentiment) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sentiment) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *Sentiment) GetCommits() []string {
	if m != nil {
		return m.Commits
	}
	return nil
}

type CommentSentimentResults struct {
	SentimentByDay map[int32]*Sentiment `protobuf:"bytes,1,rep,name=sentiment_by_day,json=sentimentByDay" json:"sentiment_by_day,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CommentSentimentResults) Reset()                    { *m = CommentSentimentResults{} }
func (m *CommentSentimentResults) String() string            { return proto.CompactTextString(m) }
func (*CommentSentimentResults) ProtoMessage()               {}
func (*CommentSentimentResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{20} }

func (m *CommentSentimentResults) GetSentimentByDay() map[int32]*Sentiment {
	if m != nil {
		return m.SentimentByDay
	}
	return nil
}

type CommitFile struct {
	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language string     `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Stats    *LineStats `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
}

func (m *CommitFile) Reset()                    { *m = CommitFile{} }
func (m *CommitFile) String() string            { return proto.CompactTextString(m) }
func (*CommitFile) ProtoMessage()               {}
func (*CommitFile) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{21} }

func (m *CommitFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommitFile) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *CommitFile) GetStats() *LineStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Commit struct {
	Hash         string        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	WhenUnixTime int64         `protobuf:"varint,2,opt,name=when_unix_time,json=whenUnixTime,proto3" json:"when_unix_time,omitempty"`
	Author       int32         `protobuf:"varint,3,opt,name=author,proto3" json:"author,omitempty"`
	Files        []*CommitFile `protobuf:"bytes,4,rep,name=files" json:"files,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{22} }

func (m *Commit) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Commit) GetWhenUnixTime() int64 {
	if m != nil {
		return m.WhenUnixTime
	}
	return 0
}

func (m *Commit) GetAuthor() int32 {
	if m != nil {
		return m.Author
	}
	return 0
}

func (m *Commit) GetFiles() []*CommitFile {
	if m != nil {
		return m.Files
	}
	return nil
}

type CommitsAnalysisResults struct {
	Commits     []*Commit `protobuf:"bytes,1,rep,name=commits" json:"commits,omitempty"`
	AuthorIndex []string  `protobuf:"bytes,2,rep,name=author_index,json=authorIndex" json:"author_index,omitempty"`
}

func (m *CommitsAnalysisResults) Reset()                    { *m = CommitsAnalysisResults{} }
func (m *CommitsAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*CommitsAnalysisResults) ProtoMessage()               {}
func (*CommitsAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{23} }

func (m *CommitsAnalysisResults) GetCommits() []*Commit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *CommitsAnalysisResults) GetAuthorIndex() []string {
	if m != nil {
		return m.AuthorIndex
	}
	return nil
}

type Typo struct {
	Wrong   string `protobuf:"bytes,1,opt,name=wrong,proto3" json:"wrong,omitempty"`
	Correct string `protobuf:"bytes,2,opt,name=correct,proto3" json:"correct,omitempty"`
	Commit  string `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	File    string `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Line    int32  `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
}

func (m *Typo) Reset()                    { *m = Typo{} }
func (m *Typo) String() string            { return proto.CompactTextString(m) }
func (*Typo) ProtoMessage()               {}
func (*Typo) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{24} }

func (m *Typo) GetWrong() string {
	if m != nil {
		return m.Wrong
	}
	return ""
}

func (m *Typo) GetCorrect() string {
	if m != nil {
		return m.Correct
	}
	return ""
}

func (m *Typo) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *Typo) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Typo) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

type TyposDataset struct {
	Typos []*Typo `protobuf:"bytes,1,rep,name=typos" json:"typos,omitempty"`
}

func (m *TyposDataset) Reset()                    { *m = TyposDataset{} }
func (m *TyposDataset) String() string            { return proto.CompactTextString(m) }
func (*TyposDataset) ProtoMessage()               {}
func (*TyposDataset) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{25} }

func (m *TyposDataset) GetTypos() []*Typo {
	if m != nil {
		return m.Typos
	}
	return nil
}

type AnalysisResults struct {
	Header *Metadata `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// the mapped values are dynamic messages which require the second parsing pass.
	Contents map[string][]byte `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *AnalysisResults) Reset()                    { *m = AnalysisResults{} }
func (m *AnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*AnalysisResults) ProtoMessage()               {}
func (*AnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{26} }

func (m *AnalysisResults) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AnalysisResults) GetContents() map[string][]byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "Metadata")
	proto.RegisterType((*BurndownSparseMatrixRow)(nil), "BurndownSparseMatrixRow")
	proto.RegisterType((*BurndownSparseMatrix)(nil), "BurndownSparseMatrix")
	proto.RegisterType((*FilesOwnership)(nil), "FilesOwnership")
	proto.RegisterType((*BurndownAnalysisResults)(nil), "BurndownAnalysisResults")
	proto.RegisterType((*CompressedSparseRowMatrix)(nil), "CompressedSparseRowMatrix")
	proto.RegisterType((*Couples)(nil), "Couples")
	proto.RegisterType((*TouchedFiles)(nil), "TouchedFiles")
	proto.RegisterType((*CouplesAnalysisResults)(nil), "CouplesAnalysisResults")
	proto.RegisterType((*UASTChange)(nil), "UASTChange")
	proto.RegisterType((*UASTChangesSaverResults)(nil), "UASTChangesSaverResults")
	proto.RegisterType((*ShotnessRecord)(nil), "ShotnessRecord")
	proto.RegisterType((*ShotnessAnalysisResults)(nil), "ShotnessAnalysisResults")
	proto.RegisterType((*FileHistory)(nil), "FileHistory")
	proto.RegisterType((*FileHistoryResultMessage)(nil), "FileHistoryResultMessage")
	proto.RegisterType((*LineStats)(nil), "LineStats")
	proto.RegisterType((*DevDay)(nil), "DevDay")
	proto.RegisterType((*DayDevs)(nil), "DayDevs")
	proto.RegisterType((*DevsAnalysisResults)(nil), "DevsAnalysisResults")
	proto.RegisterType((*Sentiment)(nil), "Sentiment")
	proto.RegisterType((*CommentSentimentResults)(nil), "CommentSentimentResults")
	proto.RegisterType((*CommitFile)(nil), "CommitFile")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterType((*CommitsAnalysisResults)(nil), "CommitsAnalysisResults")
	proto.RegisterType((*Typo)(nil), "Typo")
	proto.RegisterType((*TyposDataset)(nil), "TyposDataset")
	proto.RegisterType((*AnalysisResults)(nil), "AnalysisResults")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptorPb) }

var fileDescriptorPb = []byte{
	// 1568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xcd, 0x6e, 0xdb, 0xc6,
	0x13, 0x07, 0x25, 0x51, 0x1f, 0x23, 0x59, 0x4e, 0x36, 0xfe, 0xdb, 0x8c, 0x02, 0xe7, 0xaf, 0x10,
	0x6e, 0xe0, 0x36, 0x29, 0x13, 0x38, 0x3d, 0xa4, 0xe9, 0x25, 0xfe, 0x68, 0x10, 0x03, 0x71, 0xd3,
	0xd2, 0x4e, 0x7a, 0x8b, 0xb0, 0x16, 0xd7, 0x12, 0x5b, 0x69, 0x49, 0xec, 0x92, 0x92, 0x05, 0xb4,
	0xaf, 0xd2, 0x5b, 0x0f, 0x2d, 0xd0, 0x53, 0x5f, 0xa0, 0x87, 0x5e, 0x7a, 0xed, 0x43, 0x14, 0xe8,
	0x5b, 0x14, 0xfb, 0x45, 0x91, 0x0a, 0x9d, 0xa6, 0x37, 0xce, 0xcc, 0x6f, 0x77, 0x67, 0x7e, 0x33,
	0xb3, 0xb3, 0x84, 0x66, 0x7c, 0xee, 0xc5, 0x2c, 0x4a, 0x22, 0xf7, 0xaf, 0x0a, 0x34, 0x4f, 0x48,
	0x82, 0x03, 0x9c, 0x60, 0xe4, 0x40, 0x63, 0x46, 0x18, 0x0f, 0x23, 0xea, 0x58, 0x7d, 0x6b, 0xd7,
	0xf6, 0x8d, 0x88, 0x10, 0xd4, 0xc6, 0x98, 0x8f, 0x9d, 0x4a, 0xdf, 0xda, 0x6d, 0xf9, 0xf2, 0x1b,
	0xdd, 0x06, 0x60, 0x24, 0x8e, 0x78, 0x98, 0x44, 0x6c, 0xe1, 0x54, 0xa5, 0x25, 0xa7, 0x41, 0x77,
	0x61, 0xfd, 0x9c, 0x8c, 0x42, 0x3a, 0x48, 0x69, 0x78, 0x39, 0x48, 0xc2, 0x29, 0x71, 0x6a, 0x7d,
	0x6b, 0xb7, 0xea, 0xaf, 0x49, 0xf5, 0x2b, 0x1a, 0x5e, 0x9e, 0x85, 0x53, 0x82, 0x5c, 0x58, 0x23,
	0x34, 0xc8, 0xa1, 0x6c, 0x89, 0x6a, 0x13, 0x1a, 0x64, 0x18, 0x07, 0x1a, 0xc3, 0x68, 0x3a, 0x0d,
	0x13, 0xee, 0xd4, 0x95, 0x67, 0x5a, 0x44, 0x37, 0xa1, 0xc9, 0x52, 0xaa, 0x16, 0x36, 0xe4, 0xc2,
	0x06, 0x4b, 0xa9, 0x5c, 0xf4, 0x1c, 0xae, 0x1b, 0xd3, 0x20, 0x26, 0x6c, 0x10, 0x26, 0x64, 0xea,
	0x34, 0xfb, 0xd5, 0xdd, 0xf6, 0xde, 0xb6, 0x67, 0x82, 0xf6, 0x7c, 0x85, 0xfe, 0x92, 0xb0, 0xe3,
	0x84, 0x4c, 0x3f, 0xa7, 0x09, 0x5b, 0xf8, 0x5d, 0x56, 0x50, 0xf6, 0xf6, 0xe1, 0x46, 0x09, 0x0c,
	0x5d, 0x83, 0xea, 0xb7, 0x64, 0x21, 0xb9, 0x6a, 0xf9, 0xe2, 0x13, 0x6d, 0x80, 0x3d, 0xc3, 0x93,
	0x94, 0x48, 0xa2, 0x2c, 0x5f, 0x09, 0x4f, 0x2a, 0x8f, 0x2d, 0xf7, 0x11, 0x6c, 0x1d, 0xa4, 0x8c,
	0x06, 0xd1, 0x9c, 0x9e, 0xc6, 0x98, 0x71, 0x72, 0x82, 0x13, 0x16, 0x5e, 0xfa, 0xd1, 0x5c, 0x05,
	0x37, 0x49, 0xa7, 0x94, 0x3b, 0x56, 0xbf, 0xba, 0xbb, 0xe6, 0x1b, 0xd1, 0xfd, 0xd9, 0x82, 0x8d,
	0xb2, 0x55, 0x22, 0x1f, 0x14, 0x4f, 0x89, 0x3e, 0x5a, 0x7e, 0xa3, 0x1d, 0xe8, 0xd2, 0x74, 0x7a,
	0x4e, 0xd8, 0x20, 0xba, 0x18, 0xb0, 0x68, 0xce, 0xa5, 0x13, 0xb6, 0xdf, 0x51, 0xda, 0x97, 0x17,
	0x7e, 0x34, 0xe7, 0xe8, 0x23, 0xb8, 0xbe, 0x44, 0x99, 0x63, 0xab, 0x12, 0xb8, 0x6e, 0x80, 0x87,
	0x4a, 0x8d, 0xee, 0x43, 0x4d, 0xee, 0x53, 0x93, 0x9c, 0x39, 0xde, 0x15, 0x01, 0xf8, 0x12, 0xe5,
	0x7e, 0x07, 0xdd, 0x67, 0xe1, 0x84, 0xf0, 0x97, 0x73, 0x4a, 0x18, 0x1f, 0x87, 0x31, 0x7a, 0x68,
	0xd8, 0xb0, 0xe4, 0x06, 0x3d, 0xaf, 0x68, 0xf7, 0x5e, 0x0b, 0xa3, 0x62, 0x5c, 0x01, 0x7b, 0x8f,
	0x01, 0x96, 0xca, 0x3c, 0xbf, 0x76, 0x09, 0xbf, 0x76, 0x9e, 0xdf, 0xbf, 0x2b, 0x4b, 0x82, 0xf7,
	0x29, 0x9e, 0x2c, 0x78, 0xc8, 0x7d, 0xc2, 0xd3, 0x49, 0xc2, 0x51, 0x1f, 0xda, 0x23, 0x86, 0x69,
	0x3a, 0xc1, 0x2c, 0x4c, 0xcc, 0x7e, 0x79, 0x15, 0xea, 0x41, 0x93, 0xe3, 0x69, 0x3c, 0x09, 0xe9,
	0x48, 0x6f, 0x9d, 0xc9, 0xe8, 0x01, 0x34, 0x62, 0x16, 0x7d, 0x43, 0x86, 0x89, 0xe4, 0xa9, 0xbd,
	0xf7, 0xbf, 0x72, 0x22, 0x0c, 0x0a, 0xdd, 0x03, 0xfb, 0x42, 0x04, 0xaa, 0x79, 0xbb, 0x02, 0xae,
	0x30, 0xe8, 0x63, 0xa8, 0xc7, 0x24, 0x8a, 0x27, 0xa2, 0xec, 0xdf, 0x81, 0xd6, 0x20, 0x74, 0x0c,
	0x48, 0x7d, 0x0d, 0x42, 0x9a, 0x10, 0x86, 0x87, 0x89, 0xe8, 0xd6, 0xba, 0xf4, 0xab, 0xe7, 0x1d,
	0x46, 0xd3, 0x98, 0x11, 0xce, 0x49, 0xa0, 0x16, 0xfb, 0xd1, 0x5c, 0xaf, 0xbf, 0xae, 0x56, 0x1d,
	0x2f, 0x17, 0xa1, 0xc7, 0xb0, 0x2e, 0x5d, 0x18, 0x44, 0x26, 0x21, 0x4e, 0x43, 0xba, 0xb0, 0xbe,
	0x92, 0x27, 0xbf, 0x7b, 0x51, 0x90, 0xdd, 0x5f, 0x2d, 0xb8, 0x79, 0xe5, 0x51, 0x25, 0x75, 0x68,
	0xbd, 0x6f, 0x1d, 0x56, 0xca, 0xeb, 0x10, 0x41, 0x4d, 0xb4, 0xaa, 0x53, 0xed, 0x57, 0x77, 0xab,
	0x7e, 0xcd, 0xdc, 0x55, 0x21, 0x0d, 0xc2, 0xa1, 0xa6, 0xd9, 0xf6, 0x8d, 0x88, 0x36, 0xa1, 0x1e,
	0xd2, 0x20, 0x4e, 0x98, 0x64, 0xb4, 0xea, 0x6b, 0xc9, 0x3d, 0x85, 0xc6, 0x61, 0x94, 0xc6, 0x82,
	0xf4, 0x0d, 0xb0, 0x43, 0x1a, 0x90, 0x4b, 0x59, 0x98, 0x2d, 0x5f, 0x09, 0x68, 0x0f, 0xea, 0x53,
	0x19, 0x82, 0xf4, 0xe3, 0xdd, 0x7c, 0x6a, 0xa4, 0xbb, 0x03, 0x9d, 0xb3, 0x28, 0x1d, 0x8e, 0x49,
	0x20, 0x39, 0x13, 0x3b, 0xab, 0xdc, 0x5b, 0xd2, 0x29, 0x25, 0xb8, 0x7f, 0x58, 0xb0, 0xa9, 0xcf,
	0x5e, 0xad, 0xcd, 0x7b, 0xd0, 0x11, 0x98, 0xc1, 0x50, 0x99, 0x75, 0x2a, 0x9b, 0x9e, 0x86, 0xfb,
	0x6d, 0x61, 0x35, 0x7e, 0x3f, 0x80, 0xae, 0xce, 0xbe, 0x81, 0x37, 0x56, 0xe0, 0x6b, 0xca, 0x6e,
	0x16, 0x3c, 0x84, 0x8e, 0x5e, 0xa0, 0xbc, 0x52, 0xb7, 0xdf, 0x9a, 0x97, 0xf7, 0xd9, 0x6f, 0x2b,
	0x88, 0x0a, 0xe0, 0xff, 0xd0, 0x56, 0x55, 0x31, 0x09, 0x29, 0xe1, 0x4e, 0x4b, 0x86, 0x01, 0x52,
	0xf5, 0x42, 0x68, 0xdc, 0x1f, 0x2d, 0x80, 0x57, 0xfb, 0xa7, 0x67, 0x87, 0x63, 0x4c, 0x47, 0x04,
	0xdd, 0x82, 0x96, 0xf4, 0x3f, 0x77, 0x1d, 0x35, 0x85, 0xe2, 0x0b, 0x71, 0x25, 0x6d, 0x03, 0x70,
	0x36, 0x1c, 0x9c, 0x93, 0x8b, 0x88, 0x11, 0x3d, 0x3c, 0x5a, 0x9c, 0x0d, 0x0f, 0xa4, 0x42, 0xac,
	0x15, 0x66, 0x7c, 0x91, 0x10, 0xa6, 0x07, 0x48, 0x93, 0xb3, 0xe1, 0xbe, 0x90, 0x85, 0x23, 0x29,
	0xe6, 0x89, 0x59, 0x5c, 0x53, 0xf3, 0x45, 0xa8, 0xf4, 0xea, 0x6d, 0x90, 0x92, 0x5e, 0x6e, 0xab,
	0xcd, 0x85, 0x46, 0xae, 0x77, 0x9f, 0xc2, 0xd6, 0xd2, 0x4d, 0x7e, 0x8a, 0x67, 0x84, 0x19, 0xce,
	0x3f, 0x80, 0xc6, 0x50, 0xa9, 0xf5, 0xcd, 0xd4, 0xf6, 0x96, 0x50, 0xdf, 0xd8, 0xdc, 0xdf, 0x2d,
	0xe8, 0x9e, 0x8e, 0xa3, 0x84, 0x12, 0xce, 0x7d, 0x32, 0x8c, 0x58, 0x20, 0x2a, 0x31, 0x59, 0xc4,
	0xd9, 0xbd, 0x2b, 0xbe, 0xb3, 0xbb, 0xb8, 0x92, 0xbb, 0x8b, 0x11, 0xd4, 0x04, 0x09, 0x3a, 0x28,
	0xf9, 0x8d, 0x3e, 0x85, 0xe6, 0x30, 0x4a, 0x45, 0x03, 0x9a, 0x9b, 0x61, 0xdb, 0x2b, 0x6e, 0x2f,
	0xb2, 0x28, 0xed, 0xea, 0x4e, 0xcc, 0xe0, 0xbd, 0xcf, 0x60, 0xad, 0x60, 0xfa, 0x4f, 0x37, 0xe3,
	0x11, 0x6c, 0x99, 0x63, 0x56, 0x8b, 0xef, 0x43, 0x68, 0x30, 0x79, 0xb2, 0x21, 0x62, 0x7d, 0xc5,
	0x23, 0xdf, 0xd8, 0xdd, 0x3f, 0x2d, 0x68, 0x8b, 0x0a, 0x79, 0x1e, 0x72, 0x39, 0xdd, 0x73, 0x13,
	0x59, 0x35, 0x51, 0x36, 0x91, 0x5f, 0xc3, 0x86, 0x66, 0x70, 0x70, 0xbe, 0x18, 0x04, 0x64, 0x46,
	0x26, 0x51, 0x4c, 0x98, 0x53, 0x91, 0x27, 0xec, 0x78, 0xb9, 0x5d, 0x3c, 0x9d, 0x9d, 0x83, 0xc5,
	0x91, 0x81, 0xa9, 0xd0, 0xd1, 0xf0, 0x2d, 0x43, 0xef, 0x2b, 0xd8, 0xba, 0x02, 0x5e, 0x42, 0x47,
	0x3f, 0x4f, 0x47, 0x7b, 0x0f, 0x3c, 0x51, 0xbc, 0xa7, 0x09, 0x4e, 0x78, 0x9e, 0x9a, 0x1f, 0x2c,
	0x70, 0x72, 0xee, 0x28, 0x5a, 0x4e, 0x08, 0xe7, 0x78, 0x44, 0xd0, 0x93, 0x7c, 0x2b, 0xaf, 0x38,
	0x5e, 0x40, 0xaa, 0xeb, 0x52, 0xcf, 0x31, 0xb9, 0xa4, 0xf7, 0x0c, 0x60, 0xa9, 0x2c, 0x79, 0x27,
	0xb8, 0x45, 0xf7, 0x3a, 0x85, 0xbd, 0x73, 0x0e, 0xbe, 0x82, 0x56, 0xe6, 0xb8, 0x48, 0x31, 0x0e,
	0x02, 0x12, 0xe8, 0x38, 0x95, 0x20, 0x12, 0xc1, 0xc8, 0x34, 0x9a, 0x91, 0x40, 0xa7, 0xde, 0x88,
	0x32, 0x45, 0x92, 0xb0, 0x40, 0x0f, 0x78, 0x23, 0x8a, 0xca, 0xae, 0x1f, 0x91, 0xd9, 0x11, 0x5e,
	0xc9, 0x63, 0xe1, 0x65, 0xd5, 0x07, 0x9b, 0x8b, 0x73, 0xcb, 0x28, 0x94, 0x06, 0xf4, 0x09, 0xb4,
	0x26, 0x98, 0x8e, 0x52, 0x2c, 0x3a, 0xa9, 0x2a, 0x59, 0xda, 0xf4, 0xd4, 0xbe, 0xde, 0x0b, 0x63,
	0x50, 0xbc, 0x2c, 0x81, 0xbd, 0xe7, 0xd0, 0x2d, 0x1a, 0x4b, 0xf8, 0x79, 0xbf, 0xf4, 0x71, 0x68,
	0x1c, 0x61, 0x51, 0x0b, 0x1c, 0xdd, 0x85, 0x5a, 0x40, 0x66, 0x26, 0x57, 0xc8, 0xd3, 0x7a, 0xe1,
	0x8d, 0xf6, 0x40, 0xda, 0x7b, 0x4f, 0xa1, 0x95, 0xa9, 0x4a, 0xca, 0x66, 0xbb, 0x78, 0x6e, 0x43,
	0x47, 0x93, 0x3f, 0xf4, 0x27, 0x0b, 0x6e, 0x88, 0x2d, 0x56, 0x7b, 0x69, 0x4f, 0x0c, 0xa9, 0x85,
	0xf1, 0xe0, 0xb6, 0x57, 0x82, 0x11, 0x5e, 0x65, 0xde, 0xe0, 0x05, 0x17, 0x17, 0x60, 0x40, 0x66,
	0x03, 0x35, 0x8b, 0x2a, 0xb2, 0x8d, 0x9a, 0x01, 0x99, 0x1d, 0x0b, 0xb9, 0xb7, 0x0f, 0xad, 0x0c,
	0x5f, 0xe2, 0xea, 0xed, 0xa2, 0xab, 0x4d, 0x13, 0x72, 0xde, 0xd7, 0xaf, 0xa1, 0x75, 0x4a, 0xa8,
	0x78, 0x00, 0xd3, 0x64, 0x79, 0x43, 0x88, 0x4d, 0x2a, 0x1a, 0x26, 0x5e, 0x3e, 0x22, 0xe1, 0x84,
	0xca, 0x44, 0x4b, 0x0f, 0x8c, 0x9c, 0xaf, 0x8d, 0x6a, 0xa1, 0xc7, 0xdd, 0xdf, 0x2c, 0xd8, 0x3a,
	0x54, 0xb0, 0xec, 0x00, 0x43, 0xc4, 0x6b, 0xb8, 0xc6, 0x8d, 0x4e, 0xde, 0x00, 0x78, 0xa1, 0x49,
	0xb9, 0xef, 0x5d, 0xb1, 0xc6, 0xcb, 0x14, 0x07, 0x8b, 0x23, 0xbc, 0xd0, 0x8f, 0x70, 0x5e, 0x50,
	0xf6, 0x4e, 0xe0, 0x46, 0x09, 0xec, 0x7d, 0x7a, 0x7f, 0x79, 0x5c, 0x8e, 0x9b, 0x37, 0x00, 0x87,
	0x32, 0x1a, 0xd1, 0x7a, 0xa5, 0x0f, 0xea, 0x1e, 0x34, 0x4d, 0xd5, 0x9a, 0xe9, 0x64, 0xe4, 0x65,
	0x73, 0xd4, 0xae, 0x68, 0x0e, 0xf7, 0x7b, 0xa8, 0xab, 0xfd, 0xb3, 0x9f, 0x27, 0x2b, 0xf7, 0xf3,
	0xb4, 0x03, 0xdd, 0xf9, 0x98, 0xe4, 0xff, 0x8d, 0x2a, 0xf2, 0xe7, 0xa5, 0x23, 0xb4, 0xd9, 0x6f,
	0xcf, 0x26, 0xd4, 0x71, 0x9a, 0x8c, 0x23, 0xa6, 0x1b, 0x58, 0x4b, 0xe8, 0x4e, 0xf1, 0x85, 0xd9,
	0xf6, 0x96, 0x91, 0x98, 0x27, 0xc7, 0x1b, 0xf1, 0xe2, 0x90, 0xc9, 0x5a, 0x2d, 0xd4, 0x3b, 0xc5,
	0x9b, 0x5b, 0x54, 0xb9, 0x42, 0x2e, 0x5b, 0xff, 0x0e, 0x74, 0xd4, 0x49, 0x85, 0xd2, 0x6c, 0x2b,
	0x9d, 0xac, 0x4e, 0x77, 0x06, 0xb5, 0xb3, 0x45, 0x1c, 0x89, 0xaa, 0x9a, 0xb3, 0x88, 0x8e, 0x74,
	0x74, 0x4a, 0x50, 0x95, 0xc3, 0x98, 0x78, 0x33, 0xab, 0xb1, 0x68, 0x44, 0x11, 0x92, 0x3a, 0x45,
	0x53, 0xaa, 0xa5, 0x6c, 0x62, 0xd6, 0x72, 0x13, 0x13, 0x41, 0x4d, 0xbc, 0x42, 0xe4, 0x6c, 0xb7,
	0x7d, 0xf9, 0xed, 0xde, 0x83, 0x8e, 0x38, 0x97, 0x1f, 0xe1, 0x04, 0x73, 0x92, 0xa0, 0x5b, 0x60,
	0x27, 0x42, 0xd6, 0xb1, 0xd8, 0x9e, 0xb0, 0xfa, 0x4a, 0xe7, 0xfe, 0x62, 0xc1, 0xfa, 0xdb, 0xe1,
	0xd7, 0xc7, 0x04, 0x07, 0x84, 0x49, 0x8f, 0xdb, 0x7b, 0xad, 0xec, 0x57, 0xd0, 0xd7, 0x06, 0xf4,
	0x44, 0xf4, 0x04, 0x4d, 0xb2, 0x9e, 0x10, 0xed, 0xbc, 0xda, 0xca, 0x87, 0x1a, 0x90, 0x8d, 0x6a,
	0x25, 0xaa, 0x51, 0x9d, 0x33, 0xfd, 0xdb, 0x4f, 0x62, 0x27, 0x57, 0x93, 0xe7, 0x75, 0xf9, 0x53,
	0xfe, 0xe8, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xab, 0x56, 0x6d, 0xa0, 0x0f, 0x00, 0x00,
}
