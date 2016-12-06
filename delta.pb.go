// Code generated by protoc-gen-go.
// source: delta.proto
// DO NOT EDIT!

/*
Package delta is a generated protocol buffer package.

It is generated from these files:
	delta.proto

It has these top-level messages:
	Point
	Span
	Network
	Mark
	Marks
	Station
	Stations
	Monument
	Offset
	Equipment
	InstalledAntenna
	InstalledRadome
	Firmware
	Receiver
	DeployedReceiver
	Session
	InstalledMetSensor
*/
package delta

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	// Latitude - Geographical latitude of the point for the given datum.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude - Geographical longitude of the point for the given datum.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	// Elevation - Height in meters of the point for the given datum.
	Elevation float64 `protobuf:"fixed64,3,opt,name=elevation" json:"elevation,omitempty"`
	// Datum - Geographical reference system used for the latitude, longitude & elevation.
	Datum string `protobuf:"bytes,4,opt,name=datum" json:"datum,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Point) GetElevation() float64 {
	if m != nil {
		return m.Elevation
	}
	return 0
}

func (m *Point) GetDatum() string {
	if m != nil {
		return m.Datum
	}
	return ""
}

type Span struct {
	// Start - time in Unix seconds.
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	// End - time in Unix seconds.  A future date of 9999-01-01T00:00:00Z is used to indicate still open.
	End int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Span) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Span) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type Network struct {
	Code        string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	External    string `protobuf:"bytes,2,opt,name=external" json:"external,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Restricted  bool   `protobuf:"varint,4,opt,name=restricted" json:"restricted,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Network) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Network) GetExternal() string {
	if m != nil {
		return m.External
	}
	return ""
}

func (m *Network) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Network) GetRestricted() bool {
	if m != nil {
		return m.Restricted
	}
	return false
}

