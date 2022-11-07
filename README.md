# photo-editor-api
Api on golang for edit your photoes

# Functions
## method `resize`
Resize resizes the image to the specified width and height using the specified resampling filter and returns the transformed image. If one of width or height is 0, the image aspect ratio is preserved.
>`src` - path to your file
>`width` - image width to resize
>`height` - image height to resize
>`filter` - resampling filter

Image resizing using various resampling filters. The most notable ones:
`lanczos` - A high-quality resampling filter for photographic images yielding sharp results.
`catmullrom` - A sharp cubic filter that is faster than Lanczos filter while providing similar results.
`mitchellnetravali` - A cubic filter that produces smoother results with less ringing artifacts than CatmullRom.
`linear` - Bilinear resampling filter, produces smooth output. Faster than cubic filters.
`box` - Simple and fast averaging filter appropriate for downscaling. When upscaling it's similar to NearestNeighbor.
`nearestneighbor` - Fastest resampling filter, no antialiasing.

## method `cropanchor`
CropAnchor cuts out a rectangular region with the specified size from the image using the specified anchor point and returns the cropped image.
>`src` - path to your file
>`width` - image width to resize
>`height` - image height to resize
>`anchor` - anchor type

Anchor types:
`center`
`top`
`left`
`right`
`bottom`
`topleft`
`topright`
`bottomleft`
`bottomright`

## method `blur`
Blur produces a blurred version of the image using a Gaussian function. Sigma parameter must be positive and indicates how much the image will be blurred.
>`src` - path to your file
>`scale` - blur scale

## method `sharpening`
Sharpening produces a sharpened version of the image. Sigma parameter must be positive and indicates how much the image will be sharpened.
>`src` - path to your file
>`scale` - sharpen scale

## method `adjust/gamma`
AdjustGamma performs a gamma correction on the image and returns the adjusted image. Gamma parameter must be positive. Gamma = 1.0 gives the original image. Gamma less than 1.0 darkens the image and gamma greater than 1.0 lightens it.
>`src` - path to your file
>`scale` - gamma scale

## method `adjust/contrast`
AdjustContrast changes the contrast of the image using the percentage parameter and returns the adjusted image. The percentage must be in range (-100, 100). The percentage = 0 gives the original image. The percentage = -100 gives solid gray image.
>`src` - path to your file
>`scale` - contrast scale

## method `adjust/brightness`
AdjustBrightness changes the brightness of the image using the percentage parameter and returns the adjusted image. The percentage must be in range (-100, 100). The percentage = 0 gives the original image. The percentage = -100 gives solid black image. The percentage = 100 gives solid white image.
>`src` - path to your file
>`scale` - brightness scale

## method `adjust/saturation`
AdjustSaturation changes the saturation of the image using the percentage parameter and returns the adjusted image. The percentage must be in the range (-100, 100). The percentage = 0 gives the original image. The percentage = 100 gives the image with the saturation value doubled for each pixel. The percentage = -100 gives the image with the saturation value zeroed for each pixel (grayscale).
>`src` - path to your file
>`scale` - saturation scale

## method `grayscale`
Grayscale produces a grayscale version of the image.
>`src` - path to your file

## method `invert`
Invert produces an inverted (negated) version of the image.
>`src` - path to your file

## method `flip/vertical`
FlipVertical flips the image vertically (from top to bottom) and returns the transformed image.
>`src` - path to your file

## method `flip/horizontal`
FlipHorizontal flips the image horizontally (from left to right) and returns the transformed image.
>`src` - path to your file

##method `rotate`
Rotate rotates an image by the given angle counter-clockwise.
>`src` - path to your file
>`angle` - rotate angle
>`color` - parameter specifies the color of the uncovered zone after the rotation

Color types:
`red`
`green`
`blue`
`white`
`black`

## method `fill`
Fill creates an image with the specified dimensions and fills it with the scaled source image. To achieve the correct aspect ratio without stretching, the source image will be cropped.
>`src` - path to your file
>`width` - image width to resize
>`height` - image height to resize
>`anchor` - anchor type
>`filter` - resampling filter

## method `fit`
Fit scales down the image using the specified resample filter to fit the specified maximum width and height and returns the transformed image.
>`src` - path to your file
>`width` - image width to resize
>`height` - image height to resize
>`filter` - resampling filter