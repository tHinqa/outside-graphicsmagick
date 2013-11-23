// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package outside-graphicsmagick provides API definitions
//for accessing CORE_RL_magick_.dll.
//Based on GraphicsMagick v1.3.8.
package core

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
)

type Quantum byte //Q8

const (
	MaxTextExtent = 2053
	SignatureSize = 64
)

type (
	fix uintptr

	dirent fix
	DIR    fix
	FILE   fix

	Char              byte
	Enum              int
	IndexPacket       Quantum
	MagickBooleanType uint  // bool
	MagickSizeType    int64 // code says int64 doc says uint64
	MagickStatusType  uint  // bool
	Size              uint  // TODO(t):size_t on gcc vs msc 32 vs 64
	SSize             int   // TODO(t):ssize_t on gcc vs msc 32 vs 64
)

// Opaque types
type (
	Ascii85Info          struct{}
	BlobInfo             struct{}
	CacheInfo            struct{}
	GSMainInstance       struct{}
	ImageProfileIterator struct{}
	MagickMap            struct{}
	SemaphoreInfo        struct{}
	ThreadViewSet        struct{}
	ViewInfo             struct{}
	Void                 struct{}

	MagickTsdKey *struct{} // threading dependent

)

type (
	DecoderHandler                  func(*ImageInfo, *ExceptionInfo) *Image
	EncoderHandler                  func(*ImageInfo, *Image) uint
	MagickFreeFunc                  func(ptr *Void)
	MagickHandler                   func(*byte, Size) uint
	MagickMallocFunc                func(size Size) *Void
	MagickReallocFunc               func(ptr *Void, size Size) *Void
	MonitorHandler                  func(text string, quantum int64, span uint64, exception *ExceptionInfo) bool
	PixelIteratorMonoReadCallback   func(mutableData, immutableData *Void, constImage *Image, pixels *PixelPacket, indexes *IndexPacket, nPixels SSize, exception *ExceptionInfo) bool
	PixelIteratorMonoModifyCallback func(mutableData, immutableData *Void, constImage *Image, pixels *PixelPacket, indexes *IndexPacket, nPixels SSize, exception *ExceptionInfo) bool
	PixelIteratorDualReadCallback   func(mutableData, immutableData *Void, constImage1 *Image, pixels1 *PixelPacket, indexes1 *IndexPacket, nPixels1 SSize, constImage2 *Image, pixels2 *PixelPacket, indexes2 *IndexPacket, nPixels2 SSize, exception *ExceptionInfo) bool
	PixelIteratorDualModifyCallback func(mutableData, immutableData *Void, sourceImage *Image, sourcePixels *PixelPacket, sourceIndexes *IndexPacket, updImage *Image, updPixels *PixelPacket, updIndexes *IndexPacket, nPixels SSize, exception *ExceptionInfo) bool
	PixelIteratorDualNewCallback    PixelIteratorDualModifyCallback

	PixelIteratorTripleModifyCallback func(mutableData, immutableData *Void, source1Image *Image, source1Pixels *PixelPacket, source1Indexes *IndexPacket, source2Image *Image, source2Pixels *PixelPacket, source2Indexes *IndexPacket, updateImage *Image, updatePixels *PixelPacket, updateIndexes *IndexPacket, nPixels Size, exception *ExceptionInfo) bool
	PixelIteratorTripleNewCallback    PixelIteratorTripleModifyCallback
)

type Image struct {
	StorageClass    ClassType
	Colorspace      ColorspaceType
	Compression     CompressionType
	Dither          MagickBooleanType
	Matte           MagickBooleanType
	Columns         Size
	Rows            Size
	Colors          uint
	Depth_          uint
	Colormap        *PixelPacket
	BackgroundColor PixelPacket
	BorderColor     PixelPacket
	MatteColor      PixelPacket
	Gamma_          float64
	Chromaticity    ChromaticityInfo
	Orientation     OrientationType
	RenderingIntent RenderingIntent
	Units           ResolutionType
	Montage         *CString //TODO(t): CHECK GM vs IM -> NOTE(t): If *VString crash on i.Destroy()
	Directory       *CString
	Geometry_       *CString
	Offset          SSize
	XResolution     float64
	YResolution     float64
	Page            RectangleInfo
	TileInfo        RectangleInfo
	Blur_           float64
	Fuzz            float64
	Filter_         FilterTypes
	Interlace       InterlaceType
	Endian          EndianType
	Gravity         GravityType
	Compose         CompositeOperator
	Dispose_        DisposeType
	Scene           Size
	Delay           Size
	Iterations      Size
	TotalColors     Size
	StartLoop       SSize
	Error           ErrorInfo
	Timer           TimerInfo
	ClientData      *Void
	Filename        [MaxTextExtent]Char
	MagickFilename  [MaxTextExtent]Char
	Magick          [MaxTextExtent]Char
	MagickColumns   Size
	MagickRows      Size
	Exception_      ExceptionInfo2
	Previous        *Image
	Next            *Image
	Profiles        *Void
	IsMonochrome    uint
	IsGrayscale     uint
	Taint_          uint
	ClipMask_       *Image
	Ping            MagickBooleanType
	Cache           *CacheInfo
	DefaultViews    *ThreadViewSet
	Attributes      *ImageAttribute
	Ascii85         *Ascii85Info
	Blob            *BlobInfo
	ReferenceCount_ SSize
	Semaphore       *SemaphoreInfo
	Logging         uint
	List            *Image
	Signature_      Size
}
type ChannelType Enum

const (
	UndefinedChannel ChannelType = iota
	RedChannel
	CyanChannel
	GreenChannel
	MagentaChannel
	BlueChannel
	YellowChannel
	OpacityChannel
	BlackChannel
	MatteChannel
	AllChannels
	GrayChannel
)

type PrimaryInfo struct {
	X, Y, Z float64
}
type ChromaticityInfo struct {
	RedPrimary,
	GreenPrimary,
	BluePrimary,
	WhitePoint PrimaryInfo
}
type ClassType Enum

const (
	UndefinedClass ClassType = iota
	DirectClass
	PseudoClass
)

type ColorspaceType Enum

const (
	UndefinedColorspace ColorspaceType = iota
	RGBColorspace
	GRAYColorspace
	TransparentColorspace
	OHTAColorspace
	XYZColorspace
	YCCColorspace
	YIQColorspace
	YPbPrColorspace
	YUVColorspace
	CMYKColorspace
	SRGBColorspace
	HSLColorspace
	HWBColorspace
	LABColorspace
	CineonLogRGBColorspace
	Rec601LumaColorspace
	Rec601YCbCrColorspace
	Rec709LumaColorspace
	Rec709YCbCrColorspace
	YCbCrColorspace = Rec601YCbCrColorspace
)

type CompositeOperator Enum

const (
	UndefinedCompositeOp CompositeOperator = iota
	OverCompositeOp
	InCompositeOp
	OutCompositeOp
	AtopCompositeOp
	XorCompositeOp
	PlusCompositeOp
	MinusCompositeOp
	AddCompositeOp
	SubtractCompositeOp
	DifferenceCompositeOp
	MultiplyCompositeOp
	BumpmapCompositeOp
	CopyCompositeOp
	CopyRedCompositeOp
	CopyGreenCompositeOp
	CopyBlueCompositeOp
	CopyOpacityCompositeOp
	ClearCompositeOp
	DissolveCompositeOp
	DisplaceCompositeOp
	ModulateCompositeOp
	ThresholdCompositeOp
	NoCompositeOp
	DarkenCompositeOp
	LightenCompositeOp
	HueCompositeOp
	SaturateCompositeOp
	ColorizeCompositeOp
	LuminizeCompositeOp
	ScreenCompositeOp
	OverlayCompositeOp
	CopyCyanCompositeOp
	CopyMagentaCompositeOp
	CopyYellowCompositeOp
	CopyBlackCompositeOp
	DivideCompositeOp
)

type CompressionType Enum

const (
	UndefinedCompression CompressionType = iota
	NoCompression
	BZipCompression
	FaxCompression
	Group4Compression
	JPEGCompression
	LosslessJPEGCompression
	LZWCompression
	RLECompression
	ZipCompression
	LZMACompression
	JPEG2000Compression
	JBIG1Compression
	JBIG2Compression
	Group3Compression = FaxCompression
)

type ConfirmAccessMode Enum

const (
	UndefinedConfirmAccessMode ConfirmAccessMode = iota
	FileExecuteConfirmAccessMode
	FileReadConfirmAccessMode
	FileWriteConfirmAccessMode
	URLGetFTPConfirmAccessMode
	URLGetFileConfirmAccessMode
	URLGetHTTPConfirmAccessMode
)

type ConfirmAccessHandler func(mode ConfirmAccessMode, path string, exception *ExceptionInfo) bool
type DifferenceImageOptions struct {
	Channel        ChannelType
	HighlightStyle HighlightStyle
	HighlightColor PixelPacket
}
type HighlightStyle Enum

const (
	UndefinedHighlightStyle HighlightStyle = iota
	AssignHighlightStyle
	ThresholdHighlightStyle
	TintHighlightStyle
	XorHighlightStyle
)

type DifferenceStatistics struct {
	Red, Green, Blue, Opacity, Combined float64
}
type DisposeType Enum

const (
	UndefinedDispose DisposeType = iota
	NoneDispose
	BackgroundDispose
	PreviousDispose
)

type DrawContext struct {
	Image               *Image
	Mvg                 *VString
	MvgAlloc, MvgLength Size
	MvgWidth            uint
	PatternId           *VString
	PatternBounds       RectangleInfo
	PatternOffset       Size
	Index               uint
	GraphicContext      **DrawInfo
	FilterOff           int
	IndentDepth         uint
	PathOperation       PathOperation
	PathMode            PathMode
	Signature           Size
}

type DrawInfo struct {
	Primitive        *VString
	Geometry         *VString
	Affine           AffineMatrix
	Gravity          GravityType
	Fill             PixelPacket
	Stroke           PixelPacket
	StrokeWidth      float64
	Gradient         GradientInfo
	FillPattern      *Image
	Tile             *Image
	StrokePattern    *Image
	StrokeAntialias  MagickBooleanType // uint
	TextAntialias    MagickBooleanType // uint
	FillRule         FillRule
	Linecap          LineCap
	Linejoin         LineJoin
	Miterlimit       Size
	DashOffset       float64
	Decorate         DecorationType
	Compose          CompositeOperator
	Text             *VString
	Font             *VString
	Family           *VString
	Style            StyleType
	Stretch          StretchType
	Weight           Size
	Encoding         *VString
	Pointsize        float64
	Density          *VString
	Align            AlignType
	Undercolor       PixelPacket
	BorderColor      PixelPacket
	ServerName       *VString
	DashPattern      *float64
	ClipPath         *VString
	Bounds           SegmentInfo
	ClipUnits        ClipPathUnits
	Opacity          Quantum
	Render           MagickBooleanType // uint
	Debug            MagickBooleanType // uint
	ElementReference ElementReference
	Signature        Size
}
type AffineMatrix struct {
	Sx, Rx, Ry, Sy, Tx, Ty float64
}
type AlignType Enum

const (
	UndefinedAlign AlignType = iota
	LeftAlign
	CenterAlign
	RightAlign
)

type ClipPathUnits Enum

const (
	UserSpace ClipPathUnits = iota
	UserSpaceOnUse
	ObjectBoundingBox
)

type DecorationType Enum

const (
	NoDecoration DecorationType = iota
	UnderlineDecoration
	OverlineDecoration
	LineThroughDecoration
)

type ElementReference struct {
	Id        *VString
	Type      ReferenceType
	Gradient  GradientInfo
	Signature Size
	Previous  *ElementReference
	Next      *ElementReference
}

type EndianType Enum

const (
	UndefinedEndian EndianType = iota
	LSBEndian
	MSBEndian
	NativeEndian
)

type ErrorInfo struct {
	MeanErrorPerPixel,
	NormalizedMeanError,
	NormalizedMaximumError float64
}
type ExceptionInfo struct {
	Severity    ExceptionType
	Reason      *OVString
	Description *OVString
	ErrorNumber int
	Module      *OVString
	Function    *OVString
	Line        Size
	Signature   Size
}
type ExceptionInfo2 struct {
	Severity    ExceptionType
	Reason      *CString
	Description *CString
	ErrorNumber int
	Module      *CString
	Function    *CString
	Line        Size
	Signature   Size
}
type ExceptionType Enum

const (
	undefinedExceptionBase ExceptionType = iota
	exceptionBase
	resourceBase
	resourceLimitBase = resourceBase
)
const (
	UndefinedException ExceptionType = iota * 5
	typeBase
	optionBase
	delegateBase
	missingDelegateBase
	corruptImageBase
	fileOpenBase
	blobBase
	streamBase
	cacheBase
	coderBase
	moduleBase
	drawBase
	imageBase
	temporaryFileBase
	transformBase
	xServerBase
	monitorBase
	registryBase
	configureBase
	annotateBase  = typeBase
	renderBase    = drawBase
	wandBase      = imageBase + 2
	x11Base       = xServerBase + 1
	userBase      = xServerBase + 2
	localeBase    = monitorBase + 1
	deprecateBase = monitorBase + 2
)
const (
	EventException ExceptionType = 100

	ExceptionEvent       = EventException + exceptionBase
	ResourceEvent        = EventException + resourceBase
	ResourceLimitEvent   = EventException + resourceLimitBase
	TypeEvent            = EventException + typeBase
	AnnotateEvent        = EventException + annotateBase
	OptionEvent          = EventException + optionBase
	DelegateEvent        = EventException + delegateBase
	MissingDelegateEvent = EventException + missingDelegateBase
	CorruptImageEvent    = EventException + corruptImageBase
	FileOpenEvent        = EventException + fileOpenBase
	BlobEvent            = EventException + blobBase
	StreamEvent          = EventException + streamBase
	CacheEvent           = EventException + cacheBase
	CoderEvent           = EventException + coderBase
	ModuleEvent          = EventException + moduleBase
	DrawEvent            = EventException + drawBase
	RenderEvent          = EventException + renderBase
	ImageEvent           = EventException + imageBase
	WandEvent            = EventException + wandBase
	TemporaryFileEvent   = EventException + temporaryFileBase
	TransformEvent       = EventException + transformBase
	XServerEvent         = EventException + xServerBase
	X11Event             = EventException + x11Base
	UserEvent            = EventException + userBase
	MonitorEvent         = EventException + monitorBase
	LocaleEvent          = EventException + localeBase
	DeprecateEvent       = EventException + deprecateBase
	RegistryEvent        = EventException + registryBase
	ConfigureEvent       = EventException + configureBase

	WarningException ExceptionType = 300

	ExceptionWarning       = WarningException + exceptionBase
	ResourceWarning        = WarningException + resourceBase
	ResourceLimitWarning   = WarningException + resourceLimitBase
	TypeWarning            = WarningException + typeBase
	AnnotateWarning        = WarningException + annotateBase
	OptionWarning          = WarningException + optionBase
	DelegateWarning        = WarningException + delegateBase
	MissingDelegateWarning = WarningException + missingDelegateBase
	CorruptImageWarning    = WarningException + corruptImageBase
	FileOpenWarning        = WarningException + fileOpenBase
	BlobWarning            = WarningException + blobBase
	StreamWarning          = WarningException + streamBase
	CacheWarning           = WarningException + cacheBase
	CoderWarning           = WarningException + coderBase
	ModuleWarning          = WarningException + moduleBase
	DrawWarning            = WarningException + drawBase
	RenderWarning          = WarningException + renderBase
	ImageWarning           = WarningException + imageBase
	WandWarning            = WarningException + wandBase
	TemporaryFileWarning   = WarningException + temporaryFileBase
	TransformWarning       = WarningException + transformBase
	XServerWarning         = WarningException + xServerBase
	X11Warning             = WarningException + x11Base
	UserWarning            = WarningException + userBase
	MonitorWarning         = WarningException + monitorBase
	LocaleWarning          = WarningException + localeBase
	DeprecateWarning       = WarningException + deprecateBase
	RegistryWarning        = WarningException + registryBase
	ConfigureWarning       = WarningException + configureBase

	ErrorException ExceptionType = 400

	ExceptionError       = ErrorException + exceptionBase
	ResourceError        = ErrorException + resourceBase
	ResourceLimitError   = ErrorException + resourceLimitBase
	TypeError            = ErrorException + typeBase
	AnnotateError        = ErrorException + annotateBase
	OptionError          = ErrorException + optionBase
	DelegateError        = ErrorException + delegateBase
	MissingDelegateError = ErrorException + missingDelegateBase
	CorruptImageError    = ErrorException + corruptImageBase
	FileOpenError        = ErrorException + fileOpenBase
	BlobError            = ErrorException + blobBase
	StreamError          = ErrorException + streamBase
	CacheError           = ErrorException + cacheBase
	CoderError           = ErrorException + coderBase
	ModuleError          = ErrorException + moduleBase
	DrawError            = ErrorException + drawBase
	RenderError          = ErrorException + renderBase
	ImageError           = ErrorException + imageBase
	WandError            = ErrorException + wandBase
	TemporaryFileError   = ErrorException + temporaryFileBase
	TransformError       = ErrorException + transformBase
	XServerError         = ErrorException + xServerBase
	X11Error             = ErrorException + x11Base
	UserError            = ErrorException + userBase
	MonitorError         = ErrorException + monitorBase
	LocaleError          = ErrorException + localeBase
	DeprecateError       = ErrorException + deprecateBase
	RegistryError        = ErrorException + registryBase
	ConfigureError       = ErrorException + configureBase

	FatalErrorException ExceptionType = 700

	ExceptionFatalError       = FatalErrorException + exceptionBase
	ResourceFatalError        = FatalErrorException + resourceBase
	ResourceLimitFatalError   = FatalErrorException + resourceLimitBase
	TypeFatalError            = FatalErrorException + typeBase
	AnnotateFatalError        = FatalErrorException + annotateBase
	OptionFatalError          = FatalErrorException + optionBase
	DelegateFatalError        = FatalErrorException + delegateBase
	MissingDelegateFatalError = FatalErrorException + missingDelegateBase
	CorruptImageFatalError    = FatalErrorException + corruptImageBase
	FileOpenFatalError        = FatalErrorException + fileOpenBase
	BlobFatalError            = FatalErrorException + blobBase
	StreamFatalError          = FatalErrorException + streamBase
	CacheFatalError           = FatalErrorException + cacheBase
	CoderFatalError           = FatalErrorException + coderBase
	ModuleFatalError          = FatalErrorException + moduleBase
	DrawFatalError            = FatalErrorException + drawBase
	RenderFatalError          = FatalErrorException + renderBase
	ImageFatalError           = FatalErrorException + imageBase
	WandFatalError            = FatalErrorException + wandBase
	TemporaryFileFatalError   = FatalErrorException + temporaryFileBase
	TransformFatalError       = FatalErrorException + transformBase
	XServerFatalError         = FatalErrorException + xServerBase
	X11FatalError             = FatalErrorException + x11Base
	UserFatalError            = FatalErrorException + userBase
	MonitorFatalError         = FatalErrorException + monitorBase
	LocaleFatalError          = FatalErrorException + localeBase
	DeprecateFatalError       = FatalErrorException + deprecateBase
	RegistryFatalError        = FatalErrorException + registryBase
	ConfigureFatalError       = FatalErrorException + configureBase
)

type FillRule Enum

const (
	UndefinedRule FillRule = iota
	EvenOddRule
	NonZeroRule
)

type FilterTypes Enum

const (
	UndefinedFilter FilterTypes = iota
	PointFilter
	BoxFilter
	TriangleFilter
	HermiteFilter
	HanningFilter
	HammingFilter
	BlackmanFilter
	GaussianFilter
	QuadraticFilter
	CubicFilter
	CatromFilter
	MitchellFilter
	LanczosFilter
	BesselFilter
	SincFilter
)

type FrameInfo struct {
	Width, Height                Size
	X, Y, InnerBevel, OuterBevel SSize
}
type GradientInfo struct {
	Type           GradientType
	Color          PixelPacket
	Stop           SegmentInfo
	Length         Size
	Spread         SpreadMethod
	Signature      Size
	Previous, Next *GradientInfo
}
type GradientType Enum

const (
	UndefinedGradient GradientType = iota
	LinearGradient
	RadialGradient
)

type GravityType Enum

const (
	UndefinedGravity GravityType = iota
	NorthWestGravity
	NorthGravity
	NorthEastGravity
	WestGravity
	CenterGravity
	EastGravity
	SouthWestGravity
	SouthGravity
	SouthEastGravity
	StaticGravity
)

type HistogramColorPacket struct {
	Pixel PixelPacket
	Count Size
}
type ImageAttribute struct {
	Key      *VString
	Value    *VString
	Length   Size
	Previous *ImageAttribute
	next     *ImageAttribute
}
type ImageInfo struct {
	Compression     CompressionType
	Orientation     OrientationType
	Temporary       MagickBooleanType
	Adjoin          MagickBooleanType
	Antialias       MagickBooleanType
	Subimage        Size
	Subrange        Size
	Depth           Size
	Size            *VString
	Tile            *VString
	Page            *VString
	Interlace       InterlaceType
	Endian          EndianType
	Units           ResolutionType
	Quality         Size
	SamplingFactor  *VString
	ServerName      *VString
	Font            *VString
	Texture         *VString
	Density         *VString
	Pointsize       float64
	Fuzz            float64
	Pen             PixelPacket
	BackgroundColor PixelPacket
	BorderColor     PixelPacket
	MatteColor      PixelPacket
	Dither          MagickBooleanType
	Monochrome      MagickBooleanType
	Progress        MagickBooleanType
	Colorspace      ColorspaceType
	Type            ImageType
	Group           SSize
	Verbose         MagickBooleanType
	View            *VString
	Authenticate    *VString
	ClientData      *Void
	File            *FILE
	Magick          [MaxTextExtent]Char
	Filename        [MaxTextExtent]Char
	Cache           *Void
	Definitions     *Void
	Attributes      *Image
	Ping            MagickBooleanType
	PreviewType     PreviewType
	Affirm          MagickBooleanType
	Blob            *Void
	Length          Size
	Unique          [MaxTextExtent]Char
	Zero            [MaxTextExtent]Char
	Signature       Size
}
type ImageType Enum

const (
	UndefinedType ImageType = iota
	BilevelType
	GrayscaleType
	GrayscaleMatteType
	PaletteType
	PaletteMatteType
	TrueColorType
	TrueColorMatteType
	ColorSeparationType
	ColorSeparationMatteType
	OptimizeType
)

type InterlaceType Enum

const (
	UndefinedInterlace InterlaceType = iota
	NoInterlace
	LineInterlace
	PlaneInterlace
	PartitionInterlace
)

type LineCap Enum

const (
	UndefinedCap LineCap = iota
	ButtCap
	RoundCap
	SquareCap
)

type LineJoin Enum

const (
	UndefinedJoin LineJoin = iota
	MiterJoin
	RoundJoin
	BevelJoin
)

type MetricType Enum

const (
	UndefinedMetric MetricType = iota
	MeanAbsoluteErrorMetric
	MeanSquaredErrorMetric
	PeakAbsoluteErrorMetric
	PeakSignalToNoiseRatioMetric
	RootMeanSquaredErrorMetric
)

type OrientationType Enum

const (
	UndefinedOrientation OrientationType = iota
	TopLeftOrientation
	TopRightOrientation
	BottomRightOrientation
	BottomLeftOrientation
	LeftTopOrientation
	RightTopOrientation
	RightBottomOrientation
	LeftBottomOrientation
)

type PixelPacket struct {
	Blue, Green, Red, Opacity Quantum
}
type PixelPacketBE struct {
	Red, Green, Blue, Opacity Quantum
}
type PointInfo struct {
	X, Y float64
}
type PreviewType Enum

const (
	UndefinedPreview PreviewType = iota
	RotatePreview
	ShearPreview
	RollPreview
	HuePreview
	SaturationPreview
	BrightnessPreview
	GammaPreview
	SpiffPreview
	DullPreview
	GrayscalePreview
	QuantizePreview
	DespecklePreview
	ReduceNoisePreview
	AddNoisePreview
	SharpenPreview
	BlurPreview
	ThresholdPreview
	EdgeDetectPreview
	SpreadPreview
	SolarizePreview
	ShadePreview
	RaisePreview
	SegmentPreview
	SwirlPreview
	ImplodePreview
	WavePreview
	OilPaintPreview
	CharcoalDrawingPreview
	JPEGPreview
)

type RectangleInfo struct {
	Width, Height Size
	X, Y          SSize
}

type ReferenceType Enum

const (
	UndefinedReference ReferenceType = iota
	GradientReference
)

type RenderingIntent Enum

const (
	UndefinedIntent RenderingIntent = iota
	SaturationIntent
	PerceptualIntent
	AbsoluteIntent
	RelativeIntent
)

type ResolutionType Enum

const (
	UndefinedResolution ResolutionType = iota
	PixelsPerInchResolution
	PixelsPerCentimeterResolution
)

type SegmentInfo struct {
	X1, Y1, X2, Y2 float64
}

type SpreadMethod Enum

const (
	UndefinedSpread SpreadMethod = iota
	PadSpread
	ReflectSpread
	RepeatSpread
)

type StorageType Enum

const (
	CharPixel StorageType = iota
	ShortPixel
	IntegerPixel
	LongPixel
	FloatPixel
	DoublePixel
)

type StretchType Enum

const (
	NormalStretch StretchType = iota
	UltraCondensedStretch
	ExtraCondensedStretch
	CondensedStretch
	SemiCondensedStretch
	SemiExpandedStretch
	ExpandedStretch
	ExtraExpandedStretch
	UltraExpandedStretch
	AnyStretch
)