type Mark struct {
	// Code used to uniquely identify GNSS Mark.
	Code               string                `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name               string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Network            *Network              `protobuf:"bytes,3,opt,name=network" json:"network,omitempty"`
	Point              *Point                `protobuf:"bytes,4,opt,name=point" json:"point,omitempty"`
	Span               *Span                 `protobuf:"bytes,5,opt,name=span" json:"span,omitempty"`
	Monument           *Monument             `protobuf:"bytes,6,opt,name=monument" json:"monument,omitempty"`
	InstalledAntenna   []*InstalledAntenna   `protobuf:"bytes,7,rep,name=installed_antenna,json=installedAntenna" json:"installed_antenna,omitempty"`
	InstalledRadome    []*InstalledRadome    `protobuf:"bytes,8,rep,name=installed_radome,json=installedRadome" json:"installed_radome,omitempty"`
	DeployedReceiver   []*DeployedReceiver   `protobuf:"bytes,9,rep,name=deployed_receiver,json=deployedReceiver" json:"deployed_receiver,omitempty"`
	Session            []*Session            `protobuf:"bytes,10,rep,name=session" json:"session,omitempty"`
	InstalledMetSensor []*InstalledMetSensor `protobuf:"bytes,11,rep,name=installed_met_sensor,json=installedMetSensor" json:"installed_met_sensor,omitempty"`
}

func (m *Mark) Reset()                    { *m = Mark{} }
func (m *Mark) String() string            { return proto.CompactTextString(m) }
func (*Mark) ProtoMessage()               {}
func (*Mark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Mark) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Mark) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mark) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Mark) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Mark) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *Mark) GetMonument() *Monument {
	if m != nil {
		return m.Monument
	}
	return nil
}

func (m *Mark) GetInstalledAntenna() []*InstalledAntenna {
	if m != nil {
		return m.InstalledAntenna
	}
	return nil
}

func (m *Mark) GetInstalledRadome() []*InstalledRadome {
	if m != nil {
		return m.InstalledRadome
	}
	return nil
}

func (m *Mark) GetDeployedReceiver() []*DeployedReceiver {
	if m != nil {
		return m.DeployedReceiver
	}
	return nil
}

func (m *Mark) GetSession() []*Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Mark) GetInstalledMetSensor() []*InstalledMetSensor {
	if m != nil {
		return m.InstalledMetSensor
	}
	return nil
}

type Marks struct {
	Marks map[string]*Mark `protobuf:"bytes,1,rep,name=marks" json:"marks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Marks) Reset()                    { *m = Marks{} }
func (m *Marks) String() string            { return proto.CompactTextString(m) }
func (*Marks) ProtoMessage()               {}
func (*Marks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Marks) GetMarks() map[string]*Mark {
	if m != nil {
		return m.Marks
	}
	return nil
}

type Station struct {
	// Code used to uniquely identify the Station.
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// Name used to describe the general geographical location of the Station.
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Network *Network `protobuf:"bytes,3,opt,name=network" json:"network,omitempty"`
	Point   *Point   `protobuf:"bytes,4,opt,name=point" json:"point,omitempty"`
	Span    *Span    `protobuf:"bytes,5,opt,name=span" json:"span,omitempty"`
}

func (m *Station) Reset()                    { *m = Station{} }
func (m *Station) String() string            { return proto.CompactTextString(m) }
func (*Station) ProtoMessage()               {}
func (*Station) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Station) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Station) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Station) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Station) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Station) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Stations struct {
	Stations map[string]*Station `protobuf:"bytes,1,rep,name=stations" json:"stations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Stations) Reset()                    { *m = Stations{} }
func (m *Stations) String() string            { return proto.CompactTextString(m) }
func (*Stations) ProtoMessage()               {}
func (*Stations) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Stations) GetStations() map[string]*Station {
	if m != nil {
		return m.Stations
	}
	return nil
}

type Monument struct {
	DomesNumber        string  `protobuf:"bytes,1,opt,name=domes_number,json=domesNumber" json:"domes_number,omitempty"`
	MarkType           string  `protobuf:"bytes,2,opt,name=mark_type,json=markType" json:"mark_type,omitempty"`
	Type               string  `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	GroundRelationship float64 `protobuf:"fixed64,4,opt,name=ground_relationship,json=groundRelationship" json:"ground_relationship,omitempty"`
	FoundationType     string  `protobuf:"bytes,5,opt,name=foundation_type,json=foundationType" json:"foundation_type,omitempty"`
	FoundationDepth    float64 `protobuf:"fixed64,6,opt,name=foundation_depth,json=foundationDepth" json:"foundation_depth,omitempty"`
	Bedrock            string  `protobuf:"bytes,7,opt,name=bedrock" json:"bedrock,omitempty"`
	Geology            string  `protobuf:"bytes,8,opt,name=geology" json:"geology,omitempty"`
	Span               *Span   `protobuf:"bytes,9,opt,name=span" json:"span,omitempty"`
}

func (m *Monument) Reset()                    { *m = Monument{} }
func (m *Monument) String() string            { return proto.CompactTextString(m) }
func (*Monument) ProtoMessage()               {}
func (*Monument) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Monument) GetDomesNumber() string {
	if m != nil {
		return m.DomesNumber
	}
	return ""
}

func (m *Monument) GetMarkType() string {
	if m != nil {
		return m.MarkType
	}
	return ""
}

func (m *Monument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Monument) GetGroundRelationship() float64 {
	if m != nil {
		return m.GroundRelationship
	}
	return 0
}

func (m *Monument) GetFoundationType() string {
	if m != nil {
		return m.FoundationType
	}
	return ""
}

func (m *Monument) GetFoundationDepth() float64 {
	if m != nil {
		return m.FoundationDepth
	}
	return 0
}

func (m *Monument) GetBedrock() string {
	if m != nil {
		return m.Bedrock
	}
	return ""
}

func (m *Monument) GetGeology() string {
	if m != nil {
		return m.Geology
	}
	return ""
}

