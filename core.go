// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package graphicsmagick provides API definitions for
//accessing CORE_RL_magick_.dll.
//Based on GraphicsMagick v1.3.8.
package core

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
)

const MaxTextExtent = 2053

type (
	fix int

	AffineMatrix         fix
	AlignType            fix
	Char                 byte
	ClipPathUnits        fix
	DecorationType       fix
	ElementReference     fix
	EndianType           fix
	Enum                 int
	ErrorInfo            fix
	ExceptionInfo        fix
	ExceptionInfo2       fix
	ExceptionType        fix
	FILE                 fix
	FillRule             fix
	FilterTypes          fix
	FrameInfo            fix
	GradientInfo         fix
	GravityType          fix
	HistogramColorPacket fix
	ImageAttribute       fix
	ImageInfo            fix
	ImageType            fix
	InterlaceType        fix
	LineCap              fix
	LineJoin             fix
	MagickBooleanType    fix
	MetricType           fix
	OrientationType      fix
	PixelPacket          fix
	Quantum              fix
	RectangleInfo        fix
	RenderingIntent      fix
	ResolutionType       fix
	SegmentInfo          fix
	SemaphoreInfo        fix
	Size                 fix
	SSize                fix
	StorageType          fix
	StretchType          fix
	StyleType            fix
	TimerInfo            fix
	TypeMetric           fix
	Void                 fix
)

// Opaque types
type (
	Ascii85Info struct{}
	BlobInfo    struct{}
)

func init() {
	AddDllApis(DLL, false, allApis)
	AddDllData(DLL, false, allData)
	allApis = nil
	allData = nil
}