type StyleType Enum

const (
	NormalStyle StyleType = iota
	ItalicStyle
	ObliqueStyle
	AnyStyle
)

type Timer struct {
	Start,
	Stop,
	Total float64
}
type TimerInfo struct {
	User      Timer
	Elapsed   Timer
	State     TimerState
	Signature Size
}
type TimerState Enum

const (
	UndefinedTimerState TimerState = iota
	StoppedTimerState
	RunningTimerState
)

type TypeMetric struct {
	PixelsPerEm        PointInfo
	Ascent             float64
	Descent            float64
	Width              float64
	Height             float64
	MaxAdvance         float64
	Bounds             SegmentInfo
	UnderlinePosition  float64
	UnderlineThickness float64
}

type QuantumType Enum

const (
	UndefinedQuantum QuantumType = iota
	IndexQuantum
	GrayQuantum
	IndexAlphaQuantum
	GrayAlphaQuantum
	RedQuantum
	CyanQuantum
	GreenQuantum
	YellowQuantum
	BlueQuantum
	MagentaQuantum
	AlphaQuantum
	BlackQuantum
	RGBQuantum
	RGBAQuantum
	CMYKQuantum
	CMYKAQuantum
	CIEYQuantum
	CIEXYZQuantum
)

type ExportPixelAreaOptions struct {
	SampleType                     QuantumSampleType
	DoubleMinvalue, DoubleMaxvalue float64
	GrayscaleMiniswhite            MagickBooleanType
	PadBytes                       Size
	PadValue                       byte
	Endian                         EndianType
	Signature                      Size
}

type QuantumSampleType Enum

const (
	UndefinedQuantumSampleType QuantumSampleType = iota
	UnsignedQuantumSampleType
	FloatQuantumSampleType
)

type ExportPixelAreaInfo struct {
	BytesExported Size
}

type ImportPixelAreaInfo struct {
	BytesImported Size
}

type ImportPixelAreaOptions struct {
	SampleType                     QuantumSampleType
	DoubleMinvalue, DoubleMaxvalue float64
	GrayscaleMiniswhite            MagickBooleanType
	Endian                         EndianType
	Signature                      Size
}

type MagickInfo struct {
	Previous           *MagickInfo
	Next               *MagickInfo
	Name               *VString
	Description_       *VString
	Note               *VString
	Version            *VString
	Module             *VString
	Decoder_           *DecoderHandler
	Encoder_           *EncoderHandler
	Magick             *MagickHandler
	ClientData         *Void
	Adjoin_            MagickBooleanType
	Raw                MagickBooleanType
	Stealth            MagickBooleanType
	SeekableStream_    MagickBooleanType
	BlobSupport_       MagickBooleanType
	ThreadSupport_     MagickStatusType
	CoderClass         CoderClass
	ExtensionTreatment ExtensionTreatment
	Signature          Size
}

type CoderClass Enum

const (
	UnstableCoderClass CoderClass = iota
	StableCoderClass
	PrimaryCoderClass
)

type ExtensionTreatment Enum

const (
	HintExtensionTreatment ExtensionTreatment = iota
	ObeyExtensionTreatment
	IgnoreExtensionTreatment
)

type MontageInfo struct {
	Geometry        *VString
	Tile            *VString
	Title           *VString
	Frame           *VString
	Texture         *VString
	Font            *VString
	Pointsize       float64
	BorderWidth     Size
	Shadow          MagickBooleanType
	Fill            PixelPacket
	Stroke          PixelPacket
	BackgroundColor PixelPacket
	BorderColor     PixelPacket
	MatteColor      PixelPacket
	Gravity         GravityType
	Filename        [MaxTextExtent]Char
	Signature       Size
}

type QuantumOperator Enum

const (
	UndefinedQuantumOp QuantumOperator = iota
	AddQuantumOp
	AndQuantumOp
	AssignQuantumOp
	DivideQuantumOp
	LShiftQuantumOp
	MultiplyQuantumOp
	OrQuantumOp
	RShiftQuantumOp
	SubtractQuantumOp
	ThresholdQuantumOp
	ThresholdBlackQuantumOp
	ThresholdWhiteQuantumOp
	XorQuantumOp
	NoiseGaussianQuantumOp
	NoiseImpulseQuantumOp
	NoiseLaplacianQuantumOp
	NoiseMultiplicativeQuantumOp
	NoisePoissonQuantumOp
	NoiseUniformQuantumOp
	NegateQuantumOp
	GammaQuantumOp
	DepthQuantumOp
	LogQuantumOp
	MaxQuantumOp
	MinQuantumOp
	PowQuantumOp
	NoiseRandomQuantumOp
)

type VirtualPixelMethod Enum

const (
	UndefinedVirtualPixelMethod VirtualPixelMethod = iota
	ConstantVirtualPixelMethod
	EdgeVirtualPixelMethod
	MirrorVirtualPixelMethod
	TileVirtualPixelMethod
)

type PixelIteratorOptions struct {
	MaxThreads int
	Signature  Size
}

type QuantizeInfo struct {
	NumberColors Size
	TreeDepth    Size
	Dither       MagickBooleanType
	Colorspace   ColorspaceType
	MeasureError MagickBooleanType
	Signature    Size
}

type RegistryType Enum

const (
	UndefinedRegistryType RegistryType = iota
	ImageRegistryType
	ImageInfoRegistryType
)

type ResourceType Enum

const (
	UndefinedResource ResourceType = iota
	DiskResource
	FileResource
	MapResource
	MemoryResource
	PixelsResource
	ThreadResource
)

type SignatureInfo struct {
	Digest              [8]Size
	LowOrder, HighOrder Size
	Offset              SSize
	Message             [SignatureSize]byte
}

type ImageStatistics struct {
	red, green, blue, opacity ImageChannelStatistics
}

type ImageChannelStatistics struct {
	Maximum,
	Minimum,
	Mean,
	StandardDeviation,
	Variance float64
}

type PathOperation Enum

const (
	PathDefaultOperation PathOperation = iota
	PathCloseOperation
	PathCurveToOperation
	PathCurveToQuadraticBezierOperation
	PathCurveToQuadraticBezierSmoothOperation
	PathCurveToSmoothOperation
	PathEllipticArcOperation
	PathLineToHorizontalOperation
	PathLineToOperation
	PathLineToVerticalOperation
	PathMoveToOperation
)

type PathMode Enum

const (
	DefaultPathMode PathMode = iota
	AbsolutePathMode
	RelativePathMode
)

type PaintMethod Enum

const (
	PointMethod PaintMethod = iota
	ReplaceMethod
	FloodfillMethod
	FillToBorderMethod
	ResetMethod
)

type NoiseType Enum

const (
	UniformNoise NoiseType = iota
	GaussianNoise
	MultiplicativeGaussianNoise
	ImpulseNoise
	LaplacianNoise
	PoissonNoise
	RandomNoise
	UndefinedNoise
)

// Annotate

var AnnotateImage func(i *Image, drawInfo *DrawInfo) bool

func (i *Image) Annotate(drawInfo *DrawInfo) bool { return AnnotateImage(i, drawInfo) }

var GetTypeMetrics func(i *Image, drawInfo *DrawInfo, metrics *TypeMetric) bool

func (i *Image) TypeMetrics(drawInfo *DrawInfo, metrics *TypeMetric) bool {
	return GetTypeMetrics(i, drawInfo, metrics)
}

// Attribute

var CloneImageAttributes func(i *Image, cloneImage *Image) bool

func (i *Image) CloneAttributes(cloneImage *Image) bool {
	return CloneImageAttributes(i, cloneImage)
}

var DestroyImageAttributes func(i *Image)

func (i *Image) DestroyAttributes() { DestroyImageAttributes(i) }

var GetImageAttribute func(i *Image, key string) *ImageAttribute

func (i *Image) Attribute(key string) *ImageAttribute { return GetImageAttribute(i, key) }

var GetImageClippingPathAttribute func(i *Image) *ImageAttribute

func (i *Image) ClippingPathAttribute() *ImageAttribute {
	return GetImageClippingPathAttribute(i)
}

var SetImageAttribute func(i *Image, key, value string) bool

func (i *Image) SetAttribute(key, value string) bool { return SetImageAttribute(i, key, value) }

// Average

var AverageImages func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) AverageImages(exception *ExceptionInfo) *Image { return AverageImages(i, exception) }

// ASC CDL

var CdlImage func(i *Image, cdl string) bool

func (i *Image) Cdl(cdl string) bool { return CdlImage(i, cdl) }

// Blob

var AttachBlob func(b *BlobInfo, blob *Void, length Size)

func (b *BlobInfo) Attach(blob *Void, length Size) { AttachBlob(b, blob, length) }

var BlobIsSeekable func(i *Image) bool

func (i *Image) BlobIsSeekable() bool { return BlobIsSeekable(i) }

var BlobReserveSize func(i *Image, size int64) bool

func (i *Image) BlobReserveSize(size int64) bool { return BlobReserveSize(i, size) }

var BlobToFile func(filename string, blob *Void, length Size, exception *ExceptionInfo) bool

var BlobToImage func(i *ImageInfo, blob *Void, length Size, exception *ExceptionInfo) *Image

func (i *ImageInfo) BlobToImage(blob *Void, length Size, exception *ExceptionInfo) *Image {
	return BlobToImage(i, blob, length, exception)
}

var CloneBlobInfo func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Clone() *BlobInfo { return CloneBlobInfo(b) }

var DestroyBlob func(i *Image)

func (i *Image) DestroyBlob() { DestroyBlob(i) }

var DestroyBlobInfo func(blob *BlobInfo)

func (b *BlobInfo) Destroy() { DestroyBlobInfo(b) }

var DetachBlob func(b *BlobInfo) *byte

func (b *BlobInfo) Detach() *byte { return DetachBlob(b) }

var FileToBlob func(filename string, length Size, exception *ExceptionInfo) *byte

var GetBlobFileHandle func(i *Image) *FILE

func (i *Image) BlobFileHandle() *FILE { return GetBlobFileHandle(i) }

var GetBlobInfo func(b *BlobInfo)

func (b *BlobInfo) Get() { GetBlobInfo(b) }

var GetBlobStatus func(i *Image) int

func (i *Image) BlobStatus() int { return GetBlobStatus(i) }

var GetBlobStreamData func(i *Image) *byte

func (i *Image) BlobStreamData() *byte { return GetBlobStreamData(i) }

var GetBlobTemporary func(i *Image) bool

func (i *Image) BlobTemporary() bool { return GetBlobTemporary(i) }

var GetConfigureBlob func(filename, path string, length *Size, exception *ExceptionInfo) *Void

var ImageToBlob func(i *ImageInfo, image *Image, length *Size, exception *ExceptionInfo) *byte

func (i *ImageInfo) ImageToBlob(image *Image, length *Size, exception *ExceptionInfo) *byte {
	return ImageToBlob(i, image, length, exception)
}

var ImageToFile func(i *Image, filename string, exception *ExceptionInfo) bool

func (i *Image) ToFile(filename string, exception *ExceptionInfo) bool {
	return ImageToFile(i, filename, exception)
}

var PingBlob func(i *ImageInfo, blob *Void, length uint32, exception *ExceptionInfo) *Image

func (i *ImageInfo) PingBlob(blob *Void, length uint32, exception *ExceptionInfo) *Image {
	return PingBlob(i, blob, length, exception)
}

var ReferenceBlob func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Reference() *BlobInfo { return ReferenceBlob(b) }

var SetBlobClosable func(i *Image, closable bool)

func (i *Image) SetBlobClosable(closable bool) { SetBlobClosable(i, closable) }

var SetBlobTemporary func(i *Image, isTemporary bool)

func (i *Image) SetBlobTemporary(isTemporary bool) { SetBlobTemporary(i, isTemporary) }

// Channel

var ChannelImage func(i *Image, channel ChannelType) bool

func (i *Image) Channel(channel ChannelType) bool { return ChannelImage(i, channel) }

var ExportImageChannel func(i *Image, channel ChannelType, exception *ExceptionInfo) *Image

func (i *Image) ExportChannel(channel ChannelType, exception *ExceptionInfo) *Image {
	return ExportImageChannel(i, channel, exception)
}

var GetImageChannelDepth func(i *Image, channel ChannelType, exception *ExceptionInfo) Size

func (i *Image) ChannelDepth(channel ChannelType, exception *ExceptionInfo) Size {
	return GetImageChannelDepth(i, channel, exception)
}

var ImportImageChannel func(i, dst *Image, channel ChannelType) bool

func (i *Image) ImportChannel(dst *Image, channel ChannelType) bool {
	return ImportImageChannel(i, dst, channel)
}

var ImportImageChannelsMasked func(i, updateImage *Image, channels ChannelType) bool

func (i *Image) ImportChannelsMasked(updateImage *Image, channels ChannelType) bool {
	return ImportImageChannelsMasked(i, updateImage, channels)
}

var SetImageChannelDepth func(i *Image, channel ChannelType, depth Size) bool

func (i *Image) SetChannelDepth(channel ChannelType, depth Size) bool {
	return SetImageChannelDepth(i, channel, depth)
}

// Color

var GetColorHistogram func(i *Image, colors *Size, exception *ExceptionInfo) *HistogramColorPacket

func (i *Image) ColorHistogram(colors *Size, exception *ExceptionInfo) *HistogramColorPacket {
	return GetColorHistogram(i, colors, exception)
}

var GetNumberColors func(i *Image, file *FILE, exception *ExceptionInfo) Size

func (i *Image) NumberColors(file *FILE, exception *ExceptionInfo) Size {
	return GetNumberColors(i, file, exception)
}

var IsPaletteImage func(i *Image, exception *ExceptionInfo) bool

func (i *Image) Palette(exception *ExceptionInfo) bool { return IsPaletteImage(i, exception) }

// Colormap

var AllocateImageColormap func(i *Image, colors Size) bool

func (i *Image) AllocateColormap(colors Size) bool {
	return AllocateImageColormap(i, colors)
}

var CycleColormapImage func(i *Image, amount int) bool

func (i *Image) CycleColormap(amount int) bool { return CycleColormapImage(i, amount) }

var ReplaceImageColormap func(i *Image, colormap *PixelPacket, colors uint) bool

func (i *Image) ReplaceImageColormap(colormap *PixelPacket, colors uint) bool {
	return ReplaceImageColormap(i, colormap, colors)
}

// Compare

var DifferenceImage func(i, compareImage *Image, differenceOptions *DifferenceImageOptions, exception *ExceptionInfo) *Image

func (i *Image) Difference(compareImage *Image, differenceOptions *DifferenceImageOptions, exception *ExceptionInfo) *Image {
	return DifferenceImage(i, compareImage, differenceOptions, exception)
}

var GetImageChannelDifference func(i, compareImage *Image, metric MetricType, statistics *DifferenceStatistics, exception *ExceptionInfo) bool

func (i *Image) ChannelDifference(compareImage *Image, metric MetricType, statistics *DifferenceStatistics, exception *ExceptionInfo) bool {
	return GetImageChannelDifference(i, compareImage, metric, statistics, exception)
}

var GetImageChannelDistortion func(i, reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) bool

func (i *Image) ChannelDistortion(reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageChannelDistortion(i, reconstructImage, channel, metric, distortion, exception)
}

var GetImageDistortion func(i, reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) bool

func (i *Image) Distortion(reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageDistortion(i, reconstructImage, metric, distortion, exception)
}

var IsImagesEqual func(i *Image, reconstructImage *Image) bool

func (i *Image) Equal(reconstructImage *Image) bool {
	return IsImagesEqual(i, reconstructImage)
}

var InitializeDifferenceImageOptions func(options *DifferenceImageOptions, exception *ExceptionInfo)

var InitializeDifferenceStatistics func(differenceStatistics *DifferenceStatistics, exception *ExceptionInfo)

// Composite