func (m *Monument) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Offset struct {
	Vertical float64 `protobuf:"fixed64,1,opt,name=vertical" json:"vertical,omitempty"`
	North    float64 `protobuf:"fixed64,2,opt,name=north" json:"north,omitempty"`
	East     float64 `protobuf:"fixed64,3,opt,name=east" json:"east,omitempty"`
}

func (m *Offset) Reset()                    { *m = Offset{} }
func (m *Offset) String() string            { return proto.CompactTextString(m) }
func (*Offset) ProtoMessage()               {}
func (*Offset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Offset) GetVertical() float64 {
	if m != nil {
		return m.Vertical
	}
	return 0
}

func (m *Offset) GetNorth() float64 {
	if m != nil {
		return m.North
	}
	return 0
}

func (m *Offset) GetEast() float64 {
	if m != nil {
		return m.East
	}
	return 0
}

type Equipment struct {
	Make   string `protobuf:"bytes,1,opt,name=Make,json=make" json:"Make,omitempty"`
	Model  string `protobuf:"bytes,2,opt,name=Model,json=model" json:"Model,omitempty"`
	Serial string `protobuf:"bytes,3,opt,name=Serial,json=serial" json:"Serial,omitempty"`
}

func (m *Equipment) Reset()                    { *m = Equipment{} }
func (m *Equipment) String() string            { return proto.CompactTextString(m) }
func (*Equipment) ProtoMessage()               {}
func (*Equipment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Equipment) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Equipment) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Equipment) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

type InstalledAntenna struct {
	Equipment *Equipment `protobuf:"bytes,1,opt,name=equipment" json:"equipment,omitempty"`
	Offset    *Offset    `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	Span      *Span      `protobuf:"bytes,3,opt,name=span" json:"span,omitempty"`
	Azimuth   float64    `protobuf:"fixed64,4,opt,name=azimuth" json:"azimuth,omitempty"`
}

func (m *InstalledAntenna) Reset()                    { *m = InstalledAntenna{} }
func (m *InstalledAntenna) String() string            { return proto.CompactTextString(m) }
func (*InstalledAntenna) ProtoMessage()               {}
func (*InstalledAntenna) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *InstalledAntenna) GetEquipment() *Equipment {
	if m != nil {
		return m.Equipment
	}
	return nil
}

func (m *InstalledAntenna) GetOffset() *Offset {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *InstalledAntenna) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *InstalledAntenna) GetAzimuth() float64 {
	if m != nil {
		return m.Azimuth
	}
	return 0
}

type InstalledRadome struct {
	Equipment *Equipment `protobuf:"bytes,1,opt,name=equipment" json:"equipment,omitempty"`
	Span      *Span      `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
}

func (m *InstalledRadome) Reset()                    { *m = InstalledRadome{} }
func (m *InstalledRadome) String() string            { return proto.CompactTextString(m) }
func (*InstalledRadome) ProtoMessage()               {}
func (*InstalledRadome) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *InstalledRadome) GetEquipment() *Equipment {
	if m != nil {
		return m.Equipment
	}
	return nil
}

func (m *InstalledRadome) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Firmware struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Notes   string `protobuf:"bytes,2,opt,name=notes" json:"notes,omitempty"`
	Span    *Span  `protobuf:"bytes,3,opt,name=span" json:"span,omitempty"`
}

func (m *Firmware) Reset()                    { *m = Firmware{} }
func (m *Firmware) String() string            { return proto.CompactTextString(m) }
func (*Firmware) ProtoMessage()               {}
func (*Firmware) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Firmware) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Firmware) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Firmware) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Receiver struct {
	Equipment *Equipment  `protobuf:"bytes,1,opt,name=equipment" json:"equipment,omitempty"`
	Firmware  []*Firmware `protobuf:"bytes,2,rep,name=firmware" json:"firmware,omitempty"`
}