var DLL = "CORE_RL_magick_.dll"
var allApis = Apis{
	// {"AccessCacheViewPixels", &AccessCacheViewPixels},
	// {"AccessDefaultCacheView", &AccessDefaultCacheView},
	{"AccessDefinition", &AccessDefinition},
	// {"AccessImmutableIndexes", &AccessImmutableIndexes},
	// {"AccessMutableIndexes", &AccessMutableIndexes},
	// {"AccessMutablePixels", &AccessMutablePixels},
	// {"AccessThreadViewData", &AccessThreadViewData},
	// {"AccessThreadViewDataById", &AccessThreadViewDataById},
	// {"AcquireCacheView", &AcquireCacheView},
	// {"AcquireCacheViewIndexes", &AcquireCacheViewIndexes},
	// {"AcquireCacheViewPixels", &AcquireCacheViewPixels},
	// {"AcquireImagePixels", &AcquireImagePixels},
	// {"AcquireMagickRandomKernel", &AcquireMagickRandomKernel},
	// {"AcquireMagickResource", &AcquireMagickResource},
	// {"AcquireMemory", &AcquireMemory},
	// {"AcquireOneCacheViewPixel", &AcquireOneCacheViewPixel},
	// {"AcquireOnePixel", &AcquireOnePixel},
	// {"AcquireOnePixelByReference", &AcquireOnePixelByReference},
	// {"AcquireSemaphoreInfo", &AcquireSemaphoreInfo},
	// {"AcquireString", &AcquireString},
	// {"AcquireTemporaryFileDescriptor", &AcquireTemporaryFileDescriptor},
	// {"AcquireTemporaryFileName", &AcquireTemporaryFileName},
	// {"AcquireTemporaryFileStream", &AcquireTemporaryFileStream},
	// {"AdaptiveThresholdImage", &AdaptiveThresholdImage},
	{"AddDefinition", &AddDefinition},
	{"AddDefinitions", &AddDefinitions},
	// {"AddNoiseImage", &AddNoiseImage},
	// {"AddNoiseImageChannel", &AddNoiseImageChannel},
	// {"AffineTransformImage", &AffineTransformImage},
	{"AllocateImage", &AllocateImage},
	{"AllocateImageColormap", &AllocateImageColormap},
	// {"AllocateImageProfileIterator", &AllocateImageProfileIterator},
	{"AllocateNextImage", &AllocateNextImage},
	// {"AllocateSemaphoreInfo", &AllocateSemaphoreInfo},
	// {"AllocateString", &AllocateString},
	// {"AllocateThreadViewDataArray", &AllocateThreadViewDataArray},
	// {"AllocateThreadViewDataSet", &AllocateThreadViewDataSet},
	// {"AllocateThreadViewSet", &AllocateThreadViewSet},
	// {"AnimateImageCommand", &AnimateImageCommand},
	{"AnimateImages", &AnimateImages},
	{"AnnotateImage", &AnnotateImage},
	// {"AppendImageFormat", &AppendImageFormat},
	// {"AppendImageProfile", &AppendImageProfile},
	// {"AppendImageToList", &AppendImageToList},
	{"AppendImages", &AppendImages},
	// {"Ascii85Encode", &Ascii85Encode},
	// {"Ascii85Flush", &Ascii85Flush},
	// {"Ascii85Initialize", &Ascii85Initialize},
	// {"Ascii85WriteByteHook", &Ascii85WriteByteHook},
	// {"AssignThreadViewData", &AssignThreadViewData},
	{"AttachBlob", &AttachBlob},
	// {"AutoOrientImage", &AutoOrientImage},
	{"AverageImages", &AverageImages},
	// {"Base64Decode", &Base64Decode},
	// {"Base64Encode", &Base64Encode},
	// {"BenchmarkImageCommand", &BenchmarkImageCommand},
	// {"BlackThresholdImage", &BlackThresholdImage},
	{"BlobIsSeekable", &BlobIsSeekable},
	// {"BlobModeToString", &BlobModeToString},
	{"BlobReserveSize", &BlobReserveSize},
	{"BlobToFile", &BlobToFile},
	{"BlobToImage", &BlobToImage},
	// {"BlobWriteByteHook", &BlobWriteByteHook},
	// {"BlurImage", &BlurImage},
	// {"BlurImageChannel", &BlurImageChannel},
	{"BorderImage", &BorderImage},
	// {"CatchException", &CatchException},
	{"CatchImageException", &CatchImageException},
	{"CdlImage", &CdlImage},
	{"ChannelImage", &ChannelImage},
	// {"ChannelThresholdImage", &ChannelThresholdImage},
	// {"ChannelTypeToString", &ChannelTypeToString},
	// {"CharcoalImage", &CharcoalImage},
	// {"ChopImage", &ChopImage},
	// {"ClassTypeToString", &ClassTypeToString},
	// {"ClipImage", &ClipImage},
	{"ClipPathImage", &ClipPathImage},
	{"CloneBlobInfo", &CloneBlobInfo},
	// {"CloneDrawInfo", &CloneDrawInfo},
	{"CloneImage", &CloneImage},
	{"CloneImageAttributes", &CloneImageAttributes},
	{"CloneImageInfo", &CloneImageInfo},
	// {"CloneImageList", &CloneImageList},
	// {"CloneMemory", &CloneMemory},
	// {"CloneMontageInfo", &CloneMontageInfo},
	// {"CloneQuantizeInfo", &CloneQuantizeInfo},
	// {"CloneString", &CloneString},
	// {"CloseBlob", &CloseBlob},
	// {"CloseCacheView", &CloseCacheView},
	// {"CoalesceImages", &CoalesceImages},
	// {"ColorFloodfillImage", &ColorFloodfillImage},
	// {"ColorMatrixImage", &ColorMatrixImage},
	// {"ColorizeImage", &ColorizeImage},
	// {"ColorspaceTypeToString", &ColorspaceTypeToString},
	// {"CompareImageCommand", &CompareImageCommand},
	{"CompositeImage", &CompositeImage},
	// {"CompositeImageCommand", &CompositeImageCommand},
	// {"CompositeImageRegion", &CompositeImageRegion},
	// {"CompositeOperatorToString", &CompositeOperatorToString},
	// {"CompressImageColormap", &CompressImageColormap},
	// {"CompressionTypeToString", &CompressionTypeToString},
	// {"ConcatenateString", &ConcatenateString},
	// {"ConfirmAccessModeToString", &ConfirmAccessModeToString},
	// {"ConjureImageCommand", &ConjureImageCommand},
	{"ConstituteImage", &ConstituteImage},
	// {"ConstituteTextureImage", &ConstituteTextureImage},
	// {"ContinueTimer", &ContinueTimer},
	// {"Contrast", &Contrast},
	// {"ContrastImage", &ContrastImage},
	// {"ConvertImageCommand", &ConvertImageCommand},
	// {"ConvolveImage", &ConvolveImage},
	// {"CopyException", &CopyException},
	// {"CropImage", &CropImage},
	// {"CropImageToHBITMAP", &CropImageToHBITMAP},
	{"CycleColormapImage", &CycleColormapImage},
	// {"DeallocateImageProfileIterator", &DeallocateImageProfileIterator},
	// {"DeconstructImages", &DeconstructImages},
	// {"DefineClientName", &DefineClientName},
	// {"DefineClientPathAndName", &DefineClientPathAndName},
	// {"DeleteImageFromList", &DeleteImageFromList},
	// {"DeleteImageProfile", &DeleteImageProfile},
	// {"DeleteMagickRegistry", &DeleteMagickRegistry},
	{"DescribeImage", &DescribeImage},
	// {"DespeckleImage", &DespeckleImage},
	{"DestroyBlob", &DestroyBlob},
	{"DestroyBlobInfo", &DestroyBlobInfo},
	// {"DestroyColorInfo", &DestroyColorInfo},
	// {"DestroyConstitute", &DestroyConstitute},
	// {"DestroyDelegateInfo", &DestroyDelegateInfo},
	// {"DestroyDrawInfo", &DestroyDrawInfo},
	// {"DestroyExceptionInfo", &DestroyExceptionInfo},
	{"DestroyImage", &DestroyImage},
	{"DestroyImageAttributes", &DestroyImageAttributes},
	{"DestroyImageInfo", &DestroyImageInfo},
	// {"DestroyImageList", &DestroyImageList},
	// {"DestroyImagePixels", &DestroyImagePixels},
	// {"DestroyLogInfo", &DestroyLogInfo},
	// {"DestroyMagicInfo", &DestroyMagicInfo},
	// {"DestroyMagick", &DestroyMagick},
	// {"DestroyMagickModules", &DestroyMagickModules},
	// {"DestroyMagickResources", &DestroyMagickResources},
	// {"DestroyModuleInfo", &DestroyModuleInfo},
	// {"DestroyMontageInfo", &DestroyMontageInfo},
	// {"DestroyQuantizeInfo", &DestroyQuantizeInfo},
	// {"DestroySemaphore", &DestroySemaphore},
	// {"DestroySemaphoreInfo", &DestroySemaphoreInfo},
	// {"DestroyTemporaryFiles", &DestroyTemporaryFiles},
	// {"DestroyThreadViewDataSet", &DestroyThreadViewDataSet},
	// {"DestroyThreadViewSet", &DestroyThreadViewSet},
	// {"DestroyTypeInfo", &DestroyTypeInfo},
	{"DetachBlob", &DetachBlob},
	{"DifferenceImage", &DifferenceImage},
	{"DispatchImage", &DispatchImage},
	// {"DisplayImageCommand", &DisplayImageCommand},
	{"DisplayImages", &DisplayImages},
	// {"DrawAffine", &DrawAffine},
	// {"DrawAffineImage", &DrawAffineImage},
	// {"DrawAllocateContext", &DrawAllocateContext},
	// {"DrawAnnotation", &DrawAnnotation},
	// {"DrawArc", &DrawArc},
	// {"DrawBezier", &DrawBezier},
	// {"DrawCircle", &DrawCircle},
	// {"DrawClipPath", &DrawClipPath},
	// {"DrawColor", &DrawColor},
	// {"DrawComment", &DrawComment},
	// {"DrawComposite", &DrawComposite},
	// {"DrawDestroyContext", &DrawDestroyContext},
	// {"DrawEllipse", &DrawEllipse},
	// {"DrawGetClipPath", &DrawGetClipPath},
	// {"DrawGetClipRule", &DrawGetClipRule},
	// {"DrawGetClipUnits", &DrawGetClipUnits},
	// {"DrawGetFillColor", &DrawGetFillColor},
	// {"DrawGetFillOpacity", &DrawGetFillOpacity},
	// {"DrawGetFillRule", &DrawGetFillRule},
	// {"DrawGetFont", &DrawGetFont},
	// {"DrawGetFontFamily", &DrawGetFontFamily},
	// {"DrawGetFontSize", &DrawGetFontSize},
	// {"DrawGetFontStretch", &DrawGetFontStretch},
	// {"DrawGetFontStyle", &DrawGetFontStyle},
	// {"DrawGetFontWeight", &DrawGetFontWeight},
	// {"DrawGetGravity", &DrawGetGravity},
	// {"DrawGetStrokeAntialias", &DrawGetStrokeAntialias},
	// {"DrawGetStrokeColor", &DrawGetStrokeColor},
	// {"DrawGetStrokeDashArray", &DrawGetStrokeDashArray},
	// {"DrawGetStrokeDashOffset", &DrawGetStrokeDashOffset},
	// {"DrawGetStrokeLineCap", &DrawGetStrokeLineCap},
	// {"DrawGetStrokeLineJoin", &DrawGetStrokeLineJoin},
	// {"DrawGetStrokeMiterLimit", &DrawGetStrokeMiterLimit},
	// {"DrawGetStrokeOpacity", &DrawGetStrokeOpacity},
	// {"DrawGetStrokeWidth", &DrawGetStrokeWidth},
	// {"DrawGetTextAntialias", &DrawGetTextAntialias},
	// {"DrawGetTextDecoration", &DrawGetTextDecoration},
	// {"DrawGetTextEncoding", &DrawGetTextEncoding},
	// {"DrawGetTextUnderColor", &DrawGetTextUnderColor},
	// {"DrawImage", &DrawImage},
	// {"DrawLine", &DrawLine},
	// {"DrawMatte", &DrawMatte},
	// {"DrawPathClose", &DrawPathClose},
	// {"DrawPathCurveToAbsolute", &DrawPathCurveToAbsolute},
	// {"DrawPathCurveToQuadraticBezierAbsolute", &DrawPathCurveToQuadraticBezierAbsolute},
	// {"DrawPathCurveToQuadraticBezierRelative", &DrawPathCurveToQuadraticBezierRelative},
	// {"DrawPathCurveToQuadraticBezierSmoothAbsolute", &DrawPathCurveToQuadraticBezierSmoothAbsolute},
	// {"DrawPathCurveToQuadraticBezierSmoothRelative", &DrawPathCurveToQuadraticBezierSmoothRelative},
	// {"DrawPathCurveToRelative", &DrawPathCurveToRelative},
	// {"DrawPathCurveToSmoothAbsolute", &DrawPathCurveToSmoothAbsolute},
	// {"DrawPathCurveToSmoothRelative", &DrawPathCurveToSmoothRelative},
	// {"DrawPathEllipticArcAbsolute", &DrawPathEllipticArcAbsolute},
	// {"DrawPathEllipticArcRelative", &DrawPathEllipticArcRelative},
	// {"DrawPathFinish", &DrawPathFinish},
	// {"DrawPathLineToAbsolute", &DrawPathLineToAbsolute},
	// {"DrawPathLineToHorizontalAbsolute", &DrawPathLineToHorizontalAbsolute},
	// {"DrawPathLineToHorizontalRelative", &DrawPathLineToHorizontalRelative},
	// {"DrawPathLineToRelative", &DrawPathLineToRelative},
	// {"DrawPathLineToVerticalAbsolute", &DrawPathLineToVerticalAbsolute},
	// {"DrawPathLineToVerticalRelative", &DrawPathLineToVerticalRelative},
	// {"DrawPathMoveToAbsolute", &DrawPathMoveToAbsolute},
	// {"DrawPathMoveToRelative", &DrawPathMoveToRelative},
	// {"DrawPathStart", &DrawPathStart},
	// {"DrawPatternPath", &DrawPatternPath},
	// {"DrawPeekGraphicContext", &DrawPeekGraphicContext},
	// {"DrawPoint", &DrawPoint},
	// {"DrawPolygon", &DrawPolygon},
	// {"DrawPolyline", &DrawPolyline},
	// {"DrawPopClipPath", &DrawPopClipPath},
	// {"DrawPopDefs", &DrawPopDefs},
	// {"DrawPopGraphicContext", &DrawPopGraphicContext},
	// {"DrawPopPattern", &DrawPopPattern},
	// {"DrawPushClipPath", &DrawPushClipPath},
	// {"DrawPushDefs", &DrawPushDefs},
	// {"DrawPushGraphicContext", &DrawPushGraphicContext},
	// {"DrawPushPattern", &DrawPushPattern},
	// {"DrawRectangle", &DrawRectangle},
	// {"DrawRender", &DrawRender},
	// {"DrawRotate", &DrawRotate},
	// {"DrawRoundRectangle", &DrawRoundRectangle},
	// {"DrawScale", &DrawScale},
	// {"DrawSetClipPath", &DrawSetClipPath},
	// {"DrawSetClipRule", &DrawSetClipRule},
	// {"DrawSetClipUnits", &DrawSetClipUnits},
	// {"DrawSetFillColor", &DrawSetFillColor},
	// {"DrawSetFillColorString", &DrawSetFillColorString},
	// {"DrawSetFillOpacity", &DrawSetFillOpacity},
	// {"DrawSetFillPatternURL", &DrawSetFillPatternURL},
	// {"DrawSetFillRule", &DrawSetFillRule},
	// {"DrawSetFont", &DrawSetFont},
	// {"DrawSetFontFamily", &DrawSetFontFamily},
	// {"DrawSetFontSize", &DrawSetFontSize},
	// {"DrawSetFontStretch", &DrawSetFontStretch},
	// {"DrawSetFontStyle", &DrawSetFontStyle},
	// {"DrawSetFontWeight", &DrawSetFontWeight},
	// {"DrawSetGravity", &DrawSetGravity},
	// {"DrawSetStrokeAntialias", &DrawSetStrokeAntialias},
	// {"DrawSetStrokeColor", &DrawSetStrokeColor},
	// {"DrawSetStrokeColorString", &DrawSetStrokeColorString},
	// {"DrawSetStrokeDashArray", &DrawSetStrokeDashArray},
	// {"DrawSetStrokeDashOffset", &DrawSetStrokeDashOffset},
	// {"DrawSetStrokeLineCap", &DrawSetStrokeLineCap},
	// {"DrawSetStrokeLineJoin", &DrawSetStrokeLineJoin},
	// {"DrawSetStrokeMiterLimit", &DrawSetStrokeMiterLimit},
	// {"DrawSetStrokeOpacity", &DrawSetStrokeOpacity},
	// {"DrawSetStrokePatternURL", &DrawSetStrokePatternURL},
	// {"DrawSetStrokeWidth", &DrawSetStrokeWidth},
	// {"DrawSetTextAntialias", &DrawSetTextAntialias},
	// {"DrawSetTextDecoration", &DrawSetTextDecoration},
	// {"DrawSetTextEncoding", &DrawSetTextEncoding},
	// {"DrawSetTextUnderColor", &DrawSetTextUnderColor},
	// {"DrawSetTextUnderColorString", &DrawSetTextUnderColorString},
	// {"DrawSetViewbox", &DrawSetViewbox},
	// {"DrawSkewX", &DrawSkewX},
	// {"DrawSkewY", &DrawSkewY},
	// {"DrawTranslate", &DrawTranslate},
	// {"EOFBlob", &EOFBlob},
	// {"EdgeImage", &EdgeImage},
	// {"EmbossImage", &EmbossImage},
	// {"EndianTypeToString", &EndianTypeToString},
	// {"EnhanceImage", &EnhanceImage},
	// {"EqualizeImage", &EqualizeImage},
	// {"EscapeString", &EscapeString},
	// {"ExecuteModuleProcess", &ExecuteModuleProcess},
	// {"ExecuteStaticModuleProcess", &ExecuteStaticModuleProcess},
	// {"Exit", &Exit},
	// {"ExpandAffine", &ExpandAffine},
	// {"ExpandFilename", &ExpandFilename},
	// {"ExpandFilenames", &ExpandFilenames},
	{"ExportImageChannel", &ExportImageChannel},
	// {"ExportImagePixelArea", &ExportImagePixelArea},
	// {"ExportPixelAreaOptionsInit", &ExportPixelAreaOptionsInit},
	// {"ExportViewPixelArea", &ExportViewPixelArea},
	// {"ExtentImage", &ExtentImage},
	{"FileToBlob", &FileToBlob},
	// {"FinalizeSignature", &FinalizeSignature},
	// {"FlattenImages", &FlattenImages},
	// {"FlipImage", &FlipImage},
	// {"FlopImage", &FlopImage},
	// {"FormatSize", &FormatSize},
	// {"FormatString", &FormatString},
	// {"FormatStringList", &FormatStringList},
	{"FrameImage", &FrameImage},
	// {"FuzzyColorMatch", &FuzzyColorMatch},
	// {"GMCommand", &GMCommand},
	// {"GammaImage", &GammaImage},
	// {"GaussianBlurImage", &GaussianBlurImage},
	// {"GaussianBlurImageChannel", &GaussianBlurImageChannel},
	// {"GenerateDifferentialNoise", &GenerateDifferentialNoise},
	// {"GenerateNoise", &GenerateNoise},
	{"GetBlobFileHandle", &GetBlobFileHandle},
	{"GetBlobInfo", &GetBlobInfo},
	// {"GetBlobIsOpen", &GetBlobIsOpen},
	// {"GetBlobSize", &GetBlobSize},
	{"GetBlobStatus", &GetBlobStatus},
	{"GetBlobStreamData", &GetBlobStreamData},
	{"GetBlobTemporary", &GetBlobTemporary},
	// {"GetCacheView", &GetCacheView},
	// {"GetCacheViewArea", &GetCacheViewArea},
	// {"GetCacheViewIndexes", &GetCacheViewIndexes},
	// {"GetCacheViewPixels", &GetCacheViewPixels},
	// {"GetCacheViewRegion", &GetCacheViewRegion},
	// {"GetClientFilename", &GetClientFilename},
	// {"GetClientName", &GetClientName},
	// {"GetClientPath", &GetClientPath},
	{"GetColorHistogram", &GetColorHistogram},
	// {"GetColorInfo", &GetColorInfo},
	// {"GetColorInfoArray", &GetColorInfoArray},
	// {"GetColorList", &GetColorList},
	// {"GetColorTuple", &GetColorTuple},
	{"GetConfigureBlob", &GetConfigureBlob},
	// {"GetDelegateCommand", &GetDelegateCommand},
	// {"GetDelegateInfo", &GetDelegateInfo},
	// {"GetDrawInfo", &GetDrawInfo},
	// {"GetElapsedTime", &GetElapsedTime},
	// {"GetExceptionInfo", &GetExceptionInfo},
	// {"GetExecutionPath", &GetExecutionPath},
	// {"GetExecutionPathUsingName", &GetExecutionPathUsingName},
	// {"GetFirstImageInList", &GetFirstImageInList},
	// {"GetGeometry", &GetGeometry},
	{"GetImageAttribute", &GetImageAttribute},
	// {"GetImageBoundingBox", &GetImageBoundingBox},
	{"GetImageChannelDepth", &GetImageChannelDepth},
	{"GetImageChannelDifference", &GetImageChannelDifference},
	{"GetImageChannelDistortion", &GetImageChannelDistortion},
	// {"GetImageCharacteristics", &GetImageCharacteristics},
	{"GetImageClipMask", &GetImageClipMask},
	{"GetImageClippingPathAttribute", &GetImageClippingPathAttribute},
	// {"GetImageDepth", &GetImageDepth},
	{"GetImageDistortion", &GetImageDistortion},
	{"GetImageException", &GetImageException},
	// {"GetImageFromList", &GetImageFromList},
	// {"GetImageFromMagickRegistry", &GetImageFromMagickRegistry},
	{"GetImageGeometry", &GetImageGeometry},
	// {"GetImageIndexInList", &GetImageIndexInList},
	{"GetImageInfo", &GetImageInfo},
	// {"GetImageInfoAttribute", &GetImageInfoAttribute},
	// {"GetImageListLength", &GetImageListLength},
	// {"GetImageMagick", &GetImageMagick},
	// {"GetImagePixels", &GetImagePixels},
	// {"GetImagePixelsEx", &GetImagePixelsEx},
	// {"GetImageProfile", &GetImageProfile},
	// {"GetImageQuantizeError", &GetImageQuantizeError},
	// {"GetImageStatistics", &GetImageStatistics},
	// {"GetImageType", &GetImageType},
	// {"GetImageVirtualPixelMethod", &GetImageVirtualPixelMethod},
	// {"GetIndexes", &GetIndexes},
	// {"GetLastImageInList", &GetLastImageInList},
	// {"GetLocaleExceptionMessage", &GetLocaleExceptionMessage},
	// {"GetLocaleMessage", &GetLocaleMessage},
	// {"GetLocaleMessageFromID", &GetLocaleMessageFromID},
	// {"GetMagickCopyright", &GetMagickCopyright},
	// {"GetMagickDimension", &GetMagickDimension},
	// {"GetMagickFileFormat", &GetMagickFileFormat},
	// {"GetMagickGeometry", &GetMagickGeometry},
	// {"GetMagickInfo", &GetMagickInfo},
	// {"GetMagickInfoArray", &GetMagickInfoArray},
	// {"GetMagickRegistry", &GetMagickRegistry},
	// {"GetMagickResource", &GetMagickResource},
	// {"GetMagickResourceLimit", &GetMagickResourceLimit},
	// {"GetMagickVersion", &GetMagickVersion},
	// {"GetMagickWebSite", &GetMagickWebSite},
	// {"GetModuleInfo", &GetModuleInfo},
	// {"GetMontageInfo", &GetMontageInfo},
	// {"GetNextImageInList", &GetNextImageInList},
	{"GetNumberColors", &GetNumberColors},
	// {"GetOnePixel", &GetOnePixel},
	// {"GetOptimalKernelWidth", &GetOptimalKernelWidth},
	// {"GetOptimalKernelWidth1D", &GetOptimalKernelWidth1D},
	// {"GetOptimalKernelWidth2D", &GetOptimalKernelWidth2D},
	// {"GetPageGeometry", &GetPageGeometry},
	// {"GetPathComponent", &GetPathComponent},
	// {"GetPixelCacheArea", &GetPixelCacheArea},
	// {"GetPixels", &GetPixels},
	// {"GetPostscriptDelegateInfo", &GetPostscriptDelegateInfo},
	// {"GetPreviousImageInList", &GetPreviousImageInList},
	// {"GetQuantizeInfo", &GetQuantizeInfo},
	// {"GetSignatureInfo", &GetSignatureInfo},
	// {"GetThreadViewDataSetAllocatedViews", &GetThreadViewDataSetAllocatedViews},
	// {"GetTimerInfo", &GetTimerInfo},
	// {"GetTimerResolution", &GetTimerResolution},
	// {"GetToken", &GetToken},
	// {"GetTypeInfo", &GetTypeInfo},
	// {"GetTypeInfoByFamily", &GetTypeInfoByFamily},
	// {"GetTypeList", &GetTypeList},
	{"GetTypeMetrics", &GetTypeMetrics},
	// {"GetUserTime", &GetUserTime},
	// {"GlobExpression", &GlobExpression},
	// {"GradientImage", &GradientImage},
	// {"GrayscalePseudoClassImage", &GrayscalePseudoClassImage},
	// {"HSLTransform", &HSLTransform},
	// {"HWBTransform", &HWBTransform},
	// {"HaldClutImage", &HaldClutImage},
	// {"HighlightStyleToString", &HighlightStyleToString},
	// {"HuffmanDecodeImage", &HuffmanDecodeImage},
	// {"HuffmanEncode2Image", &HuffmanEncode2Image},
	// {"HuffmanEncodeImage", &HuffmanEncodeImage},
	// {"Hull", &Hull},
	// {"IdentifyImageCommand", &IdentifyImageCommand},
	// {"IdentityAffine", &IdentityAffine},
	// {"ImageListToArray", &ImageListToArray},
	{"ImageToBlob", &ImageToBlob},
	{"ImageToFile", &ImageToFile},
	// {"ImageToHBITMAP", &ImageToHBITMAP},
	// {"ImageToHuffman2DBlob", &ImageToHuffman2DBlob},
	// {"ImageToJPEGBlob", &ImageToJPEGBlob},
	// {"ImageTypeToString", &ImageTypeToString},
	// {"ImplodeImage", &ImplodeImage},
	{"ImportImageChannel", &ImportImageChannel},
	{"ImportImageChannelsMasked", &ImportImageChannelsMasked},
	// {"ImportImageCommand", &ImportImageCommand},
	// {"ImportImagePixelArea", &ImportImagePixelArea},
	// {"ImportPixelAreaOptionsInit", &ImportPixelAreaOptionsInit},
	// {"ImportViewPixelArea", &ImportViewPixelArea},
	{"InitializeDifferenceImageOptions", &InitializeDifferenceImageOptions},
	{"InitializeDifferenceStatistics", &InitializeDifferenceStatistics},
	// {"InitializeMagicInfo", &InitializeMagicInfo},
	// {"InitializeMagick", &InitializeMagick},
	// {"InitializeMagickClientPathAndName", &InitializeMagickClientPathAndName},
	// {"InitializeMagickModules", &InitializeMagickModules},
	// {"InitializeMagickRandomKernel", &InitializeMagickRandomKernel},
	// {"InitializeMagickResources", &InitializeMagickResources},
	// {"InitializeMagickSignalHandlers", &InitializeMagickSignalHandlers},
	// {"InitializePixelIteratorOptions", &InitializePixelIteratorOptions},
	// {"InitializeSemaphore", &InitializeSemaphore},
	// {"InsertImageInList", &InsertImageInList},
	// {"InterlaceTypeToString", &InterlaceTypeToString},
	// {"InterpolateColor", &InterpolateColor},
	// {"InterpolateViewColor", &InterpolateViewColor},
	// {"InvokeDelegate", &InvokeDelegate},
	// {"InvokePostscriptDelegate", &InvokePostscriptDelegate},
	// {"IsAccessible", &IsAccessible},
	// {"IsAccessibleAndNotEmpty", &IsAccessibleAndNotEmpty},
	// {"IsAccessibleNoLogging", &IsAccessibleNoLogging},
	// {"IsEventLogging", &IsEventLogging},
	// {"IsGeometry", &IsGeometry},
	// {"IsGlob", &IsGlob},
	// {"IsGrayImage", &IsGrayImage},
	{"IsImagesEqual", &IsImagesEqual},
	// {"IsMagickConflict", &IsMagickConflict},
	// {"IsMonochromeImage", &IsMonochromeImage},
	// {"IsOpaqueImage", &IsOpaqueImage},
	{"IsPaletteImage", &IsPaletteImage},
	// {"IsSubimage", &IsSubimage},
	{"IsTaintImage", &IsTaintImage},
	// {"IsWindows95", &IsWindows95},
	// {"IsWriteable", &IsWriteable},
	// {"LZWEncode2Image", &LZWEncode2Image},
	// {"LZWEncodeImage", &LZWEncodeImage},
	// {"LevelImage", &LevelImage},
	// {"LevelImageChannel", &LevelImageChannel},
	// {"LiberateMagickResource", &LiberateMagickResource},
	// {"LiberateMemory", &LiberateMemory},
	// {"LiberateSemaphoreInfo", &LiberateSemaphoreInfo},
	// {"LiberateTemporaryFile", &LiberateTemporaryFile},
	// {"ListColorInfo", &ListColorInfo},
	// {"ListDelegateInfo", &ListDelegateInfo},
	// {"ListFiles", &ListFiles},
	// {"ListMagicInfo", &ListMagicInfo},
	// {"ListMagickInfo", &ListMagickInfo},
	// {"ListMagickResourceInfo", &ListMagickResourceInfo},
	// {"ListModuleInfo", &ListModuleInfo},
	// {"ListModuleMap", &ListModuleMap},
	// {"ListTypeInfo", &ListTypeInfo},
	// {"LocaleCompare", &LocaleCompare},
	// {"LocaleLower", &LocaleLower},
	// {"LocaleNCompare", &LocaleNCompare},
	// {"LocaleUpper", &LocaleUpper},
	// {"LockSemaphoreInfo", &LockSemaphoreInfo},
	// {"LogMagickEvent", &LogMagickEvent},
	// {"LogMagickEventList", &LogMagickEventList},
	// {"MSBOrderLong", &MSBOrderLong},
	// {"MSBOrderShort", &MSBOrderShort},
	// {"MagickAllocFunctions", &MagickAllocFunctions},
	// {"MagickArraySize", &MagickArraySize},
	// {"MagickBitStreamInitializeRead", &MagickBitStreamInitializeRead},
	// {"MagickBitStreamInitializeWrite", &MagickBitStreamInitializeWrite},
	// {"MagickBitStreamMSBRead", &MagickBitStreamMSBRead},
	// {"MagickBitStreamMSBWrite", &MagickBitStreamMSBWrite},
	// {"MagickCloneMemory", &MagickCloneMemory},
	// {"MagickCommand", &MagickCommand},
	// {"MagickCompositeImageUnderColor", &MagickCompositeImageUnderColor},
	{"MagickConfirmAccess", &MagickConfirmAccess},
	// {"MagickConstrainColormapIndex", &MagickConstrainColormapIndex},
	// {"MagickCreateDirectoryPath", &MagickCreateDirectoryPath},
	// {"MagickError", &MagickError},
	// {"MagickFatalError", &MagickFatalError},
	// {"MagickFindRawImageMinMax", &MagickFindRawImageMinMax},
	// {"MagickFree", &MagickFree},
	// {"MagickFreeAligned", &MagickFreeAligned},
	// {"MagickGetBitRevTable", &MagickGetBitRevTable},
	// {"MagickGetMMUPageSize", &MagickGetMMUPageSize},
	// {"MagickGetQuantumSamplesPerPixel", &MagickGetQuantumSamplesPerPixel},
	// {"MagickMalloc", &MagickMalloc},
	// {"MagickMallocAligned", &MagickMallocAligned},
	// {"MagickMallocAlignedArray", &MagickMallocAlignedArray},
	// {"MagickMallocArray", &MagickMallocArray},
	// {"MagickMallocCleared", &MagickMallocCleared},
	// {"MagickMapAccessEntry", &MagickMapAccessEntry},
	// {"MagickMapAddEntry", &MagickMapAddEntry},
	// {"MagickMapAllocateIterator", &MagickMapAllocateIterator},
	// {"MagickMapAllocateMap", &MagickMapAllocateMap},
	// {"MagickMapClearMap", &MagickMapClearMap},
	// {"MagickMapCloneMap", &MagickMapCloneMap},
	// {"MagickMapCopyBlob", &MagickMapCopyBlob},
	// {"MagickMapCopyString", &MagickMapCopyString},
	// {"MagickMapDeallocateBlob", &MagickMapDeallocateBlob},
	// {"MagickMapDeallocateIterator", &MagickMapDeallocateIterator},
	// {"MagickMapDeallocateMap", &MagickMapDeallocateMap},
	// {"MagickMapDeallocateString", &MagickMapDeallocateString},
	// {"MagickMapDereferenceIterator", &MagickMapDereferenceIterator},
	// {"MagickMapIterateNext", &MagickMapIterateNext},
	// {"MagickMapIteratePrevious", &MagickMapIteratePrevious},
	// {"MagickMapIterateToBack", &MagickMapIterateToBack},
	// {"MagickMapIterateToFront", &MagickMapIterateToFront},
	// {"MagickMapRemoveEntry", &MagickMapRemoveEntry},
	// {"MagickMonitor", &MagickMonitor},
	// {"MagickMonitorFormatted", &MagickMonitorFormatted},
	// {"MagickRandNewSeed", &MagickRandNewSeed},
	// {"MagickRandReentrant", &MagickRandReentrant},
	// {"MagickRandomInteger", &MagickRandomInteger},
	// {"MagickRandomReal", &MagickRandomReal},
	// {"MagickRealloc", &MagickRealloc},
	// {"MagickReverseBits", &MagickReverseBits},
	// {"MagickSceneFileName", &MagickSceneFileName},
	{"MagickSetConfirmAccessHandler", &MagickSetConfirmAccessHandler},
	// {"MagickSizeStrToInt64", &MagickSizeStrToInt64},
	// {"MagickSpawnVP", &MagickSpawnVP},
	// {"MagickStrlCat", &MagickStrlCat},
	// {"MagickStrlCpy", &MagickStrlCpy},
	// {"MagickStrlCpyTrunc", &MagickStrlCpyTrunc},
	// {"MagickSwabArrayOfDouble", &MagickSwabArrayOfDouble},
	// {"MagickSwabArrayOfFloat", &MagickSwabArrayOfFloat},
	// {"MagickSwabArrayOfUInt16", &MagickSwabArrayOfUInt16},
	// {"MagickSwabArrayOfUInt32", &MagickSwabArrayOfUInt32},
	// {"MagickSwabDouble", &MagickSwabDouble},
	// {"MagickSwabFloat", &MagickSwabFloat},
	// {"MagickSwabUInt16", &MagickSwabUInt16},
	// {"MagickSwabUInt32", &MagickSwabUInt32},
	// {"MagickToMime", &MagickToMime},
	// {"MagickTsdGetSpecific", &MagickTsdGetSpecific},
	// {"MagickTsdKeyCreate", &MagickTsdKeyCreate},
	// {"MagickTsdKeyCreate2", &MagickTsdKeyCreate2},
	// {"MagickTsdKeyDelete", &MagickTsdKeyDelete},
	// {"MagickTsdSetSpecific", &MagickTsdSetSpecific},
	// {"MagickWarning", &MagickWarning},
	// {"MagickWordStreamInitializeRead", &MagickWordStreamInitializeRead},
	// {"MagickWordStreamInitializeWrite", &MagickWordStreamInitializeWrite},
	// {"MagickWordStreamLSBRead", &MagickWordStreamLSBRead},
	// {"MagickWordStreamLSBWrite", &MagickWordStreamLSBWrite},
	// {"MagickWordStreamLSBWriteFlush", &MagickWordStreamLSBWriteFlush},
	// {"MagnifyImage", &MagnifyImage},
	// {"MapBlob", &MapBlob},
	// {"MapImage", &MapImage},
	// {"MapImages", &MapImages},
	// {"MapModeToString", &MapModeToString},
	// {"MatteFloodfillImage", &MatteFloodfillImage},
	// {"MedianFilterImage", &MedianFilterImage},
	// {"MetricTypeToString", &MetricTypeToString},
	// {"MinifyImage", &MinifyImage},
	{"ModifyImage", &ModifyImage},
	// {"Modulate", &Modulate},
	// {"ModulateImage", &ModulateImage},
	// {"MogrifyImage", &MogrifyImage},
	// {"MogrifyImageCommand", &MogrifyImageCommand},
	// {"MogrifyImages", &MogrifyImages},
	// {"MontageImageCommand", &MontageImageCommand},
	// {"MontageImages", &MontageImages},
	// {"MorphImages", &MorphImages},
	// {"MosaicImages", &MosaicImages},
	// {"MotionBlurImage", &MotionBlurImage},
	// {"MultilineCensus", &MultilineCensus},
	// {"NTElapsedTime", &NTElapsedTime},
	// {"NTErrorHandler", &NTErrorHandler},
	// {"NTGetExecutionPath", &NTGetExecutionPath},
	// {"NTGetLastError", &NTGetLastError},
	// {"NTGetTypeList", &NTGetTypeList},
	// {"NTGhostscriptDLL", &NTGhostscriptDLL},
	// {"NTGhostscriptDLLVectors", &NTGhostscriptDLLVectors},
	// {"NTGhostscriptEXE", &NTGhostscriptEXE},
	// {"NTGhostscriptFonts", &NTGhostscriptFonts},
	// {"NTGhostscriptLoadDLL", &NTGhostscriptLoadDLL},
	// {"NTGhostscriptUnLoadDLL", &NTGhostscriptUnLoadDLL},
	// {"NTIsMagickConflict", &NTIsMagickConflict},
	// {"NTKernelAPISupported", &NTKernelAPISupported},
	// {"NTRegistryKeyLookup", &NTRegistryKeyLookup},
	// {"NTResourceToBlob", &NTResourceToBlob},
	// {"NTSystemComman", &NTSystemComman},
	// {"NTUserTime", &NTUserTime},
	// {"NTWarningHandler", &NTWarningHandler},
	// {"NTclosedir", &NTclosedir},
	// {"NTdlclose", &NTdlclose},
	// {"NTdlerror", &NTdlerror},
	// {"NTdlexit", &NTdlexit},
	// {"NTdlinit", &NTdlinit},
	// {"NTdlopen", &NTdlopen},
	// {"NTdlsetsearchpath", &NTdlsetsearchpath},
	// {"NTdlsym", &NTdlsym},
	// {"NTftruncate", &NTftruncate},
	// {"NTmmap", &NTmmap},
	// {"NTmsync", &NTmsync},
	// {"NTmunmap", &NTmunmap},
	// {"NTopendir", &NTopendir},
	// {"NTreaddir", &NTreaddir},
	// {"NTseekdir", &NTseekdir},
	// {"NTtelldir", &NTtelldir},
	// {"NegateImage", &NegateImage},
	// {"NewImageList", &NewImageList},
	// {"NextImageProfile", &NextImageProfile},
	// {"NoiseTypeToString", &NoiseTypeToString},
	// {"NormalizeImage", &NormalizeImage},
	// {"OilPaintImage", &OilPaintImage},
	// {"OpaqueImage", &OpaqueImage},
	// {"OpenBlob", &OpenBlob},
	// {"OpenCacheView", &OpenCacheView},
	// {"OpenModule", &OpenModule},
	// {"OpenModules", &OpenModules},
	// {"OrderedDitherImage", &OrderedDitherImage},
	// {"OrientationTypeToString", &OrientationTypeToString},
	// {"PackbitsEncode2Image", &PackbitsEncode2Image},
	// {"PackbitsEncodeImage", &PackbitsEncodeImage},
	// {"PersistCache", &PersistCache},
	{"PingBlob", &PingBlob},
	{"PingImage", &PingImage},
	// {"PixelIterateDualModify", &PixelIterateDualModify},
	// {"PixelIterateDualNew", &PixelIterateDualNew},
	// {"PixelIterateDualRead", &PixelIterateDualRead},
	// {"PixelIterateMonoModify", &PixelIterateMonoModify},
	// {"PixelIterateMonoRead", &PixelIterateMonoRead},
	// {"PixelIterateTripleModify", &PixelIterateTripleModify},
	// {"PixelIterateTripleNew", &PixelIterateTripleNew},
	// {"PlasmaImage", &PlasmaImage},
	// {"PopImagePixels", &PopImagePixels},
	// {"PrependImageToList", &PrependImageToList},
	// {"ProfileImage", &ProfileImage},
	// {"PurgeTemporaryFiles", &PurgeTemporaryFiles},
	// {"PushImagePixels", &PushImagePixels},
	// {"QuantizeImage", &QuantizeImage},
	// {"QuantizeImages", &QuantizeImages},
	// {"QuantumOperatorImage", &QuantumOperatorImage},
	// {"QuantumOperatorImageMultivalue", &QuantumOperatorImageMultivalue},
	// {"QuantumOperatorRegionImage", &QuantumOperatorRegionImage},
	// {"QuantumOperatorToString", &QuantumOperatorToString},
	// {"QuantumSampleTypeToString", &QuantumSampleTypeToString},
	// {"QuantumTypeToString", &QuantumTypeToString},
	// {"QueryColorDatabase", &QueryColorDatabase},
	// {"QueryColorname", &QueryColorname},
	// {"RGBTransformImage", &RGBTransformImage},
	{"RaiseImage", &RaiseImage},
	// {"RandomChannelThresholdImage", &RandomChannelThresholdImage},
	// {"ReacquireMemory", &ReacquireMemory},
	// {"ReadBlob", &ReadBlob},
	// {"ReadBlobByte", &ReadBlobByte},
	// {"ReadBlobLSBDouble", &ReadBlobLSBDouble},
	// {"ReadBlobLSBDoubles", &ReadBlobLSBDoubles},
	// {"ReadBlobLSBFloat", &ReadBlobLSBFloat},
	// {"ReadBlobLSBFloats", &ReadBlobLSBFloats},
	// {"ReadBlobLSBLong", &ReadBlobLSBLong},
	// {"ReadBlobLSBLongs", &ReadBlobLSBLongs},
	// {"ReadBlobLSBShort", &ReadBlobLSBShort},
	// {"ReadBlobLSBShorts", &ReadBlobLSBShorts},
	// {"ReadBlobMSBDouble", &ReadBlobMSBDouble},
	// {"ReadBlobMSBDoubles", &ReadBlobMSBDoubles},
	// {"ReadBlobMSBFloat", &ReadBlobMSBFloat},
	// {"ReadBlobMSBFloats", &ReadBlobMSBFloats},
	// {"ReadBlobMSBLong", &ReadBlobMSBLong},
	// {"ReadBlobMSBLongs", &ReadBlobMSBLongs},
	// {"ReadBlobMSBShort", &ReadBlobMSBShort},
	// {"ReadBlobMSBShorts", &ReadBlobMSBShorts},
	// {"ReadBlobString", &ReadBlobString},
	// {"ReadBlobZC", &ReadBlobZC},
	{"ReadImage", &ReadImage},
	{"ReadInlineImage", &ReadInlineImage},
	// {"ReduceNoiseImage", &ReduceNoiseImage},
	{"ReferenceBlob", &ReferenceBlob},
	{"ReferenceImage", &ReferenceImage},
	// {"RegisterMagickInfo", &RegisterMagickInfo},
	// {"RegisterStaticModules", &RegisterStaticModules},
	{"RemoveDefinitions", &RemoveDefinitions},
	// {"RemoveFirstImageFromList", &RemoveFirstImageFromList},
	// {"RemoveLastImageFromList", &RemoveLastImageFromList},
	{"ReplaceImageColormap", &ReplaceImageColormap},
	// {"ReplaceImageInList", &ReplaceImageInList},
	{"ResetImagePage", &ResetImagePage},
	// {"ResetTimer", &ResetTimer},
	// {"ResizeFilterToString", &ResizeFilterToString},
	// {"ResizeImage", &ResizeImage},
	// {"ReverseImageList", &ReverseImageList},
	// {"RollImage", &RollImage},
	// {"RotateImage", &RotateImage},
	// {"SampleImage", &SampleImage},
	// {"ScaleImage", &ScaleImage},
	// {"SeekBlob", &SeekBlob},
	// {"SegmentImage", &SegmentImage},
	{"SetBlobClosable", &SetBlobClosable},
	{"SetBlobTemporary", &SetBlobTemporary},
	// {"SetCacheView", &SetCacheView},
	// {"SetCacheViewPixels", &SetCacheViewPixels},
	// {"SetClientFilename", &SetClientFilename},
	// {"SetClientName", &SetClientName},
	// {"SetClientPath", &SetClientPath},
	// {"SetDelegateInfo", &SetDelegateInfo},
	// {"SetErrorHandler", &SetErrorHandler},
	// {"SetExceptionInfo", &SetExceptionInfo},
	// {"SetFatalErrorHandler", &SetFatalErrorHandler},
	// {"SetGeometry", &SetGeometry},
	{"SetImage", &SetImage},
	{"SetImageAttribute", &SetImageAttribute},
	{"SetImageChannelDepth", &SetImageChannelDepth},
	{"SetImageClipMask", &SetImageClipMask},
	{"SetImageColor", &SetImageColor},
	{"SetImageColorRegion", &SetImageColorRegion},
	{"SetImageDepth", &SetImageDepth},
	// {"SetImageInfo", &SetImageInfo},
	{"SetImageOpacity", &SetImageOpacity},
	// {"SetImagePixels", &SetImagePixels},
	// {"SetImagePixelsEx", &SetImagePixelsEx},
	// {"SetImageProfile", &SetImageProfile},
	{"SetImageType", &SetImageType},
	// {"SetImageVirtualPixelMethod", &SetImageVirtualPixelMethod},
	// {"SetLogEventMask", &SetLogEventMask},
	// {"SetLogFormat", &SetLogFormat},
	// {"SetMagickInfo", &SetMagickInfo},
	// {"SetMagickRegistry", &SetMagickRegistry},
	// {"SetMagickResourceLimit", &SetMagickResourceLimit},
	// {"SetMonitorHandler", &SetMonitorHandler},
	// {"SetWarningHandler", &SetWarningHandler},
	// {"ShadeImage", &ShadeImage},
	// {"SharpenImage", &SharpenImage},
	// {"SharpenImageChannel", &SharpenImageChannel},
	// {"ShaveImage", &ShaveImage},
	// {"ShearImage", &ShearImage},
	// {"SignatureImage", &SignatureImage},
	// {"SolarizeImage", &SolarizeImage},
	// {"SortColormapByIntensity", &SortColormapByIntensity},
	// {"SpliceImageIntoList", &SpliceImageIntoList},
	// {"SplitImageList", &SplitImageList},
	// {"SpreadImage", &SpreadImage},
	// {"SteganoImage", &SteganoImage},
	// {"StereoImage", &StereoImage},
	// {"StorageTypeToString", &StorageTypeToString},
	// {"StretchTypeToString", &StretchTypeToString},
	// {"StringToArgv", &StringToArgv},
	// {"StringToChannelType", &StringToChannelType},
	// {"StringToColorspaceType", &StringToColorspaceType},
	// {"StringToCompositeOperator", &StringToCompositeOperator},
	// {"StringToCompressionType", &StringToCompressionType},
	// {"StringToDouble", &StringToDouble},
	// {"StringToEndianType", &StringToEndianType},
	// {"StringToFilterTypes", &StringToFilterTypes},
	// {"StringToGravityType", &StringToGravityType},
	// {"StringToHighlightStyle", &StringToHighlightStyle},
	// {"StringToImageType", &StringToImageType},
	// {"StringToInterlaceType", &StringToInterlaceType},
	// {"StringToList", &StringToList},
	// {"StringToMetricType", &StringToMetricType},
	// {"StringToNoiseType", &StringToNoiseType},
	// {"StringToOrientationType", &StringToOrientationType},
	// {"StringToPreviewType", &StringToPreviewType},
	// {"StringToQuantumOperator", &StringToQuantumOperator},
	// {"StringToResourceType", &StringToResourceType},
	// {"StringToVirtualPixelMethod", &StringToVirtualPixelMethod},
	// {"Strip", &Strip},
	{"StripImage", &StripImage},
	// {"StyleTypeToString", &StyleTypeToString},
	// {"SubstituteString", &SubstituteString},
	// {"SwirlImage", &SwirlImage},
	// {"SyncCacheView", &SyncCacheView},
	// {"SyncCacheViewPixels", &SyncCacheViewPixels},
	// {"SyncImage", &SyncImage},
	// {"SyncImagePixels", &SyncImagePixels},
	// {"SyncImagePixelsEx", &SyncImagePixelsEx},
	// {"SyncNextImageInList", &SyncNextImageInList},
	// {"SystemCommand", &SystemCommand},
	// {"TellBlob", &TellBlob},
	// {"TextureImage", &TextureImage},
	// {"ThresholdImage", &ThresholdImage},
	// {"ThrowException", &ThrowException},
	// {"ThrowLoggedException", &ThrowLoggedException},
	// {"ThumbnailImage", &ThumbnailImage},
	// {"TimeImageCommand", &TimeImageCommand},
	// {"Tokenizer", &Tokenizer},
	// {"TransformColorspace", &TransformColorspace},
	// {"TransformHSL", &TransformHSL},
	// {"TransformHWB", &TransformHWB},
	// {"TransformImage", &TransformImage},
	// {"TransformRGBImage", &TransformRGBImage},
	// {"TransformSignature", &TransformSignature},
	// {"TranslateText", &TranslateText},
	// {"TranslateTextEx", &TranslateTextEx},
	// {"TransparentImage", &TransparentImage},
	// {"UnlockSemaphoreInfo", &UnlockSemaphoreInfo},
	// {"UnmapBlob", &UnmapBlob},
	// {"UnregisterMagickInfo", &UnregisterMagickInfo},
	// {"UnregisterStaticModules", &UnregisterStaticModules},
	// {"UnsharpMaskImage", &UnsharpMaskImage},
	// {"UnsharpMaskImageChannel", &UnsharpMaskImageChannel},
	// {"UpdateSignature", &UpdateSignature},
	// {"WaveImage", &WaveImage},
	// {"WhiteThresholdImage", &WhiteThresholdImage},
	// {"WriteBlob", &WriteBlob},
	// {"WriteBlobByte", &WriteBlobByte},
	// {"WriteBlobFile", &WriteBlobFile},
	// {"WriteBlobLSBLong", &WriteBlobLSBLong},
	// {"WriteBlobLSBShort", &WriteBlobLSBShort},
	// {"WriteBlobMSBLong", &WriteBlobMSBLong},
	// {"WriteBlobMSBShort", &WriteBlobMSBShort},
	// {"WriteBlobString", &WriteBlobString},
	{"WriteImage", &WriteImage},
	{"WriteImages", &WriteImages},
	{"WriteImagesFile", &WriteImagesFile},
	// {"ZoomImage", &ZoomImage},
	// {"_Gm_convert_fp16_to_fp32", &Gm_convert_fp16_to_fp32},
	// {"_Gm_convert_fp24_to_fp32", &Gm_convert_fp24_to_fp32},
	// {"_Gm_convert_fp32_to_fp16", &Gm_convert_fp32_to_fp16},
	// {"_Gm_convert_fp32_to_fp24", &Gm_convert_fp32_to_fp24},
	// {"_MagickError", &MagickError2},
	// {"_MagickFatalError", &MagickFatalError2},
	// {"_MagickWarning", &MagickWarning2},
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
	Cache           *Void // *CacheInfo
	DefaultViews    *Void // *ThreadViewSet
	Attributes      *Void //  *ImageAttribute
	Ascii85         *Ascii85Info
	Blob            *Void // *BlobInfo
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

func (i *Image) GetBlobTemporary() bool { return GetBlobTemporary(i) }

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

var ChannelImage func(i *Image, channel ChannelType) uint

func (i *Image) Channel(channel ChannelType) uint { return ChannelImage(i, channel) }

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

var DifferenceImage func(i, compareImage *Image, difference_options *DifferenceImageOptions, exception *ExceptionInfo) *Image

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

var DispatchImage func(i *Image, xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) uint

func (i *Image) Dispatch(xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) uint {
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