var CompositeImage func(i *Image, compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool

func (i *Image) CompositeImage(compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool {
	return CompositeImage(i, compose, compositeImage, xOffset, yOffset)
}

// Confirm Access

var MagickConfirmAccess func(mode ConfirmAccessMode, path string, exception *ExceptionInfo) bool

var MagickSetConfirmAccessHandler func(handler ConfirmAccessHandler) ConfirmAccessHandler

// Constitute

var ConstituteImage func(columns, rows Size, map_ string, storage StorageType, pixels *Void, exception *ExceptionInfo) *Image

var DispatchImage func(i *Image, xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) bool

func (i *Image) Dispatch(xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) bool {
	return DispatchImage(i, xOffset, yOffset, columns, rows, map_, type_, pixels, exception)
}

var PingImage func(i *ImageInfo, exception *ExceptionInfo) *Image

func (i *ImageInfo) PingImage(exception *ExceptionInfo) *Image { return PingImage(i, exception) }

var ReadImage func(i *ImageInfo, exception *ExceptionInfo) *Image

func (i *ImageInfo) ReadImage(exception *ExceptionInfo) *Image { return ReadImage(i, exception) }

var ReadInlineImage func(i *ImageInfo, content string, exception *ExceptionInfo) *Image

func (i *ImageInfo) ReadInlineImage(content string, exception *ExceptionInfo) *Image {
	return ReadInlineImage(i, content, exception)
}

var WriteImage func(i *ImageInfo, image *Image) bool

func (i *ImageInfo) WriteImage(image *Image) bool { return WriteImage(i, image) }

var WriteImages func(i *ImageInfo, images *Image, filename string, exception *ExceptionInfo) bool

func (i *ImageInfo) WriteImages(images *Image, filename string, exception *ExceptionInfo) bool {
	return WriteImages(i, images, filename, exception)
}

var WriteImagesFile func(i *ImageInfo, image *Image, file *FILE, exception *ExceptionInfo) bool

func (i *ImageInfo) WriteImagesFile(image *Image, file *FILE, exception *ExceptionInfo) bool {
	return WriteImagesFile(i, image, file, exception)
}

// Decorate

var BorderImage func(i *Image, borderInfo *RectangleInfo, exception *ExceptionInfo) *Image

func (i *Image) Border(borderInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return BorderImage(i, borderInfo, exception)
}

var FrameImage func(i *Image, f *FrameInfo, e *ExceptionInfo) *Image

func (i *Image) Frame(frameInfo *FrameInfo, exception *ExceptionInfo) *Image {
	return FrameImage(i, frameInfo, exception)
}

var RaiseImage func(i *Image, raiseInfo *RectangleInfo, raise bool) bool

func (i *Image) RaiseImage(raiseInfo *RectangleInfo, raise bool) bool {
	return RaiseImage(i, raiseInfo, raise)
}

// Describe

var DescribeImage func(i *Image, file *FILE, verbose bool) bool

func (i *Image) Describe(file *FILE, verbose bool) bool {
	return DescribeImage(i, file, verbose)
}

// Draw

var DrawAnnotation func(d *DrawContext, x, y float64, text *byte)

func (d *DrawContext) Annotation(x, y float64, text *byte) { DrawAnnotation(d, x, y, text) }

var DrawAffine func(d *DrawContext, a *AffineMatrix)

func (d *DrawContext) Affine(a *AffineMatrix) { DrawAffine(d, a) }

var DrawAllocateContext func(drawInfo *DrawInfo, image *Image) *DrawContext

var DrawArc func(d *DrawContext, sx, sy, ex, ey, sd, ed float64)

func (d *DrawContext) Arc(sx, sy, ex, ey, sd, ed float64) { DrawArc(d, sx, sy, ex, ey, sd, ed) }

var DrawBezier func(d *DrawContext, numCoords Size, coordinates *PointInfo)

func (d *DrawContext) Bezier(numCoords Size, coordinates *PointInfo) {
	DrawBezier(d, numCoords, coordinates)
}

var DrawCircle func(d *DrawContext, ox, oy, px, py float64)

func (d *DrawContext) Circle(ox, oy, px, py float64) { DrawCircle(d, ox, oy, px, py) }

var DrawGetClipPath func(d *DrawContext) string

func (d *DrawContext) ClipPath() string { return DrawGetClipPath(d) }

var DrawGetClipRule func(d *DrawContext) FillRule

func (d *DrawContext) ClipRule() FillRule { return DrawGetClipRule(d) }

var DrawGetClipUnits func(d *DrawContext) ClipPathUnits

func (d *DrawContext) ClipUnits() ClipPathUnits { return DrawGetClipUnits(d) }

var DrawSetClipPath func(d *DrawContext, clipPath string)

func (d *DrawContext) SetClipPath(clipPath string) { DrawSetClipPath(d, clipPath) }

var DrawSetClipRule func(d *DrawContext, fillRule FillRule)

func (d *DrawContext) SetClipRule(fillRule FillRule) { DrawSetClipRule(d, fillRule) }

var DrawSetClipUnits func(d *DrawContext, clipUnits ClipPathUnits)

func (d *DrawContext) SetClipUnits(clipUnits ClipPathUnits) { DrawSetClipUnits(d, clipUnits) }

var DrawColor func(d *DrawContext, x, y float64, paintMethod PaintMethod)

func (d *DrawContext) Color(x, y float64, paintMethod PaintMethod) { DrawColor(d, x, y, paintMethod) }

var DrawComment func(d *DrawContext, comment string)

func (d *DrawContext) Comment(comment string) { DrawComment(d, comment) }

var DrawDestroyContext func(d *DrawContext)

func (d *DrawContext) Destroy() { DrawDestroyContext(d) }

var DrawEllipse func(d *DrawContext, ox, oy, rx, ry, start, end float64)

func (d *DrawContext) Ellipse(ox, oy, rx, ry, start, end float64) {
	DrawEllipse(d, ox, oy, rx, ry, start, end)
}

var DrawGetFillColor func(d *DrawContext) PixelPacket

func (d *DrawContext) FillColor() PixelPacket { return DrawGetFillColor(d) }

var DrawSetFillColor func(d *DrawContext, fillColor *PixelPacket)

func (d *DrawContext) SetFillColor(fillColor *PixelPacket) { DrawSetFillColor(d, fillColor) }

var DrawSetFillColorString func(d *DrawContext, fillColor string)

func (d *DrawContext) SetFillColorString(fillColor string) { DrawSetFillColorString(d, fillColor) }

var DrawSetFillPatternURL func(d *DrawContext, fillUrl string)

func (d *DrawContext) SetFillPatternURL(fillUrl string) { DrawSetFillPatternURL(d, fillUrl) }

var DrawGetFillOpacity func(d *DrawContext) float64

func (d *DrawContext) FillOpacity() float64 { return DrawGetFillOpacity(d) }

var DrawSetFillOpacity func(d *DrawContext, fillOpacity float64)

func (d *DrawContext) SetFillOpacity(fillOpacity float64) { DrawSetFillOpacity(d, fillOpacity) }

var DrawGetFillRule func(d *DrawContext) FillRule

func (d *DrawContext) FillRule() FillRule { return DrawGetFillRule(d) }

var DrawSetFillRule func(d *DrawContext, fillRule FillRule)

func (d *DrawContext) SetFillRule(fillRule FillRule) { DrawSetFillRule(d, fillRule) }

var DrawGetFont func(d *DrawContext) string

func (d *DrawContext) Font() string { return DrawGetFont(d) }

var DrawGetFontFamily func(d *DrawContext) string

var DrawSetFont func(d *DrawContext, fontName string)

func (d *DrawContext) SetFont(fontName string) { DrawSetFont(d, fontName) }

func (d *DrawContext) FontFamily() string { return DrawGetFontFamily(d) }

var DrawSetFontFamily func(d *DrawContext, fontFamily string)

func (d *DrawContext) SetFontFamily(fontFamily string) { DrawSetFontFamily(d, fontFamily) }

var DrawGetFontSize func(d *DrawContext) float64

func (d *DrawContext) FontSize() float64 { return DrawGetFontSize(d) }

var DrawGetFontStretch func(d *DrawContext) StretchType

func (d *DrawContext) FontStretch() StretchType { return DrawGetFontStretch(d) }

var DrawGetFontStyle func(d *DrawContext) StyleType

func (d *DrawContext) FontStyle() StyleType { return DrawGetFontStyle(d) }

var DrawGetFontWeight func(d *DrawContext) Size

func (d *DrawContext) FontWeight() Size { return DrawGetFontWeight(d) }

var DrawSetFontSize func(d *DrawContext, fontPointsize float64)

func (d *DrawContext) SetFontSize(fontPointsize float64) { DrawSetFontSize(d, fontPointsize) }

var DrawSetFontStretch func(d *DrawContext, fontStretch StretchType)

func (d *DrawContext) SetFontStretch(fontStretch StretchType) { DrawSetFontStretch(d, fontStretch) }

var DrawSetFontStyle func(d *DrawContext, fontStyle StyleType)

func (d *DrawContext) SetFontStyle(fontStyle StyleType) { DrawSetFontStyle(d, fontStyle) }

var DrawSetFontWeight func(d *DrawContext, fontWeight Size)

func (d *DrawContext) SetFontWeight(fontWeight Size) { DrawSetFontWeight(d, fontWeight) }

var DrawGetGravity func(d *DrawContext) GravityType

func (d *DrawContext) Gravity() GravityType { return DrawGetGravity(d) }

var DrawSetGravity func(d *DrawContext, gravity GravityType)

func (d *DrawContext) SetGravity(gravity GravityType) { DrawSetGravity(d, gravity) }

var DrawComposite func(d *DrawContext, compositeOperator CompositeOperator, x, y, width, height float64, image *Image)

func (d *DrawContext) Composite(compositeOperator CompositeOperator, x, y, width, height float64, image *Image) {
	DrawComposite(d, compositeOperator, x, y, width, height, image)
}

var DrawLine func(d *DrawContext, sx, sy, ex, ey float64)

func (d *DrawContext) Line(sx, sy, ex, ey float64) { DrawLine(d, sx, sy, ex, ey) }

var DrawMatte func(d *DrawContext, x, y float64, paintMethod PaintMethod)

func (d *DrawContext) Matte(x, y float64, paintMethod PaintMethod) { DrawMatte(d, x, y, paintMethod) }

var DrawPathClose func(d *DrawContext)

func (d *DrawContext) PathClose() { DrawPathClose(d) }

var DrawPathCurveToAbsolute func(d *DrawContext, x1, y1, x2, y2, x, y float64)

func (d *DrawContext) PathCurveToAbsolute(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToAbsolute(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToRelative func(d *DrawContext, x1, y1, x2, y2, x, y float64)

func (d *DrawContext) PathCurveToRelative(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToRelative(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToQuadraticBezierAbsolute func(d *DrawContext, x1, y1, x, y float64)

func (d *DrawContext) PathCurveToQuadraticBezierAbsolute(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierAbsolute(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierRelative func(d *DrawContext, x1, y1, x, y float64)

func (d *DrawContext) PathCurveToQuadraticBezierRelative(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierRelative(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothAbsolute func(d *DrawContext, x, y float64)

func (d *DrawContext) PathCurveToQuadraticBezierSmoothAbsolute(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothAbsolute(d, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothRelative func(d *DrawContext, x, y float64)

func (d *DrawContext) PathCurveToQuadraticBezierSmoothRelative(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothRelative(d, x, y)
}

var DrawPathCurveToSmoothAbsolute func(d *DrawContext, x2, y2, x, y float64)

func (d *DrawContext) PathCurveToSmoothAbsolute(x2, y2, x, y float64) {
	DrawPathCurveToSmoothAbsolute(d, x2, y2, x, y)
}

var DrawPathCurveToSmoothRelative func(d *DrawContext, x2, y2, x, y float64)

func (d *DrawContext) PathCurveToSmoothRelative(x2, y2, x, y float64) {
	DrawPathCurveToSmoothRelative(d, x2, y2, x, y)
}

var DrawPathEllipticArcAbsolute func(d *DrawContext, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag uint, x, y float64)

func (d *DrawContext) PathEllipticArcAbsolute(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag uint, x, y float64) {
	DrawPathEllipticArcAbsolute(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathEllipticArcRelative func(d *DrawContext, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag uint, x, y float64)

func (d *DrawContext) PathEllipticArcRelative(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag uint, x, y float64) {
	DrawPathEllipticArcRelative(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathFinish func(d *DrawContext)

func (d *DrawContext) PathFinish() { DrawPathFinish(d) }

var DrawPathLineToAbsolute func(d *DrawContext, x, y float64)

func (d *DrawContext) PathLineToAbsolute(x, y float64) { DrawPathLineToAbsolute(d, x, y) }

var DrawPathLineToRelative func(d *DrawContext, x, y float64)

func (d *DrawContext) PathLineToRelative(x, y float64) { DrawPathLineToRelative(d, x, y) }

var DrawPathLineToHorizontalAbsolute func(d *DrawContext, x float64)

func (d *DrawContext) PathLineToHorizontalAbsolute(x float64) { DrawPathLineToHorizontalAbsolute(d, x) }

var DrawPathLineToHorizontalRelative func(d *DrawContext, x float64)

func (d *DrawContext) PathLineToHorizontalRelative(x float64) { DrawPathLineToHorizontalRelative(d, x) }

var DrawPathLineToVerticalAbsolute func(d *DrawContext, y float64)

func (d *DrawContext) PathLineToVerticalAbsolute(y float64) { DrawPathLineToVerticalAbsolute(d, y) }

var DrawPathLineToVerticalRelative func(d *DrawContext, y float64)

func (d *DrawContext) PathLineToVerticalRelative(y float64) { DrawPathLineToVerticalRelative(d, y) }

var DrawPathMoveToAbsolute func(d *DrawContext, x, y float64)

func (d *DrawContext) PathMoveToAbsolute(x, y float64) { DrawPathMoveToAbsolute(d, x, y) }

var DrawPathMoveToRelative func(d *DrawContext, x, y float64)

func (d *DrawContext) PathMoveToRelative(x, y float64) { DrawPathMoveToRelative(d, x, y) }

var DrawPathStart func(d *DrawContext)

func (d *DrawContext) PathStart() { DrawPathStart(d) }

var DrawPeekGraphicContext func(d *DrawContext) *DrawInfo

func (d *DrawContext) PeekGraphicContext() *DrawInfo { return DrawPeekGraphicContext(d) }

var DrawPoint func(d *DrawContext, x, y float64)

func (d *DrawContext) Point(x, y float64) { DrawPoint(d, x, y) }

var DrawPolygon func(d *DrawContext, numCoords Size, coordinates *PointInfo)

func (d *DrawContext) Polygon(numCoords Size, coordinates *PointInfo) {
	DrawPolygon(d, numCoords, coordinates)
}

var DrawPolyline func(d *DrawContext, numCoords Size, coordinates *PointInfo)

func (d *DrawContext) Polyline(numCoords Size, coordinates *PointInfo) {
	DrawPolyline(d, numCoords, coordinates)
}

var DrawPopClipPath func(d *DrawContext)

func (d *DrawContext) PopClipPath() { DrawPopClipPath(d) }

var DrawPopDefs func(d *DrawContext)

func (d *DrawContext) PopDefs() { DrawPopDefs(d) }

var DrawPopGraphicContext func(d *DrawContext)

func (d *DrawContext) PopGraphicContext() { DrawPopGraphicContext(d) }

var DrawPopPattern func(d *DrawContext)

func (d *DrawContext) PopPattern() { DrawPopPattern(d) }

var DrawPushClipPath func(d *DrawContext, clipPathId string)

func (d *DrawContext) PushClipPath(clipPathId string) { DrawPushClipPath(d, clipPathId) }

var DrawPushDefs func(d *DrawContext)

func (d *DrawContext) PushDefs() { DrawPushDefs(d) }

var DrawPushGraphicContext func(d *DrawContext)

func (d *DrawContext) PushGraphicContext() { DrawPushGraphicContext(d) }

var DrawPushPattern func(d *DrawContext, patternId string, x, y, width, height float64)

func (d *DrawContext) PushPattern(patternId string, x, y, width, height float64) {
	DrawPushPattern(d, patternId, x, y, width, height)
}

var DrawRectangle func(d *DrawContext, x1, y1, x2, y2 float64)

func (d *DrawContext) Rectangle(x1, y1, x2, y2 float64) { DrawRectangle(d, x1, y1, x2, y2) }

var DrawRender func(d *DrawContext) int

func (d *DrawContext) Render() int { return DrawRender(d) }

var DrawRotate func(d *DrawContext, degrees float64)

func (d *DrawContext) Rotate(degrees float64) { DrawRotate(d, degrees) }

var DrawRoundRectangle func(d *DrawContext, x1, y1, x2, y2, rx, ry float64)

func (d *DrawContext) RoundRectangle(x1, y1, x2, y2, rx, ry float64) {
	DrawRoundRectangle(d, x1, y1, x2, y2, rx, ry)
}

var DrawScale func(d *DrawContext, x, y float64)

func (d *DrawContext) Scale(x, y float64) { DrawScale(d, x, y) }

var DrawSkewX func(d *DrawContext, degrees float64)

func (d *DrawContext) SkewX(degrees float64) { DrawSkewX(d, degrees) }

var DrawSkewY func(d *DrawContext, degrees float64)

func (d *DrawContext) SkewY(degrees float64) { DrawSkewY(d, degrees) }

//NOTE(t): DrawSetStopColor missing from dll.

var DrawGetStrokeColor func(d *DrawContext) PixelPacket

func (d *DrawContext) StrokeColor() PixelPacket { return DrawGetStrokeColor(d) }

var DrawSetStrokeColor func(d *DrawContext, strokeColor *PixelPacket)

func (d *DrawContext) SetStrokeColor(strokeColor *PixelPacket) { DrawSetStrokeColor(d, strokeColor) }

var DrawSetStrokeColorString func(d *DrawContext, strokeColor string)

func (d *DrawContext) SetStrokeColorString(strokeColor string) {
	DrawSetStrokeColorString(d, strokeColor)
}

var DrawSetStrokePatternURL func(d *DrawContext, strokeUrl string)

func (d *DrawContext) SetStrokePatternURL(strokeUrl string) { DrawSetStrokePatternURL(d, strokeUrl) }

var DrawGetStrokeAntialias func(d *DrawContext) bool

func (d *DrawContext) StrokeAntialias() bool { return DrawGetStrokeAntialias(d) }

var DrawSetStrokeAntialias func(d *DrawContext, trueFalse bool)

func (d *DrawContext) SetStrokeAntialias(trueFalse bool) { DrawSetStrokeAntialias(d, trueFalse) }

var DrawGetStrokeDashArray func(d *DrawContext, numElems *Size) []float64

func (d *DrawContext) StrokeDashArray(numElems *Size) []float64 {
	return DrawGetStrokeDashArray(d, numElems)
}

var DrawSetStrokeDashArray func(d *DrawContext, numElems Size, dasharray []float64)

func (d *DrawContext) SetStrokeDashArray(numElems Size, dasharray []float64) {
	DrawSetStrokeDashArray(d, numElems, dasharray)
}

var DrawGetStrokeDashOffset func(d *DrawContext) float64

func (d *DrawContext) StrokeDashOffset() float64 { return DrawGetStrokeDashOffset(d) }

var DrawSetStrokeDashOffset func(d *DrawContext, dashoffset float64)

func (d *DrawContext) SetStrokeDashOffset(dashoffset float64) { DrawSetStrokeDashOffset(d, dashoffset) }

var DrawGetStrokeLineCap func(d *DrawContext) LineCap

func (d *DrawContext) StrokeLineCap() LineCap { return DrawGetStrokeLineCap(d) }

var DrawSetStrokeLineCap func(d *DrawContext, linecap LineCap)

func (d *DrawContext) SetStrokeLineCap(linecap LineCap) { DrawSetStrokeLineCap(d, linecap) }

var DrawGetStrokeLineJoin func(d *DrawContext) LineJoin

func (d *DrawContext) StrokeLineJoin() LineJoin { return DrawGetStrokeLineJoin(d) }

var DrawSetStrokeLineJoin func(d *DrawContext, linejoin LineJoin)

func (d *DrawContext) SetStrokeLineJoin(linejoin LineJoin) { DrawSetStrokeLineJoin(d, linejoin) }

var DrawGetStrokeMiterLimit func(d *DrawContext) Size

func (d *DrawContext) StrokeMiterLimit() Size { return DrawGetStrokeMiterLimit(d) }

var DrawSetStrokeMiterLimit func(d *DrawContext, miterlimit Size)

func (d *DrawContext) SetStrokeMiterLimit(miterlimit Size) {
	DrawSetStrokeMiterLimit(d, miterlimit)
}

var DrawGetStrokeOpacity func(d *DrawContext) float64

func (d *DrawContext) StrokeOpacity() float64 { return DrawGetStrokeOpacity(d) }

var DrawSetStrokeOpacity func(d *DrawContext, opacity float64)

func (d *DrawContext) SetStrokeOpacity(opacity float64) { DrawSetStrokeOpacity(d, opacity) }

var DrawGetStrokeWidth func(d *DrawContext) float64

func (d *DrawContext) StrokeWidth() float64 { return DrawGetStrokeWidth(d) }

var DrawSetStrokeWidth func(d *DrawContext, width float64)

func (d *DrawContext) SetStrokeWidth(width float64) { DrawSetStrokeWidth(d, width) }

var DrawGetTextAntialias func(d *DrawContext) bool

func (d *DrawContext) TextAntialias() bool { return DrawGetTextAntialias(d) }

var DrawSetTextAntialias func(d *DrawContext, trueFalse bool)

func (d *DrawContext) SetTextAntialias(trueFalse bool) { DrawSetTextAntialias(d, trueFalse) }

var DrawGetTextDecoration func(d *DrawContext) DecorationType

func (d *DrawContext) TextDecoration() DecorationType { return DrawGetTextDecoration(d) }

var DrawSetTextDecoration func(d *DrawContext, decoration DecorationType)

func (d *DrawContext) SetTextDecoration(decoration DecorationType) {
	DrawSetTextDecoration(d, decoration)
}

var DrawGetTextEncoding func(d *DrawContext) string

func (d *DrawContext) TextEncoding() string { return DrawGetTextEncoding(d) }

var DrawSetTextEncoding func(d *DrawContext, encoding string)

func (d *DrawContext) SetTextEncoding(encoding string) { DrawSetTextEncoding(d, encoding) }

var DrawGetTextUnderColor func(d *DrawContext) PixelPacket

func (d *DrawContext) TextUnderColor() PixelPacket { return DrawGetTextUnderColor(d) }

var DrawSetTextUnderColor func(d *DrawContext, color *PixelPacket)

func (d *DrawContext) SetTextUnderColor(color *PixelPacket) { DrawSetTextUnderColor(d, color) }

var DrawSetTextUnderColorString func(d *DrawContext, underColor string)

func (d *DrawContext) SetTextUnderColorString(underColor string) {
	DrawSetTextUnderColorString(d, underColor)
}

var DrawTranslate func(d *DrawContext, x, y float64)

func (d *DrawContext) Translate(x, y float64) { DrawTranslate(d, x, y) }

var DrawSetViewbox func(d *DrawContext, x1, y1, x2, y2 Size)

func (d *DrawContext) SetViewbox(x1, y1, x2, y2 Size) { DrawSetViewbox(d, x1, y1, x2, y2) }

// Effects

var AdaptiveThresholdImage func(i *Image, width, height Size, offset SSize, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveThreshold(width, height Size, offset SSize, exception *ExceptionInfo) *Image {
	return AdaptiveThresholdImage(i, width, height, offset, exception)
}

var AddNoiseImage func(i *Image, noiseType NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoise(noiseType NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImage(i, noiseType, exception)
}

var AddNoiseImageChannel func(i *Image, channel ChannelType, noiseType NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoiseChannel(channel ChannelType, noiseType NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImageChannel(i, channel, noiseType, exception)
}

var BlackThresholdImage func(i *Image, threshold string) bool

func (i *Image) BlackThreshold(threshold string) bool { return BlackThresholdImage(i, threshold) }

var BlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) Blur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImage(i, radius, sigma, exception)
}

var BlurImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) BlurChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImageChannel(i, channel, radius, sigma, exception)
}

var ChannelThresholdImage func(i *Image, level string) bool

func (i *Image) ChannelThreshold(level string) bool { return ChannelThresholdImage(i, level) }

var ConvolveImage func(i *Image, order Size, kernel *float64, exception *ExceptionInfo) *Image

func (i *Image) Convolve(order Size, kernel *float64, exception *ExceptionInfo) *Image {
	return ConvolveImage(i, order, kernel, exception)
}

var DespeckleImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Despeckle(exception *ExceptionInfo) *Image { return DespeckleImage(i, exception) }

var EdgeImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

func (i *Image) Edge(radius float64, exception *ExceptionInfo) *Image {
	return EdgeImage(i, radius, exception)
}

var EmbossImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) Emboss(radius, sigma float64, exception *ExceptionInfo) *Image {
	return EmbossImage(i, radius, sigma, exception)
}

var EnhanceImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Enhance(exception *ExceptionInfo) *Image { return EnhanceImage(i, exception) }

var GaussianBlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) GaussianBlur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImage(i, radius, sigma, exception)
}

var GaussianBlurImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) GaussianBlurChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImageChannel(i, channel, radius, sigma, exception)
}

var MedianFilterImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

func (i *Image) MedianFilter(radius float64, exception *ExceptionInfo) *Image {
	return MedianFilterImage(i, radius, exception)
}

var MotionBlurImage func(i *Image, radius, sigma, angle float64, exception *ExceptionInfo) *Image

func (i *Image) MotionBlur(radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return MotionBlurImage(i, radius, sigma, angle, exception)
}

var RandomChannelThresholdImage func(i *Image, channel, thresholds string, exception *ExceptionInfo) bool

func (i *Image) RandomChannelThresholdImage(channel, thresholds string, exception *ExceptionInfo) bool {
	return RandomChannelThresholdImage(i, channel, thresholds, exception)
}

var ReduceNoiseImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

func (i *Image) ReduceNoise(radius float64, exception *ExceptionInfo) *Image {
	return ReduceNoiseImage(i, radius, exception)
}

var ShadeImage func(i *Image, gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image

func (i *Image) Shade(gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image {
	return ShadeImage(i, gray, azimuth, elevation, exception)
}

var SharpenImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) Sharpen(radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImage(i, radius, sigma, exception)
}

var SharpenImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) SharpenChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImageChannel(i, channel, radius, sigma, exception)
}

var SpreadImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

func (i *Image) Spread(radius float64, exception *ExceptionInfo) *Image {
	return SpreadImage(i, radius, exception)
}

var ThresholdImage func(i *Image, threshold float64) bool

func (i *Image) Threshold(threshold float64) bool { return ThresholdImage(i, threshold) }

var UnsharpMaskImage func(i *Image, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image

func (i *Image) UnsharpMask(radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image {
	return UnsharpMaskImage(i, radius, sigma, amount, threshold, exception)
}

var UnsharpMaskImageChannel func(i *Image, channel ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image

func (i *Image) UnsharpMaskChannel(channel ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image {
	return UnsharpMaskImageChannel(i, channel, radius, sigma, amount, threshold, exception)
}

var WhiteThresholdImage func(i *Image, threshold string) bool

func (i *Image) WhiteThreshold(threshold string) bool { return WhiteThresholdImage(i, threshold) }

// Enhance

var ContrastImage func(i *Image, sharpen bool) bool

func (i *Image) Contrast(sharpen bool) bool { return ContrastImage(i, sharpen) }

var EqualizeImage func(i *Image) bool

func (i *Image) Equalize() bool { return EqualizeImage(i) }

var GammaImage func(i *Image, level string) bool

func (i *Image) Gamma(level string) bool { return GammaImage(i, level) }

var LevelImage func(i *Image, levels string) bool

func (i *Image) Level(levels string) bool { return LevelImage(i, levels) }

var LevelImageChannel func(i *Image, channel ChannelType, blackPoint, whitePoint, gamma float64) bool

func (i *Image) LevelChannel(channel ChannelType, blackPoint, whitePoint, gamma float64) bool {
	return LevelImageChannel(i, channel, blackPoint, whitePoint, gamma)
}

var ModulateImage func(i *Image, modulate string) bool

func (i *Image) Modulate(modulate string) bool { return ModulateImage(i, modulate) }

var NegateImage func(i *Image, grayscale bool) bool

func (i *Image) Negate(grayscale bool) bool { return NegateImage(i, grayscale) }

var NormalizeImage func(i *Image) bool

func (i *Image) Normalize() bool { return NormalizeImage(i) }

// Error

var CatchException func(e *ExceptionInfo)

func (e *ExceptionInfo) Catch() { CatchException(e) }

var CopyException func(e *ExceptionInfo, original *ExceptionInfo)

func (e *ExceptionInfo) Copy(original *ExceptionInfo) { CopyException(e, original) }

var DestroyExceptionInfo func(e *ExceptionInfo) *ExceptionInfo

func (e *ExceptionInfo) Destroy() *ExceptionInfo { return DestroyExceptionInfo(e) }

var GetExceptionInfo func(e *ExceptionInfo)

func (e *ExceptionInfo) Get() { GetExceptionInfo(e) }

var GetLocaleExceptionMessage func(severity ExceptionType, tag string) string

var MagickError func(err ExceptionType, reason, description string)

//MagickError2 is the export name for GraphhicsMagick name _MagickError
var MagickError2 func(err ExceptionType, reason, description string)

var MagickFatalError func(err ExceptionType, reason, description string)

//MagickFatalError2 is the export name for GraphhicsMagick name _MagickFatalError
var MagickFatalError2 func(err ExceptionType, reason, description string)

var MagickWarning func(warning ExceptionType, reason, description string)

//MagickWarning2 is the export name for GraphhicsMagick name _MagickWarning
var MagickWarning2 func(warning ExceptionType, reason, description string)

var SetErrorHandler func(handler ErrorHandler) ErrorHandler

type ErrorHandler func(ExceptionType, string, string)

var SetExceptionInfo func(e *ExceptionInfo, severity ExceptionType)

func (e *ExceptionInfo) Set(severity ExceptionType) { SetExceptionInfo(e, severity) }

var SetFatalErrorHandler func(handler FatalErrorHandler) FatalErrorHandler

type FatalErrorHandler func(ExceptionType, string, string)

var SetWarningHandler func(handler WarningHandler) WarningHandler

type WarningHandler func(ExceptionType, string, string)

var ThrowException func(e *ExceptionInfo, severity ExceptionType, reason, description string)

func (e *ExceptionInfo) Throw(severity ExceptionType, reason, description string) {
	ThrowException(e, severity, reason, description)
}

var ThrowLoggedException func(e *ExceptionInfo, severity ExceptionType, reason, description, module, function string, line Size)

func (e *ExceptionInfo) ThrowLogged(severity ExceptionType, reason, description, module, function string, line Size) {
	ThrowLoggedException(e, severity, reason, description, module, function, line)
}

// Export

var ExportImagePixelArea func(i *Image, quantumType QuantumType, quantumSize uint, destination *byte, options *ExportPixelAreaOptions, exportInfo *ExportPixelAreaInfo) bool

func (i *Image) ExportPixelArea(quantumType QuantumType, quantumSize uint, destination *byte, options *ExportPixelAreaOptions, exportInfo *ExportPixelAreaInfo) bool {
	return ExportImagePixelArea(i, quantumType, quantumSize, destination, options, exportInfo)
}

var ExportViewPixelArea func(v *ViewInfo, quantumType QuantumType, quantumSize uint, destination *byte, options *ExportPixelAreaOptions, exportInfo *ExportPixelAreaInfo) bool

func (v *ViewInfo) ExportPixelArea(quantumType QuantumType, quantumSize uint, destination *byte, options *ExportPixelAreaOptions, exportInfo *ExportPixelAreaInfo) bool {
	return ExportViewPixelArea(v, quantumType, quantumSize, destination, options, exportInfo)
}

var ExportPixelAreaOptionsInit func(options *ExportPixelAreaOptions)

// Fx

var CharcoalImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) Charcoal(radius, sigma float64, exception *ExceptionInfo) *Image {
	return CharcoalImage(i, radius, sigma, exception)
}

var ColorizeImage func(i *Image, opacity string, colorize PixelPacket, exception *ExceptionInfo) *Image

func (i *Image) Colorize(opacity string, colorize PixelPacket, exception *ExceptionInfo) *Image {
	return ColorizeImage(i, opacity, colorize, exception)
}

var ColorMatrixImage func(i *Image, colorMatrix []float64, exception *ExceptionInfo) *Image

func (i *Image) ColorMatrix(colorMatrix []float64, exception *ExceptionInfo) *Image {
	return ColorMatrixImage(i, colorMatrix, exception)
}

var ImplodeImage func(i *Image, amount float64, exception *ExceptionInfo) *Image

func (i *Image) Implode(amount float64, exception *ExceptionInfo) *Image {
	return ImplodeImage(i, amount, exception)
}

var MorphImages func(i *Image, numberFrames Size, exception *ExceptionInfo) *Image

func (i *Image) MorphImages(numberFrames Size, exception *ExceptionInfo) *Image {
	return MorphImages(i, numberFrames, exception)
}

var OilPaintImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

func (i *Image) OilPaint(radius float64, exception *ExceptionInfo) *Image {
	return OilPaintImage(i, radius, exception)
}

var SolarizeImage func(i *Image, threshold float64) bool

func (i *Image) Solarize(threshold float64) bool { return SolarizeImage(i, threshold) }

var SteganoImage func(i *Image, watermark *Image, exception *ExceptionInfo) *Image

func (i *Image) Stegano(watermark *Image, exception *ExceptionInfo) *Image {
	return SteganoImage(i, watermark, exception)
}

var StereoImage func(i *Image, offsetImage *Image, exception *ExceptionInfo) *Image

func (i *Image) Stereo(offsetImage *Image, exception *ExceptionInfo) *Image {
	return StereoImage(i, offsetImage, exception)
}

var SwirlImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image

func (i *Image) Swirl(degrees float64, exception *ExceptionInfo) *Image {
	return SwirlImage(i, degrees, exception)
}

var WaveImage func(i *Image, amplitude, waveLength float64, exception *ExceptionInfo) *Image

func (i *Image) Wave(amplitude, waveLength float64, exception *ExceptionInfo) *Image {
	return WaveImage(i, amplitude, waveLength, exception)
}

// Hald CLUT

var HaldClutImage func(i, haldImage *Image) bool

func (i *Image) HaldClut(haldImage *Image) bool { return HaldClutImage(i, haldImage) }

// Image

var AccessDefinition func(i *ImageInfo, magick, key string) string

func (i *ImageInfo) AccessDefinition(magick, key string) string {
	return AccessDefinition(i, magick, key)
}

var AddDefinition func(i *ImageInfo, magick, key, value string, exception *ExceptionInfo) bool

func (i *ImageInfo) AddDefinition(magick, key, value string, exception *ExceptionInfo) bool {
	return AddDefinition(i, magick, key, value, exception)
}

//NOTE(t): Documentation wrong (missing exception)
var AddDefinitions func(i *ImageInfo, options string, exception *ExceptionInfo) bool

func (i *ImageInfo) AddDefinitions(options string, exception *ExceptionInfo) bool {
	return AddDefinitions(i, options, exception)
}

var AllocateImage func(i *ImageInfo) *Image

func (i *ImageInfo) AllocateImage() *Image { return AllocateImage(i) }

var AllocateNextImage func(i *ImageInfo, image *Image)

func (i *ImageInfo) AllocateNextImage(image *Image) { AllocateNextImage(i, image) }

var AnimateImages func(i *ImageInfo, images *Image) bool

func (i *ImageInfo) AnimateImages(images *Image) bool { return AnimateImages(i, images) }

var AppendImages func(i *Image, stack bool, exception *ExceptionInfo) *Image

func (i *Image) AppendImages(stack bool, exception *ExceptionInfo) *Image {
	return AppendImages(i, stack, exception)
}

var CatchImageException func(i *Image) ExceptionType

func (i *Image) CatchException() ExceptionType { return CatchImageException(i) }

var ClipPathImage func(i *Image, pathname string, inside bool) bool

func (i *Image) ClipPathImage(pathname string, inside bool) bool {
	return ClipPathImage(i, pathname, inside)
}

var CloneImage func(i *Image, columns, rows Size, orphan bool, exception *ExceptionInfo) *Image

func (i *Image) Clone(columns, rows Size, orphan bool, exception *ExceptionInfo) *Image {
	return CloneImage(i, columns, rows, orphan, exception)
}

var CloneImageInfo func(i *ImageInfo) *ImageInfo

func (i *ImageInfo) CloneInfo() *ImageInfo { return CloneImageInfo(i) }

var DestroyImage func(i *Image) *Image

func (i *Image) Destroy() *Image { return DestroyImage(i) }

var DestroyImageInfo func(i *ImageInfo) *ImageInfo

func (i *ImageInfo) Destroy() *ImageInfo { return DestroyImageInfo(i) }

var DisplayImages func(i *ImageInfo, images *Image) bool

func (i *ImageInfo) DisplayImages(images *Image) bool { return DisplayImages(i, images) }

var GetImageClipMask func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) ClipMask(exception *ExceptionInfo) *Image {
	return GetImageClipMask(i, exception)
}

var GetImageException func(i *Image, exception *ExceptionInfo)

func (i *Image) Exception(exception *ExceptionInfo) { GetImageException(i, exception) }

var GetImageGeometry func(i *Image, geometry string, sizeToFit uint, regionInfo *RectangleInfo) int

func (i *Image) Geometry(geometry string, sizeToFit uint, regionInfo *RectangleInfo) int {
	return GetImageGeometry(i, geometry, sizeToFit, regionInfo)
}

var GetImageInfo func(i *ImageInfo)

func (i *ImageInfo) Get() { GetImageInfo(i) }

var IsTaintImage func(i *Image) bool

func (i *Image) Taint() bool { return IsTaintImage(i) }

var ModifyImage func(image **Image, exception *ExceptionInfo) bool

var ReferenceImage func(i *Image) *Image

func (i *Image) Reference() *Image { return ReferenceImage(i) }

var RemoveDefinitions func(i *ImageInfo, options string) bool

func (i *ImageInfo) RemoveDefinitions(options string) bool { return RemoveDefinitions(i, options) }

var ResetImagePage func(i *Image, page string) bool

func (i *Image) ResetPage(page string) bool { return ResetImagePage(i, page) }

var SetImage func(i *Image, opacity Quantum)

func (i *Image) SetImage(opacity Quantum) { SetImage(i, opacity) }

var SetImageColor func(i *Image, color *PixelPacket) bool

func (i *Image) SetColor(color *PixelPacket) bool { return SetImageColor(i, color) }

var SetImageColorRegion func(i *Image, x, y SSize, width, height Size, pixel *PixelPacket) bool

func (i *Image) SetColorRegion(x, y SSize, width, height Size, pixel *PixelPacket) bool {
	return SetImageColorRegion(i, x, y, width, height, pixel)
}

var SetImageClipMask func(i *Image, clipMask *Image) bool

func (i *Image) SetClipMask(clipMask *Image) bool { return SetImageClipMask(i, clipMask) }

var SetImageDepth func(i *Image, depth Size) bool

func (i *Image) SetDepth(depth Size) bool { return SetImageDepth(i, depth) }

var SetImageOpacity func(i *Image, opacity Quantum) bool

func (i *Image) SetOpacity(opacity Quantum) bool { return SetImageOpacity(i, opacity) }

var SetImageType func(i *Image, imageType ImageType) bool

func (i *Image) SetType(imageType ImageType) bool { return SetImageType(i, imageType) }

var StripImage func(i *Image) bool

func (i *Image) Strip() bool { return StripImage(i) }

var ImportImagePixelArea func(i *Image, quantumType QuantumType, quantumSize uint, source *byte, options *ImportPixelAreaOptions, importInfo *ImportPixelAreaInfo) bool

func (i *Image) ImportPixelArea(quantumType QuantumType, quantumSize uint, source *byte, options *ImportPixelAreaOptions, importInfo *ImportPixelAreaInfo) bool {
	return ImportImagePixelArea(i, quantumType, quantumSize, source, options, importInfo)
}

var ImportImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) ImportImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return ImportImageCommand(i, argc, argv, metadata, exception)
}

var ImportViewPixelArea func(v *ViewInfo, quantumType QuantumType, quantumSize uint, source *byte, options *ImportPixelAreaOptions, importInfo *ImportPixelAreaInfo) bool

func (v *ViewInfo) ImportPixelArea(quantumType QuantumType, quantumSize uint, source *byte, options *ImportPixelAreaOptions, importInfo *ImportPixelAreaInfo) bool {
	return ImportViewPixelArea(v, quantumType, quantumSize, source, options, importInfo)
}

// List

//TODO(t): Sometime *list others **list; harmonise Go-wise
var AppendImageToList func(images []*Image, image *Image)
var CloneImageList func(images *Image, exception *ExceptionInfo) *Image
var DeleteImageFromList func(images []*Image)
var DestroyImageList func(images *Image) *Image
var GetFirstImageInList func(images *Image) *Image
var GetImageFromList func(images *Image, index SSize) *Image
var GetImageIndexInList func(images *Image) SSize
var GetImageListLength func(images *Image) Size
var GetLastImageInList func(images *Image) *Image
var GetNextImageInList func(images *Image) *Image
var GetPreviousImageInList func(images *Image) *Image
var ImageListToArray func(images *Image, exception *ExceptionInfo) []*Image
var InsertImageInList func(images []*Image, image *Image)
var NewImageList func() *Image
var PrependImageToList func(images []*Image, image *Image)
var RemoveFirstImageFromList func(images []*Image) *Image
var RemoveLastImageFromList func(images []*Image) *Image
var ReplaceImageInList func(images []*Image, image *Image)
var ReverseImageList func(images []*Image)
var SpliceImageIntoList func(images []*Image, length Size, splice *Image) *Image
var SplitImageList func(images *Image) *Image

// Magick
//NOTE(t): DestroyMagicInfoList is static C

var DestroyMagick func()
var DestroyMagicInfo func()
var GetImageMagick func(magick *byte, length uint32) string
var GetMagickInfo func(name string, exception *ExceptionInfo) *MagickInfo
var GetMagickInfoArray func(exception *ExceptionInfo) **MagickInfo
var InitializeMagick func(path string)
var IsMagickConflict func(magick string) bool
var ListMagickInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagickInfo(exception *ExceptionInfo) bool { return ListMagickInfo(f, exception) }

var ListModuleMap func(file *FILE, exception *ExceptionInfo) bool
var MagickToMime func(magick string) string
var RegisterMagickInfo func(m *MagickInfo) *MagickInfo

func (m *MagickInfo) Register() *MagickInfo { return RegisterMagickInfo(m) }

var SetMagickInfo func(name string) *MagickInfo
var UnregisterMagickInfo func(name string) bool

// Memory

var MagickAllocFunctions func(freeFunc MagickFreeFunc, mallocFunc MagickMallocFunc, reallocFunc MagickReallocFunc)
var MagickMalloc func(size uint32) *Void
var MagickMallocAligned func(alignment, size uint32) *Void
var MagickMallocCleared func(size uint32) *Void
var MagickCloneMemory func(destination, source *Void, size uint32) *Void
var MagickRealloc func(memory *Void, size uint32) *Void
var MagickFree func(memory *Void)
var MagickFreeAligned func(memory *Void)

//NOTE(t): MagickMallocAlignedArray undocumented
var MagickMallocAlignedArray func(alignment, count, size uint32) *Void
var MagickMallocArray func(count, size uint32) *Void

// Monitor

var MagickMonitor func(text string, offset int64, span uint64, clientData *Void) bool
var MagickMonitorFormatted func(quantum int64, span uint64, exception *ExceptionInfo, format string, va ...VArg) bool
var SetMonitorHandler func(handler MonitorHandler) MonitorHandler

// Montage

var CloneMontageInfo func(i *ImageInfo, montageInfo *MontageInfo) *MontageInfo
var DestroyMontageInfo func(montageInfo *MontageInfo) *MontageInfo

func (i *ImageInfo) CloneMontageInfo(montageInfo *MontageInfo) *MontageInfo {
	return CloneMontageInfo(i, montageInfo)
}

var GetMontageInfo func(i *ImageInfo, montageInfo *MontageInfo)

func (i *ImageInfo) MontageInfo(montageInfo *MontageInfo) { GetMontageInfo(i, montageInfo) }

var MontageImages func(images *Image, montageInfo *MontageInfo, exception *ExceptionInfo) *Image

// Operator

var QuantumOperatorImage func(i *Image, channel ChannelType, quantumOperator QuantumOperator, rvalue float64, exception *ExceptionInfo) bool

func (i *Image) QuantumOperator(channel ChannelType, quantumOperator QuantumOperator, rvalue float64, exception *ExceptionInfo) bool {
	return QuantumOperatorImage(i, channel, quantumOperator, rvalue, exception)
}

var QuantumOperatorRegionImage func(i *Image, x, y SSize, columns, rows Size, channel ChannelType, quantumOperator QuantumOperator, rvalue float64, exception *ExceptionInfo) bool

func (i *Image) QuantumOperatorRegion(x, y SSize, columns, rows Size, channel ChannelType, quantumOperator QuantumOperator, rvalue float64, exception *ExceptionInfo) bool {
	return QuantumOperatorRegionImage(i, x, y, columns, rows, channel, quantumOperator, rvalue, exception)
}

// Paint

var ColorFloodfillImage func(i *Image, drawInfo *DrawInfo, target PixelPacket, xOffset, yOffset SSize, method PaintMethod) bool

func (i *Image) ColorFloodfill(drawInfo *DrawInfo, target PixelPacket, xOffset, yOffset SSize, method PaintMethod) bool {
	return ColorFloodfillImage(i, drawInfo, target, xOffset, yOffset, method)
}

var MatteFloodfillImage func(i *Image, target PixelPacket, opacity Quantum, xOffset, yOffset SSize, method PaintMethod) bool

func (i *Image) MatteFloodfill(target PixelPacket, opacity Quantum, xOffset, yOffset SSize, method PaintMethod) bool {
	return MatteFloodfillImage(i, target, opacity, xOffset, yOffset, method)
}

var OpaqueImage func(i *Image, target, fill PixelPacket) bool

func (i *Image) Opaque(target, fill PixelPacket) bool { return OpaqueImage(i, target, fill) }

var TransparentImage func(i *Image, target PixelPacket, opacity Quantum) bool

func (i *Image) Transparent(target PixelPacket, opacity Quantum) bool {
	return TransparentImage(i, target, opacity)
}

// Pixel Cache

var AccessCacheViewPixels func(v *ViewInfo) *PixelPacket

func (v *ViewInfo) AccessCachePixels() *PixelPacket { return AccessCacheViewPixels(v) }

var GetCacheViewArea func(v *ViewInfo) int64

func (v *ViewInfo) CacheArea() int64 { return GetCacheViewArea(v) }

var GetCacheViewRegion func(v *ViewInfo) RectangleInfo

func (v *ViewInfo) CacheRegion() RectangleInfo { return GetCacheViewRegion(v) }

var AccessImmutableIndexes func(i *Image) *IndexPacket

func (i *Image) AccessImmutableIndexes() *IndexPacket { return AccessImmutableIndexes(i) }

var AccessMutableIndexes func(i *Image) *IndexPacket

func (i *Image) AccessMutableIndexes() *IndexPacket { return AccessMutableIndexes(i) }

var AccessMutablePixels func(i *Image) *PixelPacket

func (i *Image) AccessMutablePixels() *PixelPacket { return AccessMutablePixels(i) }

var AcquireCacheViewPixels func(c *ViewInfo, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

func (c *ViewInfo) AcquirePixels(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return AcquireCacheViewPixels(c, x, y, columns, rows, exception)
}

var AcquireImagePixels func(i *Image, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

func (i *Image) AcquireImagePixels(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return AcquireImagePixels(i, x, y, columns, rows, exception)
}

var AcquireCacheViewIndexes func(c *ViewInfo) *IndexPacket

func (c *ViewInfo) AcquireIndexes() *IndexPacket { return AcquireCacheViewIndexes(c) }

var AcquireOneCacheViewPixel func(c *ViewInfo, x, y SSize, exception *ExceptionInfo) PixelPacket

func (c *ViewInfo) AcquirePixel(x, y SSize, exception *ExceptionInfo) PixelPacket {
	return AcquireOneCacheViewPixel(c, x, y, exception)
}

var AcquireOnePixel func(i *Image, x, y SSize, exception *ExceptionInfo) PixelPacket

func (i *Image) AcquireOnePixel(x, y SSize, exception *ExceptionInfo) PixelPacket {
	return AcquireOnePixel(i, x, y, exception)
}

var AcquireOnePixelByReference func(i *Image, pixel *PixelPacket, x, y SSize, exception *ExceptionInfo) bool

func (i *Image) AcquirePixelByReference(pixel *PixelPacket, x, y SSize, exception *ExceptionInfo) bool {
	return AcquireOnePixelByReference(i, pixel, x, y, exception)
}

//NOTE(t): DestroyCacheInfo not exported

var GetCacheViewPixels func(c *ViewInfo, x, y SSize, columns, rows Size) *PixelPacket

func (c *ViewInfo) Pixels(x, y SSize, columns, rows Size) *PixelPacket {
	return GetCacheViewPixels(c, x, y, columns, rows)
}

//NOTE(t): GetCacheViewImage not exported

var GetCacheViewIndexes func(c *ViewInfo) *IndexPacket

func (c *ViewInfo) ViewIndexes() *IndexPacket { return GetCacheViewIndexes(c) }

var GetImagePixels func(i *Image, x, y SSize, columns, rows Size) *PixelPacket

func (i *Image) ImagePixels(x, y SSize, columns, rows Size) *PixelPacket {
	return GetImagePixels(i, x, y, columns, rows)
}

var GetImagePixelsEx func(i *Image, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

func (i *Image) PixelsEx(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return GetImagePixelsEx(i, x, y, columns, rows, exception)
}

var GetImageVirtualPixelMethod func(i *Image) VirtualPixelMethod

func (i *Image) VirtualPixelMethod() VirtualPixelMethod { return GetImageVirtualPixelMethod(i) }

var GetIndexes func(i *Image) *IndexPacket

func (i *Image) Indexes() *IndexPacket { return GetIndexes(i) }

var GetOnePixel func(i *Image, x, y SSize) PixelPacket

func (i *Image) OnePixel(x, y SSize) PixelPacket { return GetOnePixel(i, x, y) }

var GetPixels func(i *Image) *PixelPacket

func (i *Image) Pixels() *PixelPacket { return GetPixels(i) }

//NOTE(t): ModifyCache not exported

var OpenCacheView func(i *Image) *ViewInfo

func (i *Image) OpenCacheView() *ViewInfo { return OpenCacheView(i) }

//NOTE(t): ReferenceCache not exported

var SetCacheViewPixels func(c *ViewInfo, x, y int32, columns, rows uint32) *PixelPacket

func (c *ViewInfo) SetPixels(x, y int32, columns, rows uint32) *PixelPacket {
	return SetCacheViewPixels(c, x, y, columns, rows)
}

var SetImagePixels func(i *Image, x, y SSize, columns, rows Size) *PixelPacket

func (i *Image) SetPixels(x, y SSize, columns, rows Size) *PixelPacket {
	return SetImagePixels(i, x, y, columns, rows)
}

var SetImagePixelsEx func(i *Image, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

func (i *Image) SetPixelsEx(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return SetImagePixelsEx(i, x, y, columns, rows, exception)
}

var SetImageVirtualPixelMethod func(i *Image, virtualPixelMethod VirtualPixelMethod) VirtualPixelMethod

func (i *Image) SetVirtualPixelMethod(virtualPixelMethod VirtualPixelMethod) VirtualPixelMethod {
	return SetImageVirtualPixelMethod(i, virtualPixelMethod)
}

var SyncCacheViewPixels func(c *ViewInfo) bool

func (c *ViewInfo) SyncPixels() bool { return SyncCacheViewPixels(c) }

var SyncImagePixels func(i *Image) bool

func (i *Image) SyncPixels() bool { return SyncImagePixels(i) }

var SyncImagePixelsEx func(i *Image, exception *ExceptionInfo) bool

func (i *Image) SyncPixelsEx(exception *ExceptionInfo) bool { return SyncImagePixelsEx(i, exception) }

// Pixel Iterator

var InitializePixelIteratorOptions func(options *PixelIteratorOptions, exception *ExceptionInfo)
var PixelIterateMonoRead func(callBack PixelIteratorMonoReadCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, x, y SSize, columns, rows Size, image *Image, exception *ExceptionInfo) bool
var PixelIterateMonoModify func(callBack PixelIteratorMonoModifyCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, x, y SSize, columns, rows Size, image *Image, exception *ExceptionInfo) bool
var PixelIterateDualRead func(callBack PixelIteratorDualReadCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, columns, rows Size, firstImage *Image, firstX, firstY SSize, secondImage *Image, secondX, secondY SSize, exception *ExceptionInfo) bool
var PixelIterateDualModify func(callBack PixelIteratorDualModifyCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, columns, rows Size, sourceImage *Image, sourceX, sourceY SSize, updateImage *Image, updateX, updateY SSize, exception *ExceptionInfo) bool
var PixelIterateDualNew func(callBack PixelIteratorDualNewCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, columns, rows Size, sourceImage *Image, sourceX, sourceY SSize, newImage *Image, newX, newY SSize, exception *ExceptionInfo) bool
var PixelIterateTripleModify func(callBack PixelIteratorTripleModifyCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, columns, rows Size, source1Image, source2Image *Image, sourceX, sourceY SSize, updateImage *Image, updateX, updateY SSize, exception *ExceptionInfo) bool
var PixelIterateTripleNew func(callBack PixelIteratorTripleNewCallback, options *PixelIteratorOptions, description string, mutableData, immutableData *Void, columns, rows Size, source1Image, source2Image *Image, sourceX, sourceY SSize, newImage *Image, newX, newY SSize, exception *ExceptionInfo) bool

// Plasma

var PlasmaImage func(i *Image, segment *SegmentInfo, attenuate, depth Size) bool

func (i *Image) Plasma(segment *SegmentInfo, attenuate, depth Size) bool {
	return PlasmaImage(i, segment, attenuate, depth)
}

// Profile

var AllocateImageProfileIterator func(i *Image) *ImageProfileIterator

func (i *Image) AllocateImageProfileIterator() *ImageProfileIterator {
	return AllocateImageProfileIterator(i)
}

var AppendImageProfile func(i *Image, name string, profileChunk *byte, chunkLength uint32) bool

func (i *Image) AppendProfile(name string, profileChunk *byte, chunkLength uint32) bool {
	return AppendImageProfile(i, name, profileChunk, chunkLength)
}

var DeallocateImageProfileIterator func(profileIterator *ImageProfileIterator)

var DeleteImageProfile func(i *Image, name string) bool

func (i *Image) DeleteProfile(name string) bool { return DeleteImageProfile(i, name) }

var GetImageProfile func(i *Image, name string, length *Size) []byte

func (i *Image) Profile(name string, length *Size) []byte { return GetImageProfile(i, name, length) }

var GetNextImageProfile func(i *ImageProfileIterator, name **Char, profile **byte, length *Size) string

var ProfileImage func(i *Image, name string, profile []byte, length Size, clone bool) bool

func (i *Image) ProfileImage(name string, profile []byte, length Size, clone bool) bool {
	return ProfileImage(i, name, profile, length, clone)
}

var SetImageProfile func(i *Image, name string, profile []byte, length Size) bool

func (i *Image) SetProfile(name string, profile []byte, length Size) bool {
	return SetImageProfile(i, name, profile, length)
}

// Quantize

var CloneQuantizeInfo func(q *QuantizeInfo) *QuantizeInfo

func (q *QuantizeInfo) Clone() *QuantizeInfo { return CloneQuantizeInfo(q) }

var CompressImageColormap func(i *Image)

func (i *Image) CompressColormap() { CompressImageColormap(i) }

var DestroyQuantizeInfo func(q *QuantizeInfo)

func (q *QuantizeInfo) Destroy() { DestroyQuantizeInfo(q) }

var GetImageQuantizeError func(i *Image) bool

func (i *Image) QuantizeError() bool { return GetImageQuantizeError(i) }

var GetQuantizeInfo func(q *QuantizeInfo)

func (q *QuantizeInfo) Get() { GetQuantizeInfo(q) }

var GrayscalePseudoClassImage func(*Image)

var MapImage func(i *Image, mapImage *Image, dither bool) bool

func (i *Image) Map(mapImage *Image, dither bool) bool { return MapImage(i, mapImage, dither) }

var MapImages func(images, mapImage *Image, dither bool) bool

var OrderedDitherImage func(i *Image) bool

func (i *Image) OrderedDither() bool { return OrderedDitherImage(i) }

var QuantizeImage func(q *QuantizeInfo, image *Image) bool

func (q *QuantizeInfo) QuantizeImage(image *Image) bool { return QuantizeImage(q, image) }

var QuantizeImages func(q *QuantizeInfo, images *Image) bool

func (q *QuantizeInfo) QuantizeImages(images *Image) bool { return QuantizeImages(q, images) }

// Registry

var DeleteMagickRegistry func(id Size) bool
var GetImageFromMagickRegistry func(name string, id *Size, exception *ExceptionInfo) *Image
var GetMagickRegistry func(id Size, type_ *RegistryType, length *Size, exception *ExceptionInfo) *Void
var SetMagickRegistry func(type_ RegistryType, blob *Void, length Size, exception *ExceptionInfo) SSize

// Resize

var MagnifyImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Magnify(exception *ExceptionInfo) *Image { return MagnifyImage(i, exception) }

var MinifyImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Minify(exception *ExceptionInfo) *Image { return MinifyImage(i, exception) }

var ResizeImage func(i *Image, columns, rows Size, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image

func (i *Image) Resize(columns, rows Size, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image {
	return ResizeImage(i, columns, rows, filter, blur, exception)
}

var SampleImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image

func (i *Image) Sample(columns, rows Size, exception *ExceptionInfo) *Image {
	return SampleImage(i, columns, rows, exception)
}

var ScaleImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image

func (i *Image) Scale(columns, rows Size, exception *ExceptionInfo) *Image {
	return ScaleImage(i, columns, rows, exception)
}

var ThumbnailImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image

func (i *Image) Thumbnail(columns, rows Size, exception *ExceptionInfo) *Image {
	return ThumbnailImage(i, columns, rows, exception)
}

// Resource

var AcquireMagickResource func(r ResourceType, size MagickSizeType) bool

func (r ResourceType) AcquireResource(size MagickSizeType) bool {
	return AcquireMagickResource(r, size)
}

var GetMagickResource func(r ResourceType) MagickSizeType

func (r ResourceType) Resource() MagickSizeType { return GetMagickResource(r) }

var GetMagickResourceLimit func(r ResourceType) MagickSizeType

func (r ResourceType) Limit() MagickSizeType { return GetMagickResourceLimit(r) }

var LiberateMagickResource func(r ResourceType, size int64)

var ListMagickResourceInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagickResourceInfo(exception *ExceptionInfo) bool {
	return ListMagickResourceInfo(f, exception)
}

//NOTE(t): Doc says unsigned long, code int64
var SetMagickResourceLimit func(r ResourceType, limit MagickSizeType) bool

func (r ResourceType) SetLimit(limit MagickSizeType) bool { return SetMagickResourceLimit(r, limit) }

// Segment

var SegmentImage func(i *Image, colorspace ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool

func (i *Image) Segment(colorspace ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool {
	return SegmentImage(i, colorspace, verbose, clusterThreshold, smoothThreshold)
}

// Shear

var AffineTransformImage func(i *Image, affineMatrix *AffineMatrix, exception *ExceptionInfo) *Image

func (i *Image) AffineTransform(affineMatrix *AffineMatrix, exception *ExceptionInfo) *Image {
	return AffineTransformImage(i, affineMatrix, exception)
}

var AutoOrientImage func(i *Image, o OrientationType, e *ExceptionInfo) *Image

func (i *Image) AutoOrient(o OrientationType, e *ExceptionInfo) *Image {
	return AutoOrientImage(i, o, e)
}

var RotateImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image

func (i *Image) Rotate(degrees float64, exception *ExceptionInfo) *Image {
	return RotateImage(i, degrees, exception)
}

var ShearImage func(i *Image, xShear, yShear float64, exception *ExceptionInfo) *Image

func (i *Image) Shear(xShear, yShear float64, exception *ExceptionInfo) *Image {
	return ShearImage(i, xShear, yShear, exception)
}

// Signature

var FinalizeSignature func(s *SignatureInfo)

func (s *SignatureInfo) Finalize() { FinalizeSignature(s) }

var GetSignatureInfo func(s *SignatureInfo)

func (s *SignatureInfo) Get() { GetSignatureInfo(s) }

var SignatureImage func(i *Image) bool

func (i *Image) Signature() bool { return SignatureImage(i) }

var TransformSignature func(s *SignatureInfo)

func (s *SignatureInfo) Transform() { TransformSignature(s) }

var UpdateSignature func(s *SignatureInfo, message *byte, length uint32)

func (s *SignatureInfo) Update(message *byte, length uint32) { UpdateSignature(s, message, length) }

// Statistics

var GetImageStatistics func(i *Image, statistics *ImageStatistics, exception *ExceptionInfo) bool

func (i *Image) Statistics(statistics *ImageStatistics, exception *ExceptionInfo) bool {
	return GetImageStatistics(i, statistics, exception)
}

// Texture

var ConstituteTextureImage func(columns, rows Size, texture *Image, exception *ExceptionInfo) *Image

var TextureImage func(i, texture *Image) bool

func (i *Image) Texture(texture *Image) bool { return TextureImage(i, texture) }

// Transform

var ChopImage func(i *Image, chopInfo *RectangleInfo, exception *ExceptionInfo) *Image

func (i *Image) Chop(chopInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return ChopImage(i, chopInfo, exception)
}

var CoalesceImages func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) CoalesceImages(exception *ExceptionInfo) *Image { return CoalesceImages(i, exception) }

var CropImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image

func (i *Image) Crop(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
	return CropImage(i, geometry, exception)
}

var DeconstructImages func(images *Image, exception *ExceptionInfo) *Image

var ExtentImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image

func (i *Image) Extent(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
	return ExtentImage(i, geometry, exception)
}

var FlattenImages func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) FlattenImages(exception *ExceptionInfo) *Image { return FlattenImages(i, exception) }

var FlipImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Flip(exception *ExceptionInfo) *Image { return FlipImage(i, exception) }

var FlopImage func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) Flop(exception *ExceptionInfo) *Image { return FlopImage(i, exception) }

var MosaicImages func(i *Image, exception *ExceptionInfo) *Image

func (i *Image) MosaicImages(exception *ExceptionInfo) *Image { return MosaicImages(i, exception) }

var RollImage func(i *Image, xOffset, yOffset SSize, exception *ExceptionInfo) *Image

func (i *Image) Roll(xOffset, yOffset SSize, exception *ExceptionInfo) *Image {
	return RollImage(i, xOffset, yOffset, exception)
}

var ShaveImage func(i *Image, shaveInfo *RectangleInfo, exception *ExceptionInfo) *Image

func (i *Image) Shave(shaveInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return ShaveImage(i, shaveInfo, exception)
}

var TransformImage func(image **Image, cropGeometry, imageGeometry string) bool

//============== Undocumented ==========================

type (
	fp16bits [2]byte
	fp24bits [3]byte

	MagickMapObjectClone       func(object *Void, objectSize Size) *Void
	MagickMapObjectDeallocator func(object *Void)
	MagickTextTranslate        func(dst, src string, size Size) Size
	WordStreamReadFunc         func(readFuncState *Void) Size
	WordStreamWriteFunc        func(writeFuncState *Void, value Size) Size
	WriteByteHook              func(*Image, uint8, *Void) uint
)

type ThreadViewDataSet struct {
	ViewData   **Void
	Destructor MagickFreeFunc
	NViews     uint
}

type MagickRandomKernel struct {
	z, w uint32
}

type FileIOMode Enum

const (
	BinaryFileIOMode FileIOMode = iota
	TextFileIOMode
)

type BlobMode Enum

const (
	UndefinedBlobMode BlobMode = iota
	ReadBlobMode
	ReadBinaryBlobMode
	WriteBlobMode
	WriteBinaryBlobMode
)

type CompositeOptions struct {
	PercentBrightness, Amount, Threshold float64
}

type ColorInfo struct {
	Path           *VString
	Name           *VString
	Compliance     ComplianceType
	Color          PixelPacket
	Stealth        uint
	Signature      Size
	Previous, Next *ColorInfo
}
type ComplianceType Enum

const (
	SVGCompliance ComplianceType = 1 << iota
	X11Compliance
	XPMCompliance
	UndefinedCompliance ComplianceType = 0x0000
	NoCompliance                       = UndefinedCompliance
	AllCompliance       ComplianceType = 0xffff
)

type MagickPixelPacket struct {
	Colorspace                ColorspaceType
	Matte                     uint
	Red, Green, Blue, Opacity Quantum
	Index                     IndexPacket
}

type DelegateInfo struct {
	Path, Decode, Encode, Commands *VString
	Mode                           int
	Stealth                        MagickBooleanType
	Signature                      Size
	Previous, Next                 *DelegateInfo
}

type ImageCharacteristics struct {
	Cmyk,
	Grayscale,
	Monochrome,
	Opaque,
	Palette MagickBooleanType
}

type ModuleInfo struct {
	Path, Magick, Name *VString
	Stealth            uint
	Signature          Size
	Previous, Next     *ModuleInfo
}

type PathType Enum

const (
	RootPath PathType = iota
	HeadPath
	TailPath
	BasePath
	ExtensionPath
	MagickPath
	SubImagePath
	FullPath
)

type TypeInfo struct {
	Path, Name, Description, Family            *VString
	Style                                      StyleType
	Stretch                                    StretchType
	Weight                                     Size
	Encoding, Foundry, Format, Metrics, Glyphs *VString
	Stealth                                    uint
	Signature                                  Size
	Previous, Next                             *TypeInfo
}

type LogEventType Enum

const (
	ConfigureEventMask LogEventType = 1 << iota
	AnnotateEventMask
	RenderEventMask
	TransformEventMask
	LocaleEventMask
	CoderEventMask
	X11EventMask
	CacheEventMask
	BlobEventMask
	DeprecateEventMask
	UserEventMask
	ResourceEventMask
	TemporaryFileEventMask
	_
	OptionEventMask
	InformationEventMask
	WarningEventMask
	ErrorEventMask
	FatalErrorEventMask
	UndefinedEventMask LogEventType = 0
	NoEventsMask                    = UndefinedEventMask
	AllEventsMask      LogEventType = 0x7FFFFFFF

	ExceptionEventMask = WarningEventMask | ErrorEventMask | FatalErrorEventMask
)

type MagickMapIterator struct {
	Map       MagickMap
	Member    *MagickMapObject
	Position  MagickMapIteratorPosition
	Signature Size
}

type MagickMapIteratorPosition Enum

const (
	InListPosition MagickMapIteratorPosition = iota
	FrontPosition
	BackPosition
)

type MagickMapObject struct {
	Key                *VString
	Object             *Void
	ObjectSize         Size
	CloneFunction      MagickMapObjectClone
	DeallocateFunction MagickMapObjectDeallocator
	ReferenceCount     SSize
	Previous, Next     *MagickMapObject
	Signature          Size
}

type BitStreamReadHandle struct {
	Bytes         *byte
	BitsRemaining uint
}

type BitStreamWriteHandle struct {
	Bytes         *byte
	BitsRemaining uint
}

type WordStreamReadHandle struct {
	Word          uint32
	BitsRemaining uint
	ReadFunc      WordStreamReadFunc
	ReadFuncState *Void
}

type WordStreamWriteHandle struct {
	Word           uint32
	BitsRemaining  uint
	WriteFunc      WordStreamWriteFunc
	WriteFuncState *Void
}

type MapMode Enum

const (
	ReadMode MapMode = iota
	WriteMode
	IOMode
)

type GhostscriptVectors struct {
	Exit           func(i *GSMainInstance) int
	DeleteInstance func(i *GSMainInstance)
	InitWithArgs   func(i *GSMainInstance, argc int, argv []string) int
	NewInstance    func(i **GSMainInstance, callerHandle *Void) int
	RunString      func(i *GSMainInstance, str string, userErrors int, exitCode *int) int
}

type TokenInfo struct {
	State  int
	Flag   uint
	Offset SSize
	Quote  Char
}

var AccessDefaultCacheView func(i *Image) *ViewInfo

func (i *Image) AccessDefaultCacheView() *ViewInfo { return AccessDefaultCacheView(i) }

var AccessThreadViewData func(t *ThreadViewDataSet) *Void

func (t *ThreadViewDataSet) Access() *Void { return AccessThreadViewData(t) }

var AccessThreadViewDataById func(t *ThreadViewDataSet, index uint) *Void

func (t *ThreadViewDataSet) AccessById(index uint) *Void {
	return AccessThreadViewDataById(t, index)
}

var AcquireCacheView func(i *Image) *ViewInfo

func (i *Image) AcquireCacheView() *ViewInfo { return AcquireCacheView(i) }

var AcquireMagickRandomKernel func() *MagickRandomKernel
var AcquireMemory func(size Size) *Void
var AcquireSemaphoreInfo func(s **SemaphoreInfo)
var AcquireString func(source string) string
var AcquireTemporaryFileDescriptor func(filename string) int
var AcquireTemporaryFileName func(filename string) bool
var AcquireTemporaryFileStream func(filename string, mode FileIOMode) *FILE
var AllocateSemaphoreInfo func() *SemaphoreInfo

var AllocateString func(source string) string
var AllocateThreadViewDataArray func(i *Image, exception *ExceptionInfo, count, size uint32) *ThreadViewDataSet

func (i *Image) AllocateThreadViewDataArray(exception *ExceptionInfo, count, size uint32) *ThreadViewDataSet {
	return AllocateThreadViewDataArray(i, exception, count, size)
}

var AllocateThreadViewDataSet func(destructor MagickFreeFunc, image *Image, exception *ExceptionInfo) *ThreadViewDataSet
var AllocateThreadViewSet func(i *Image, exception *ExceptionInfo) *ThreadViewSet

func (i *Image) AllocateThreadViewSet(exception *ExceptionInfo) *ThreadViewSet {
	return AllocateThreadViewSet(i, exception)
}

var AnimateImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) AnimateImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return AnimateImageCommand(i, argc, argv, metadata, exception)
}

var AppendImageFormat func(format, filename string)

var Ascii85Encode func(i *Image, code byte)

func (i *Image) Ascii85Encode(code byte) { Ascii85Encode(i, code) }

var Ascii85Flush func(i *Image)

func (i *Image) Ascii85Flush() { Ascii85Flush(i) }

var Ascii85Initialize func(i *Image)

func (i *Image) Ascii85Initialize() { Ascii85Initialize(i) }

var Ascii85WriteByteHook func(i *Image, code byte, info *Void) uint

func (i *Image) Ascii85WriteByteHook(code byte, info *Void) uint {
	return Ascii85WriteByteHook(i, code, info)
}

var AssignThreadViewData func(t *ThreadViewDataSet, index uint, data *Void)

func (t *ThreadViewDataSet) Assign(index uint, data *Void) { AssignThreadViewData(t, index, data) }

var Base64Decode func(source string, length *Size) *byte
var Base64Encode func(blob *byte, blobLength Size, encodeLength *Size) string
var BenchmarkImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) BenchmarkImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return BenchmarkImageCommand(i, argc, argv, metadata, exception)
}

var BlobModeToString func(blobMode BlobMode) string

var BlobWriteByteHook func(i *Image, code byte, info *Void) uint

func (i *Image) BlobWriteByteHook(code byte, info *Void) uint {
	return BlobWriteByteHook(i, code, info)
}

var ChannelTypeToString func(channel ChannelType) string
var ClassTypeToString func(classType ClassType) string
var ClipImage func(i *Image) bool

func (i *Image) Clip() bool { return ClipImage(i) }

var CloneDrawInfo func(i *ImageInfo, drawInfo *DrawInfo) *DrawInfo

func (i *ImageInfo) CloneDrawInfo(drawInfo *DrawInfo) *DrawInfo { return CloneDrawInfo(i, drawInfo) }

var CloneMemory func(destination *Void, source *Void, size Size) *Void
var CloneString func(destination **Char, source string) string
var CloseBlob func(i *Image) bool

func (i *Image) CloseBlob() bool { return CloseBlob(i) }

var CloseCacheView func(c *ViewInfo) *ViewInfo

func (c *ViewInfo) Close() *ViewInfo { return CloseCacheView(c) }

var ColorspaceTypeToString func(colorspace ColorspaceType) string
var CompareImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) CompareImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return CompareImageCommand(i, argc, argv, metadata, exception)
}

var CompositeImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) CompositeImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return CompositeImageCommand(i, argc, argv, metadata, exception)
}

var CompositeImageRegion func(compose CompositeOperator, options *CompositeOptions, columns, rows Size, updateImage *Image, updateX, updateY SSize, canvasImage *Image, canvasX, canvasY SSize, exception *ExceptionInfo) bool
var CompositeOperatorToString func(compositeOp CompositeOperator) string
var CompressionTypeToString func(compressionType CompressionType) string
var ConcatenateString func(destination **Char, source string) bool
var ConfirmAccessModeToString func(accessMode ConfirmAccessMode) string
var ConjureImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) ConjureImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return ConjureImageCommand(i, argc, argv, metadata, exception)
}

var ContinueTimer func(t *TimerInfo) bool

func (t *TimerInfo) Continue() bool { return ContinueTimer(t) }

var Contrast func(int, *Quantum, *Quantum, *Quantum)
var ConvertImageCommand func(*ImageInfo, int, []string, []string, *ExceptionInfo) bool

func (i *ImageInfo) ConvertImageCommand(a2 int, a3, a4 []string, a5 *ExceptionInfo) bool {
	return ConvertImageCommand(i, a2, a3, a4, a5)
}

var CropImageToHBITMAP func(i *Image, r *RectangleInfo, e *ExceptionInfo) *Void

func (i *Image) CropToHBITMAP(r *RectangleInfo, e *ExceptionInfo) *Void {
	return CropImageToHBITMAP(i, r, e)
}

var DefineClientName func(string)
var DefineClientPathAndName func(string)
var DestroyColorInfo func()
var DestroyConstitute func()
var DestroyDelegateInfo func()
var DestroyDrawInfo func(d *DrawInfo) *DrawInfo

func (d *DrawInfo) Destroy() *DrawInfo { return DestroyDrawInfo(d) }

var DestroyImagePixels func(i *Image)

func (i *Image) DestroyPixels() { DestroyImagePixels(i) }

var DestroyLogInfo func()
var DestroyMagickModules func()
var DestroyMagickResources func()
var DestroyModuleInfo func()
var DestroySemaphore func()
var DestroySemaphoreInfo func(s *SemaphoreInfo) *SemaphoreInfo

func (s *SemaphoreInfo) Destroy() *SemaphoreInfo { return DestroySemaphoreInfo(s) }

var DestroyTemporaryFiles func()
var DestroyThreadViewDataSet func(t *ThreadViewDataSet)

func (t *ThreadViewDataSet) Destroy() { DestroyThreadViewDataSet(t) }

var DestroyThreadViewSet func(viewSet *ThreadViewSet)
var DestroyTypeInfo func()
var DisplayImageCommand func(*ImageInfo, int, []string, []string, *ExceptionInfo) bool

func (i *ImageInfo) DisplayImageCommand(a2 int, a3, a4 []string, a5 *ExceptionInfo) bool {
	return DisplayImageCommand(i, a2, a3, a4, a5)
}

var DrawAffineImage func(i *Image, source *Image, affine *AffineMatrix) bool

func (i *Image) DrawAffine(source *Image, affine *AffineMatrix) bool {
	return DrawAffineImage(i, source, affine)
}

var DrawClipPath func(i *Image, drawInfo *DrawInfo, name string) bool

func (i *Image) DrawClipPath(drawInfo *DrawInfo, name string) bool {
	return DrawClipPath(i, drawInfo, name)
}

var DrawImage func(i *Image, drawInfo *DrawInfo) bool

func (i *Image) Draw(drawInfo *DrawInfo) bool { return DrawImage(i, drawInfo) }

var DrawPatternPath func(i *Image, drawInfo *DrawInfo, name string, pattern **Image) bool

func (i *Image) DrawPatternPath(drawInfo *DrawInfo, name string, pattern **Image) bool {
	return DrawPatternPath(i, drawInfo, name, pattern)
}

var EOFBlob func(i *Image) int

func (i *Image) EOFBlob() int { return EOFBlob(i) }

var EndianTypeToString func(endianType EndianType) string
var EscapeString func(source string, escape int8) string
var ExecuteModuleProcess func(tag string, image **Image, argc int, argv []string) bool
var ExecuteStaticModuleProcess func(string, **Image, int, []string) bool
var Exit func(int) int
var ExpandAffine func(affine *AffineMatrix) float64
var ExpandFilename func(path string)
var ExpandFilenames func(argc *int, argv *[]string) bool
var FormatSize func(size int64, format string)
var FormatString func(str, format string, va ...VArg)
var FormatStringList func(str, format string, operands VAList)
var FuzzyColorMatch func(p, q *PixelPacket, fuzz float64) uint
var GMCommand func(argc int, argv []string) int
var GenerateDifferentialNoise func(quantumPixel Quantum, noiseType NoiseType, kernel *MagickRandomKernel) float64
var GenerateNoise func(Quantum, NoiseType) Quantum
var GetBlobIsOpen func(i *Image) bool

func (i *Image) BlobIsOpen() bool { return GetBlobIsOpen(i) }

var GetBlobSize func(i *Image) MagickSizeType

func (i *Image) BlobSize() MagickSizeType { return GetBlobSize(i) }

var GetCacheView func(c *ViewInfo, x, y SSize, columns, rows Size) *PixelPacket

func (c *ViewInfo) Get(x, y SSize, columns, rows Size) *PixelPacket {
	return GetCacheView(c, x, y, columns, rows)
}

var GetClientFilename func() string
var GetClientName func() string
var GetClientPath func() string
var GetColorInfo func(name string, exception *ExceptionInfo) *ColorInfo
var GetColorInfoArray func(exception *ExceptionInfo) []*ColorInfo
var GetColorList func(pattern string, numberColors *Size, exception *ExceptionInfo) []string
var GetColorTuple func(m *MagickPixelPacket, hex bool, tuple string)

func (m *MagickPixelPacket) ColorTuple(hex bool, tuple string) { GetColorTuple(m, hex, tuple) }

var GetDelegateCommand func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) string

func (i *ImageInfo) DelegateCommand(image *Image, decode, encode string, exception *ExceptionInfo) string {
	return GetDelegateCommand(i, image, decode, encode, exception)
}

var GetDelegateInfo func(decode, encode string, exception *ExceptionInfo) *DelegateInfo
var GetDrawInfo func(i *ImageInfo, drawInfo *DrawInfo)

func (i *ImageInfo) DrawInfo(drawInfo *DrawInfo) { GetDrawInfo(i, drawInfo) }

var GetElapsedTime func(t *TimerInfo) float64

func (t *TimerInfo) ElapsedTime() float64 { return GetElapsedTime(t) }

var GetExecutionPath func(path string, extent uint32) bool
var GetExecutionPathUsingName func(string) bool
var GetGeometry func(geometry string, x, y *SSize, width, height *Size) MagickStatusType
var GetImageBoundingBox func(i *Image, exception *ExceptionInfo) RectangleInfo

func (i *Image) BoundingBox(exception *ExceptionInfo) RectangleInfo {
	return GetImageBoundingBox(i, exception)
}

var GetImageCharacteristics func(i *Image, characteristics *ImageCharacteristics, optimize bool, exception *ExceptionInfo) bool

func (i *Image) Characteristics(characteristics *ImageCharacteristics, optimize bool, exception *ExceptionInfo) bool {
	return GetImageCharacteristics(i, characteristics, optimize, exception)
}

var GetImageDepth func(i *Image, exception *ExceptionInfo) SSize

func (i *Image) Depth(exception *ExceptionInfo) SSize {
	return GetImageDepth(i, exception)
}

var GetImageType func(i *Image, exception *ExceptionInfo) ImageType

func (i *Image) Type(exception *ExceptionInfo) ImageType { return GetImageType(i, exception) }

var GetLocaleMessage func(tag string) string
var GetLocaleMessageFromID func(int) string
var GetMagickCopyright func() string
var GetImageInfoAttribute func(i *ImageInfo, image *Image, key string) *ImageAttribute

func (i *ImageInfo) ImageInfoAttribute(image *Image, key string) *ImageAttribute {
	return GetImageInfoAttribute(i, image, key)
}

var GetMagickDimension func(str string, width, height, xoff, yoff *float64) int
var GetMagickFileFormat func(header *byte, headerLength uint32, format string, formatLength uint32, exception *ExceptionInfo) bool
var GetMagickGeometry func(geometry string, x, y *SSize, width, height *Size) uint
var GetMagickVersion func(version *SSize) string
var GetMagickWebSite func() string
var GetModuleInfo func(tag string, exception *ExceptionInfo) *ModuleInfo
var GetOptimalKernelWidth func(radius, sigma float64) Size
var GetOptimalKernelWidth1D func(radius, sigma float64) Size
var GetOptimalKernelWidth2D func(radius, sigma float64) Size
var GetPageGeometry func(pageGeometry string) string
var GetPathComponent func(path string, type_ PathType, component string)
var GetPixelCacheArea func(i *Image) int64
var GetPostscriptDelegateInfo func(i *ImageInfo, antialias *uint, exception *ExceptionInfo) *DelegateInfo

func (i *ImageInfo) PostscriptDelegateInfo(antialias *uint, exception *ExceptionInfo) *DelegateInfo {
	return GetPostscriptDelegateInfo(i, antialias, exception)
}

var GetThreadViewDataSetAllocatedViews func(t *ThreadViewDataSet) uint

func (t *ThreadViewDataSet) AllocatedViews() uint { return GetThreadViewDataSetAllocatedViews(t) }

var GetTimerInfo func(t *TimerInfo)

func (t *TimerInfo) TimerInfo() { GetTimerInfo(t) }

var GetTimerResolution func() float64
var GetToken func(string, **Char, *Char)
var GetTypeInfo func(name string, exception *ExceptionInfo) *TypeInfo
var GetTypeInfoByFamily func(family string, style StyleType, stretch StretchType, weight Size, exception *ExceptionInfo) *TypeInfo
var GetTypeList func(pattern string, numberFonts *Size, exception *ExceptionInfo) []string
var GetUserTime func(t *TimerInfo) float64

func (t *TimerInfo) UserTime() float64 { return GetUserTime(t) }

var GlobExpression func(expression, pattern string, caseInsensitive bool) bool
var GradientImage func(i *Image, startColor, stopColor *PixelPacket) bool

func (i *Image) Gradient(startColor, stopColor *PixelPacket) bool {
	return GradientImage(i, startColor, stopColor)
}

var HSLTransform func(hue, saturation, lightness float64, red, green, blue *Quantum)
var HWBTransform func(float64, float64, float64, *Quantum, *Quantum, *Quantum)
var HighlightStyleToString func(differenceAlgorithm HighlightStyle) string
var HuffmanDecodeImage func(i *Image) bool

func (i *Image) HuffmanDecode() bool { return HuffmanDecodeImage(i) }

var HuffmanEncode2Image func(i *ImageInfo, image *Image, writeByte WriteByteHook, info *Void) bool

func (i *ImageInfo) HuffmanEncode2(image *Image, writeByte WriteByteHook, info *Void) bool {
	return HuffmanEncode2Image(i, image, writeByte, info)
}

var HuffmanEncodeImage func(i *ImageInfo, image *Image) bool

func (i *ImageInfo) HuffmanEncode(image *Image) bool { return HuffmanEncodeImage(i, image) }

var Hull func(SSize, SSize, Size, Size, *Quantum, *Quantum, int)
var IdentifyImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) IdentifyImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return IdentifyImageCommand(i, argc, argv, metadata, exception)
}

var IdentityAffine func(affine *AffineMatrix)
var ImageToHBITMAP func(i *Image) **Void

func (i *Image) HBITMAP() **Void { return ImageToHBITMAP(i) }

var ImageToHuffman2DBlob func(i *Image, imageInfo *ImageInfo, length *uint32, exception *ExceptionInfo) *byte

func (i *Image) Huffman2DBlob(imageInfo *ImageInfo, length *uint32, exception *ExceptionInfo) *byte {
	return ImageToHuffman2DBlob(i, imageInfo, length, exception)
}

var ImageToJPEGBlob func(i *Image, imageInfo *ImageInfo, length *uint32, exception *ExceptionInfo) *byte

func (i *Image) JPEGBlob(imageInfo *ImageInfo, length *uint32, exception *ExceptionInfo) *byte {
	return ImageToJPEGBlob(i, imageInfo, length, exception)
}

var ImageTypeToString func(imageType ImageType) string
var ImportPixelAreaOptionsInit func(options *ImportPixelAreaOptions)
var InitializeMagicInfo func() bool
var InitializeMagickClientPathAndName func(string)
var InitializeMagickModules func()
var InitializeMagickRandomKernel func(kernel *MagickRandomKernel)
var InitializeMagickResources func()
var InitializeMagickSignalHandlers func()
var InitializeSemaphore func()
var InterlaceTypeToString func(interlaceType InterlaceType) string
var InterpolateColor func(i *Image, xOffset, yOffset float64, exception *ExceptionInfo) PixelPacket

func (i *Image) InterpolateColor(xOffset, yOffset float64, exception *ExceptionInfo) PixelPacket {
	return InterpolateColor(i, xOffset, yOffset, exception)
}

var InterpolateViewColor func(v *ViewInfo, color *PixelPacket, xOffset, yOffset float64, exception *ExceptionInfo)

func (v *ViewInfo) InterpolateColor(color *PixelPacket, xOffset, yOffset float64, exception *ExceptionInfo) {
	InterpolateViewColor(v, color, xOffset, yOffset, exception)
}

var InvokeDelegate func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) bool