func (m *Receiver) Reset()                    { *m = Receiver{} }
func (m *Receiver) String() string            { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()               {}
func (*Receiver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Receiver) GetEquipment() *Equipment {
	if m != nil {
		return m.Equipment
	}
	return nil
}

func (m *Receiver) GetFirmware() []*Firmware {
	if m != nil {
		return m.Firmware
	}
	return nil
}

type DeployedReceiver struct {
	Receiver *Receiver `protobuf:"bytes,1,opt,name=receiver" json:"receiver,omitempty"`
	Span     *Span     `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
}

func (m *DeployedReceiver) Reset()                    { *m = DeployedReceiver{} }
func (m *DeployedReceiver) String() string            { return proto.CompactTextString(m) }
func (*DeployedReceiver) ProtoMessage()               {}
func (*DeployedReceiver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DeployedReceiver) GetReceiver() *Receiver {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *DeployedReceiver) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Session struct {
	Operator        string `protobuf:"bytes,1,opt,name=operator" json:"operator,omitempty"`
	Agency          string `protobuf:"bytes,2,opt,name=agency" json:"agency,omitempty"`
	Model           string `protobuf:"bytes,3,opt,name=model" json:"model,omitempty"`
	SatelliteSystem string `protobuf:"bytes,4,opt,name=satellite_system,json=satelliteSystem" json:"satellite_system,omitempty"`
	// Interval - sample interval in nanoseconds
	Interval      int64   `protobuf:"varint,5,opt,name=interval" json:"interval,omitempty"`
	ElevationMask float64 `protobuf:"fixed64,6,opt,name=elevation_mask,json=elevationMask" json:"elevation_mask,omitempty"`
	HeaderComment string  `protobuf:"bytes,7,opt,name=header_comment,json=headerComment" json:"header_comment,omitempty"`
	Span          *Span   `protobuf:"bytes,8,opt,name=span" json:"span,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Session) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *Session) GetAgency() string {
	if m != nil {
		return m.Agency
	}
	return ""
}

func (m *Session) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Session) GetSatelliteSystem() string {
	if m != nil {
		return m.SatelliteSystem
	}
	return ""
}

func (m *Session) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *Session) GetElevationMask() float64 {
	if m != nil {
		return m.ElevationMask
	}
	return 0
}

func (m *Session) GetHeaderComment() string {
	if m != nil {
		return m.HeaderComment
	}
	return ""
}

