# EXIF Scrubber

By default, all modern digital cameras and photo processing systems
add special metadata to JPEG files. This metadata, known as
[EXIF](http://en.wikipedia.org/wiki/Exchangeable_image_file_format),
can contain such information as the exact time and location the photo
was taken, specific data (such as the model and serial number) of the
camera used to take the photo, and details of any processing software
used on the photo after it was taken.

EXIF metadata has been used for benign purposes: geotagging photos,
recovering lost/stolen cameras, and so on. That said, there are
several ways it could be used for ill: a stalker could determine the
exact location of their target, a potential burglar could discover the
owner of a fancy camera, and abusive governments could track down
surveillance targets. (Notably, fugitive antivirus author and alleged
murderer John McAfee's location was
[discovered through EXIF metadata](http://nakedsecurity.sophos.com/2012/12/03/john-mcafee-location-exif/).)

If you'd rather not leak this information, this simple utility will
strip EXIF metadata from target photos.

## Requirements

To build from source, you need **go** -- install instructions [here](http://golang.org/doc/install). Simply run `make`.

## Usage

`./exifscrub <in.jpg> <out.jpg>`

You can test whether it's working with Jeffrey Friedl's excellent online [EXIF viewer](http://regex.info/exif.cgi).