func (i *ImageInfo) InvokeDelegate(image *Image, decode, encode string, exception *ExceptionInfo) bool {
	return InvokeDelegate(i, image, decode, encode, exception)
}

var InvokePostscriptDelegate func(verbose uint, command string, exception *ExceptionInfo) bool
var IsAccessible func(path string) bool
var IsAccessibleAndNotEmpty func(string) bool
var IsAccessibleNoLogging func(string) bool
var IsEventLogging func() bool
var IsGeometry func(geometry string) bool
var IsGlob func(path string) bool
var IsGrayImage func(i *Image, exception *ExceptionInfo) bool

func (i *Image) Gray(exception *ExceptionInfo) bool { return IsGrayImage(i, exception) }

var IsMonochromeImage func(i *Image, exception *ExceptionInfo) bool

func (i *Image) Monochrome(exception *ExceptionInfo) bool {
	return IsMonochromeImage(i, exception)
}

var IsOpaqueImage func(i *Image, exception *ExceptionInfo) bool

func (i *Image) IsOpaque(exception *ExceptionInfo) bool { return IsOpaqueImage(i, exception) }

var IsSubimage func(geometry string, pedantic uint) uint
var IsWindows95 func() bool
var IsWriteable func(string) bool
var LZWEncode2Image func(i *Image, length uint32, pixels *byte, writeByte WriteByteHook, info *Void) bool

