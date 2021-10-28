# photoprism-backup
Simple / quick and dirty golang script to get a list of absolute file paths for all images and videos in the PhotoPrism Database (only works with Maria or MySQL) with a given tag applied to them, or in an album.

##Usage:

```./ppbackup /volume1/Pictures '<user>:<password>@tcp(<IP>:<PORT>)/photoprism' tag tag1 tag 2...```

```./ppbackup /volume1/Pictures '<user>:<password>@tcp(<IP>:<PORT>)/photoprism' album album1 album2...```

**Note:** the sql uses the 'slug' columns for the album and tag names

The above connects to the PhotoPrism database and finds all the images and videos.
It then outputs each absolute file path to `backupfiles.txt` in the current directory.

This list can then be used to with `rsync --files-from=/tmp/ppbackup/backupfiles.txt` to copy the files to another directory setup with Asustor - DataSync Center (eg. Google Drive) so the selected file get backed up to the cloud.

The above can also be automated using bash and cron if required.

##Compiling:

Compiling this for the Asustor NAS (AS5304T with ADM 4.0) can be tricky, well at least on MacOS it was for me. So I used docker to cross compile it.

``` docker run --rm -v $PWD:/full/path/to/ppbackup/ -w /full/path/to/ppbackup/ golang:1.16.7-stretch go build -v```