func (m *Session) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type InstalledMetSensor struct {
	Equipment  *Equipment `protobuf:"bytes,1,opt,name=equipment" json:"equipment,omitempty"`
	Span       *Span      `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
	Point      *Point     `protobuf:"bytes,3,opt,name=point" json:"point,omitempty"`
	IMSComment string     `protobuf:"bytes,4,opt,name=iMS_comment,json=iMSComment" json:"iMS_comment,omitempty"`
}

func (m *InstalledMetSensor) Reset()                    { *m = InstalledMetSensor{} }
func (m *InstalledMetSensor) String() string            { return proto.CompactTextString(m) }
func (*InstalledMetSensor) ProtoMessage()               {}
func (*InstalledMetSensor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *InstalledMetSensor) GetEquipment() *Equipment {
	if m != nil {
		return m.Equipment
	}
	return nil
}

func (m *InstalledMetSensor) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *InstalledMetSensor) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *InstalledMetSensor) GetIMSComment() string {
	if m != nil {
		return m.IMSComment
	}
	return ""
}

func init() {
	proto.RegisterType((*Point)(nil), "delta.Point")
	proto.RegisterType((*Span)(nil), "delta.Span")
	proto.RegisterType((*Network)(nil), "delta.Network")
	proto.RegisterType((*Mark)(nil), "delta.Mark")
	proto.RegisterType((*Marks)(nil), "delta.Marks")
	proto.RegisterType((*Station)(nil), "delta.Station")
	proto.RegisterType((*Stations)(nil), "delta.Stations")
	proto.RegisterType((*Monument)(nil), "delta.Monument")
	proto.RegisterType((*Offset)(nil), "delta.Offset")
	proto.RegisterType((*Equipment)(nil), "delta.Equipment")
	proto.RegisterType((*InstalledAntenna)(nil), "delta.InstalledAntenna")
	proto.RegisterType((*InstalledRadome)(nil), "delta.InstalledRadome")
	proto.RegisterType((*Firmware)(nil), "delta.Firmware")
	proto.RegisterType((*Receiver)(nil), "delta.Receiver")
	proto.RegisterType((*DeployedReceiver)(nil), "delta.DeployedReceiver")
	proto.RegisterType((*Session)(nil), "delta.Session")
	proto.RegisterType((*InstalledMetSensor)(nil), "delta.InstalledMetSensor")
}

func init() { proto.RegisterFile("delta.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1061 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x8e, 0x23, 0x35,
	0x10, 0x56, 0xe7, 0xb7, 0x53, 0xd9, 0x99, 0x09, 0x66, 0xb5, 0x34, 0xcb, 0xdf, 0x6c, 0x8b, 0x15,
	0xb3, 0x5a, 0x31, 0x48, 0xc3, 0x05, 0xb8, 0x2d, 0xcc, 0x20, 0xa1, 0x55, 0x06, 0xe4, 0x70, 0x82,
	0x43, 0xf0, 0xa4, 0x6b, 0x32, 0xad, 0x74, 0xb7, 0x1b, 0xdb, 0xc9, 0x12, 0x10, 0x0f, 0xc0, 0x0d,
	0xf1, 0x04, 0x1c, 0x38, 0xf1, 0x04, 0x3c, 0x1e, 0x72, 0xd9, 0xed, 0xfc, 0x2c, 0xc3, 0x6a, 0x25,
	0x0e, 0x5c, 0x22, 0xd7, 0xf7, 0x55, 0xbb, 0xca, 0x9f, 0xcb, 0x55, 0x81, 0x61, 0x86, 0x85, 0x11,
	0xa7, 0xb5, 0x92, 0x46, 0xb2, 0x2e, 0x19, 0xe9, 0x12, 0xba, 0x5f, 0xc9, 0xbc, 0x32, 0xec, 0x3e,
	0xc4, 0x85, 0x30, 0xb9, 0x59, 0x66, 0x98, 0x44, 0xc7, 0xd1, 0x49, 0xc4, 0x83, 0xcd, 0xde, 0x84,
	0x41, 0x21, 0xab, 0xb9, 0x23, 0x5b, 0x44, 0x6e, 0x00, 0xcb, 0x62, 0x81, 0x2b, 0x61, 0x72, 0x59,
	0x25, 0x6d, 0xc7, 0x06, 0x80, 0xdd, 0x85, 0x6e, 0x26, 0xcc, 0xb2, 0x4c, 0x3a, 0xc7, 0xd1, 0xc9,
	0x80, 0x3b, 0x23, 0x3d, 0x85, 0xce, 0xa4, 0x16, 0xc4, 0x6a, 0x23, 0x94, 0xa1, 0x90, 0x6d, 0xee,
	0x0c, 0x36, 0x82, 0x36, 0x56, 0x19, 0x45, 0x6a, 0x73, 0xbb, 0x4c, 0x7f, 0x82, 0xfe, 0x25, 0x9a,
	0x67, 0x52, 0x2d, 0x18, 0x83, 0xce, 0x4c, 0xfa, 0x24, 0x07, 0x9c, 0xd6, 0x36, 0x79, 0xfc, 0xc1,
	0xa0, 0xaa, 0x44, 0x41, 0x5f, 0x0d, 0x78, 0xb0, 0xd9, 0xb1, 0x3d, 0xb7, 0x9e, 0xa9, 0xbc, 0x0e,
	0x09, 0x0e, 0xf8, 0x36, 0xc4, 0xde, 0x06, 0x50, 0xa8, 0x8d, 0xca, 0x67, 0x06, 0x33, 0xca, 0x33,
	0xe6, 0x5b, 0x48, 0xfa, 0x6b, 0x07, 0x3a, 0x63, 0x71, 0x4b, 0x68, 0x06, 0x9d, 0x4a, 0x94, 0xe8,
	0xc3, 0xd2, 0x9a, 0x9d, 0x40, 0xbf, 0x72, 0xd9, 0x52, 0xb8, 0xe1, 0xd9, 0xe1, 0xa9, 0x93, 0xde,
	0x9f, 0x81, 0x37, 0x34, 0x4b, 0xa1, 0x5b, 0x5b, 0xf9, 0x29, 0xea, 0xf0, 0xec, 0x8e, 0xf7, 0xa3,
	0x2b, 0xe1, 0x8e, 0x62, 0xef, 0x40, 0x47, 0xd7, 0xa2, 0x4a, 0xba, 0xe4, 0x32, 0xf4, 0x2e, 0x56,
	0x3e, 0x4e, 0x04, 0x7b, 0x0c, 0x71, 0x29, 0xab, 0x65, 0x89, 0x95, 0x49, 0x7a, 0xe4, 0x74, 0xe4,
	0x9d, 0xc6, 0x1e, 0xe6, 0xc1, 0x81, 0x9d, 0xc3, 0x2b, 0x79, 0xa5, 0x8d, 0x28, 0x0a, 0xcc, 0xa6,
	0xa2, 0x32, 0x58, 0x55, 0x22, 0xe9, 0x1f, 0xb7, 0x4f, 0x86, 0x67, 0xaf, 0xf9, 0xaf, 0xbe, 0x68,
	0xf8, 0x27, 0x8e, 0xe6, 0xa3, 0x7c, 0x0f, 0x61, 0x4f, 0x60, 0x83, 0x4d, 0x95, 0xc8, 0x64, 0x89,
	0x49, 0x4c, 0x9b, 0xdc, 0xdb, 0xdf, 0x84, 0x13, 0xcb, 0x8f, 0xf2, 0x5d, 0xc0, 0x26, 0x92, 0x61,
	0x5d, 0xc8, 0xb5, 0xdd, 0x01, 0x67, 0x98, 0xaf, 0x50, 0x25, 0x83, 0x9d, 0x44, 0xce, 0x3d, 0xcf,
	0x3d, 0xcd, 0x47, 0xd9, 0x1e, 0x62, 0xa5, 0xd6, 0xa8, 0xb5, 0xbd, 0x59, 0xa0, 0x6f, 0x1b, 0xa9,
	0x27, 0x0e, 0xe5, 0x0d, 0xcd, 0x9e, 0xc2, 0xdd, 0x4d, 0xca, 0x25, 0x9a, 0xa9, 0xc6, 0x4a, 0x4b,
	0x95, 0x0c, 0xe9, 0xb3, 0xd7, 0xf7, 0xd3, 0x1e, 0xa3, 0x99, 0x90, 0x03, 0x67, 0xf9, 0x73, 0x58,
	0xfa, 0x33, 0x74, 0x6d, 0x45, 0x68, 0xf6, 0x3e, 0x74, 0x4b, 0xbb, 0x48, 0xa2, 0x9d, 0xcc, 0x89,
	0x74, 0xbf, 0x17, 0x95, 0x51, 0x6b, 0xee, 0xbc, 0xee, 0x5f, 0x00, 0x6c, 0x40, 0x5b, 0xe7, 0x0b,
	0x5c, 0xfb, 0x72, 0xb2, 0x4b, 0xf6, 0x00, 0xba, 0x2b, 0x51, 0x2c, 0x5d, 0x39, 0x6d, 0x2e, 0xdb,
	0x7e, 0xc3, 0x1d, 0xf3, 0x49, 0xeb, 0xa3, 0x28, 0xfd, 0x3d, 0x82, 0xfe, 0xc4, 0xb8, 0x07, 0xf6,
	0xff, 0x2c, 0xca, 0xf4, 0xb7, 0x08, 0x62, 0x9f, 0xa2, 0x66, 0x1f, 0x43, 0xac, 0xfd, 0xda, 0x0b,
	0xf5, 0x56, 0xf3, 0x85, 0x87, 0xc3, 0xc2, 0xc9, 0x15, 0xdc, 0xef, 0x3f, 0x85, 0x83, 0x1d, 0xea,
	0x1f, 0x44, 0x7b, 0x77, 0x57, 0xb4, 0xc3, 0xdd, 0xad, 0xb7, 0x75, 0xfb, 0xab, 0x05, 0x71, 0xf3,
	0x26, 0xd8, 0x03, 0xb8, 0x63, 0x0b, 0x51, 0x4f, 0xab, 0x65, 0x79, 0x85, 0xca, 0xef, 0x38, 0x24,
	0xec, 0x92, 0x20, 0xf6, 0x06, 0x0c, 0xec, 0xbd, 0x4d, 0xcd, 0xba, 0x6e, 0xc4, 0x8c, 0x2d, 0xf0,
	0xf5, 0xba, 0x26, 0x91, 0x09, 0x77, 0x1d, 0x85, 0xd6, 0xec, 0x03, 0x78, 0x75, 0xae, 0xe4, 0xb2,
	0xb2, 0x25, 0x5d, 0xb8, 0xac, 0x6f, 0xf2, 0x9a, 0x84, 0x8c, 0x38, 0x73, 0x14, 0xdf, 0x62, 0xd8,
	0x7b, 0x70, 0x74, 0x6d, 0x41, 0x42, 0x5c, 0x9c, 0x2e, 0xed, 0x77, 0xb8, 0x81, 0x29, 0xda, 0x23,
	0x18, 0x6d, 0x39, 0x66, 0x58, 0x9b, 0x1b, 0x7a, 0xec, 0x11, 0xdf, 0xda, 0xe0, 0xdc, 0xc2, 0x2c,
	0x81, 0xfe, 0x15, 0x66, 0x4a, 0xce, 0x16, 0x49, 0x9f, 0xf6, 0x6a, 0x4c, 0xcb, 0xcc, 0x51, 0x16,
	0x72, 0xbe, 0x4e, 0x62, 0xc7, 0x78, 0x33, 0xdc, 0xe7, 0xe0, 0xb6, 0xfb, 0xbc, 0x84, 0xde, 0x97,
	0xd7, 0xd7, 0x1a, 0x69, 0x52, 0xac, 0x50, 0x99, 0x7c, 0x26, 0x8a, 0x66, 0x52, 0x34, 0xb6, 0xed,
	0xe7, 0x95, 0x54, 0xe6, 0xc6, 0x4f, 0x09, 0x67, 0x58, 0xa5, 0x50, 0x68, 0xe3, 0x87, 0x03, 0xad,
	0xd3, 0x31, 0x0c, 0x2e, 0xbe, 0x5f, 0xe6, 0x35, 0x5d, 0x05, 0xb3, 0x0d, 0x76, 0x11, 0x6a, 0xb8,
	0x14, 0x0b, 0xb4, 0x5b, 0x8d, 0x65, 0x86, 0x4d, 0x43, 0xef, 0x96, 0xd6, 0x60, 0xf7, 0xa0, 0x37,
	0x41, 0x95, 0x8b, 0xc2, 0xcb, 0xde, 0xd3, 0x64, 0xa5, 0x7f, 0x44, 0x30, 0xda, 0xef, 0x5b, 0xec,
	0x14, 0x06, 0xd8, 0xc4, 0xa0, 0xbd, 0x87, 0x67, 0x23, 0x7f, 0xb2, 0x10, 0x9b, 0x6f, 0x5c, 0xd8,
	0x43, 0xe8, 0x49, 0x3a, 0xa3, 0xaf, 0xa4, 0x03, 0xef, 0xec, 0x0e, 0xce, 0x3d, 0x19, 0xb4, 0x6a,
	0xdf, 0xd6, 0x90, 0x13, 0xe8, 0x8b, 0x1f, 0xf3, 0x72, 0x69, 0x6e, 0xfc, 0xcd, 0x37, 0x66, 0x7a,
	0x05, 0x47, 0x7b, 0x8d, 0xf1, 0xa5, 0x93, 0x6c, 0xa2, 0xb7, 0x6e, 0xbb, 0xa9, 0x6f, 0x21, 0xfe,
	0x3c, 0x57, 0xe5, 0x33, 0xa1, 0xd0, 0x66, 0xb2, 0x42, 0x45, 0xed, 0xd1, 0x69, 0xdb, 0x98, 0xee,
	0xa6, 0x0c, 0xea, 0x46, 0x5e, 0x32, 0x5e, 0x78, 0xb4, 0x74, 0x0e, 0x71, 0xe8, 0xbd, 0x2f, 0x9b,
	0xf9, 0x63, 0x88, 0xaf, 0x7d, 0x62, 0x49, 0x8b, 0xba, 0x40, 0x33, 0xa7, 0x9a, 0x7c, 0x79, 0x70,
	0x48, 0xbf, 0x83, 0xd1, 0x7e, 0xfb, 0xb7, 0x1b, 0x84, 0x49, 0x11, 0xed, 0x0c, 0xba, 0x30, 0x21,
	0x82, 0xc3, 0x8b, 0x75, 0xfa, 0xa5, 0x05, 0x7d, 0x3f, 0x25, 0x6c, 0x4d, 0xcb, 0x1a, 0x95, 0x30,
	0xb2, 0xe9, 0x03, 0xc1, 0xb6, 0x25, 0x27, 0xe6, 0x58, 0xcd, 0xd6, 0x5e, 0x2a, 0x6f, 0x59, 0x05,
	0xa9, 0x26, 0x7d, 0x25, 0xfa, 0x02, 0x7d, 0x04, 0x23, 0x2d, 0x0c, 0x16, 0x45, 0x6e, 0x70, 0xaa,
	0xd7, 0xda, 0x60, 0xf3, 0xd7, 0xe7, 0x28, 0xe0, 0x13, 0x82, 0x6d, 0xd0, 0xbc, 0x32, 0xa8, 0x56,
	0xa2, 0xa0, 0x47, 0xdf, 0xe6, 0xc1, 0x66, 0x0f, 0xe1, 0x30, 0xfc, 0x87, 0x9a, 0x96, 0x42, 0x2f,
	0xfc, 0x63, 0x3f, 0x08, 0xe8, 0x58, 0xe8, 0x85, 0x75, 0xbb, 0x41, 0x91, 0xa1, 0x9a, 0xce, 0x64,
	0x49, 0xf7, 0xe0, 0x5e, 0xfc, 0x81, 0x43, 0x3f, 0x73, 0x60, 0xd0, 0x22, 0xbe, 0x4d, 0x8b, 0x3f,
	0x23, 0x60, 0xcf, 0x8f, 0xbe, 0xff, 0xbc, 0x36, 0x37, 0xa3, 0xa5, 0xfd, 0x6f, 0xa3, 0x65, 0x98,
	0x8f, 0x27, 0xe1, 0x40, 0x4e, 0x3c, 0xc8, 0xc7, 0x13, 0x7f, 0x9a, 0x4f, 0xfb, 0xdf, 0xb8, 0x3f,
	0xaf, 0x57, 0x3d, 0xfa, 0x2b, 0xfb, 0xe1, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x1d, 0x4e,
	0x47, 0xd9, 0x0a, 0x00, 0x00,
}