func (i *Image) LZWEncode2(length uint32, pixels *byte, writeByte WriteByteHook, info *Void) bool {
	return LZWEncode2Image(i, length, pixels, writeByte, info)
}

var LZWEncodeImage func(i *Image, length uint32, pixels *byte) bool

func (i *Image) LZWEncode(length uint32, pixels *byte) bool {
	return LZWEncodeImage(i, length, pixels)
}

var LiberateMemory func(memory **Void)
var LiberateSemaphoreInfo func(s **SemaphoreInfo)
var LiberateTemporaryFile func(filename string) bool
var ListColorInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ColorInfo(exception *ExceptionInfo) bool { return ListColorInfo(f, exception) }

var ListDelegateInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) DelegateInfo(exception *ExceptionInfo) bool { return ListDelegateInfo(f, exception) }

var ListFiles func(directory, pattern string, numberEntries *Size) []string
var ListMagicInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagicInfo(exception *ExceptionInfo) bool { return ListMagicInfo(f, exception) }

var ListModuleInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ModuleInfo(exception *ExceptionInfo) bool { return ListModuleInfo(f, exception) }

var ListTypeInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) TypeInfo(exception *ExceptionInfo) bool { return ListTypeInfo(f, exception) }

var LocaleCompare func(p, q string) SSize
var LocaleLower func(str string)
var LocaleNCompare func(p, q string, length uint32) SSize
var LocaleUpper func(str string)
var LockSemaphoreInfo func(s *SemaphoreInfo) bool

