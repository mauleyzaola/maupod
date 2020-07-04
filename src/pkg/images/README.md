## Images

This package requires `imagemagick` to be installed locally

If you are not developing, then nothing is needed to be done

### Imagemagick Install

```
wget https://www.imagemagick.org/download/ImageMagick.tar.gz
tar xvzf ImageMagick.tar.gz
cd ImageMagick-7.0.10-23/
./configure 
make
sudo make install
sudo ldconfig /usr/local/lib
```

After installation you should be running latest version
```
convert --version
Version: ImageMagick 7.0.10-23 Q16 x86_64 2020-07-04 https://imagemagick.org
Copyright: © 1999-2020 ImageMagick Studio LLC
License: https://imagemagick.org/script/license.php
Features: Cipher DPC HDRI OpenMP(4.5) 
Delegates (built-in): 
```