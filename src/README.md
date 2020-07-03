# SERVER

Maupod server has various services described below

## Artwork

TODO

## Audioscan

TODO

## Mediainfo

TODO

## Restapi

Restful web server

### Examples

Get all media from performer *depeche mode* in album *ultra*

```
curl -XGET -G "http://localhost:8000/media" \
 --data-urlencode 'performer=depeche mode' \
 --data-urlencode 'album=ultra' \
 | jq '. | length'
```