func (s *SemaphoreInfo) Lock() bool { return LockSemaphoreInfo(s) }

var LogMagickEvent func(type_ LogEventType, module, function string, line Size, format string, va ...VArg) bool
var LogMagickEventList func(type_ LogEventType, module, function string, line Size, format string, operands VAList) bool
var MSBOrderLong func(buffer *byte, length uint32)
var MSBOrderShort func(p *byte, length uint32)
var MagickArraySize func(count, size uint32) uint32
var MagickBitStreamInitializeRead func(bitStream *BitStreamReadHandle, bytes *byte)
var MagickBitStreamInitializeWrite func(bitStream *BitStreamWriteHandle, bytes *byte)
var MagickBitStreamMSBRead func(bitStream *BitStreamReadHandle, requestedBits uint) uint
var MagickBitStreamMSBWrite func(bitStream *BitStreamWriteHandle, requestedBits, quantum uint)
var MagickCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) MagickCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return MagickCommand(i, argc, argv, metadata, exception)
}

var MagickCompositeImageUnderColor func(i *Image, undercolor *PixelPacket, exception *ExceptionInfo) bool

func (i *Image) CompositeUnderColor(undercolor *PixelPacket, exception *ExceptionInfo) bool {
	return MagickCompositeImageUnderColor(i, undercolor, exception)
}

var MagickConstrainColormapIndex func(i *Image, index uint) uint

func (i *Image) ConstrainColormapIndex(index uint) uint {
	return MagickConstrainColormapIndex(i, index)
}

var MagickCreateDirectoryPath func(dir string, exception *ExceptionInfo) bool
var MagickFindRawImageMinMax func(i *Image, endian EndianType, width, height Size, type_ StorageType, scanlineOctets uint, scanlineBuffer *Void, min, max *float64) bool

func (i *Image) FindRawMinMax(endian EndianType, width, height Size, type_ StorageType, scanlineOctets uint, scanlineBuffer *Void, min, max *float64) bool {
	return MagickFindRawImageMinMax(i, endian, width, height, type_, scanlineOctets, scanlineBuffer, min, max)
}

var MagickGetBitRevTable func(int) *byte
var MagickGetMMUPageSize func() SSize
var MagickGetQuantumSamplesPerPixel func(quantumType QuantumType) uint
var MagickMapAccessEntry func(map_ *MagickMap, key string, objectSize *uint32) *Void
var MagickMapAddEntry func(map_ *MagickMap, key string, object *Void, objectSize uint32, exception *ExceptionInfo) uint
var MagickMapAllocateIterator func(map_ *MagickMap) *MagickMapIterator
var MagickMapAllocateMap func(clone MagickMapObjectClone, deallocate MagickMapObjectDeallocator) *MagickMap
var MagickMapClearMap func(map_ *MagickMap)
var MagickMapCloneMap func(map_ *MagickMap, exception *ExceptionInfo) MagickMap
var MagickMapCopyBlob func(blob *Void, size uint32) *Void
var MagickMapCopyString func(str *Void, size uint32) *Void
var MagickMapDeallocateMap func(map_ *MagickMap)
var MagickMapDeallocateBlob func(blob *Void)
var MagickMapDeallocateIterator func(iterator *MagickMapIterator)
var MagickMapDeallocateString func(str *Void)
var MagickMapDereferenceIterator func(iterator *MagickMapIterator, objectSize *uint32) *Void
var MagickMapIterateNext func(iterator *MagickMapIterator, key []string) uint
var MagickMapIteratePrevious func(iterator *MagickMapIterator, key []string) uint
var MagickMapIterateToBack func(iterator *MagickMapIterator)
var MagickMapIterateToFront func(iterator *MagickMapIterator)
var MagickMapRemoveEntry func(map_ *MagickMap, key string) uint
var MagickRandNewSeed func() uint
var MagickRandReentrant func(seed *uint) int
var MagickRandomInteger func() uint32
var MagickRandomReal func() float64
var MagickReverseBits func(cp *byte, n uint32)
var MagickSceneFileName func(filename, filenameTemplate, sceneTemplate string, force bool, scene Size) bool
var MagickSizeStrToInt64 func(str string, kilo uint) int64
var MagickSpawnVP func(verbose uint, file string, argv []string) int
var MagickStrlCat func(dst string, src string, size uint32) uint32
var MagickStrlCpy func(dst string, src string, size uint32) uint32
var MagickStrlCpyTrunc func(dst string, src string, size uint32) uint32
var MagickSwabArrayOfDouble func(dp *float64, n uint32)
var MagickSwabArrayOfFloat func(fp *float32, n uint32)
var MagickSwabArrayOfUInt16 func(wp *uint16, n uint32)
var MagickSwabArrayOfUInt32 func(lp *uint32, n uint32)
var MagickSwabDouble func(dp *float64)
var MagickSwabFloat func(fp *float32)
var MagickSwabUInt16 func(wp *uint16)
var MagickSwabUInt32 func(lp *uint32)
var MagickTsdGetSpecific func(key MagickTsdKey) *Void
var MagickTsdKeyCreate func(key *MagickTsdKey) bool
var MagickTsdKeyCreate2 func(key *MagickTsdKey, destructor MagickFreeFunc) bool
var MagickTsdKeyDelete func(key MagickTsdKey) bool
var MagickTsdSetSpecific func(key MagickTsdKey, value *Void) bool
var MagickWordStreamInitializeRead func(wordStream *WordStreamReadHandle, readFunc WordStreamReadFunc, readFuncState *Void)
var MagickWordStreamInitializeWrite func(wordStream *WordStreamWriteHandle, writeFunc WordStreamWriteFunc, writeFuncState *Void)
var MagickWordStreamLSBRead func(wordStream *WordStreamReadHandle, requestedBits uint) uint
var MagickWordStreamLSBWrite func(wordStream *WordStreamWriteHandle, requestedBits uint, quantum uint)
var MagickWordStreamLSBWriteFlush func(wordStream *WordStreamWriteHandle)
var MapBlob func(file int, mode MapMode, offset int64, length uint32) *byte
var MapModeToString func(mapMode MapMode) string
var MetricTypeToString func(metric MetricType) string
var Modulate func(float64, float64, float64, *Quantum, *Quantum, *Quantum)
var MogrifyImage func(*ImageInfo, int, []string, **Image, *ExceptionInfo) uint

func (i *ImageInfo) MogrifyImage(a2 int, a3 []string, a4 **Image, a5 *ExceptionInfo) uint {
	return MogrifyImage(i, a2, a3, a4, a5)
}

var MogrifyImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) MogrifyImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return MogrifyImageCommand(i, argc, argv, metadata, exception)
}

var MogrifyImages func(*ImageInfo, int, []string, **Image) uint

func (i *ImageInfo) MogrifyImages(a2 int, a3 []string, a4 **Image) uint {
	return MogrifyImages(i, a2, a3, a4)
}

var MontageImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) MontageImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return MontageImageCommand(i, argc, argv, metadata, exception)
}

var MultilineCensus func(label string) Size
var NTElapsedTime func() float64
var NTErrorHandler func(ExceptionType, string, string)
var NTGetExecutionPath func(path *Char) bool
var NTGetLastError func() string
var NTGetTypeList func() *TypeInfo
var NTGhostscriptDLL func(path *Char, pathLength int) int
var NTGhostscriptDLLVectors func() *GhostscriptVectors
var NTGhostscriptEXE func(string, int) int
var NTGhostscriptFonts func(string, int) int
var NTGhostscriptLoadDLL func() int
var NTGhostscriptUnLoadDLL func() int
var NTIsMagickConflict func(string) bool
var NTKernelAPISupported func(name string) bool
var NTRegistryKeyLookup func(string) string
var NTResourceToBlob func(string) *byte
var NTSystemCommand func(string) int
var NTUserTime func() float64
var NTWarningHandler func(ExceptionType, string, string) ***Void //*** Return?
var NTclosedir func(*DIR) int
var NTdlclose func(handle *Void) int
var NTdlerror func() string
var NTdlexit func() int
var NTdlinit func() int
var NTdlopen func(filename string) *Void
var NTdlsetsearchpath func(string) int
var NTdlsym func(handle *Void, name string) *Void
var NTftruncate func(filedes int, length int64) int
var NTmmap func(address string, length uint32, protection int, access int, file int, offset int64) *Void
var NTmsync func(addr *Void, len uint32, flags int) int
var NTmunmap func(addr *Void, len uint32) int
var NTopendir func(string) *DIR
var NTreaddir func(*DIR) *dirent
var NTseekdir func(*DIR, SSize)
var NTtelldir func(*DIR) SSize
var NextImageProfile func(profileIterator ImageProfileIterator, name *string, profile **byte, length *uint32) bool
var NoiseTypeToString func(noiseType NoiseType) string
var OpenBlob func(i *ImageInfo, image *Image, mode BlobMode, exception *ExceptionInfo) bool

func (i *ImageInfo) OpenBlob(image *Image, mode BlobMode, exception *ExceptionInfo) bool {
	return OpenBlob(i, image, mode, exception)
}

var OpenModule func(module string, exception *ExceptionInfo) bool
var OpenModules func(exception *ExceptionInfo) bool
var OrientationTypeToString func(orientationType OrientationType) string
var PackbitsEncode2Image func(i *Image, length uint32, pixels *byte, writeByte WriteByteHook, info *Void) bool

func (i *Image) PackbitsEncode2(length uint32, pixels *byte, writeByte WriteByteHook, info *Void) bool {
	return PackbitsEncode2Image(i, length, pixels, writeByte, info)
}

var PackbitsEncodeImage func(i *Image, length uint32, pixels *byte) bool

func (i *Image) PackbitsEncode(length uint32, pixels *byte) bool {
	return PackbitsEncodeImage(i, length, pixels)
}

var PersistCache func(i *Image, filename string, attach bool, offset *int64, exception *ExceptionInfo) bool
var PopImagePixels func(i *Image, quantum QuantumType, destination *byte) bool

func (i *Image) PopPixels(quantum QuantumType, destination *byte) bool {
	return PopImagePixels(i, quantum, destination)
}

var PurgeTemporaryFiles func()
var PushImagePixels func(i *Image, quantum QuantumType, source *byte) bool

func (i *Image) PushPixels(quantum QuantumType, source *byte) bool {
	return PushImagePixels(i, quantum, source)
}

var QuantumOperatorImageMultivalue func(i *Image, quantumOperator QuantumOperator, values string) bool

func (i *Image) QuantumOperatorMultivalue(quantumOperator QuantumOperator, values string) bool {
	return QuantumOperatorImageMultivalue(i, quantumOperator, values)
}

var QuantumOperatorToString func(quantumOperator QuantumOperator) string
var QuantumSampleTypeToString func(sampleType QuantumSampleType) string
var QuantumTypeToString func(quantumType QuantumType) string
var QueryColorDatabase func(name string, color *PixelPacket, exception *ExceptionInfo) bool
var QueryColorname func(i *Image, color *PixelPacket, compliance ComplianceType, name string, exception *ExceptionInfo) bool

func (i *Image) QueryColorname(color *PixelPacket, compliance ComplianceType, name string, exception *ExceptionInfo) bool {
	return QueryColorname(i, color, compliance, name, exception)
}

var RGBTransformImage func(i *Image, colorspace ColorspaceType) bool

func (i *Image) RGBTransform(colorspace ColorspaceType) bool {
	return RGBTransformImage(i, colorspace)
}

var ReacquireMemory func(memory **Void, size uint32)
var ReadBlob func(i *Image, length uint32, data *byte) int32

func (i *Image) ReadBlob(length uint32, data *byte) int32 { return ReadBlob(i, length, data) }

var ReadBlobByte func(i *Image) int

func (i *Image) ReadBlobByte() int { return ReadBlobByte(i) }

var ReadBlobLSBDouble func(i *Image) float64

func (i *Image) ReadBlobLSBDouble() float64 { return ReadBlobLSBDouble(i) }

var ReadBlobLSBDoubles func(i *Image, octets uint32, data *float64) uint32

func (i *Image) ReadBlobLSBDoubles(octets uint32, data *float64) uint32 {
	return ReadBlobLSBDoubles(i, octets, data)
}

var ReadBlobLSBFloat func(i *Image) float32

func (i *Image) ReadBlobLSBFloat() float32 { return ReadBlobLSBFloat(i) }

var ReadBlobLSBFloats func(i *Image, octets uint32, data *float32) uint32

func (i *Image) ReadBlobLSBFloats(octets uint32, data *float32) uint32 {
	return ReadBlobLSBFloats(i, octets, data)
}

var ReadBlobLSBLong func(i *Image) uint32

func (i *Image) ReadBlobLSBLong() uint32 { return ReadBlobLSBLong(i) }

var ReadBlobLSBLongs func(i *Image, octets uint32, data *uint32) uint32

func (i *Image) ReadBlobLSBLongs(octets uint32, data *uint32) uint32 {
	return ReadBlobLSBLongs(i, octets, data)
}

var ReadBlobLSBShort func(i *Image) uint16

func (i *Image) ReadBlobLSBShort() uint16 { return ReadBlobLSBShort(i) }

var ReadBlobLSBShorts func(i *Image, octets uint32, data *uint16) uint32

func (i *Image) ReadBlobLSBShorts(octets uint32, data *uint16) uint32 {
	return ReadBlobLSBShorts(i, octets, data)
}

var ReadBlobMSBDouble func(i *Image) float64

func (i *Image) ReadBlobMSBDouble() float64 { return ReadBlobMSBDouble(i) }

var ReadBlobMSBDoubles func(i *Image, octets uint32, data *float64) uint32

func (i *Image) ReadBlobMSBDoubles(octets uint32, data *float64) uint32 {
	return ReadBlobMSBDoubles(i, octets, data)
}

var ReadBlobMSBFloat func(i *Image) float32

func (i *Image) ReadBlobMSBFloat() float32 { return ReadBlobMSBFloat(i) }

var ReadBlobMSBFloats func(i *Image, octets uint32, data *float32) uint32

func (i *Image) ReadBlobMSBFloats(octets uint32, data *float32) uint32 {
	return ReadBlobMSBFloats(i, octets, data)
}

var ReadBlobMSBLong func(i *Image) uint32

func (i *Image) ReadBlobMSBLong() uint32 { return ReadBlobMSBLong(i) }

var ReadBlobMSBLongs func(i *Image, octets uint32, data *uint32) uint32

func (i *Image) ReadBlobMSBLongs(octets uint32, data *uint32) uint32 {
	return ReadBlobMSBLongs(i, octets, data)
}

var ReadBlobMSBShort func(i *Image) uint16

func (i *Image) ReadBlobMSBShort() uint16 { return ReadBlobMSBShort(i) }

var ReadBlobMSBShorts func(i *Image, octets uint32, data *uint16) uint32

func (i *Image) ReadBlobMSBShorts(octets uint32, data *uint16) uint32 {
	return ReadBlobMSBShorts(i, octets, data)
}

var ReadBlobString func(i *Image, str string) string

func (i *Image) ReadBlobString(str string) string { return ReadBlobString(i, str) }

var ReadBlobZC func(i *Image, length uint32, data **Void) uint32

func (i *Image) ReadBlobZC(length uint32, data **Void) uint32 { return ReadBlobZC(i, length, data) }

var RegisterStaticModules func()
var ResetTimer func(t *TimerInfo)

func (t *TimerInfo) Reset() { ResetTimer(t) }

var ResizeFilterToString func(filter FilterTypes) string
var SeekBlob func(i *Image, offset int64, whence int) int64

func (i *Image) SeekBlob(offset int64, whence int) int64 {
	return SeekBlob(i, offset, whence)
}

var SetCacheView func(c *ViewInfo, x, y SSize, columns, rows Size) *PixelPacket

func (c *ViewInfo) Set(x, y SSize, columns, rows Size) *PixelPacket {
	return SetCacheView(c, x, y, columns, rows)
}

var SetClientFilename func(string) string
var SetClientName func(name string) string
var SetClientPath func(path string) string
var SetDelegateInfo func(*DelegateInfo) *DelegateInfo
var SetGeometry func(i *Image, geometry *RectangleInfo)

func (i *Image) SetGeometry(geometry *RectangleInfo) { SetGeometry(i, geometry) }

var SetImageInfo func(i *ImageInfo, rectify bool, exception *ExceptionInfo) bool

func (i *ImageInfo) SetImageInfo(rectify bool, exception *ExceptionInfo) bool {
	return SetImageInfo(i, rectify, exception)
}

var SetLogEventMask func(events string) LogEventType
var SetLogFormat func(format string)
var SortColormapByIntensity func(i *Image) bool

func (i *Image) SortColormapByIntensity() bool { return SortColormapByIntensity(i) }

var StorageTypeToString func(storageType StorageType) string
var StretchTypeToString func(stretch StretchType) string
var StringToArgv func(text string, argc *int) []string
var StringToChannelType func(option string) ChannelType
var StringToColorspaceType func(colorspaceString string) ColorspaceType
var StringToCompositeOperator func(option string) CompositeOperator
var StringToCompressionType func(option string) CompressionType
var StringToDouble func(*string, float64) float64
var StringToEndianType func(option string) EndianType
var StringToFilterTypes func(option string) FilterTypes
var StringToGravityType func(option string) GravityType
var StringToHighlightStyle func(option string) HighlightStyle
var StringToImageType func(option string) ImageType
var StringToInterlaceType func(option string) InterlaceType
var StringToList func(text string) []string
var StringToMetricType func(option string) MetricType
var StringToNoiseType func(option string) NoiseType
var StringToOrientationType func(option string) OrientationType
var StringToPreviewType func(option string) PreviewType
var StringToQuantumOperator func(option string) QuantumOperator
var StringToResourceType func(option string) ResourceType
var StringToVirtualPixelMethod func(option string) VirtualPixelMethod
var Strip func(message string)
var StyleTypeToString func(style StyleType) string
var SubstituteString func(buffer []string, search, replace string) bool
var SyncCacheView func(c *ViewInfo) bool

func (c *ViewInfo) Sync() bool { return SyncCacheView(c) }

var SyncImage func(i *Image) bool

func (i *Image) Sync() bool { return SyncImage(i) }

var SyncNextImageInList func(images *Image) *Image
var SystemCommand func(verbose bool, command string) int
var TellBlob func(i *Image) int64

func (i *Image) TellBlob() int64 { return TellBlob(i) }

var TimeImageCommand func(i *ImageInfo, argc int, argv, metadata []string, exception *ExceptionInfo) uint

func (i *ImageInfo) TimeImageCommand(argc int, argv, metadata []string, exception *ExceptionInfo) uint {
	return TimeImageCommand(i, argc, argv, metadata, exception)
}

var Tokenizer func(t *TokenInfo, flag uint, token string, maxTokenLength uint32, line, white, breakSet, quote string, escape int8, breaker string, next *int, quoted string) int

func (t *TokenInfo) Tokenizer(flag uint, token string, maxTokenLength uint32, line, white, breakSet, quote string, escape int8, breaker string, next *int, quoted string) int {
	return Tokenizer(t, flag, token, maxTokenLength, line, white, breakSet, quote, escape, breaker, next, quoted)
}

var TransformColorspace func(i *Image, colorspace ColorspaceType) uint

func (i *Image) TransformColorspace(colorspace ColorspaceType) uint {
	return TransformColorspace(i, colorspace)
}

var TransformHSL func(red, green, blue Quantum, hue, saturation, lightness *float64)
var TransformHWB func(Quantum, Quantum, Quantum, *float64, *float64, *float64)
var TransformRGBImage func(i *Image, colorspace ColorspaceType) bool

func (i *Image) TransformRGB(colorspace ColorspaceType) bool {
	return TransformRGBImage(i, colorspace)
}

var TranslateText func(i *ImageInfo, image *Image, embedText string) string

func (i *ImageInfo) TranslateText(image *Image, embedText string) string {
	return TranslateText(i, image, embedText)
}

var TranslateTextEx func(*ImageInfo, *Image, string, MagickTextTranslate) string

func (i *ImageInfo) TranslateTextEx(a2 *Image, a3 string, a4 MagickTextTranslate) string {
	return TranslateTextEx(i, a2, a3, a4)
}

var UnlockSemaphoreInfo func(s *SemaphoreInfo) bool

func (s *SemaphoreInfo) Unlock() bool { return UnlockSemaphoreInfo(s) }

var UnmapBlob func(map_ *Void, length uint32) bool
var UnregisterStaticModules func()
var WriteBlob func(i *Image, length uint32, data *byte) int32

func (i *Image) WriteBlob(length uint32, data *byte) int32 { return WriteBlob(i, length, data) }

var WriteBlobByte func(i *Image, value byte) int32

func (i *Image) WriteBlobByte(value byte) int32 { return WriteBlobByte(i, value) }

var WriteBlobFile func(i *Image, filename string) bool

func (i *Image) WriteBlobFile(filename string) bool { return WriteBlobFile(i, filename) }

var WriteBlobLSBLong func(i *Image, value uint32) int32

func (i *Image) WriteBlobLSBLong(value uint32) int32 { return WriteBlobLSBLong(i, value) }

var WriteBlobLSBShort func(i *Image, value uint16) int32

func (i *Image) WriteBlobLSBShort(value uint16) int32 { return WriteBlobLSBShort(i, value) }

var WriteBlobMSBLong func(i *Image, value uint32) int32

func (i *Image) WriteBlobMSBLong(value uint32) int32 { return WriteBlobMSBLong(i, value) }

var WriteBlobMSBShort func(i *Image, value uint16) int32

func (i *Image) WriteBlobMSBShort(value uint16) int32 { return WriteBlobMSBShort(i, value) }

var WriteBlobString func(i *Image, str string) int32

func (i *Image) WriteBlobString(str string) int32 { return WriteBlobString(i, str) }

var ZoomImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image

func (i *Image) Zoom(columns, rows Size, exception *ExceptionInfo) *Image {
	return ZoomImage(i, columns, rows, exception)
}

var GmConvertFp16toFp32 func(fp16 *fp16bits, fp32 *float32) int
var GmConvertFp24toFp32 func(fp24 *fp24bits, fp32 *float32, mode int) int
var GmConvertFp32toFp16 func(fp32 *float32, fp16 *fp16bits, mode int) int
var GmConvertFp32toFp24 func(fp32 *float32, fp24 *fp24bits, mode int) int

func init() {
	AddDllApis(DLL, false, allApis)
	AddDllData(DLL, false, allData)
	allApis = nil
	allData = nil
}

var DLL = "CORE_RL_magick_.dll"
var allApis = Apis{
	{"AccessCacheViewPixels", &AccessCacheViewPixels},
	{"AccessDefaultCacheView", &AccessDefaultCacheView},
	{"AccessDefinition", &AccessDefinition},
	{"AccessImmutableIndexes", &AccessImmutableIndexes},
	{"AccessMutableIndexes", &AccessMutableIndexes},
	{"AccessMutablePixels", &AccessMutablePixels},
	{"AccessThreadViewData", &AccessThreadViewData},
	{"AccessThreadViewDataById", &AccessThreadViewDataById},
	{"AcquireCacheView", &AcquireCacheView},
	{"AcquireCacheViewIndexes", &AcquireCacheViewIndexes},
	{"AcquireCacheViewPixels", &AcquireCacheViewPixels},
	{"AcquireImagePixels", &AcquireImagePixels},
	{"AcquireMagickRandomKernel", &AcquireMagickRandomKernel},
	{"AcquireMagickResource", &AcquireMagickResource},
	{"AcquireMemory", &AcquireMemory},
	{"AcquireOneCacheViewPixel", &AcquireOneCacheViewPixel},
	{"AcquireOnePixel", &AcquireOnePixel},
	{"AcquireOnePixelByReference", &AcquireOnePixelByReference},
	{"AcquireSemaphoreInfo", &AcquireSemaphoreInfo},
	{"AcquireString", &AcquireString},
	{"AcquireTemporaryFileDescriptor", &AcquireTemporaryFileDescriptor},
	{"AcquireTemporaryFileName", &AcquireTemporaryFileName},
	{"AcquireTemporaryFileStream", &AcquireTemporaryFileStream},
	{"AdaptiveThresholdImage", &AdaptiveThresholdImage},
	{"AddDefinition", &AddDefinition},
	{"AddDefinitions", &AddDefinitions},
	{"AddNoiseImage", &AddNoiseImage},
	{"AddNoiseImageChannel", &AddNoiseImageChannel},
	{"AffineTransformImage", &AffineTransformImage},
	{"AllocateImage", &AllocateImage},
	{"AllocateImageColormap", &AllocateImageColormap},
	{"AllocateImageProfileIterator", &AllocateImageProfileIterator},
	{"AllocateNextImage", &AllocateNextImage},
	{"AllocateSemaphoreInfo", &AllocateSemaphoreInfo},
	{"AllocateString", &AllocateString},
	{"AllocateThreadViewDataArray", &AllocateThreadViewDataArray},
	{"AllocateThreadViewDataSet", &AllocateThreadViewDataSet},
	{"AllocateThreadViewSet", &AllocateThreadViewSet},
	{"AnimateImageCommand", &AnimateImageCommand},
	{"AnimateImages", &AnimateImages},
	{"AnnotateImage", &AnnotateImage},
	{"AppendImageFormat", &AppendImageFormat},
	{"AppendImageProfile", &AppendImageProfile},
	{"AppendImageToList", &AppendImageToList},
	{"AppendImages", &AppendImages},
	{"Ascii85Encode", &Ascii85Encode},
	{"Ascii85Flush", &Ascii85Flush},
	{"Ascii85Initialize", &Ascii85Initialize},
	{"Ascii85WriteByteHook", &Ascii85WriteByteHook},
	{"AssignThreadViewData", &AssignThreadViewData},
	{"AttachBlob", &AttachBlob},
	{"AutoOrientImage", &AutoOrientImage},
	{"AverageImages", &AverageImages},
	{"Base64Decode", &Base64Decode},
	{"Base64Encode", &Base64Encode},
	{"BenchmarkImageCommand", &BenchmarkImageCommand},
	{"BlackThresholdImage", &BlackThresholdImage},
	{"BlobIsSeekable", &BlobIsSeekable},
	{"BlobModeToString", &BlobModeToString},
	{"BlobReserveSize", &BlobReserveSize},
	{"BlobToFile", &BlobToFile},
	{"BlobToImage", &BlobToImage},
	{"BlobWriteByteHook", &BlobWriteByteHook},
	{"BlurImage", &BlurImage},
	{"BlurImageChannel", &BlurImageChannel},
	{"BorderImage", &BorderImage},
	{"CatchException", &CatchException},
	{"CatchImageException", &CatchImageException},
	{"CdlImage", &CdlImage},
	{"ChannelImage", &ChannelImage},
	{"ChannelThresholdImage", &ChannelThresholdImage},
	{"ChannelTypeToString", &ChannelTypeToString},
	{"CharcoalImage", &CharcoalImage},
	{"ChopImage", &ChopImage},
	{"ClassTypeToString", &ClassTypeToString},
	{"ClipImage", &ClipImage},
	{"ClipPathImage", &ClipPathImage},
	{"CloneBlobInfo", &CloneBlobInfo},
	{"CloneDrawInfo", &CloneDrawInfo},
	{"CloneImage", &CloneImage},
	{"CloneImageAttributes", &CloneImageAttributes},
	{"CloneImageInfo", &CloneImageInfo},
	{"CloneImageList", &CloneImageList},
	{"CloneMemory", &CloneMemory},
	{"CloneMontageInfo", &CloneMontageInfo},
	{"CloneQuantizeInfo", &CloneQuantizeInfo},
	{"CloneString", &CloneString},
	{"CloseBlob", &CloseBlob},
	{"CloseCacheView", &CloseCacheView},
	{"CoalesceImages", &CoalesceImages},
	{"ColorFloodfillImage", &ColorFloodfillImage},
	{"ColorMatrixImage", &ColorMatrixImage},
	{"ColorizeImage", &ColorizeImage},
	{"ColorspaceTypeToString", &ColorspaceTypeToString},
	{"CompareImageCommand", &CompareImageCommand},
	{"CompositeImage", &CompositeImage},
	{"CompositeImageCommand", &CompositeImageCommand},
	{"CompositeImageRegion", &CompositeImageRegion},
	{"CompositeOperatorToString", &CompositeOperatorToString},
	{"CompressImageColormap", &CompressImageColormap},
	{"CompressionTypeToString", &CompressionTypeToString},
	{"ConcatenateString", &ConcatenateString},
	{"ConfirmAccessModeToString", &ConfirmAccessModeToString},
	{"ConjureImageCommand", &ConjureImageCommand},
	{"ConstituteImage", &ConstituteImage},
	{"ConstituteTextureImage", &ConstituteTextureImage},
	{"ContinueTimer", &ContinueTimer},
	{"Contrast", &Contrast},
	{"ContrastImage", &ContrastImage},
	{"ConvertImageCommand", &ConvertImageCommand},
	{"ConvolveImage", &ConvolveImage},
	{"CopyException", &CopyException},
	{"CropImage", &CropImage},
	{"CropImageToHBITMAP", &CropImageToHBITMAP},
	{"CycleColormapImage", &CycleColormapImage},
	{"DeallocateImageProfileIterator", &DeallocateImageProfileIterator},
	{"DeconstructImages", &DeconstructImages},
	{"DefineClientName", &DefineClientName},
	{"DefineClientPathAndName", &DefineClientPathAndName},
	{"DeleteImageFromList", &DeleteImageFromList},
	{"DeleteImageProfile", &DeleteImageProfile},
	{"DeleteMagickRegistry", &DeleteMagickRegistry},
	{"DescribeImage", &DescribeImage},
	{"DespeckleImage", &DespeckleImage},
	{"DestroyBlob", &DestroyBlob},
	{"DestroyBlobInfo", &DestroyBlobInfo},
	{"DestroyColorInfo", &DestroyColorInfo},
	{"DestroyConstitute", &DestroyConstitute},
	{"DestroyDelegateInfo", &DestroyDelegateInfo},
	{"DestroyDrawInfo", &DestroyDrawInfo},
	{"DestroyExceptionInfo", &DestroyExceptionInfo},
	{"DestroyImage", &DestroyImage},
	{"DestroyImageAttributes", &DestroyImageAttributes},
	{"DestroyImageInfo", &DestroyImageInfo},
	{"DestroyImageList", &DestroyImageList},
	{"DestroyImagePixels", &DestroyImagePixels},
	{"DestroyLogInfo", &DestroyLogInfo},
	{"DestroyMagicInfo", &DestroyMagicInfo},
	{"DestroyMagick", &DestroyMagick},
	{"DestroyMagickModules", &DestroyMagickModules},
	{"DestroyMagickResources", &DestroyMagickResources},
	{"DestroyModuleInfo", &DestroyModuleInfo},
	{"DestroyMontageInfo", &DestroyMontageInfo},
	{"DestroyQuantizeInfo", &DestroyQuantizeInfo},
	{"DestroySemaphore", &DestroySemaphore},
	{"DestroySemaphoreInfo", &DestroySemaphoreInfo},
	{"DestroyTemporaryFiles", &DestroyTemporaryFiles},
	{"DestroyThreadViewDataSet", &DestroyThreadViewDataSet},
	{"DestroyThreadViewSet", &DestroyThreadViewSet},
	{"DestroyTypeInfo", &DestroyTypeInfo},
	{"DetachBlob", &DetachBlob},
	{"DifferenceImage", &DifferenceImage},
	{"DispatchImage", &DispatchImage},
	{"DisplayImageCommand", &DisplayImageCommand},
	{"DisplayImages", &DisplayImages},
	{"DrawAffine", &DrawAffine},
	{"DrawAffineImage", &DrawAffineImage},
	{"DrawAllocateContext", &DrawAllocateContext},
	{"DrawAnnotation", &DrawAnnotation},
	{"DrawArc", &DrawArc},
	{"DrawBezier", &DrawBezier},
	{"DrawCircle", &DrawCircle},
	{"DrawClipPath", &DrawClipPath},
	{"DrawColor", &DrawColor},
	{"DrawComment", &DrawComment},
	{"DrawComposite", &DrawComposite},
	{"DrawDestroyContext", &DrawDestroyContext},
	{"DrawEllipse", &DrawEllipse},
	{"DrawGetClipPath", &DrawGetClipPath},
	{"DrawGetClipRule", &DrawGetClipRule},
	{"DrawGetClipUnits", &DrawGetClipUnits},
	{"DrawGetFillColor", &DrawGetFillColor},
	{"DrawGetFillOpacity", &DrawGetFillOpacity},
	{"DrawGetFillRule", &DrawGetFillRule},
	{"DrawGetFont", &DrawGetFont},
	{"DrawGetFontFamily", &DrawGetFontFamily},
	{"DrawGetFontSize", &DrawGetFontSize},
	{"DrawGetFontStretch", &DrawGetFontStretch},
	{"DrawGetFontStyle", &DrawGetFontStyle},
	{"DrawGetFontWeight", &DrawGetFontWeight},
	{"DrawGetGravity", &DrawGetGravity},
	{"DrawGetStrokeAntialias", &DrawGetStrokeAntialias},
	{"DrawGetStrokeColor", &DrawGetStrokeColor},
	{"DrawGetStrokeDashArray", &DrawGetStrokeDashArray},
	{"DrawGetStrokeDashOffset", &DrawGetStrokeDashOffset},
	{"DrawGetStrokeLineCap", &DrawGetStrokeLineCap},
	{"DrawGetStrokeLineJoin", &DrawGetStrokeLineJoin},
	{"DrawGetStrokeMiterLimit", &DrawGetStrokeMiterLimit},
	{"DrawGetStrokeOpacity", &DrawGetStrokeOpacity},
	{"DrawGetStrokeWidth", &DrawGetStrokeWidth},
	{"DrawGetTextAntialias", &DrawGetTextAntialias},
	{"DrawGetTextDecoration", &DrawGetTextDecoration},
	{"DrawGetTextEncoding", &DrawGetTextEncoding},
	{"DrawGetTextUnderColor", &DrawGetTextUnderColor},
	{"DrawImage", &DrawImage},
	{"DrawLine", &DrawLine},
	{"DrawMatte", &DrawMatte},
	{"DrawPathClose", &DrawPathClose},
	{"DrawPathCurveToAbsolute", &DrawPathCurveToAbsolute},
	{"DrawPathCurveToQuadraticBezierAbsolute", &DrawPathCurveToQuadraticBezierAbsolute},
	{"DrawPathCurveToQuadraticBezierRelative", &DrawPathCurveToQuadraticBezierRelative},
	{"DrawPathCurveToQuadraticBezierSmoothAbsolute", &DrawPathCurveToQuadraticBezierSmoothAbsolute},
	{"DrawPathCurveToQuadraticBezierSmoothRelative", &DrawPathCurveToQuadraticBezierSmoothRelative},
	{"DrawPathCurveToRelative", &DrawPathCurveToRelative},
	{"DrawPathCurveToSmoothAbsolute", &DrawPathCurveToSmoothAbsolute},
	{"DrawPathCurveToSmoothRelative", &DrawPathCurveToSmoothRelative},
	{"DrawPathEllipticArcAbsolute", &DrawPathEllipticArcAbsolute},
	{"DrawPathEllipticArcRelative", &DrawPathEllipticArcRelative},
	{"DrawPathFinish", &DrawPathFinish},
	{"DrawPathLineToAbsolute", &DrawPathLineToAbsolute},
	{"DrawPathLineToHorizontalAbsolute", &DrawPathLineToHorizontalAbsolute},
	{"DrawPathLineToHorizontalRelative", &DrawPathLineToHorizontalRelative},
	{"DrawPathLineToRelative", &DrawPathLineToRelative},
	{"DrawPathLineToVerticalAbsolute", &DrawPathLineToVerticalAbsolute},
	{"DrawPathLineToVerticalRelative", &DrawPathLineToVerticalRelative},
	{"DrawPathMoveToAbsolute", &DrawPathMoveToAbsolute},
	{"DrawPathMoveToRelative", &DrawPathMoveToRelative},
	{"DrawPathStart", &DrawPathStart},
	{"DrawPatternPath", &DrawPatternPath},
	{"DrawPeekGraphicContext", &DrawPeekGraphicContext},
	{"DrawPoint", &DrawPoint},
	{"DrawPolygon", &DrawPolygon},
	{"DrawPolyline", &DrawPolyline},
	{"DrawPopClipPath", &DrawPopClipPath},
	{"DrawPopDefs", &DrawPopDefs},
	{"DrawPopGraphicContext", &DrawPopGraphicContext},
	{"DrawPopPattern", &DrawPopPattern},
	{"DrawPushClipPath", &DrawPushClipPath},
	{"DrawPushDefs", &DrawPushDefs},
	{"DrawPushGraphicContext", &DrawPushGraphicContext},
	{"DrawPushPattern", &DrawPushPattern},
	{"DrawRectangle", &DrawRectangle},
	{"DrawRender", &DrawRender},
	{"DrawRotate", &DrawRotate},
	{"DrawRoundRectangle", &DrawRoundRectangle},
	{"DrawScale", &DrawScale},
	{"DrawSetClipPath", &DrawSetClipPath},
	{"DrawSetClipRule", &DrawSetClipRule},
	{"DrawSetClipUnits", &DrawSetClipUnits},
	{"DrawSetFillColor", &DrawSetFillColor},
	{"DrawSetFillColorString", &DrawSetFillColorString},
	{"DrawSetFillOpacity", &DrawSetFillOpacity},
	{"DrawSetFillPatternURL", &DrawSetFillPatternURL},
	{"DrawSetFillRule", &DrawSetFillRule},
	{"DrawSetFont", &DrawSetFont},
	{"DrawSetFontFamily", &DrawSetFontFamily},
	{"DrawSetFontSize", &DrawSetFontSize},
	{"DrawSetFontStretch", &DrawSetFontStretch},
	{"DrawSetFontStyle", &DrawSetFontStyle},
	{"DrawSetFontWeight", &DrawSetFontWeight},
	{"DrawSetGravity", &DrawSetGravity},
	{"DrawSetStrokeAntialias", &DrawSetStrokeAntialias},
	{"DrawSetStrokeColor", &DrawSetStrokeColor},
	{"DrawSetStrokeColorString", &DrawSetStrokeColorString},
	{"DrawSetStrokeDashArray", &DrawSetStrokeDashArray},
	{"DrawSetStrokeDashOffset", &DrawSetStrokeDashOffset},
	{"DrawSetStrokeLineCap", &DrawSetStrokeLineCap},
	{"DrawSetStrokeLineJoin", &DrawSetStrokeLineJoin},
	{"DrawSetStrokeMiterLimit", &DrawSetStrokeMiterLimit},
	{"DrawSetStrokeOpacity", &DrawSetStrokeOpacity},
	{"DrawSetStrokePatternURL", &DrawSetStrokePatternURL},
	{"DrawSetStrokeWidth", &DrawSetStrokeWidth},
	{"DrawSetTextAntialias", &DrawSetTextAntialias},
	{"DrawSetTextDecoration", &DrawSetTextDecoration},
	{"DrawSetTextEncoding", &DrawSetTextEncoding},
	{"DrawSetTextUnderColor", &DrawSetTextUnderColor},
	{"DrawSetTextUnderColorString", &DrawSetTextUnderColorString},
	{"DrawSetViewbox", &DrawSetViewbox},
	{"DrawSkewX", &DrawSkewX},
	{"DrawSkewY", &DrawSkewY},
	{"DrawTranslate", &DrawTranslate},
	{"EOFBlob", &EOFBlob},
	{"EdgeImage", &EdgeImage},
	{"EmbossImage", &EmbossImage},
	{"EndianTypeToString", &EndianTypeToString},
	{"EnhanceImage", &EnhanceImage},
	{"EqualizeImage", &EqualizeImage},
	{"EscapeString", &EscapeString},
	{"ExecuteModuleProcess", &ExecuteModuleProcess},
	{"ExecuteStaticModuleProcess", &ExecuteStaticModuleProcess},
	{"Exit", &Exit},
	{"ExpandAffine", &ExpandAffine},
	{"ExpandFilename", &ExpandFilename},
	{"ExpandFilenames", &ExpandFilenames},
	{"ExportImageChannel", &ExportImageChannel},
	{"ExportImagePixelArea", &ExportImagePixelArea},
	{"ExportPixelAreaOptionsInit", &ExportPixelAreaOptionsInit},
	{"ExportViewPixelArea", &ExportViewPixelArea},
	{"ExtentImage", &ExtentImage},
	{"FileToBlob", &FileToBlob},
	{"FinalizeSignature", &FinalizeSignature},
	{"FlattenImages", &FlattenImages},
	{"FlipImage", &FlipImage},
	{"FlopImage", &FlopImage},
	{"FormatSize", &FormatSize},
	{"FormatString", &FormatString},
	{"FormatStringList", &FormatStringList},
	{"FrameImage", &FrameImage},
	{"FuzzyColorMatch", &FuzzyColorMatch},
	{"GMCommand", &GMCommand},
	{"GammaImage", &GammaImage},
	{"GaussianBlurImage", &GaussianBlurImage},
	{"GaussianBlurImageChannel", &GaussianBlurImageChannel},
	{"GenerateDifferentialNoise", &GenerateDifferentialNoise},
	{"GenerateNoise", &GenerateNoise},
	{"GetBlobFileHandle", &GetBlobFileHandle},
	{"GetBlobInfo", &GetBlobInfo},
	{"GetBlobIsOpen", &GetBlobIsOpen},
	{"GetBlobSize", &GetBlobSize},
	{"GetBlobStatus", &GetBlobStatus},
	{"GetBlobStreamData", &GetBlobStreamData},
	{"GetBlobTemporary", &GetBlobTemporary},
	{"GetCacheView", &GetCacheView},
	{"GetCacheViewArea", &GetCacheViewArea},
	{"GetCacheViewIndexes", &GetCacheViewIndexes},
	{"GetCacheViewPixels", &GetCacheViewPixels},
	{"GetCacheViewRegion", &GetCacheViewRegion},
	{"GetClientFilename", &GetClientFilename},
	{"GetClientName", &GetClientName},
	{"GetClientPath", &GetClientPath},
	{"GetColorHistogram", &GetColorHistogram},
	{"GetColorInfo", &GetColorInfo},
	{"GetColorInfoArray", &GetColorInfoArray},
	{"GetColorList", &GetColorList},
	{"GetColorTuple", &GetColorTuple},
	{"GetConfigureBlob", &GetConfigureBlob},
	{"GetDelegateCommand", &GetDelegateCommand},
	{"GetDelegateInfo", &GetDelegateInfo},
	{"GetDrawInfo", &GetDrawInfo},
	{"GetElapsedTime", &GetElapsedTime},
	{"GetExceptionInfo", &GetExceptionInfo},
	{"GetExecutionPath", &GetExecutionPath},
	{"GetExecutionPathUsingName", &GetExecutionPathUsingName},
	{"GetFirstImageInList", &GetFirstImageInList},
	{"GetGeometry", &GetGeometry},
	{"GetImageAttribute", &GetImageAttribute},
	{"GetImageBoundingBox", &GetImageBoundingBox},
	{"GetImageChannelDepth", &GetImageChannelDepth},
	{"GetImageChannelDifference", &GetImageChannelDifference},
	{"GetImageChannelDistortion", &GetImageChannelDistortion},
	{"GetImageCharacteristics", &GetImageCharacteristics},
	{"GetImageClipMask", &GetImageClipMask},
	{"GetImageClippingPathAttribute", &GetImageClippingPathAttribute},
	{"GetImageDepth", &GetImageDepth},
	{"GetImageDistortion", &GetImageDistortion},
	{"GetImageException", &GetImageException},
	{"GetImageFromList", &GetImageFromList},
	{"GetImageFromMagickRegistry", &GetImageFromMagickRegistry},
	{"GetImageGeometry", &GetImageGeometry},
	{"GetImageIndexInList", &GetImageIndexInList},
	{"GetImageInfo", &GetImageInfo},
	{"GetImageInfoAttribute", &GetImageInfoAttribute},
	{"GetImageListLength", &GetImageListLength},
	{"GetImageMagick", &GetImageMagick},
	{"GetImagePixels", &GetImagePixels},
	{"GetImagePixelsEx", &GetImagePixelsEx},
	{"GetImageProfile", &GetImageProfile},
	{"GetImageQuantizeError", &GetImageQuantizeError},
	{"GetImageStatistics", &GetImageStatistics},
	{"GetImageType", &GetImageType},
	{"GetImageVirtualPixelMethod", &GetImageVirtualPixelMethod},
	{"GetIndexes", &GetIndexes},
	{"GetLastImageInList", &GetLastImageInList},
	{"GetLocaleExceptionMessage", &GetLocaleExceptionMessage},
	{"GetLocaleMessage", &GetLocaleMessage},
	{"GetLocaleMessageFromID", &GetLocaleMessageFromID},
	{"GetMagickCopyright", &GetMagickCopyright},
	{"GetMagickDimension", &GetMagickDimension},
	{"GetMagickFileFormat", &GetMagickFileFormat},
	{"GetMagickGeometry", &GetMagickGeometry},
	{"GetMagickInfo", &GetMagickInfo},
	{"GetMagickInfoArray", &GetMagickInfoArray},
	{"GetMagickRegistry", &GetMagickRegistry},
	{"GetMagickResource", &GetMagickResource},
	{"GetMagickResourceLimit", &GetMagickResourceLimit},
	{"GetMagickVersion", &GetMagickVersion},
	{"GetMagickWebSite", &GetMagickWebSite},
	{"GetModuleInfo", &GetModuleInfo},
	{"GetMontageInfo", &GetMontageInfo},
	{"GetNextImageInList", &GetNextImageInList},
	{"GetNumberColors", &GetNumberColors},
	{"GetOnePixel", &GetOnePixel},
	{"GetOptimalKernelWidth", &GetOptimalKernelWidth},
	{"GetOptimalKernelWidth1D", &GetOptimalKernelWidth1D},
	{"GetOptimalKernelWidth2D", &GetOptimalKernelWidth2D},
	{"GetPageGeometry", &GetPageGeometry},
	{"GetPathComponent", &GetPathComponent},
	{"GetPixelCacheArea", &GetPixelCacheArea},
	{"GetPixels", &GetPixels},
	{"GetPostscriptDelegateInfo", &GetPostscriptDelegateInfo},
	{"GetPreviousImageInList", &GetPreviousImageInList},
	{"GetQuantizeInfo", &GetQuantizeInfo},
	{"GetSignatureInfo", &GetSignatureInfo},
	{"GetThreadViewDataSetAllocatedViews", &GetThreadViewDataSetAllocatedViews},
	{"GetTimerInfo", &GetTimerInfo},
	{"GetTimerResolution", &GetTimerResolution},
	{"GetToken", &GetToken},
	{"GetTypeInfo", &GetTypeInfo},
	{"GetTypeInfoByFamily", &GetTypeInfoByFamily},
	{"GetTypeList", &GetTypeList},
	{"GetTypeMetrics", &GetTypeMetrics},
	{"GetUserTime", &GetUserTime},
	{"GlobExpression", &GlobExpression},
	{"GradientImage", &GradientImage},
	{"GrayscalePseudoClassImage", &GrayscalePseudoClassImage},
	{"HSLTransform", &HSLTransform},
	{"HWBTransform", &HWBTransform},
	{"HaldClutImage", &HaldClutImage},
	{"HighlightStyleToString", &HighlightStyleToString},
	{"HuffmanDecodeImage", &HuffmanDecodeImage},
	{"HuffmanEncode2Image", &HuffmanEncode2Image},
	{"HuffmanEncodeImage", &HuffmanEncodeImage},
	{"Hull", &Hull},
	{"IdentifyImageCommand", &IdentifyImageCommand},
	{"IdentityAffine", &IdentityAffine},
	{"ImageListToArray", &ImageListToArray},
	{"ImageToBlob", &ImageToBlob},
	{"ImageToFile", &ImageToFile},
	{"ImageToHBITMAP", &ImageToHBITMAP},
	{"ImageToHuffman2DBlob", &ImageToHuffman2DBlob},
	{"ImageToJPEGBlob", &ImageToJPEGBlob},
	{"ImageTypeToString", &ImageTypeToString},
	{"ImplodeImage", &ImplodeImage},
	{"ImportImageChannel", &ImportImageChannel},
	{"ImportImageChannelsMasked", &ImportImageChannelsMasked},
	{"ImportImageCommand", &ImportImageCommand},
	{"ImportImagePixelArea", &ImportImagePixelArea},
	{"ImportPixelAreaOptionsInit", &ImportPixelAreaOptionsInit},
	{"ImportViewPixelArea", &ImportViewPixelArea},
	{"InitializeDifferenceImageOptions", &InitializeDifferenceImageOptions},
	{"InitializeDifferenceStatistics", &InitializeDifferenceStatistics},
	{"InitializeMagicInfo", &InitializeMagicInfo},
	{"InitializeMagick", &InitializeMagick},
	{"InitializeMagickClientPathAndName", &InitializeMagickClientPathAndName},
	{"InitializeMagickModules", &InitializeMagickModules},
	{"InitializeMagickRandomKernel", &InitializeMagickRandomKernel},
	{"InitializeMagickResources", &InitializeMagickResources},
	{"InitializeMagickSignalHandlers", &InitializeMagickSignalHandlers},
	{"InitializePixelIteratorOptions", &InitializePixelIteratorOptions},
	{"InitializeSemaphore", &InitializeSemaphore},
	{"InsertImageInList", &InsertImageInList},
	{"InterlaceTypeToString", &InterlaceTypeToString},
	{"InterpolateColor", &InterpolateColor},
	{"InterpolateViewColor", &InterpolateViewColor},
	{"InvokeDelegate", &InvokeDelegate},
	{"InvokePostscriptDelegate", &InvokePostscriptDelegate},
	{"IsAccessible", &IsAccessible},
	{"IsAccessibleAndNotEmpty", &IsAccessibleAndNotEmpty},
	{"IsAccessibleNoLogging", &IsAccessibleNoLogging},
	{"IsEventLogging", &IsEventLogging},
	{"IsGeometry", &IsGeometry},
	{"IsGlob", &IsGlob},
	{"IsGrayImage", &IsGrayImage},
	{"IsImagesEqual", &IsImagesEqual},
	{"IsMagickConflict", &IsMagickConflict},
	{"IsMonochromeImage", &IsMonochromeImage},
	{"IsOpaqueImage", &IsOpaqueImage},
	{"IsPaletteImage", &IsPaletteImage},
	{"IsSubimage", &IsSubimage},
	{"IsTaintImage", &IsTaintImage},
	{"IsWindows95", &IsWindows95},
	{"IsWriteable", &IsWriteable},
	{"LZWEncode2Image", &LZWEncode2Image},
	{"LZWEncodeImage", &LZWEncodeImage},
	{"LevelImage", &LevelImage},
	{"LevelImageChannel", &LevelImageChannel},
	{"LiberateMagickResource", &LiberateMagickResource},
	{"LiberateMemory", &LiberateMemory},
	{"LiberateSemaphoreInfo", &LiberateSemaphoreInfo},
	{"LiberateTemporaryFile", &LiberateTemporaryFile},
	{"ListColorInfo", &ListColorInfo},
	{"ListDelegateInfo", &ListDelegateInfo},
	{"ListFiles", &ListFiles},
	{"ListMagicInfo", &ListMagicInfo},
	{"ListMagickInfo", &ListMagickInfo},
	{"ListMagickResourceInfo", &ListMagickResourceInfo},
	{"ListModuleInfo", &ListModuleInfo},
	{"ListModuleMap", &ListModuleMap},
	{"ListTypeInfo", &ListTypeInfo},
	{"LocaleCompare", &LocaleCompare},
	{"LocaleLower", &LocaleLower},
	{"LocaleNCompare", &LocaleNCompare},
	{"LocaleUpper", &LocaleUpper},
	{"LockSemaphoreInfo", &LockSemaphoreInfo},
	{"LogMagickEvent", &LogMagickEvent},
	{"LogMagickEventList", &LogMagickEventList},
	{"MSBOrderLong", &MSBOrderLong},
	{"MSBOrderShort", &MSBOrderShort},
	{"MagickAllocFunctions", &MagickAllocFunctions},
	{"MagickArraySize", &MagickArraySize},
	{"MagickBitStreamInitializeRead", &MagickBitStreamInitializeRead},
	{"MagickBitStreamInitializeWrite", &MagickBitStreamInitializeWrite},
	{"MagickBitStreamMSBRead", &MagickBitStreamMSBRead},
	{"MagickBitStreamMSBWrite", &MagickBitStreamMSBWrite},
	{"MagickCloneMemory", &MagickCloneMemory},
	{"MagickCommand", &MagickCommand},
	{"MagickCompositeImageUnderColor", &MagickCompositeImageUnderColor},
	{"MagickConfirmAccess", &MagickConfirmAccess},
	{"MagickConstrainColormapIndex", &MagickConstrainColormapIndex},
	{"MagickCreateDirectoryPath", &MagickCreateDirectoryPath},
	{"MagickError", &MagickError},
	{"MagickFatalError", &MagickFatalError},
	{"MagickFindRawImageMinMax", &MagickFindRawImageMinMax},
	{"MagickFree", &MagickFree},
	{"MagickFreeAligned", &MagickFreeAligned},
	{"MagickGetBitRevTable", &MagickGetBitRevTable},
	{"MagickGetMMUPageSize", &MagickGetMMUPageSize},
	{"MagickGetQuantumSamplesPerPixel", &MagickGetQuantumSamplesPerPixel},
	{"MagickMalloc", &MagickMalloc},
	{"MagickMallocAligned", &MagickMallocAligned},
	{"MagickMallocAlignedArray", &MagickMallocAlignedArray},
	{"MagickMallocArray", &MagickMallocArray},
	{"MagickMallocCleared", &MagickMallocCleared},
	{"MagickMapAccessEntry", &MagickMapAccessEntry},
	{"MagickMapAddEntry", &MagickMapAddEntry},
	{"MagickMapAllocateIterator", &MagickMapAllocateIterator},
	{"MagickMapAllocateMap", &MagickMapAllocateMap},
	{"MagickMapClearMap", &MagickMapClearMap},
	{"MagickMapCloneMap", &MagickMapCloneMap},
	{"MagickMapCopyBlob", &MagickMapCopyBlob},
	{"MagickMapCopyString", &MagickMapCopyString},
	{"MagickMapDeallocateBlob", &MagickMapDeallocateBlob},
	{"MagickMapDeallocateIterator", &MagickMapDeallocateIterator},
	{"MagickMapDeallocateMap", &MagickMapDeallocateMap},
	{"MagickMapDeallocateString", &MagickMapDeallocateString},
	{"MagickMapDereferenceIterator", &MagickMapDereferenceIterator},
	{"MagickMapIterateNext", &MagickMapIterateNext},
	{"MagickMapIteratePrevious", &MagickMapIteratePrevious},
	{"MagickMapIterateToBack", &MagickMapIterateToBack},
	{"MagickMapIterateToFront", &MagickMapIterateToFront},
	{"MagickMapRemoveEntry", &MagickMapRemoveEntry},
	{"MagickMonitor", &MagickMonitor},
	{"MagickMonitorFormatted", &MagickMonitorFormatted},
	{"MagickRandNewSeed", &MagickRandNewSeed},
	{"MagickRandReentrant", &MagickRandReentrant},
	{"MagickRandomInteger", &MagickRandomInteger},
	{"MagickRandomReal", &MagickRandomReal},
	{"MagickRealloc", &MagickRealloc},
	{"MagickReverseBits", &MagickReverseBits},
	{"MagickSceneFileName", &MagickSceneFileName},
	{"MagickSetConfirmAccessHandler", &MagickSetConfirmAccessHandler},
	{"MagickSizeStrToInt64", &MagickSizeStrToInt64},
	{"MagickSpawnVP", &MagickSpawnVP},
	{"MagickStrlCat", &MagickStrlCat},
	{"MagickStrlCpy", &MagickStrlCpy},
	{"MagickStrlCpyTrunc", &MagickStrlCpyTrunc},
	{"MagickSwabArrayOfDouble", &MagickSwabArrayOfDouble},
	{"MagickSwabArrayOfFloat", &MagickSwabArrayOfFloat},
	{"MagickSwabArrayOfUInt16", &MagickSwabArrayOfUInt16},
	{"MagickSwabArrayOfUInt32", &MagickSwabArrayOfUInt32},
	{"MagickSwabDouble", &MagickSwabDouble},
	{"MagickSwabFloat", &MagickSwabFloat},
	{"MagickSwabUInt16", &MagickSwabUInt16},
	{"MagickSwabUInt32", &MagickSwabUInt32},
	{"MagickToMime", &MagickToMime},
	{"MagickTsdGetSpecific", &MagickTsdGetSpecific},
	{"MagickTsdKeyCreate", &MagickTsdKeyCreate},
	{"MagickTsdKeyCreate2", &MagickTsdKeyCreate2},
	{"MagickTsdKeyDelete", &MagickTsdKeyDelete},
	{"MagickTsdSetSpecific", &MagickTsdSetSpecific},
	{"MagickWarning", &MagickWarning},
	{"MagickWordStreamInitializeRead", &MagickWordStreamInitializeRead},
	{"MagickWordStreamInitializeWrite", &MagickWordStreamInitializeWrite},
	{"MagickWordStreamLSBRead", &MagickWordStreamLSBRead},
	{"MagickWordStreamLSBWrite", &MagickWordStreamLSBWrite},
	{"MagickWordStreamLSBWriteFlush", &MagickWordStreamLSBWriteFlush},
	{"MagnifyImage", &MagnifyImage},
	{"MapBlob", &MapBlob},
	{"MapImage", &MapImage},
	{"MapImages", &MapImages},
	{"MapModeToString", &MapModeToString},
	{"MatteFloodfillImage", &MatteFloodfillImage},
	{"MedianFilterImage", &MedianFilterImage},
	{"MetricTypeToString", &MetricTypeToString},
	{"MinifyImage", &MinifyImage},
	{"ModifyImage", &ModifyImage},
	{"Modulate", &Modulate},
	{"ModulateImage", &ModulateImage},
	{"MogrifyImage", &MogrifyImage},
	{"MogrifyImageCommand", &MogrifyImageCommand},
	{"MogrifyImages", &MogrifyImages},
	{"MontageImageCommand", &MontageImageCommand},
	{"MontageImages", &MontageImages},
	{"MorphImages", &MorphImages},
	{"MosaicImages", &MosaicImages},
	{"MotionBlurImage", &MotionBlurImage},
	{"MultilineCensus", &MultilineCensus},
	{"NTElapsedTime", &NTElapsedTime},
	{"NTErrorHandler", &NTErrorHandler},
	{"NTGetExecutionPath", &NTGetExecutionPath},
	{"NTGetLastError", &NTGetLastError},
	{"NTGetTypeList", &NTGetTypeList},
	{"NTGhostscriptDLL", &NTGhostscriptDLL},
	{"NTGhostscriptDLLVectors", &NTGhostscriptDLLVectors},
	{"NTGhostscriptEXE", &NTGhostscriptEXE},
	{"NTGhostscriptFonts", &NTGhostscriptFonts},
	{"NTGhostscriptLoadDLL", &NTGhostscriptLoadDLL},
	{"NTGhostscriptUnLoadDLL", &NTGhostscriptUnLoadDLL},
	{"NTIsMagickConflict", &NTIsMagickConflict},
	{"NTKernelAPISupported", &NTKernelAPISupported},
	{"NTRegistryKeyLookup", &NTRegistryKeyLookup},
	{"NTResourceToBlob", &NTResourceToBlob},
	{"NTSystemComman", &NTSystemCommand},
	{"NTUserTime", &NTUserTime},
	{"NTWarningHandler", &NTWarningHandler},
	{"NTclosedir", &NTclosedir},
	{"NTdlclose", &NTdlclose},
	{"NTdlerror", &NTdlerror},
	{"NTdlexit", &NTdlexit},
	{"NTdlinit", &NTdlinit},
	{"NTdlopen", &NTdlopen},
	{"NTdlsetsearchpath", &NTdlsetsearchpath},
	{"NTdlsym", &NTdlsym},
	{"NTftruncate", &NTftruncate},
	{"NTmmap", &NTmmap},
	{"NTmsync", &NTmsync},
	{"NTmunmap", &NTmunmap},
	{"NTopendir", &NTopendir},
	{"NTreaddir", &NTreaddir},
	{"NTseekdir", &NTseekdir},
	{"NTtelldir", &NTtelldir},
	{"NegateImage", &NegateImage},
	{"NewImageList", &NewImageList},
	{"NextImageProfile", &NextImageProfile},
	{"NoiseTypeToString", &NoiseTypeToString},
	{"NormalizeImage", &NormalizeImage},
	{"OilPaintImage", &OilPaintImage},
	{"OpaqueImage", &OpaqueImage},
	{"OpenBlob", &OpenBlob},
	{"OpenCacheView", &OpenCacheView},
	{"OpenModule", &OpenModule},
	{"OpenModules", &OpenModules},
	{"OrderedDitherImage", &OrderedDitherImage},
	{"OrientationTypeToString", &OrientationTypeToString},
	{"PackbitsEncode2Image", &PackbitsEncode2Image},
	{"PackbitsEncodeImage", &PackbitsEncodeImage},
	{"PersistCache", &PersistCache},
	{"PingBlob", &PingBlob},
	{"PingImage", &PingImage},
	{"PixelIterateDualModify", &PixelIterateDualModify},
	{"PixelIterateDualNew", &PixelIterateDualNew},
	{"PixelIterateDualRead", &PixelIterateDualRead},
	{"PixelIterateMonoModify", &PixelIterateMonoModify},
	{"PixelIterateMonoRead", &PixelIterateMonoRead},
	{"PixelIterateTripleModify", &PixelIterateTripleModify},
	{"PixelIterateTripleNew", &PixelIterateTripleNew},
	{"PlasmaImage", &PlasmaImage},
	{"PopImagePixels", &PopImagePixels},
	{"PrependImageToList", &PrependImageToList},
	{"ProfileImage", &ProfileImage},
	{"PurgeTemporaryFiles", &PurgeTemporaryFiles},
	{"PushImagePixels", &PushImagePixels},
	{"QuantizeImage", &QuantizeImage},
	{"QuantizeImages", &QuantizeImages},
	{"QuantumOperatorImage", &QuantumOperatorImage},
	{"QuantumOperatorImageMultivalue", &QuantumOperatorImageMultivalue},
	{"QuantumOperatorRegionImage", &QuantumOperatorRegionImage},
	{"QuantumOperatorToString", &QuantumOperatorToString},
	{"QuantumSampleTypeToString", &QuantumSampleTypeToString},
	{"QuantumTypeToString", &QuantumTypeToString},
	{"QueryColorDatabase", &QueryColorDatabase},
	{"QueryColorname", &QueryColorname},
	{"RGBTransformImage", &RGBTransformImage},
	{"RaiseImage", &RaiseImage},
	{"RandomChannelThresholdImage", &RandomChannelThresholdImage},
	{"ReacquireMemory", &ReacquireMemory},
	{"ReadBlob", &ReadBlob},
	{"ReadBlobByte", &ReadBlobByte},
	{"ReadBlobLSBDouble", &ReadBlobLSBDouble},
	{"ReadBlobLSBDoubles", &ReadBlobLSBDoubles},
	{"ReadBlobLSBFloat", &ReadBlobLSBFloat},
	{"ReadBlobLSBFloats", &ReadBlobLSBFloats},
	{"ReadBlobLSBLong", &ReadBlobLSBLong},
	{"ReadBlobLSBLongs", &ReadBlobLSBLongs},
	{"ReadBlobLSBShort", &ReadBlobLSBShort},
	{"ReadBlobLSBShorts", &ReadBlobLSBShorts},
	{"ReadBlobMSBDouble", &ReadBlobMSBDouble},
	{"ReadBlobMSBDoubles", &ReadBlobMSBDoubles},
	{"ReadBlobMSBFloat", &ReadBlobMSBFloat},
	{"ReadBlobMSBFloats", &ReadBlobMSBFloats},
	{"ReadBlobMSBLong", &ReadBlobMSBLong},
	{"ReadBlobMSBLongs", &ReadBlobMSBLongs},
	{"ReadBlobMSBShort", &ReadBlobMSBShort},
	{"ReadBlobMSBShorts", &ReadBlobMSBShorts},
	{"ReadBlobString", &ReadBlobString},
	{"ReadBlobZC", &ReadBlobZC},
	{"ReadImage", &ReadImage},
	{"ReadInlineImage", &ReadInlineImage},
	{"ReduceNoiseImage", &ReduceNoiseImage},
	{"ReferenceBlob", &ReferenceBlob},
	{"ReferenceImage", &ReferenceImage},
	{"RegisterMagickInfo", &RegisterMagickInfo},
	{"RegisterStaticModules", &RegisterStaticModules},
	{"RemoveDefinitions", &RemoveDefinitions},
	{"RemoveFirstImageFromList", &RemoveFirstImageFromList},
	{"RemoveLastImageFromList", &RemoveLastImageFromList},
	{"ReplaceImageColormap", &ReplaceImageColormap},
	{"ReplaceImageInList", &ReplaceImageInList},
	{"ResetImagePage", &ResetImagePage},
	{"ResetTimer", &ResetTimer},
	{"ResizeFilterToString", &ResizeFilterToString},
	{"ResizeImage", &ResizeImage},
	{"ReverseImageList", &ReverseImageList},
	{"RollImage", &RollImage},
	{"RotateImage", &RotateImage},
	{"SampleImage", &SampleImage},
	{"ScaleImage", &ScaleImage},
	{"SeekBlob", &SeekBlob},
	{"SegmentImage", &SegmentImage},
	{"SetBlobClosable", &SetBlobClosable},
	{"SetBlobTemporary", &SetBlobTemporary},
	{"SetCacheView", &SetCacheView},
	{"SetCacheViewPixels", &SetCacheViewPixels},
	{"SetClientFilename", &SetClientFilename},
	{"SetClientName", &SetClientName},
	{"SetClientPath", &SetClientPath},
	{"SetDelegateInfo", &SetDelegateInfo},
	{"SetErrorHandler", &SetErrorHandler},
	{"SetExceptionInfo", &SetExceptionInfo},
	{"SetFatalErrorHandler", &SetFatalErrorHandler},
	{"SetGeometry", &SetGeometry},
	{"SetImage", &SetImage},
	{"SetImageAttribute", &SetImageAttribute},
	{"SetImageChannelDepth", &SetImageChannelDepth},
	{"SetImageClipMask", &SetImageClipMask},
	{"SetImageColor", &SetImageColor},
	{"SetImageColorRegion", &SetImageColorRegion},
	{"SetImageDepth", &SetImageDepth},
	{"SetImageInfo", &SetImageInfo},
	{"SetImageOpacity", &SetImageOpacity},
	{"SetImagePixels", &SetImagePixels},
	{"SetImagePixelsEx", &SetImagePixelsEx},
	{"SetImageProfile", &SetImageProfile},
	{"SetImageType", &SetImageType},
	{"SetImageVirtualPixelMethod", &SetImageVirtualPixelMethod},
	{"SetLogEventMask", &SetLogEventMask},
	{"SetLogFormat", &SetLogFormat},
	{"SetMagickInfo", &SetMagickInfo},
	{"SetMagickRegistry", &SetMagickRegistry},
	{"SetMagickResourceLimit", &SetMagickResourceLimit},
	{"SetMonitorHandler", &SetMonitorHandler},
	{"SetWarningHandler", &SetWarningHandler},
	{"ShadeImage", &ShadeImage},
	{"SharpenImage", &SharpenImage},
	{"SharpenImageChannel", &SharpenImageChannel},
	{"ShaveImage", &ShaveImage},
	{"ShearImage", &ShearImage},
	{"SignatureImage", &SignatureImage},
	{"SolarizeImage", &SolarizeImage},
	{"SortColormapByIntensity", &SortColormapByIntensity},
	{"SpliceImageIntoList", &SpliceImageIntoList},
	{"SplitImageList", &SplitImageList},
	{"SpreadImage", &SpreadImage},
	{"SteganoImage", &SteganoImage},
	{"StereoImage", &StereoImage},
	{"StorageTypeToString", &StorageTypeToString},
	{"StretchTypeToString", &StretchTypeToString},
	{"StringToArgv", &StringToArgv},
	{"StringToChannelType", &StringToChannelType},
	{"StringToColorspaceType", &StringToColorspaceType},
	{"StringToCompositeOperator", &StringToCompositeOperator},
	{"StringToCompressionType", &StringToCompressionType},
	{"StringToDouble", &StringToDouble},
	{"StringToEndianType", &StringToEndianType},
	{"StringToFilterTypes", &StringToFilterTypes},
	{"StringToGravityType", &StringToGravityType},
	{"StringToHighlightStyle", &StringToHighlightStyle},
	{"StringToImageType", &StringToImageType},
	{"StringToInterlaceType", &StringToInterlaceType},
	{"StringToList", &StringToList},
	{"StringToMetricType", &StringToMetricType},
	{"StringToNoiseType", &StringToNoiseType},
	{"StringToOrientationType", &StringToOrientationType},
	{"StringToPreviewType", &StringToPreviewType},
	{"StringToQuantumOperator", &StringToQuantumOperator},
	{"StringToResourceType", &StringToResourceType},
	{"StringToVirtualPixelMethod", &StringToVirtualPixelMethod},
	{"Strip", &Strip},
	{"StripImage", &StripImage},
	{"StyleTypeToString", &StyleTypeToString},
	{"SubstituteString", &SubstituteString},
	{"SwirlImage", &SwirlImage},
	{"SyncCacheView", &SyncCacheView},
	{"SyncCacheViewPixels", &SyncCacheViewPixels},
	{"SyncImage", &SyncImage},
	{"SyncImagePixels", &SyncImagePixels},
	{"SyncImagePixelsEx", &SyncImagePixelsEx},
	{"SyncNextImageInList", &SyncNextImageInList},
	{"SystemCommand", &SystemCommand},
	{"TellBlob", &TellBlob},
	{"TextureImage", &TextureImage},
	{"ThresholdImage", &ThresholdImage},
	{"ThrowException", &ThrowException},
	{"ThrowLoggedException", &ThrowLoggedException},
	{"ThumbnailImage", &ThumbnailImage},
	{"TimeImageCommand", &TimeImageCommand},
	{"Tokenizer", &Tokenizer},
	{"TransformColorspace", &TransformColorspace},
	{"TransformHSL", &TransformHSL},
	{"TransformHWB", &TransformHWB},
	{"TransformImage", &TransformImage},
	{"TransformRGBImage", &TransformRGBImage},
	{"TransformSignature", &TransformSignature},
	{"TranslateText", &TranslateText},
	{"TranslateTextEx", &TranslateTextEx},
	{"TransparentImage", &TransparentImage},
	{"UnlockSemaphoreInfo", &UnlockSemaphoreInfo},
	{"UnmapBlob", &UnmapBlob},
	{"UnregisterMagickInfo", &UnregisterMagickInfo},
	{"UnregisterStaticModules", &UnregisterStaticModules},
	{"UnsharpMaskImage", &UnsharpMaskImage},
	{"UnsharpMaskImageChannel", &UnsharpMaskImageChannel},
	{"UpdateSignature", &UpdateSignature},
	{"WaveImage", &WaveImage},
	{"WhiteThresholdImage", &WhiteThresholdImage},
	{"WriteBlob", &WriteBlob},
	{"WriteBlobByte", &WriteBlobByte},
	{"WriteBlobFile", &WriteBlobFile},
	{"WriteBlobLSBLong", &WriteBlobLSBLong},
	{"WriteBlobLSBShort", &WriteBlobLSBShort},
	{"WriteBlobMSBLong", &WriteBlobMSBLong},
	{"WriteBlobMSBShort", &WriteBlobMSBShort},
	{"WriteBlobString", &WriteBlobString},
	{"WriteImage", &WriteImage},
	{"WriteImages", &WriteImages},
	{"WriteImagesFile", &WriteImagesFile},
	{"ZoomImage", &ZoomImage},
	{"_Gm_convert_fp16_to_fp32", &GmConvertFp16toFp32},
	{"_Gm_convert_fp24_to_fp32", &GmConvertFp24toFp32},
	{"_Gm_convert_fp32_to_fp16", &GmConvertFp32toFp16},
	{"_Gm_convert_fp32_to_fp24", &GmConvertFp32toFp24},
	{"_MagickError", &MagickError2},
	{"_MagickFatalError", &MagickFatalError2},
	{"_MagickWarning", &MagickWarning2},
}

var allData = Data{
	{"BackgroundColor", (*byte)(nil)},
	{"BorderColor", (*byte)(nil)},
	{"DefaultCompressionQuality", (*uint32)(nil)}, // unsigned long
	{"DefaultTileFrame", (*byte)(nil)},
	{"DefaultTileGeometry", (*byte)(nil)},
	{"DefaultTileLabel", (*byte)(nil)},
	{"ForegroundColor", (*byte)(nil)},
	{"HighlightColor", (*byte)(nil)},
	{"LoadImageText", (*byte)(nil)},  //Undocumented
	{"LoadImagesText", (*byte)(nil)}, //Undocumented
	{"MatteColor", (*byte)(nil)},
	{"PSDensityGeometry", (*byte)(nil)},
	{"PSPageGeometry", (*byte)(nil)},
	{"SaveImageText", (*byte)(nil)},  //Undocumented
	{"SaveImagesText", (*byte)(nil)}, //Undocumented
}
